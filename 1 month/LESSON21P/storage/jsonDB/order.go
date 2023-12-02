package jsonDB

import (
	"app/config"
	"app/models"
	"app/pkg/helpers"
	"app/storage"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type orderRepo struct {
	orderFile string
}

func NewOrderRepo(cfg *config.Config) storage.OrederRepoI {

	return &orderRepo{orderFile: cfg.FilePath + cfg.File.OrderFile}
}

func (u *orderRepo) Create(req *models.CreateOrder) (*models.Order, error) {

	data, err := helpers.Read(u.orderFile)
	if err != nil {
		return nil, err
	}

	var resp = &models.Order{
		Id:        uuid.New().String(),
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Sum:       req.Sum,
		SumCount:  req.SumCount,
		DateTime:  req.DateTime,
		Status:    req.Status,
	}

	data = append(data, resp)

	err = helpers.Write(u.orderFile, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *orderRepo) GetByID(req *models.PrimeryKeyOrder) (*models.Order, error) {

	data, err := helpers.Read(c.orderFile)
	if err != nil {
		return nil, err
	}

	for _, orderObj := range data {
		var order = cast.ToStringMap(orderObj)
		if cast.ToString(order["id"]) == req.Id {

			return &models.Order{
				Id:        cast.ToString(order["id"]),
				UserId:    cast.ToString(order["user_id"]),
				ProductId: cast.ToString(order["product_id"]),
				Sum:       cast.ToInt(order["sum"]),
				SumCount:  cast.ToInt(order["sum_count"]),
				DateTime:  cast.ToString(order["date_time"]),
				Status:    cast.ToString(order["status"]),
			}, nil
		}
	}

	return nil, errors.New("data is no rows")
}

func (c *orderRepo) GetList(req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {
	data, err := helpers.Read(c.orderFile)
	if err != nil {
		return nil, err
	}
	// fmt.Println(data)
	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = &models.GetListOrderResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var order = cast.ToStringMap(data[i])
			fmt.Println(i+1, order)
			resp.Orders = append(resp.Orders, &models.Order{
				Id:        cast.ToString(order["id"]),
				UserId:    cast.ToString(order["user_id"]),
				ProductId: cast.ToString(order["product_id"]),
				Sum:       cast.ToInt(order["sum"]),
				SumCount:  cast.ToInt(order["sum_count"]),
				DateTime:  cast.ToString(order["date_time"]),
				Status:    cast.ToString(order["status"]),
			})
		} else {
			break
		}
	}
	resp.Count = len(data)

	return resp, nil
}

func (c *orderRepo) Update(req *models.UpadetOrder) (*models.Order, error) {

	data, err := helpers.Read(c.orderFile)
	if err != nil {
		return nil, err
	}

	for _, orderObj := range data {
		var order = cast.ToStringMap(orderObj)

		if cast.ToString(order["id"]) == req.Id {
			order["id"] = req.Id
			order["user_id"] = req.UserId
			order["product_id"] = req.ProductId
			order["sum"] = req.Sum
			order["sum_count"] = req.SumCount
			order["date_time"] = req.DateTime
			order["status"] = req.Status

			// data[ind] = order
			break
		}
	}

	err = helpers.Write(c.orderFile, data)
	if err != nil {
		return nil, err
	}

	return c.GetByID(&models.PrimeryKeyOrder{Id: req.Id})
}

func (c *orderRepo) Delete(req *models.PrimeryKeyOrder) error {

	data, err := helpers.Read(c.orderFile)
	if err != nil {
		return err
	}

	for ind, orderObj := range data {
		var order = cast.ToStringMap(orderObj)
		if cast.ToString(order["id"]) == req.Id {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = helpers.Write(c.orderFile, data)
	if err != nil {
		return err
	}

	return nil
}
