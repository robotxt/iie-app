package api

import (
	"encoding/json"
	"github/robotxt/iie-app/src/logging"
	"github/robotxt/iie-app/src/service"
	"net/http"
)

type itemCreateApiResponse struct {
	Message string
	UID     string
}

type ItemData struct {
	Name        string
	Description string
	Tags        string
	Amount      float64
}

// CreateItemApi
func (api *ApiV1) CreateItemApi(w http.ResponseWriter, r *http.Request) {
	rd := &ItemData{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	authUser := r.Context().Value(UserCtxKey("authUser")).(service.UserType)

	item := service.ItemType{}
	item.UserUID = authUser.UID
	item.Name = rd.Name
	item.Tags = rd.Tags
	item.Description = rd.Description
	item.CreateItem(api.Ctx)

	respondJSON(w, http.StatusOK, &itemCreateApiResponse{
		Message: "Successfully created.",
		UID:     item.UID,
	})
}

type fetchItemsResponse struct {
	Message string
	Items   []service.ItemType
}

func (api *ApiV1) FetchItems(w http.ResponseWriter, r *http.Request) {
	authUser := r.Context().Value(UserCtxKey("authUser")).(service.UserType)
	item := service.ItemType{}
	item.UserUID = authUser.UID

	allItems, err := item.GetUserItems(api.Ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logging.Info("query result: ", allItems)

	respondJSON(w, http.StatusOK, &fetchItemsResponse{
		Message: "wasad",
		Items:   allItems,
	})
}
