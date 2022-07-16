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

// SelectResponse for selects
type SelectResponse struct {
	Name string `json: "name"`
}

// GetClinics gets search results
func (c *ClinicController) GetClinics(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("page"))
	limitQuery, _ := strconv.Atoi(ctx.Query("limit"))
	benefit := strings.ToUpper(ctx.Query("benefit"))
	city := strings.ToUpper(ctx.Query("city"))
	address := strings.ToUpper(ctx.Query("address"))
	voivodeship := strings.ToUpper(ctx.Query("voivodeship"))
	children := ctx.Query("benefits_for_children")
	private_name := strings.ToUpper(ctx.Query("private_name"))

	offset, limit := pagination.BuildPaginationQuery(int64(page), int64(limitQuery))

	var benefits models.Benefit
	var result []helpers.ClinicInfoResponse
	var paginationResponse *pagination.Result

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

	if where == nil && benefit == "" && (children == "" || children == "false") {
		c.DB.Raw("SELECT * FROM clinics LIMIT " + fmt.Sprint(limit) + " OFFSET " + fmt.Sprint(offset)).Scan(&result)
		paginationResponse = pagination.PaginationResponseBuilder(c.DB, pagination.Param{Page: int64(page), Limit: limit, Offset: offset}, "clinicsAll", result)

	}

	if benefit != "" {
		rows := c.DB.Select("id,name").Where("name = ?", benefit).First(&benefits).RowsAffected
		if rows == 0 {
			ctx.JSON(http.StatusNotFound, helpers.Response{
				Code:    404,
				Message: "No benefit found with this name",
				Data:    result,
			})
			return
		}

		bID := fmt.Sprintf("'%s'", benefits.ID)
		if whereBenefit != nil {
			c.DB.Raw("SELECT clinics.*, clinic_benefits.awaiting,clinic_benefits.visit_date,clinic_benefits.average_period FROM clinics FULL OUTER JOIN clinic_benefits on clinic_benefits.clinic_id = clinics.id WHERE clinic_benefits.benefit_id = "+bID+" AND "+strings.Join(whereBenefit, " AND ")+" LIMIT "+fmt.Sprint(limit)+" OFFSET "+fmt.Sprint(offset), values...).Scan(&result)
			paginationResponse = pagination.PaginationResponseBuilder(c.DB, pagination.Param{Page: int64(page), Limit: limit, Offset: offset, Values: values, Where: whereBenefit, BenefitID: bID}, "clinicsBenefits", result)
		} else {
			c.DB.Raw("SELECT clinics.*, clinic_benefits.awaiting,clinic_benefits.visit_date,clinic_benefits.average_period FROM clinics FULL OUTER JOIN clinic_benefits on clinic_benefits.clinic_id = clinics.id WHERE clinic_benefits.benefit_id = " + bID + " LIMIT " + fmt.Sprint(limit) + " OFFSET " + fmt.Sprint(offset)).Scan(&result)
			paginationResponse = pagination.PaginationResponseBuilder(c.DB, pagination.Param{Page: int64(page), Limit: limit, Offset: offset, BenefitID: bID}, "clinicsBenefits", result)
		}

	}
	if where != nil {
		c.DB.Raw("SELECT * FROM clinics WHERE "+strings.Join(where, " AND ")+" LIMIT "+fmt.Sprint(limit)+" OFFSET "+fmt.Sprint(offset), values...).Scan(&result)
		paginationResponse = pagination.PaginationResponseBuilder(c.DB, pagination.Param{Page: int64(page), Limit: limit, Offset: offset, Values: values, Where: where}, "clinicsWithWhere", result)

	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Success",
		Data:    paginationResponse,
	})
	return
}

// GetClinic gets data for exac clinic
func (c *ClinicController) GetClinic(ctx *gin.Context) {
	id := ctx.Param("id")
	var clinic models.Clinic
	var clinicsbenefits []helpers.BenefitResponse

	c.DB.Table("clinics").Where("id = ?", id).Scan(&clinic)
	c.DB.Raw("SELECT benefits.name, clinic_benefits.awaiting,clinic_benefits.visit_date,clinic_benefits.average_period FROM benefits FULL OUTER JOIN clinic_benefits on clinic_benefits.benefit_id = benefits.id WHERE clinic_benefits.clinic_id = " + "'" + id + "'").Scan(&clinicsbenefits)

	result := helpers.ClinicResponse{
		ClinicInfo: clinic,
		Benefits:   clinicsbenefits,
	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Clinic retrived succesfull",
		Data:    result,
	})
	return

}

// GetBenefits gets all benefits avaiable from NFZ (limit 20)
func (c *ClinicController) GetBenefits(ctx *gin.Context) {
	benefitName := strings.ToUpper(ctx.Query("name"))
	var result []SelectResponse
	c.DB.Table("benefits").Select([]string{"name"}).Where("name LIKE ?", helpers.LikeStatement(benefitName)).Limit(20).Find(&result)
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Benefits rertived successfully",
		Data:    result,
	})
	return

}

// GetCity gets all city's names from NFZ (limit 20)
func (c *ClinicController) GetCity(ctx *gin.Context) {
	cityName := strings.ToUpper(ctx.Query("name"))
	var SelectResponse []SelectResponse
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
	var SelectResponse []SelectResponse
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
	var SelectResponse []SelectResponse
	c.DB.Table("clinics").Select([]string{"address"}).Where("address LIKE ?", helpers.LikeStatement(addressName)).Limit(20).Find(&SelectResponse)
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Address rertived successfully",
		Data:    SelectResponse,
	})
	return
}
