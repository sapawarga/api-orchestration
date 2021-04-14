package phonebook

import (
	"context"

	model "github.com/sapawarga/api-orchestration/model/phonebook"
	"github.com/sapawarga/api-orchestration/usecase"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetList(ctx context.Context, fs usecase.PhonebookI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetListRequest)

		response, err = fs.GetList(ctx, &model.GetListRequest{
			Search:     req.Search,
			RegencyID:  req.RegencyID,
			DistrictID: req.DistrictID,
			VillageID:  req.VillageID,
			Status:     req.Status,
			Limit:      req.Limit,
			Page:       req.Page,
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
		})
		if err != nil {
			return nil, err
		}

		return response, nil
	}
}

func MakeGetDetail(ctx context.Context, fs usecase.PhonebookI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetDetailRequest)

		response, err = fs.GetDetail(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return response, nil
	}
}
