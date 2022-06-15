package controllers

import (
	"encoding/json"
	"fmt"
	"gin-boilerplate/pkg/helpers/external_response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	http "net/http"
	"strconv"
)

type ClinicController struct {
	//rep repository.Repository
	DB *gorm.DB
}

const (
	DolnySlask = "01"
)

func (c *ClinicController) AddClinics(ctx *gin.Context) {
	var response external_response.NFZResponse

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "25"))
	province := ctx.DefaultQuery("province", DolnySlask)
	benefitForChildren := ctx.DefaultQuery("benefitForChildren", "false")
	format := ctx.DefaultQuery("format", "json")
	version := ctx.DefaultQuery("api-version", "1.3")
	url := fmt.Sprintf("https://api.nfz.gov.pl/app-itl-api/queues?page=%d&limit=%d&case=1&province=%s&benefitForChildren=%s&api-version=%s&format=%s", page, limit, province, benefitForChildren, format, version)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("error: ", err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error: ", err)
	}

	json.Unmarshal(responseData, &response)

	ctx.JSON(http.StatusOK, response.Data)

}
