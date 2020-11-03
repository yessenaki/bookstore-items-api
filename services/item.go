package services

import (
	"github.com/yesseneon/bookstore_items_api/domain/item"
	"github.com/yesseneon/bookstore_utils/errors"
)

var ItemService itemServiceInterface = &itemService{}

type itemServiceInterface interface {
	Create(item.Item) (*item.Item, *errors.RESTError)
	Get(string) (*item.Item, *errors.RESTError)
}

type itemService struct {
}

func (srv itemService) Create(item.Item) (*item.Item, *errors.RESTError) {
	return nil, errors.InternalServerError()
}

func (srv itemService) Get(string) (*item.Item, *errors.RESTError) {
	return nil, errors.InternalServerError()
}
