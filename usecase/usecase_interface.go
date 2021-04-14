package usecase

import (
	"context"

	modelPb "github.com/sapawarga/api-orchestration/model/phonebook"
)

type PhonebookI interface {
	GetList(ctx context.Context, req *modelPb.GetListRequest) (*modelPb.GetListResponse, error)
	GetDetail(ctx context.Context, id int64) (*modelPb.PhoneDetail, error)
	Insert(ctx context.Context, req *modelPb.AddPhonebookRequest) error
	Update(ctx context.Context, req *modelPb.UpdatePhonebookRequest) error
	Delete(ctx context.Context, id int64) error
}
