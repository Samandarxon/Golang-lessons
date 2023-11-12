package controller

import (
	"app/models"
	"fmt"
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
func (c *CategoryController) GenerateCategory(limit int) CategoryController {
	newCategoryController := NewCategoryController([]models.Category{})
	for i := 0; i < limit; i++ {

		newCategoryController.CreateCategory(models.CreateCategoryRequest{
			Name: faker.FirstName(),
		})
	}
	return newCategoryController
}

/* ***********************************************Generate Category Method*********************************************** */
/* ***********************************************Create Category Method*********************************************** */
func (c *CategoryController) CreateCategory(req models.CreateCategoryRequest) models.Category {
	newCategory := models.Category{
		Id:        uuid.New(),
		Name:      req.Name,
		CreatedAt: time.Now().Format("2006-01-01"),
		UpdatedAt: time.Now().Format("2006-01-01"),
	}
	c.Categorys = append(c.Categorys, newCategory)
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

	for i, el := range c.Categorys {
		if el.Id == req.Id {
			c.Categorys[i].Name = req.Name
			c.Categorys[i].UpdatedAt = time.Now().Format("2006-01-01")
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
			return "DELETED", del_Category
		}
	}
	return fmt.Sprintf("%s shu ID ega Category Topilmadi", req.Id.String()), models.Category{}
}

/* ***********************************************Category Delete Method*********************************************** */
