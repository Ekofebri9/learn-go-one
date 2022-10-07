package repository

import "github.com/Ekofebri9/learn-go-one/v2/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
