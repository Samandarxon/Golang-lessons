package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type CountryRepo struct {
	db *sql.DB
}

func NewCountryRepo(db *sql.DB) *CountryRepo {
	return &CountryRepo{
		db: db,
	}
}

func (p *CountryRepo) Create(req models.CreateCountry) (*models.Country, error) {
	_, err := p.db.Exec(`INSERT INTO country(id,guid,title,code,continent,updated_at) VALUES ($1,$2,$3,4$,%5,NOW())`, uuid.New().String(), req.Guid, req.Title, req.Code, req.Continent)
	if err != nil {
		return &models.Country{}, err
	}
	return &models.Country{}, nil
}

func (c *CountryRepo) GetById(req models.CountryPrimaryKey) (*models.Country, error) {
	var country = models.Country{}
	resp := c.db.QueryRow(`SELECT id,guid,title,code,continent,created_at,updated_at FROM country`, req.Id)
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&country.Id,
		&country.Guid,
		&country.Title,
		&country.Code,
		&country.Code,
		&country.CreatedAt,
		&country.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (c *CountryRepo) GetList(req models.GetListCountryRequest) (*models.GetListCountryResponse, error) {
	var countries = models.GetListCountryResponse{}
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
	rows, err := c.db.Query(`SELECT COUNT(*) OVER(), id,guid,title,code,continent,created_at,updated_at FROM country %s %s`, offset, limit)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var country models.Country
		rows.Scan(
			&country.Id,
			&country.Guid,
			&country.Title,
			&country.Code,
			&country.Code,
			&country.CreatedAt,
			&country.UpdatedAt,
		)
		countries.Countries = append(countries.Countries, country)
	}
	rows.Close()
	return &countries, nil
}

func (c *CountryRepo) Update(req models.UpdateCountry) (*models.Country, error) {

	_, err := c.db.Exec(`UPDATE country SET guid=$1,title=$2,code=$3,continent$4,updated_at=NOW() WHERE id = $5`, req.Id, req.Guid, req.Title, req.Code, req.Continent)
	if err != nil {
		return &models.Country{}, err
	}

	return &models.Country{}, nil
}

func (c *CountryRepo) Delete(req models.CountryPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM country WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CountryRepo) Search(req models.Country) (*models.GetListCountryResponse, error) {
	var countries models.GetListCountryResponse
	rows, err := c.db.Query(`SELECT COUNT(*) OVER(),id,guid,title,code, continent,created_at,updated_at FROM country WHERE title = $1`, req.Title)
	if err != nil {
		return &models.GetListCountryResponse{}, err
	}
	for rows.Next() {
		var country models.Country
		rows.Scan(
			&country.Id,
			&country.Guid,
			&country.Title,
			&country.Code,
			&country.Code,
			&country.CreatedAt,
			&country.UpdatedAt,
		)
		countries.Countries = append(countries.Countries, country)
	}
	rows.Close()
	return &countries, nil
}
