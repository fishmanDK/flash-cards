package service

import (
	"github.com/fishmanDK/anki_telegram/internal/db"
)

type ApiService struct {
	db db.Api
}

func NewApiService(db db.Api) *ApiService {
	return &ApiService{
		db: db,
	}
}

func (a *ApiService) GetDataAboutUser(id string) (db.ResultInfoAboutUser, error) {
	userData, err := a.db.GetUserData(id)
	if err != nil {
		return userData, err
	}

	return userData, nil
}

func (a *ApiService) GetCategories(id string) (db.UserCategories, error) {
	userCategories, err := a.db.GetUserCategories(id)
	if err != nil {
		return userCategories, err
	}

	return userCategories, nil
}