package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	Error = "error >>> "
	Info  = "info >>> "
	Log   = "log >>> "
)

type Config struct {
	ClinicPath  string
	BranchPath  string
	CashierPath string
	DoctorPath  string
	PatientPath string
}

func Load() Config {

	if err := godotenv.Load(".env_prod"); err != nil {
		log.Println("not found env")
	}

	var cfg Config

	cfg.ClinicPath = cast.ToString(getValueOrDefault("CLINIC_PATH", "./data/clinic.json"))
	cfg.BranchPath = cast.ToString(getValueOrDefault("BRANCH_PATH", "./data/branch.json"))
	cfg.CashierPath = cast.ToString(getValueOrDefault("CASHIER_PATH", "./data/cashier.json"))
	cfg.DoctorPath = cast.ToString(getValueOrDefault("DOCTOR_PATH", "./data/doctor.json"))
	cfg.PatientPath = cast.ToString(getValueOrDefault("PATIENT_PATH", "./data/patient.json"))

	return cfg
}

func getValueOrDefault(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)
	// fmt.Println(key, val)
	if exists {
		return val
	}

	return defaultValue
}
