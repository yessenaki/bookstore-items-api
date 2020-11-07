package services

import (
	"github.com/yesseneon/bookstore-items-api/domain/item"
	"github.com/yesseneon/bookstore-utils/errors"
)

var ItemService itemServiceInterface = &itemService{}

type itemServiceInterface interface {
	Create(item.Item) (*item.Item, *errors.RESTError)
	Get(string) (*item.Item, *errors.RESTError)
}

type itemService struct{}

func (svc itemService) Create(i item.Item) (*item.Item, *errors.RESTError) {
	if err := i.Save(); err != nil {
		return nil, err
	}

	return *i, nil
}

func (svc itemService) Get(string) (*item.Item, *errors.RESTError) {
	return nil, errors.InternalServerError()
}
