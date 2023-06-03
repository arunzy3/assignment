package product

import (
	"Assignment/models"
	"Assignment/models/tables"
)

type ProductInterface interface {
	ReadProduct(pid string) (tables.Products, models.Errors)
	ReadProducts(category string, isCategorySpecified bool) ([]tables.Products, models.Errors)
	AddProduct(reqBody models.AddProductRequest) (tables.Products, models.Errors)
}
