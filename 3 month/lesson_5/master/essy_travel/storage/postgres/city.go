package postgres

import (
	"database/sql"
	"essy_travel/models"
	"essy_travel/pkg/helpers"
	"fmt"

	"github.com/google/uuid"
)

type CityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (p *CityRepo) Create(req models.CreateCity) (*models.City, error) {

	query := `
		INSERT INTO city(
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"updated_at"
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())`

	id := uuid.New().String()
	_, err := p.db.Exec(query,
		id,
		req.Title,
		helpers.NewNullString(req.CountryId),
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		helpers.NewNullString(req.TimezoneId),
		req.CountryName,
	)
	if err != nil {
		return &models.City{}, err
	}

	return p.GetById(models.CityPrimaryKey{Id: id})
}

func (c *CityRepo) GetById(req models.CityPrimaryKey) (*models.City, error) {
	query := `
		SELECT
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"created_at",
			"updated_at"
		FROM city
		WHERE guid = $1
	`

	var (
		Guid        sql.NullString
		Title       sql.NullString
		CountryId   sql.NullString
		CityCode    sql.NullString
		Latitude    sql.NullString
		Longitude   sql.NullString
		Offset      sql.NullString
		TimezoneId  sql.NullString
		CountryName sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)

	err := c.db.QueryRow(query, req.Id).Scan(
		&Guid,
		&Title,
		&CountryId,
		&CityCode,
		&Latitude,
		&Longitude,
		&Offset,
		&TimezoneId,
		&CountryName,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.City{
		Guid:        Guid.String,
		Title:       Title.String,
		CountryId:   CountryId.String,
		CityCode:    CityCode.String,
		Latitude:    Latitude.String,
		Longitude:   Longitude.String,
		Offset:      Offset.String,
		TimezoneId:  TimezoneId.String,
		CountryName: CountryName.String,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}, nil
}

func (c *CityRepo) GetList(req models.GetListCityRequest) (*models.GetListCityResponse, error) {
	var (
		resp   = models.GetListCityResponse{}
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"created_at",
			"updated_at"
		FROM city
	`
	query += where + limit + offset

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			Guid        sql.NullString
			Title       sql.NullString
			CountryId   sql.NullString
			CityCode    sql.NullString
			Latitude    sql.NullString
			Longitude   sql.NullString
			Offset      sql.NullString
			TimezoneId  sql.NullString
			CountryName sql.NullString
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&Guid,
			&Title,
			&CountryId,
			&CityCode,
			&Latitude,
			&Longitude,
			&Offset,
			&TimezoneId,
			&CountryName,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Cities = append(resp.Cities, models.City{
			Guid:        Guid.String,
			Title:       Title.String,
			CountryId:   CountryId.String,
			CityCode:    CityCode.String,
			Latitude:    Latitude.String,
			Longitude:   Longitude.String,
			Offset:      Offset.String,
			TimezoneId:  TimezoneId.String,
			CountryName: CountryName.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}

	return &resp, nil
}

func (c *CityRepo) Update(req models.UpdateCity) (*models.City, error) {

	_, err := c.db.Exec(`UPDATE city SET id=$1,guid=$2,title=$3,country_id=$4,city_code=$5,latitude=$6,longitude=$7,offset=$8,timezone_id=$9,country_name=$10,updated_at=NOW() WHERE id = $11`, req.Id)
	if err != nil {
		return &models.City{}, err
	}

	return &models.City{}, nil
}

func (c *CityRepo) Delete(req models.CityPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM city WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CityRepo) Search(req models.City) (*models.GetListCityResponse, error) {
	var cities models.GetListCityResponse
	rows, err := c.db.Query(`SELECT COUNT(*) OVER(),id,guid,title,country_id,city_code,latitude,longitude,offset,timezone_id,country_name,created_at,updated_at FROM City WHERE title = $1`, req.Title)
	if err != nil {
		return &models.GetListCityResponse{}, err
	}
	for rows.Next() {
		var city models.City
		rows.Scan()
		cities.Cities = append(cities.Cities, city)
	}
	rows.Close()
	return &cities, nil
}
