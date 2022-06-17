package controllers

import (
	"errors"
	"fmt"
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/helpers"
	http "net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClinicController struct {
	//rep repository.Repository
	DB *gorm.DB
}

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

func (c *ClinicController) GetClinics(ctx *gin.Context) {

	//page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	//limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "25"))
	benefit := ctx.Query("benefit")
	city := ctx.Query("city") + "%"
	address := ctx.Query("address")
	voivodeship := ctx.Query("voivodeship")
	children := ctx.Query("children")
	private_name := ctx.Query("private_name")
	var result []Response
	var query = map[string]string{
		"benefit":      benefit,
		"city":         city + "%",
		"address":      address + "%",
		"voivodeship":  voivodeship + "%",
		"children":     children,
		"private_name": private_name + "%",
	}

	sql, err := c.buildSQL(query)
	if err != nil {
		ctx.JSON(http.StatusOK, helpers.Response{
			Code:    404,
			Message: "Resource not found",
			Data:    sql,
		})
		return
	}

	c.DB.Raw(sql).Scan(&result)

	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Succesfull added and updated",
		Data:    result,
	})
	return

}

func (c *ClinicController) buildSQL(query map[string]string) (string, error) {
	var benefits models.Benefit
	var q []string
	var benefitq []string
	var whereClause string
	var fullSQL string
	var customWhereClause string
	var customBenefitWhereClause string
	var beginSQL string

	fullSQL = "SELECT * FROM clinics WHERE " + whereClause
	for key, value := range query {

		if value != "" && key != "benefit" {
			customWhereClause = fmt.Sprintf("%s LIKE '%s'", key, value)

			if key == "children" {
				customWhereClause = fmt.Sprintf("%s = '%s'", "benefits_for_children", value)
			}

			q = append(q, customWhereClause)

		}
		if key == "benefit" {
			rows := c.DB.Select("id,name").Where("name = ?", query["benefit"]).First(&benefits).RowsAffected
			if rows == 0 {
				err := errors.New("Benefit not found")
				return "", err
			}
			beginSQL = fmt.Sprintf("SELECT clinics.*, clinic_benefits.visit_date FROM clinics FULL OUTER JOIN clinic_benefits on clinic_benefits.clinic_id = clinics.id WHERE clinic_benefits.benefit_id = '%s' AND ", benefits.ID)
			if value != "" {
				customBenefitWhereClause = fmt.Sprintf("%s LIKE '%s'", key, value)

				if key == "children" {
					customBenefitWhereClause = fmt.Sprintf("%s = '%s'", "benefits_for_children", value)
				}

				benefitq = append(benefitq, customBenefitWhereClause)
			}
			whereClause = strings.Join(benefitq, ` AND `)
			fullSQL = beginSQL + whereClause
		}
		whereClause = strings.Join(q, ` AND `)
	}

	return fullSQL, nil

}

func (c *ClinicController) PrivateNameLike(value string) *gorm.DB {
	return c.DB.Where("private_name LIKE ?", value)
}
func (c *ClinicController) AddressLike(value string) *gorm.DB {
	return c.DB.Where("address LIKE ?", value)
}
func (c *ClinicController) VoivodeshipLike(value string) *gorm.DB {
	return c.DB.Where("voivodeship LIKE ?", value)
}
func (c *ClinicController) CityLike(value string) *gorm.DB {
	return c.DB.Where("city LIKE ?", value)
}
func (c *ClinicController) IsChildren(value string) *gorm.DB {
	if value == "" {
		value = "false"
	}
	return c.DB.Where("benefits_for_children = ?", value)
}
