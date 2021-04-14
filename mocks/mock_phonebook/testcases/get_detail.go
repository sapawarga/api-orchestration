package testcases

import (
	"errors"

	"github.com/sapawarga/api-orchestration/helper"
	model "github.com/sapawarga/api-orchestration/model/phonebook"
)

type GetDetailRepository struct {
	Result *model.PhoneDetail
	Error  error
}

type GetDetailUsecase struct {
	Result *model.PhoneDetail
	Error  error
}

type GetDetailPhonebook struct {
	Description      string
	UsecaseParams    int64
	RepositoryParams int64
	MockRepository   GetDetailRepository
	MockUsecase      GetDetailUsecase
}

var (
	currentTime, _ = helper.GetCurrentTimeUTC()
	phoneDetail    = &model.PhoneDetail{
		ID:           1,
		PhoneNumbers: "[{\"type\":\"phone\", \"phone\":\"+628773827392\"}]",
		Description:  "description",
		Name:         "test",
		Address:      "address",
		Latitude:     "1.1302",
		Longitude:    "2.12343",
		Status:       10,
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
		CategoryID:   1,
		CategoryName: "category",
		Distance:     "10",
		RegencyID:    1,
		RegencyName:  "regency",
		DistrictID:   1,
		DistrictName: "district",
		VillageID:    1,
		VillageName:  "village",
	}
)

var GetDetailPhonebookData = []GetDetailPhonebook{
	{
		Description:      "success_get_detail",
		UsecaseParams:    1,
		RepositoryParams: 1,
		MockRepository: GetDetailRepository{
			Result: phoneDetail,
			Error:  nil,
		},
		MockUsecase: GetDetailUsecase{
			Result: phoneDetail,
			Error:  nil,
		},
	}, {
		Description:      "failed_get_detail",
		UsecaseParams:    1,
		RepositoryParams: 1,
		MockRepository: GetDetailRepository{
			Result: nil,
			Error:  errors.New("error_get_detail"),
		},
		MockUsecase: GetDetailUsecase{
			Result: nil,
			Error:  errors.New("error_get_detail"),
		},
	},
}

func PhonebookDetailDescription() []string {
	var arr = []string{}
	for _, data := range GetDetailPhonebookData {
		arr = append(arr, data.Description)
	}
	return arr
}
