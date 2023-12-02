package main

import (
	"fmt"
	"log"
	"simple-clinic/config"
	"simple-clinic/controller"
	"simple-clinic/models"
	"simple-clinic/storage/jsondb"

	"github.com/bxcodec/faker/v3"
)

func main() {

	var cfg = config.Load()

	jsondb, err := jsondb.NewJsonDbConnection(&cfg)
	if err != nil {
		panic(err)
	}

	con := controller.NewController(&cfg, jsondb)

	_, err = con.CreateClinic(
		&models.CreateClinic{
			Name: "Akfa clinic",
			Logo: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.vectorstock.com%2Froyalty-free-vector%2Fclinic-care-logo-icon-design-vector-22560856&psig=AOvVaw0rmUsoAItEKyY5eJKams9h&ust=1696684557507000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCIiQ1NTA4YEDFQAAAAAdAAAAABAE",
			City: "Tashkent",
		},
	)
	if err != nil {
		log.Println("Create clinic:", err.Error())
		return
	}

	// fmt.Println(resp)
	_, err = con.UpdateClinic(
		&models.UpdateClinic{
			ClinicID: "d69795c7-51b8-4212-97c4-76648a710047",
			Logo:     "https://img.freepik.com/free-vector/bird-colorful-logo-gradient-vector_343694-1365.jpg",
			Name:     "Nasaf clinic",
			City:     "Qarshi",
		})

	if err != nil {
		log.Println("Update clinic:", err.Error())
		return
	}
	clinic, err := con.GetByIDClinic(
		&models.PrimeryKeyClinicID{
			ClinicID: "34d55ad5-705b-445d-b2ea-1f5963876eb4",
		})

	if err != nil {
		log.Println("Update clinic:", err.Error())
		return
	}
	fmt.Println(clinic)
	del, err := con.DeleteClinic(
		&models.PrimeryKeyClinicID{
			ClinicID: "cd2bcdfd-8b20-4292-b055-0e58283a7ef1",
		})

	if err != nil {
		log.Println("Update clinic:", err.Error())
		return
	}
	fmt.Println(del)

	/******************************************BranchOffice******************************************/
	con = controller.NewController(&cfg, jsondb)

	_, err = con.CreateBranchOffice(
		&models.CreateBranchOffice{
			ClinicID: "d69795c7-51b8-4212-97c4-76648a710047",
			Addres:   "Xorazim",
			Phone:    faker.Phonenumber(),
		},
	)
	if err != nil {
		log.Println("Create clinic:", err.Error())
		return
	}

	// fmt.Println(resp)
	_, err = con.UpdateBranchOffice(
		&models.UpdateBranchOffice{
			BranchOfficeID: "8bba3318-83da-4197-bd2d-b8374cbb13c3",
			ClinicID:       "34d55ad5-705b-445d-b2ea-1f5963876eb4",
			Addres:         "Surxandaryo",
			Phone:          faker.Phonenumber(),
		})

	if err != nil {
		log.Println("Update BranchOffice ", err.Error())
		return
	}
	branch, err := con.GetByIDBranchOffice(&models.PrimeryKeyBranchOfficeID{BranchOfficeID: "5b9af49e-9e61-499e-9e93-69fa6d6fcbf1"})

	if err != nil {
		log.Println("GetByID BranchOffice:", err.Error())
		return
	}
	fmt.Println(branch)
	del, err = con.DeleteBranchOffice(
		&models.PrimeryKeyBranchOfficeID{
			BranchOfficeID: "e612e88b-3ca6-4797-b883-fe19f8471ee5",
		})

	if err != nil {
		log.Println("Update clinic:", err.Error())
		return
	}
	fmt.Println(del)

}
