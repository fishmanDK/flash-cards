package service

import (
	"github.com/fishmanDK/anki_telegram/internal/db"
)

type CategoryService struct {
	db db.CategoryMethonds
}

func NewCategoryService(db db.CategoryMethonds) *CategoryService {
	return &CategoryService{
		db: db,
	}
}

// func (a *ApiService) GetDataAboutUser(id string) (db.ResultInfoAboutUser, error) {
// 	userData, err := a.db.GetUserData(id)
// 	if err != nil {
// 		return userData, err
// 	}

// 	return userData, nil
// }

func (a *CategoryService) GetCategories(id string) (db.UserCategories, error) {
	return a.db.GetUserCategories(id)
}

func (a *CategoryService) GetCategoryById(userId, categoryName string) (db.Category, error) {
	return a.db.GetUserCategoryById(userId, categoryName)
}

func (a *CategoryService) CreateCategory(userId, categoryName string) (string, error) {
	return a.db.CreateCategory(userId, categoryName)
}

func (a *CategoryService) UpdateTitle(userId, oldName, newName string) error {
	return a.db.UpdateTitle(userId, oldName, newName)
}

func (a *CategoryService) DeleteCategory(userId, categoryName string) error{
	return a.db.DeleteCategory(userId, categoryName)
}