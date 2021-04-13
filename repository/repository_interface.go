package repository

import (
	"context"

	model "github.com/sapawarga/api-orchestration/model/phonebook"
)

type PhonebookI interface {
	GetList(ctx context.Context, req *model.GetListRequest) (*model.GetListResponse, error)
	GetDetail(ctx context.Context, id int64) (*model.PhoneDetail, error)
	Insert(ctx context.Context, req *model.AddPhonebookRequest) error
	Update(ctx context.Context, req *model.UpdatePhonebookRequest) error
	Delete(ctx context.Context, id int64) error
}
