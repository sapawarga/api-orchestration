package testcases

import "errors"

type DeletePhonebook struct {
	Description     string
	UsecaseParam    int64
	RepositoryParam int64
	MockUsecase     error
	MockRepository  error
}

var DeletePhonebookData = []DeletePhonebook{
	{
		Description:     "success_delete",
		UsecaseParam:    1,
		RepositoryParam: 1,
		MockUsecase:     nil,
		MockRepository:  nil,
	}, {
		Description:     "failed_delete",
		UsecaseParam:    1,
		RepositoryParam: 1,
		MockUsecase:     errors.New("failed_delete"),
		MockRepository:  errors.New("failed_delete"),
	},
}

func DeletePhonebookDescription() []string {
	var arr = []string{}
	for _, data := range DeletePhonebookData {
		arr = append(arr, data.Description)
	}
	return arr
}
