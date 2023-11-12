package models

type ClinicPrimaryKey struct {
	ID string `json:"id"`
}

type CreateClinic struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
	City string `json:"city"`
}

type Clinic struct {
	ClinicID string `json:"clinic_id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	City     string `json:"city"`
}

type UpdateClinic struct {
	ClinicID string `json:"clinic_id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	City     string `json:"city"`
}

type GetListClinicRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Search string `json:"search"`
}

type GetListClinicResponse struct {
	Count   int       `json:"count"`
	Clinics []*Clinic `json:"clinics"`
}
