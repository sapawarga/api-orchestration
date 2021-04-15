package phonebook

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	endpoint "github.com/sapawarga/api-orchestration/endpoint/phonebook"
	"github.com/sapawarga/api-orchestration/helper"
)

func decodeGetListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	stringRegID := r.URL.Query().Get("regency_id")
	stringDisID := r.URL.Query().Get("district_id")
	stringVilID := r.URL.Query().Get("village_id")
	stringStatus := r.URL.Query().Get("status")
	stringLimit := r.URL.Query().Get("limit")
	stringPage := r.URL.Query().Get("page")
	regencyID, err := helper.ParseFromStringToInt64(stringRegID)
	if err != nil {
		return nil, err
	}
	districtID, err := helper.ParseFromStringToInt64(stringDisID)
	if err != nil {
		return nil, err
	}
	villageID, err := helper.ParseFromStringToInt64(stringVilID)
	if err != nil {
		return nil, err
	}
	status, err := helper.ParseFromStringToInt64(stringStatus)
	if err != nil {
		return nil, err
	}
	limit, err := helper.ParseFromStringToInt64(stringLimit)
	if err != nil {
		return nil, err
	}
	page, err := helper.ParseFromStringToInt64(stringPage)
	if err != nil {
		return nil, err
	}

	request := &endpoint.GetListRequest{
		Search:     r.URL.Query().Get("search"),
		RegencyID:  helper.GetInt64FromPointer(regencyID),
		DistrictID: helper.GetInt64FromPointer(districtID),
		VillageID:  helper.GetInt64FromPointer(villageID),
		Status:     status,
		Limit:      helper.GetInt64FromPointer(limit),
		Page:       helper.GetInt64FromPointer(page),
		Latitude:   r.URL.Query().Get("latitude"),
		Longitude:  r.URL.Query().Get("longitude"),
	}
	return request, nil
}

func decodeGetDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	idInt, err := helper.ParseFromStringToInt64(params["id"])
	if err != nil {
		return nil, err
	}

	return &endpoint.GetDetailRequest{
		ID: helper.GetInt64FromPointer(idInt),
	}, nil
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	reqBody := &endpoint.CreatePhonebook{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return nil, err
	}

	return reqBody, nil
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id, err := helper.ParseFromStringToInt64(params["id"])
	if err != nil {
		return nil, err
	}

	reqBody := &endpoint.UpdatePhonebook{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return nil, err
	}

	reqBody.ID = helper.GetInt64FromPointer(id)
	return reqBody, nil
}
