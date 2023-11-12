package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

type CategoryController struct {
	Categorys []models.Category
}

/* ***********************************************New Category Controller Function*********************************************** */

func NewCategoryController(Categorys []models.Category) CategoryController {
	return CategoryController{
		Categorys: Categorys,
	}
}

/* ***********************************************New Category Controller Function*********************************************** */
/* ***********************************************Generate Category Method*********************************************** */
func (c *CategoryController) GenerateCategory(limit int) *CategoryController {
	// newCategoryController := NewCategoryController([]models.Category{})
	for i := 0; i < limit; i++ {

		c.CreateCategory(models.CreateCategory{
			Name: faker.FirstName(),
		})
	}
	return c
}
func (c *CategoryController) WriteFile(path string) {
	respons, err := json.MarshalIndent(c.Categorys, " ", "   ")
	// resp,err := io.ReadAll()
	if err != nil {
		log.Println(err)
	}
	// fmt.Printf(" %s\n", string(respons))

	err = os.WriteFile(path, respons, 0644)
	if err != nil {
		log.Println(err)
	}

}
func (c *CategoryController) ReadAll(path string) {
	resp, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(resp, &c.Categorys)
	// resp,err := io.ReadAll()
	if err != nil {
		log.Println(err)
	}
}

/* ***********************************************Generate Category Method*********************************************** */
/* ***********************************************Create Category Method*********************************************** */
func (c *CategoryController) CreateCategory(req models.CreateCategory) models.Category {
	newCategory := models.Category{
		Id:        uuid.New().String(),
		Name:      req.Name,
		CreatedAt: time.Now().Format("2006-01-01"),
		UpdatedAt: time.Now().Format("2006-01-01"),
	}
	c.ReadAll("JsonDB/category.json")
	c.Categorys = append(c.Categorys, newCategory)
	c.WriteFile("JsonDB/category.json")
	return newCategory
}

/* ***********************************************Create Category Method*********************************************** */
/* ***********************************************Category Get List Method*********************************************** */
func (c *CategoryController) CategoryGetList(req models.CategoryGetListRequest) models.CategoryGetListRespons {
	fmt.Println(" ********************************************CategoryGetList Started******************************************** ")
	newLimitCategorys := []models.Category{}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.OffSet+req.Limit > len(c.Categorys) {
		fmt.Println("Counnt ", len(c.Categorys))
		for i, el := range c.Categorys {
			fmt.Println(i+1, "  ", el)
		}

		fmt.Println(" ********************************************CategoryGetList Finished******************************************** ")
		return models.CategoryGetListRespons{
			Count:     len(c.Categorys),
			Categorys: c.Categorys,
		}
	}
	fmt.Println("Counnt ", len(c.Categorys))
	for i := req.OffSet; i < req.Limit+req.OffSet; i++ {
		fmt.Println(i+1, c.Categorys[i])
	}
	newLimitCategorys = append(newLimitCategorys, c.Categorys...)
	fmt.Println(" ********************************************CategoryGetList Finished******************************************** ")
	return models.CategoryGetListRespons{
		Count:     len(c.Categorys),
		Categorys: newLimitCategorys,
	}
}

/* ***********************************************Category Get List Method*********************************************** */
/* ***********************************************Category Get By Id Method*********************************************** */
func (c *CategoryController) CategoryGetById(req models.CategoryPrimaryKey) models.Category {
	for _, el := range c.Categorys {
		if el.Id == req.Id {
			return el
		}
	}
	return models.Category{}
}

/* ***********************************************Category Get By Id Method*********************************************** */
/* ***********************************************Category Update Method*********************************************** */
func (c *CategoryController) CategoryUpdate(req models.CategoryUpdateRequest) models.Category {
	// path := "JsonDB/category.json"

	resp, err := os.ReadFile("JsonDB/category.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(resp, &c.Categorys)
	// resp,err := io.ReadAll()
	if err != nil {
		log.Println(err)
	}

	// c.Categorys = append(c.Categorys,)

	for i, el := range c.Categorys {
		if el.Id == req.Id {
			c.Categorys[i].Name = req.Name
			c.Categorys[i].UpdatedAt = time.Now().Format("2006-01-01")

			c.WriteFile("JsonDB/category.json")
			return c.Categorys[i]
		}
	}

	return models.Category{}
}

/* ***********************************************Category Update Method*********************************************** */
/* ***********************************************Category Delete Method*********************************************** */
func (c *CategoryController) CategoryDelete(req models.CategoryPrimaryKey) (string, models.Category) {
	for index, el := range c.Categorys {
		if el.Id == req.Id {
			del_Category := c.Categorys[index]
			c.Categorys = append(c.Categorys[:index], c.Categorys[index+1:]...)
			c.WriteFile("JsonDB/category.json")
			return "DELETED", del_Category
		}
	}
	return fmt.Sprintf("%s shu ID ega Category Topilmadi", req.Id), models.Category{}
}

/* ***********************************************Category Delete Method*********************************************** */
