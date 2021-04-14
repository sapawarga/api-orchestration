package testcases

import (
	"errors"

	"github.com/sapawarga/api-orchestration/helper"
	model "github.com/sapawarga/api-orchestration/model/phonebook"
)

type GetListResponseRepository struct {
	Result *model.GetListResponse
	Error  error
}

type GetListResponseUsecase struct {
	Result *model.GetListResponse
	Error  error
}

type GetListPhonebook struct {
	Description             string
	UsecaseParams           *model.GetListRequest
	GetListRepoParams       *model.GetListRequest
	MockUsecase             GetListResponseUsecase
	MockPhonebookRepository GetListResponseRepository
}

var (
	request = &model.GetListRequest{
		Search:     "test",
		RegencyID:  1,
		DistrictID: 1,
		VillageID:  1,
		Status:     helper.SetPointerInt64(10),
		Limit:      10,
		Page:       1,
		Latitude:   "42.1234134",
		Longitude:  "-0.2342415",
	}
	data = []*model.Phone{
		{
			ID:           1,
			Name:         "kantor",
			PhoneNumbers: `[{"phone_number": "022123"}]`,
			Description:  "kantor cabang MCD",
		},
		{
			ID:           2,
			Name:         "kantor",
			PhoneNumbers: `[{"phone_number": "423443"}]`,
			Description:  "kantor makanan",
		},
	}
	metadata = &model.Metadata{
		Page:      1,
		Total:     10,
		TotalPage: 1,
	}
	response = &model.GetListResponse{
		Data:     data,
		Metadata: metadata,
	}
)

var GetListPhonebookData = []GetListPhonebook{
	{
		Description:       "success_get_detail",
		UsecaseParams:     request,
		GetListRepoParams: request,

		MockUsecase: GetListResponseUsecase{
			Result: response,
			Error:  nil,
		},
		MockPhonebookRepository: GetListResponseRepository{
			Result: response,
			Error:  nil,
		},
	}, {
		Description:       "failed_get_list",
		UsecaseParams:     request,
		GetListRepoParams: request,

		MockUsecase: GetListResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_response"),
		},
		MockPhonebookRepository: GetListResponseRepository{
			Result: nil,
			Error:  errors.New("failed_get_response"),
		},
	},
}

func PhoneBookListDescription() []string {
	var arr = []string{}
	for _, data := range GetListPhonebookData {
		arr = append(arr, data.Description)
	}
	return arr
}
