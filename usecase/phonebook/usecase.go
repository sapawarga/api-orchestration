package phonebook

import (
	"context"

	model "github.com/sapawarga/api-orchestration/model/phonebook"
	"github.com/sapawarga/api-orchestration/repository"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Usecase struct {
	Repo   repository.PhonebookI
	Logger kitlog.Logger
}

func NewUsecase(repo repository.PhonebookI, logger kitlog.Logger) *Usecase {
	return &Usecase{
		Repo:   repo,
		Logger: logger,
	}
}

func (svc *Usecase) GetList(ctx context.Context, req *model.GetListRequest) (*model.GetListResponse, error) {
	logger := kitlog.With(svc.Logger, "method", "GetList")
	resp, err := svc.Repo.GetList(ctx, req)
	if err != nil {
		level.Error(logger).Log("error_get_list", err)
		return nil, err
	}

	return resp, nil
}

func (svc *Usecase) GetDetail(ctx context.Context, id int64) (*model.PhoneDetail, error) {
	logger := kitlog.With(svc.Logger, "method", "GetDetail")
	resp, err := svc.Repo.GetDetail(ctx, id)
	if err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return nil, err
	}

	return resp, nil
}

func (svc *Usecase) Insert(ctx context.Context, req *model.AddPhonebookRequest) error {
	logger := kitlog.With(svc.Logger, "method", "Insert")
	if err := svc.Repo.Insert(ctx, req); err != nil {
		level.Error(logger).Log("error_insert", err)
		return err
	}
	return nil
}

func (svc *Usecase) Update(ctx context.Context, req *model.UpdatePhonebookRequest) error {
	logger := kitlog.With(svc.Logger, "method", "Update")
	if err := svc.Repo.Update(ctx, req); err != nil {
		level.Error(logger).Log("error_update", err)
		return err
	}
	return nil
}

func (svc *Usecase) Delete(ctx context.Context, id int64) error {
	logger := kitlog.With(svc.Logger, "method", "Delete")
	if err := svc.Repo.Delete(ctx, id); err != nil {
		level.Error(logger).Log("error_delete", err)
		return err
	}
	return nil
}
