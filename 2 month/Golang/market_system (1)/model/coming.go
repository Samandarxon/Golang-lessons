package model

type Coming struct {
	Id       string
	DateTime string
	Employee string
	Adresss  string
	Status   string
}

type ComingPrimaryKey struct {
	Id string
}

type ComingUpdate struct {
	Id       string
	DateTime string
	Employee string
	Adresss  string
	Status   string
}

type GetListComingRequest struct {
	Offset int
	Limit  int
}

type GetListComingResponse struct {
	Count   int
	Comings []Coming
}
