package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/helpers"
	"gin-boilerplate/pkg/helpers/converter"
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

// BenefitResponse for selects
type SelectResponse struct {
	Value string `json:"value"`
}

// GetClinics gets search results
func (c *ClinicController) GetClinics(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("page"))
	limitQuery, _ := strconv.Atoi(ctx.Query("limit"))
	offset, limit := pagination.BuildPaginationQuery(int64(page), int64(limitQuery))
	benefit := strings.ToUpper(ctx.Query("benefit"))

	var queryOptions = map[string]string{
		"city":                  strings.ToUpper(ctx.Query("city")),
		"address":               strings.ToUpper(ctx.Query("address")),
		"voivodeship":           strings.ToUpper(ctx.Query("voivodeship")),
		"benefits_for_children": ctx.Query("benefits_for_children"),
		"private_name":          strings.ToUpper(ctx.Query("private_name")),
	}

	var result []helpers.ClinicInfoResponse
	var paginationResponse *pagination.Result
	var count int64

	sql := helpers.BuildQuery(queryOptions, benefit)
	paginationSQL := helpers.BuildQueryWithPagination(sql, limit, offset)

	c.DB.Raw(paginationSQL).Scan(&result)
	c.DB.Raw(sql).Count(&count)

	paginationResponse = pagination.PaginationResponseBuilder(pagination.Param{Page: int64(page), Limit: limit, Offset: offset}, result, count)

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
	var names []string
	var result []SelectResponse

	c.DB.Table("benefits").Select([]string{"name"}).Where("name LIKE ?", helpers.LikeStatement(benefitName)).Limit(20).Find(&names)
	for _, v := range names {
		result = append(result, SelectResponse{Value: v})
	}
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
	var names []string
	var result []SelectResponse

	c.DB.Table("clinics").Distinct([]string{"city"}).Where("city LIKE ?", helpers.LikeStatement(cityName)).Limit(20).Find(&names)
	for _, v := range names {
		result = append(result, SelectResponse{Value: v})
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "City retrived successfully",
		Data:    result,
	})
	return
}

// GetPrivateName gets all privateNames (limit 20)
func (c *ClinicController) GetPrivateName(ctx *gin.Context) {
	privateName := strings.ToUpper(ctx.Query("name"))
	var names []string
	var result []SelectResponse

	c.DB.Table("clinics").Select([]string{"private_name"}).Where("private_name LIKE ?", helpers.LikeStatement(privateName)).Limit(20).Find(&names)
	for _, v := range names {
		result = append(result, SelectResponse{Value: v})
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Private name rertived successfully",
		Data:    result,
	})
	return
}

// GetStreet gets all streets from NFZ (limit 20)
func (c *ClinicController) GetAddress(ctx *gin.Context) {
	addressName := strings.ToUpper(ctx.Query("name"))
	var names []string
	var result []SelectResponse

	c.DB.Table("clinics").Select([]string{"address"}).Where("address LIKE ?", helpers.LikeStatement(addressName)).Limit(20).Find(&names)
	for _, v := range names {
		result = append(result, SelectResponse{Value: v})
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Address rertived successfully",
		Data:    result,
	})
	return
}

// Getvoivodeship gets all voivodeship from NFZ (limit 20)
func (c *ClinicController) GetVoivodeship(ctx *gin.Context) {

	var names = converter.Voievodship
	var result []SelectResponse

	for index, _ := range names {
		result = append(result, SelectResponse{Value: index})
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    200,
		Message: "Voivodeship rertived successfully",
		Data:    result,
	})
	return
}
