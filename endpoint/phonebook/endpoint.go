package phonebook

import (
	"context"
	"encoding/json"

	"github.com/sapawarga/api-orchestration/helper"
	model "github.com/sapawarga/api-orchestration/model/phonebook"
	"github.com/sapawarga/api-orchestration/usecase"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetList(ctx context.Context, fs usecase.PhonebookI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetListRequest)

		resp, err := fs.GetList(ctx, &model.GetListRequest{
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

		data := make([]*Phone, 0)
		for _, v := range resp.Data {
			phoneJson := make([]map[string]interface{}, 0)
			err := json.Unmarshal([]byte(v.PhoneNumbers), &phoneJson)
			if err != nil {
				return nil, err
			}
			result := &Phone{
				ID:           v.ID,
				PhoneNumbers: phoneJson,
				Description:  v.Description,
				Name:         v.Name,
				Address:      v.Address,
				Latitude:     v.Latitude,
				Longitude:    v.Longitude,
				Status:       v.Status,
				Category:     v.Category,
				Distance:     v.Distance,
			}
			data = append(data, result)
		}

		meta := &Metadata{
			Page:      resp.Metadata.Page,
			Total:     resp.Metadata.Total,
			TotalPage: resp.Metadata.TotalPage,
		}

		return &GetListResponse{
			Data:     data,
			Metadata: meta,
		}, nil
	}
}

func MakeGetDetail(ctx context.Context, fs usecase.PhonebookI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetDetailRequest)

		resp, err := fs.GetDetail(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		phoneJson := make([]map[string]interface{}, 0)
		err = json.Unmarshal([]byte(resp.PhoneNumbers), &phoneJson)
		if err != nil {
			return nil, err
		}

		return &PhoneDetail{
			ID:           resp.ID,
			PhoneNumbers: phoneJson,
			Description:  resp.Description,
			Name:         resp.Name,
			Address:      resp.Address,
			Latitude:     resp.Latitude,
			Longitude:    resp.Longitude,
			Status:       resp.Status,
			CreatedAt:    resp.CreatedAt,
			UpdatedAt:    resp.UpdatedAt,
			CategoryID:   resp.CategoryID,
			CategoryName: resp.CategoryName,
			Distance:     resp.Distance,
			RegencyID:    resp.RegencyID,
			RegencyName:  resp.RegencyName,
			DistrictID:   resp.DistrictID,
			DistrictName: resp.DistrictName,
			VillageID:    resp.VillageID,
			VillageName:  resp.VillageName,
		}, nil
	}
}

func MakeCreatePhonebook(ctx context.Context, fs usecase.PhonebookI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreatePhonebook)
		if err = Validation(req); err != nil {
			return nil, err
		}

		phoneString, _ := json.Marshal(req.PhoneNumbers)

		err = fs.Insert(ctx, &model.AddPhonebookRequest{
			Name:           req.Name,
			PhoneNumbers:   string(phoneString),
			Address:        req.Address,
			Description:    req.Description,
			RegencyID:      helper.GetInt64FromPointer(req.RegencyID),
			DistrictID:     helper.GetInt64FromPointer(req.DistrictID),
			VillageID:      helper.GetInt64FromPointer(req.VillageID),
			Latitude:       req.Latitude,
			Longitude:      req.Longitude,
			CoverImagePath: req.CoverImagePath,
			Status:         helper.GetInt64FromPointer(req.Status),
			CategoryID:     req.CategoryID,
		})
		if err != nil {
			return nil, err
		}
		return nil, err
	}
}
