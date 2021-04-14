package testcases

import (
	"errors"

	model "github.com/sapawarga/api-orchestration/model/phonebook"
)

type UpdatePhonebook struct {
	Description     string
	UsecaseParam    *model.UpdatePhonebookRequest
	RepositoryParam *model.UpdatePhonebookRequest
	MockUsecase     error
	MockRepository  error
}

var (
	updatePhonebook = &model.UpdatePhonebookRequest{
		ID:           1,
		Name:         "test",
		PhoneNumbers: "[{\"type\":\"phone\", \"phone\":\"+628773827392\"}]",
		Address:      "addrss",
		Description:  "description",
		RegencyID:    1,
		DistrictID:   1,
		VillageID:    1,
	}
)

var UpdatePhonebookData = []UpdatePhonebook{
	{
		Description:     "success_update",
		UsecaseParam:    updatePhonebook,
		RepositoryParam: updatePhonebook,
		MockUsecase:     nil,
		MockRepository:  nil,
	}, {
		Description:     "failed_update",
		UsecaseParam:    updatePhonebook,
		RepositoryParam: updatePhonebook,
		MockUsecase:     errors.New("failed_update"),
		MockRepository:  errors.New("failed_update"),
	},
}

func UpdatePhonebookDescription() []string {
	var arr = []string{}
	for _, data := range UpdatePhonebookData {
		arr = append(arr, data.Description)
	}
	return arr
}
