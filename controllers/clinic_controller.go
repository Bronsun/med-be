package controllers

import (
	"fmt"
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/helpers"
	"gin-boilerplate/pkg/helpers/pagination"
	http "net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ClinicController struct
type ClinicController struct {
	DB *gorm.DB
}

// Response for clinic search
type Response struct {
	ID                  string  `json:"id"`
	PrivateName         string  `json:"private_name"`
	NfzName             string  `json:"nfz_name"`
	Address             string  `json:"address"`
	City                string  `json:"city"`
	Voivodeship         string  `json:"voivodeship"`
	Phone               string  `json:"phone"`
	RegistryNumber      string  `json:"registry_number"`
	BenefitsForChildren bool    `json:"benefits_for_children"`
	Covid19             bool    `json:"covid-19"`
	Toilet              bool    `json:"toilet"`
	Ramp                bool    `json:"ramp"`
	CarPark             bool    `json:"car_park"`
	Elevator            bool    `json:"elevator"`
	Latitude            float32 `json:"latitude"`
	Longitude           float32 `json:"longitude"`
	VisitDate           string  `json:"visit_date"`
}

// SelectResponse for selects
var SelectResponse []string

// GetClinics gets search results
func (c *ClinicController) GetClinics(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limitQuery, _ := strconv.Atoi(ctx.DefaultQuery("limit", "25"))
	benefit := ctx.Query("benefit")
	city := ctx.Query("city")
	address := ctx.Query("address")
	voivodeship := ctx.Query("voivodeship")
	children := ctx.Query("benefits_for_children")
	private_name := ctx.Query("private_name")

	offset, limit := pagination.BuildPaginationQuery(int64(page), int64(limitQuery))

	var benefits models.Benefit
	var result []Response

	var query = map[string]string{
		"city":         city,
		"address":      address,
		"voivodeship":  voivodeship,
		"private_name": private_name,
	}
	values, where, whereBenefit := helpers.DynamicWhereLikeBuilder(query, "clinics")

	if children == "true" {
		where = append(where, "benefits_for_children = true")
		whereBenefit = append(where, "clinics.benefits_for_children = true")
	}

	if where == nil && benefit == "" && children == "" {
		c.DB.Raw("SELECT * FROM clinics LIMIT " + fmt.Sprint(limit) + " OFFSET " + fmt.Sprint(offset)).Scan(&result)
	}

	if benefit != "" {
		rows := c.DB.Select("id,name").Where("name = ?", benefit).First(&benefits).RowsAffected
		if rows == 0 {
			ctx.JSON(http.StatusOK, helpers.Response{
				Code:    404,
				Message: "No benefit found with this name",
				Data:    result,
			})
			return
		}
		bID := fmt.Sprintf("'%s'", benefits.ID)

		c.DB.Raw("SELECT clinics.*, clinic_benefits.visit_date FROM clinics FULL OUTER JOIN clinic_benefits on clinic_benefits.clinic_id = clinics.id WHERE clinic_benefits.benefit_id = "+bID+strings.Join(whereBenefit, " AND "), values...).Scan(&result)

	} else {
		c.DB.Raw("SELECT * FROM clinics WHERE "+strings.Join(where, " AND ")+" LIMIT "+fmt.Sprint(limit)+" OFFSET "+fmt.Sprint(offset), values...).Scan(&result)
	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Success ",
		Data:    result,
	})
	return
}

// GetBenefits gets all benefits avaiable from NFZ (limit 20)
func (c *ClinicController) GetBenefits(ctx *gin.Context) {
	benefitName := strings.ToUpper(ctx.Query("name"))

	c.DB.Table("benefits").Select([]string{"name"}).Where("name LIKE ?", helpers.LikeStatement(benefitName)).Limit(20).Find(&SelectResponse)
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Benefits rertived successfully",
		Data:    SelectResponse,
	})
	return

}

// GetCity gets all city's names from NFZ (limit 20)
func (c *ClinicController) GetCity(ctx *gin.Context) {
	cityName := strings.ToUpper(ctx.Query("name"))

	c.DB.Table("clinics").Select([]string{"city"}).Where("city LIKE ?", helpers.LikeStatement(cityName)).Limit(20).Find(&SelectResponse)
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "City retrived successfully",
		Data:    SelectResponse,
	})
	return
}

// GetPrivateName gets all privateNames (limit 20)
func (c *ClinicController) GetPrivateName(ctx *gin.Context) {
	privateName := strings.ToUpper(ctx.Query("name"))

	c.DB.Table("clinics").Select([]string{"private_name"}).Where("private_name LIKE ?", privateName).Limit(20).Find(&SelectResponse)
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Private name rertived successfully",
		Data:    SelectResponse,
	})
	return
}

// GetStreet gets all streets from NFZ (limit 20)
func (c *ClinicController) GetAddress(ctx *gin.Context) {
	addressName := strings.ToUpper(ctx.Query("name"))

	c.DB.Table("clinics").Select([]string{"address"}).Where("address LIKE ?", helpers.LikeStatement(addressName)).Limit(20).Find(&SelectResponse)
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Address rertived successfully",
		Data:    SelectResponse,
	})
	return
}
