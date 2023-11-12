package models

type CreateBranchOffice struct {
	ClinicID string `json:"clinic_id"`
	Addres   string `json:"addres"`
	Phone    string `json:"phone"`
}

type BranchOffice struct {
	BranchOfficeID string `json:"branch_office_id"`
	ClinicID       string `json:"clinic_id"`
	Addres         string `json:"addres"`
	Phone          string `json:"phone"`
}
type PrimeryKeyBranchOfficeID struct {
	BranchOfficeID string `json:"branch_office_id"`
}
type UpdateBranchOffice struct {
	BranchOfficeID string `json:"branch_office_id"`
	ClinicID       string `json:"clinic_id"`
	Addres         string `json:"addres"`
	Phone          string `json:"phone"`
}
