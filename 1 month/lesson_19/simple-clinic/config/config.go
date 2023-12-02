package config

type Config struct {
	ClinicPath       string
	BranchOfficePath string
}

func Load() Config {

	var cfg Config

	cfg.ClinicPath = "./data/clinic.json"
	cfg.BranchOfficePath = "./data/branchOffice.json"
	return cfg
}
