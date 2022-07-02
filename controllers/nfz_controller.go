package controllers

import (
	"encoding/json"
	"fmt"
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/helpers"
	"gin-boilerplate/pkg/helpers/converter"
	"gin-boilerplate/pkg/helpers/external_response"
	"io/ioutil"
	http "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type NFZController struct {
	//rep repository.Repository
	DB *gorm.DB
}

// SaveNFZClinics is endpoint logic for saving data from api to database
func (c *NFZController) SaveNFZClinics(ctx *gin.Context) {

	province := ctx.DefaultQuery("province", "01")
	benefitForChildren := ctx.DefaultQuery("benefitForChildren", "false")
	//caseBenefit := ctx.DefaultQuery("case", "1")
	format := ctx.DefaultQuery("format", "json")
	version := ctx.DefaultQuery("api-version", "1.3")
	endpoint := fmt.Sprintf("/app-itl-api/queues?page=%d&limit=%d&case=%d&province=%s&benefitForChildren=%s&api-version=%s&format=%s", 1, 25, 1, province, benefitForChildren, version, format)

	resp, err := c.saveNFZData(endpoint, province)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, helpers.Response{

		Data: resp,
	})

}

func (c *NFZController) saveNFZData(endpoint string, province string) (int64, error) {
	url := viper.GetString("NFZ_URL")

	var response external_response.NFZResponse

	resp, err := http.Get(url + endpoint)
	if err != nil {
		return 0, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	json.Unmarshal(responseData, &response)
	responseFromDB, err := c.saveToDB(&response, province)
	time.Sleep(500 * time.Millisecond) // Securing from blocking
	if err != nil {
		return 0, err
	}
	fmt.Println(response.Meta.Count)

	if response.Links.Next != "" {

		c.saveNFZData(response.Links.Next, province)
	}

	return responseFromDB, nil
}

func (c *NFZController) saveToDB(resp *external_response.NFZResponse, province string) (int64, error) {
	var results int64
	for _, data := range resp.Data {

		clinics := models.Clinic{
			PrivateName:         data.Attributes.Provider,
			ProviderCode:        data.Attributes.ProviderCode,
			Regon:               data.Attributes.RegonProvider,
			Nip:                 data.Attributes.NipProvider,
			NfzName:             data.Attributes.Place,
			Address:             data.Attributes.Address,
			City:                data.Attributes.Locality,
			Voivodeship:         converter.VoivodeshipConverter(province),
			Phone:               data.Attributes.Phone,
			RegistryNumber:      data.Attributes.RegistryNumber,
			BenefitsForChildren: converter.BoolConverter(data.Attributes.BenefitsForChildren),
			Covid19:             converter.BoolConverter(data.Attributes.Covid19),
			Toilet:              converter.BoolConverter(data.Attributes.Toilet),
			Ramp:                converter.BoolConverter(data.Attributes.Ramp),
			CarPark:             converter.BoolConverter(data.Attributes.CarPark),
			Elevator:            converter.BoolConverter(data.Attributes.Elevator),
			Latitude:            data.Attributes.Latitude,
			Longitude:           data.Attributes.Longitude,
			CreatedAt:           time.Now(),
		}

		if c.DB.Model(&clinics).Where("nip = ?", data.Attributes.NipProvider).Updates(&clinics).RowsAffected == 0 {
			c.DB.Create(&clinics)
		} else {
			c.DB.Where("nip=?", data.Attributes.NipProvider).First(&clinics)
		}

		benefits := models.Benefit{
			Name:      data.Attributes.Benefit,
			CreatedAt: time.Now(),
		}
		if c.DB.Model(&benefits).Where("name = ?", data.Attributes.Benefit).Updates(&benefits).RowsAffected == 0 {
			c.DB.Create(&benefits)
		} else {
			c.DB.Where("name=?", data.Attributes.Benefit).First(&benefits)
		}

		clinicsbenefits := models.ClinicBenefit{
			BenefitID:     benefits.ID,
			ClinicID:      clinics.ID,
			Awaiting:      data.Attributes.Statistics.ProviderData.Awaiting,
			Removed:       data.Attributes.Statistics.ProviderData.Removed,
			AveragePeriod: data.Attributes.Statistics.ProviderData.AveragePeriod,
			VisitDate:     data.Attributes.Dates.Date,
			DateUpdatedAt: data.Attributes.Dates.DateSituationAsAt,
		}

		if c.DB.Model(&clinicsbenefits).Where("clinic_id = ?", clinics.ID).Where("benefit_id", benefits.ID).Updates(&clinicsbenefits).RowsAffected == 0 {
			c.DB.Create(&clinicsbenefits)
		}

		results = 0

	}

	return results, nil
}
