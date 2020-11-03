package controllers

import (
	"fmt"
	"net/http"

	"github.com/yesseneon/bookstore_items_api/domain/item"
	"github.com/yesseneon/bookstore_items_api/services"
	"github.com/yesseneon/bookstore_oauth_lib/oauth"
)

var ItemController itemControllerInterface = &itemController{}

type itemControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemController struct{}

func (ctr *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateUser(r); restErr != nil {
		return
	}

	i := item.Item{
		Seller: oauth.GetCallerID(r),
	}

	res, restErr := services.ItemService.Create(i)
	if restErr != nil {
		return
	}

	fmt.Println(res)
}

func (ctr *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
