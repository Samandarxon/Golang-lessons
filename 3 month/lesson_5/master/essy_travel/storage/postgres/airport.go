package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type AirportRepo struct {
	db *sql.DB
}

func NewAirportRepo(db *sql.DB) *AirportRepo {
	return &AirportRepo{
		db: db,
	}
}

func (p *AirportRepo) Create(req models.CreateAirport) (*models.Airport, error) {
	_, err := p.db.Exec(`INSERT INTO aiport(id,title,country_id,cityId,longitude,radius,image,adress,timezone_id,country,city,search_text,code,product_count,gmt,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,NOW())`,
		uuid.New().String(),
		req.Title,
		req.CountryId,
		req.CityId,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Adress,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return &models.Airport{}, err
	}
	return &models.Airport{}, nil
}

func (c *AirportRepo) GetById(req models.AirportPrimaryKey) (*models.Airport, error) {
	var airport = models.Airport{}
	resp := c.db.QueryRow(`SELECT id,title,country_id,cityId,longitude,radius,image,adress,timezone_id,country,city,search_text,code,product_count,gmt,created_at,updated_at FROM airport`, req.Id)
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&airport.Id,
		&airport.Title,
		&airport.CountryId,
		&airport.CityId,
		&airport.Longitude,
		&airport.Radius,
		&airport.Image,
		&airport.Adress,
		&airport.TimezoneId,
		&airport.Country,
		&airport.City,
		&airport.SearchText,
		&airport.Code,
		&airport.ProductCount,
		&airport.Gmt,
		&airport.CreatedAt,
		&airport.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &airport, nil
}

func (c *AirportRepo) GetList(req models.GetListAirportRequest) (*models.GetListAirportResponse, error) {
	var airports = models.GetListAirportResponse{}
	var offset string = " offset"
	var limit string = " limit"
	if req.Offset <= 0 {
		offset += " 0"
	} else if req.Offset > 0 {
		offset += fmt.Sprintf(" %d", req.Offset)
	}

	if req.Limit <= 0 {
		limit += " 10"
	} else if req.Limit > 0 {
		limit += fmt.Sprintf(" %d", req.Limit)
	}
	rows, err := c.db.Query(`SELECT COUNT(*) OVER(), id,title,country_id,cityId,longitude,radius,image,adress,timezone_id,country,city,search_text,code,product_count,gmt,created_at,updated_at FROM airport`, offset, limit)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var airport models.Airport
		rows.Scan(
			&airports.Count,
			&airport.Id,
			&airport.Title,
			&airport.CountryId,
			&airport.CityId,
			&airport.Longitude,
			&airport.Radius,
			&airport.Image,
			&airport.Adress,
			&airport.TimezoneId,
			&airport.Country,
			&airport.City,
			&airport.SearchText,
			&airport.Code,
			&airport.ProductCount,
			&airport.Gmt,
			&airport.CreatedAt,
			&airport.UpdatedAt,
		)
		airports.Airports = append(airports.Airports, airport)
	}
	rows.Close()
	return &airports, nil
}

func (c *AirportRepo) Update(req models.UpdateAirport) (*models.Airport, error) {

	_, err := c.db.Exec(`UPDATE airports SET id=$1,title=$2,country_id=$3,cityId=$4,longitude=$5,radius=$6,image=$7,adress=$8,timezone_id=$9,country=$10,city=$11,search_text=$12,code=$13,product_count=$14,gmt=$15,updated_at = NOW() WHERE id = $16`, req.Id)
	if err != nil {
		return &models.Airport{}, err
	}

	return &models.Airport{}, nil
}

func (c *AirportRepo) Delete(req models.AirportPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM airports WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *AirportRepo) Search(req models.Airport) (*models.GetListAirportResponse, error) {
	var airports models.GetListAirportResponse
	rows, err := c.db.Query(`SELECT COUNT(*) OVER(),id,title,country_id,cityId,longitude,radius,image,adress,timezone_id,country,city,search_text,code,product_count,gmt,created_at,updated_at FROM airport WHERE title = $1`, req.Title)
	if err != nil {
		return &models.GetListAirportResponse{}, err
	}
	for rows.Next() {
		var airport models.Airport
		rows.Scan(
			&airports.Count,
			&airport.Id,
			&airport.Title,
			&airport.CountryId,
			&airport.CityId,
			&airport.Longitude,
			&airport.Radius,
			&airport.Image,
			&airport.Adress,
			&airport.TimezoneId,
			&airport.Country,
			&airport.City,
			&airport.SearchText,
			&airport.Code,
			&airport.ProductCount,
			&airport.Gmt,
			&airport.CreatedAt,
			&airport.UpdatedAt,
		)
		airports.Airports = append(airports.Airports, airport)
	}
	rows.Close()
	return &airports, nil
}
