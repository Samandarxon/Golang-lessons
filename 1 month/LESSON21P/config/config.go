package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	FilePath string
	File     File
}
type File struct {
	UserFile     string
	CategoryFile string
	OrderFile    string
	ProductFile  string
}

func ConnectionDataBases() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("not found env")
	}

	var cfg Config

	cfg.FilePath = cast.ToString(getValueOrDefault("FILE_PATH", "./DataBase/"))
	cfg.File.UserFile = cast.ToString(getValueOrDefault("USER_FILE", "user.json"))
	cfg.File.CategoryFile = cast.ToString(getValueOrDefault("CATEGORY_FILE", "category.json"))
	cfg.File.OrderFile = cast.ToString(getValueOrDefault("ORDER_FILE", "order.json"))
	cfg.File.ProductFile = cast.ToString(getValueOrDefault("PRODUCT_FILE", "product.json"))

	return cfg
}

func getValueOrDefault(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)
	// fmt.Println(key, val)
	if exists {
		return val
	}

	return defaultValue
}
