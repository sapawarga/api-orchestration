package phonebook

import (
	"context"
	"math"

	"github.com/sapawarga/api-orchestration/helper"
	model "github.com/sapawarga/api-orchestration/model/phonebook"

	proto "github.com/sapawarga/proto-file/phonebook"
	"google.golang.org/grpc"
)

type Provider struct {
	client proto.PhoneBookHandlerClient
}

func NewProvider(c *grpc.ClientConn) *Provider {
	return &Provider{
		client: proto.NewPhoneBookHandlerClient(c),
	}
}

func (c *Provider) GetList(ctx context.Context, req *model.GetListRequest) (*model.GetListResponse, error) {
	filter := &proto.GetListRequest{
		Search:     req.Search,
		RegencyId:  req.RegencyID,
		DistrictId: req.DistrictID,
		VillageId:  req.VillageID,
		Status: &proto.NullInt64{
			Value:  helper.GetInt64FromPointer(req.Status),
			IsNull: false,
		},
		Limit:     req.Limit,
		Page:      req.Page,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	resp, err := c.client.GetList(ctx, filter)
	if err != nil {
		return nil, err
	}

	data := make([]*model.Phone, 0)
	for _, v := range resp.Data {
		res := &model.Phone{
			ID:           v.GetId(),
			PhoneNumbers: v.GetPhoneNumbers(),
			Description:  v.GetDescription(),
			Name:         v.GetName(),
			Address:      v.GetAddress(),
			Latitude:     v.GetLatitude(),
			Longitude:    v.GetLongitude(),
			Status:       v.GetStatus(),
			Category:     v.GetCategory(),
			Distance:     v.GetDistance(),
		}
		data = append(data, res)
	}

	totalPage := int64(math.Floor(float64(resp.Metadata.Total) / float64(resp.Metadata.Page)))

	meta := &model.Metadata{
		Page:      resp.Metadata.GetPage(),
		Total:     resp.Metadata.GetTotal(),
		TotalPage: totalPage,
	}

	return &model.GetListResponse{
		Data:     data,
		Metadata: meta,
	}, nil
}

func (c *Provider) GetDetail(ctx context.Context, id int64) (*model.PhoneDetail, error) {
	resp, err := c.client.GetDetail(ctx, &proto.GetDetailRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &model.PhoneDetail{
		ID:           resp.GetId(),
		PhoneNumbers: resp.GetPhoneNumbers(),
		Description:  resp.GetDescription(),
		Name:         resp.GetName(),
		Address:      resp.GetAddress(),
		Latitude:     resp.GetLatitude(),
		Longitude:    resp.GetLongitude(),
		Status:       resp.GetStatus(),
		CategoryID:   resp.GetCategoryId(),
		CategoryName: resp.GetCategoryName(),
		Distance:     resp.GetDistance(),
		RegencyID:    resp.GetRegencyId(),
		RegencyName:  resp.GetRegencyName(),
		DistrictID:   resp.GetDistrictId(),
		DistrictName: resp.GetDistrictName(),
		VillageID:    resp.GetVillageId(),
		VillageName:  resp.GetVillageName(),
		CreatedAt:    helper.FormatTimeFromString(resp.GetCreatedAt()),
		UpdatedAt:    helper.FormatTimeFromString(resp.GetUpdatedAt()),
	}, nil
}

func (c *Provider) Insert(ctx context.Context, req *model.AddPhonebookRequest) error {
	request := &proto.AddPhonebookRequest{
		Name:           req.Name,
		PhoneNumbers:   req.PhoneNumbers,
		Address:        req.Address,
		Description:    req.Description,
		RegencyId:      req.RegencyID,
		DistrictId:     req.DistrictID,
		VillageId:      req.VillageID,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		CoverImagePath: req.CoverImagePath,
		Status:         req.Status,
		CategoryId:     req.CategoryID,
	}
	if _, err := c.client.AddPhonebook(ctx, request); err != nil {
		return err
	}

	return nil
}

func (c *Provider) Update(ctx context.Context, req *model.UpdatePhonebookRequest) error {
	nullStatus := &proto.NullInt64{
		Value:  0,
		IsNull: true,
	}

	if req.Status != nil {
		nullStatus.Value = helper.GetInt64FromPointer(req.Status)
		nullStatus.IsNull = false
	}

	request := &proto.UpdatePhonebookRequest{
		Id:             req.ID,
		Name:           req.Name,
		PhoneNumbers:   req.PhoneNumbers,
		Address:        req.Address,
		Description:    req.Description,
		RegencyId:      req.RegencyID,
		DistrictId:     req.DistrictID,
		VillageId:      req.VillageID,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		CoverImagePath: req.CoverImagePath,
		CategoryId:     req.CategoryID,
		Status:         nullStatus,
	}

	if _, err := c.client.UpdatePhonebook(ctx, request); err != nil {
		return err
	}
	return nil
}

func (c *Provider) Delete(ctx context.Context, id int64) error {
	if _, err := c.client.DeletePhonebook(ctx, &proto.GetDetailRequest{
		Id: id,
	}); err != nil {
		return err
	}

	return nil
}
