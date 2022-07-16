package models

import (
	"time"

	uuid "github.com/google/uuid"
)

// Clinic model structure
type Clinic struct {
	ID                  uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4();uniqueIndex"`
	PrivateName         string          `json:"private_name"`
	ProviderCode        string          `json:"provider_code"`
	Regon               string          `json:"regon" gorm:"uniqueIndex"`
	Nip                 string          `json:"nip" gorm:"uniqueIndex"`
	NfzName             string          `json:"nfz_name"`
	Address             string          `json:"address"`
	City                string          `json:"city"`
	Voivodeship         string          `json:"voivodeship"`
	Phone               string          `json:"phone"`
	RegistryNumber      string          `json:"registry_number"`
	BenefitsForChildren bool            `json:"benefits_for_children"`
	Covid19             bool            `json:"covid_19"`
	Toilet              bool            `json:"toilet"`
	Ramp                bool            `json:"ramp"`
	CarPark             bool            `json:"car_park"`
	Elevator            bool            `json:"elevator"`
	Latitude            float32         `json:"latitude"`
	Longitude           float32         `json:"longitude"`
	CreatedAt           time.Time       `json:"created_at"`
	Benefits            []ClinicBenefit `json:",omitempty"`
}

// Benefit model is model that contains all valid benefits for patients
type Benefit struct {
	ID        uuid.UUID `json:"id"  gorm:"type:uuid;primary_key;default:uuid_generate_v4();uniqueIndex"`
	Name      string    `json:"name" gorm:"uniqueIndex"`
	CreatedAt time.Time `json:"created_at"`
	Clinics   []ClinicBenefit
}

// ClinicBenefit model is model that conatians statistic Many-to-many
type ClinicBenefit struct {
	BenefitID     uuid.UUID
	Benefit       Benefit `gorm:"foreignKey:BenefitID;references:ID"`
	ClinicID      uuid.UUID
	Clinic        Clinic `gorm:"foreignKey:ClinicID;references:ID"`
	Awaiting      int    `json:"awaiting"`
	Removed       int    `json:"removed"`
	AveragePeriod int    `json:"average_period"`
	VisitDate     string `json:"visit_date"`
	DateUpdatedAt string `json:"date_updated_at"`
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
