package services

import (
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
	"net/http"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
