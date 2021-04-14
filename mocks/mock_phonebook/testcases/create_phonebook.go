package testcases

import (
	"errors"

	model "github.com/sapawarga/api-orchestration/model/phonebook"
)

type CreatePhonebook struct {
	Description     string
	UsecaseParam    *model.AddPhonebookRequest
	RepositoryParam *model.AddPhonebookRequest
	MockUsecase     error
	MockRepository  error
}

var (
	createPhone = &model.AddPhonebookRequest{
		Name:           "test",
		PhoneNumbers:   "[{\"type\":\"phone\", \"phone\":\"+628773827392\"}]",
		Address:        "address",
		Description:    "description",
		RegencyID:      1,
		DistrictID:     1,
		VillageID:      1,
		Latitude:       "1.2345",
		Longitude:      "2.1345",
		CoverImagePath: "image_path",
		Status:         10,
		CategoryID:     1,
	}
)

var CreatePhonebookData = []CreatePhonebook{
	{
		Description:     "success_created",
		UsecaseParam:    createPhone,
		RepositoryParam: createPhone,
		MockUsecase:     nil,
		MockRepository:  nil,
	}, {
		Description:     "failed_created",
		UsecaseParam:    createPhone,
		RepositoryParam: createPhone,
		MockUsecase:     errors.New("failed_created"),
		MockRepository:  errors.New("failed_created"),
	},
}

func InsertPhonebookDescription() []string {
	var arr = []string{}
	for _, data := range CreatePhonebookData {
		arr = append(arr, data.Description)
	}
	return arr
}
