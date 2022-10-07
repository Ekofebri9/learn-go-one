package service

import (
	"testing"

	"github.com/Ekofebri9/learn-go-one/v2/entity"
	"github.com/Ekofebri9/learn-go-one/v2/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	catagory, err := categoryService.Get("1")
	assert.Nil(t, catagory)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	data := entity.Category{
		Id:   "2",
		Name: "Tech",
	}
	categoryRepository.Mock.On("FindById", "2").Return(data)
	catagory, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, catagory)
	assert.Equal(t, data.Id, catagory.Id)
	assert.Equal(t, data.Name, catagory.Name)
}
