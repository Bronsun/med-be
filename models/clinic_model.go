package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Clinic model structure
type Clinic struct {
	ID                  uuid.UUID `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	NfzID               string    `json:"nfz_id"`
	PrivateName         string    `json:"private_name"`
	ProviderCode        string    `json:"provider-code"`
	Regon               string    `json:"regon"`
	Nip                 string    `json:"nip"`
	NfzName             string    `json:"nfz_name"`
	Address             string    `json:"address"`
	City                string    `json:"city"`
	Voivodeship         string    `json:"voivodeship"`
	Phone               string    `json:"phone"`
	RegistryNumber      string    `json:"registry_number"`
	BenefitsForChildren bool      `json:"benefits_for_children"`
	Covid19             bool      `json:"covid-19"`
	Toilet              bool      `json:"toilet"`
	Ramp                bool      `json:"ramp"`
	CarPark             bool      `json:"car_park"`
	Elevator            bool      `json:"elevator"`
	Latitude            float32   `json:"latitude"`
	Longitude           float32   `json:"longitude"`
	Benefits            []ClinicBenefit
}

// Benefit model is model that contains all valid benefits for patients
type Benefit struct {
	ID      uuid.UUID `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	Clinics []ClinicBenefit
}

// ClinicBenefit model is model that conatians statistic Many-to-many
type ClinicBenefit struct {
	BenefitID     uuid.UUID
	Benefit       Benefit `gorm:"foreignKey:BenefitID;references:ID"`
	ClinicID      uuid.UUID
	Clinic        Clinic     `gorm:"foreignKey:ClinicID;rreferences:ID"`
	Awaiting      int        `json:"awaiting"`
	Removed       int        `json:"removed"`
	AveragePeriod int        `json:"average_period"`
	VisitDate     *time.Time `json:"visit_date"`
	DateUpdatedAt *time.Time `json:"date_updated_at"`
}

// ClinicName returns DB table name
func (e *Clinic) ClinicName() string {
	return "clinic"
}

// BenefitName returns DB table name
func (e *Clinic) BenefitName() string {
	return "benefit"
}

// StatisticName returns DB table name
func (e *Clinic) StatisticName() string {
	return "statistic"
}
