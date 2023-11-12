package psqlDB

import (
	"app/config"
	"app/models"
	"app/storage"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type categoryRepo struct {
	psql_url string
}

func NewCategoryRepo(cfg *config.Config_DB) storage.CategoryRepoI {
	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)
	return &categoryRepo{psql_url: connect}
}

func (u *categoryRepo) Create(req *models.CreateCategory) (*models.Category, error) {

	data, err := sql.Open("postgres", u.psql_url)
	if err != nil {
		return nil, err
	}

	var resp = &models.Category{
		Id:   uuid.New().String(),
		Name: req.Name,
	}
	query := `INSERT INTO "category"("id","name") VALUES ($1,$2)`
	result, err := data.Exec(query, resp.Id, resp.Name)

	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	return resp, nil
}

// func (c *categoryRepo) GetByID(req *models.PrimeryKeyCategory) (*models.Category, error) {

// 	data, err := helpers.Read(c.categoryFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, categoryObj := range data {
// 		var category = cast.ToStringMap(categoryObj)
// 		if cast.ToString(category["id"]) == req.Id {

// 			return &models.Category{
// 				Id:   cast.ToString(category["id"]),
// 				Name: cast.ToString(category["name"]),
// 			}, nil
// 		}
// 	}

// 	return nil, errors.New("data is no rows")
// }

// func (c *categoryRepo) GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {
// 	data, err := helpers.Read(c.categoryFile)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// fmt.Println(data)
// 	if req.Limit == 0 {
// 		req.Limit = 10
// 	}

// 	var resp = &models.GetListCategoryResponse{}

// 	for i := req.Offset; i < req.Limit+req.Offset; i++ {
// 		if len(data) > i {
// 			var category = cast.ToStringMap(data[i])
// 			fmt.Println(i+1, category)
// 			resp.Categorys = append(resp.Categorys, &models.Category{
// 				Id:   cast.ToString(category["id"]),
// 				Name: cast.ToString(category["name"]),
// 			})
// 		} else {
// 			break
// 		}
// 	}
// 	resp.Count = len(data)

// 	return resp, nil
// }

// func (c *categoryRepo) Update(req *models.UpadetCategory) (*models.Category, error) {

// 	data, err := helpers.Read(c.categoryFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, categoryObj := range data {
// 		var category = cast.ToStringMap(categoryObj)

// 		if cast.ToString(category["id"]) == req.Id {
// 			category["id"] = req.Id
// 			category["name"] = req.Name
// 			break
// 		}
// 	}

// 	err = helpers.Write(c.categoryFile, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return c.GetByID(&models.PrimeryKeyCategory{Id: req.Id})
// }

// func (c *categoryRepo) Delete(req *models.PrimeryKeyCategory) error {

// 	data, err := helpers.Read(c.categoryFile)
// 	if err != nil {
// 		return err
// 	}

// 	for ind, categoryObj := range data {
// 		var category = cast.ToStringMap(categoryObj)
// 		if cast.ToString(category["id"]) == req.Id {
// 			data = append(data[:ind], data[ind+1:]...)
// 			break
// 		}
// 	}

// 	err = helpers.Write(c.categoryFile, data)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
