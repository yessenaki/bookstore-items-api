package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yesseneon/bookstore-items-api/domain/item"
	"github.com/yesseneon/bookstore-items-api/services"
	"github.com/yesseneon/bookstore-items-api/utils/handlers"
	"github.com/yesseneon/bookstore-oauth-lib/oauth"
	"github.com/yesseneon/bookstore-utils/errors"
)

var ItemController itemControllerInterface = &itemController{}

type itemControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemController struct{}

func (ctr *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateUser(r); restErr != nil {
		// handlers.RespondError(w, restErr)
		return
	}

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		restErr := errors.BadRequest("Invalid request body") // unauthorized error 401
		handlers.RespondError(w, *restErr)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := errors.BadRequest("Invalid JSON body")
		handlers.RespondError(w, restErr)
		return
	}
	defer r.Body.Close()

	var i item.Item
	if err := json.Unmarshal(bytes, &i); err != nil {
		restErr := errors.BadRequest("Invalid JSON body")
		handlers.RespondError(w, restErr)
		return
	}

	i.Seller = sellerID

	res, restErr := services.ItemService.Create(i)
	if restErr != nil {
		handlers.RespondError(w, restErr)
		return
	}

	handlers.RespondJson(w, http.StatusCreated, res)
}

func (ctr *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
