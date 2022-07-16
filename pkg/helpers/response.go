package helpers

import "gin-boilerplate/models"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"response"`
}

// Response for clinic search
type ClinicInfoResponse struct {
	ID                  string  `json:"id"`
	PrivateName         string  `json:"private_name"`
	ProviderCode        string  `json:"provider_code"`
	Regon               string  `json:"regon" gorm:"uniqueIndex"`
	Nip                 string  `json:"nip" gorm:"uniqueIndex"`
	NfzName             string  `json:"nfz_name"`
	Address             string  `json:"address"`
	City                string  `json:"city"`
	Voivodeship         string  `json:"voivodeship"`
	Phone               string  `json:"phone"`
	RegistryNumber      string  `json:"registry_number"`
	BenefitsForChildren bool    `json:"benefits_for_children"`
	Covid19             bool    `json:"covid_19"`
	Toilet              bool    `json:"toilet"`
	Ramp                bool    `json:"ramp"`
	CarPark             bool    `json:"car_park"`
	Elevator            bool    `json:"elevator"`
	Latitude            float32 `json:"latitude"`
	Longitude           float32 `json:"longitude"`
	AveragePeriod       int     `json:"average_period"`
	Awaiting            int     `json:"awaiting"`
	VisitDate           string  `json:"visit_date"`
}

type ClinicResponse struct {
	ClinicInfo models.Clinic     `json:"clinic_info"`
	Benefits   []BenefitResponse `json:"clinic_benefits"`
}

type BenefitResponse struct {
	Name          string `json:"name"`
	AveragePeriod int    `json:"average_period"`
	Awaiting      int    `json:"awaiting"`
	VisitDate     string `json:"visit_date"`
}
