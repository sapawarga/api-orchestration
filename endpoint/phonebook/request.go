package phonebook

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sapawarga/api-orchestration/helper"
)

type GetListRequest struct {
	Search     string `httpquery:"search,omitempty"`
	RegencyID  int64  `httpquery:"regency_id,omitempty"`
	DistrictID int64  `httpquery:"district_id,omitempty"`
	VillageID  int64  `httpquery:"village_id,omitempty"`
	Status     *int64 `httpquery:"status,omitempty"`
	Limit      int64  `httpquery:"limit,omitempty"`
	Page       int64  `httpquery:"page,omitempty"`
	Latitude   string `httpquery:"latitude,omiempty"`
	Longitude  string `httpquery:"longitude,omitempty"`
}

type GetDetailRequest struct {
	ID int64 `httpparam:"id"`
}

type CreatePhonebook struct {
	Name           string         `json:"name"`
	Address        string         `json:"address"`
	Description    string         `json:"description"`
	PhoneNumbers   []*PhoneNumber `json:"phone_numbers"`
	CategoryID     int64          `json:"category_id"`
	RegencyID      *int64         `json:"regency_id"`
	DistrictID     *int64         `json:"district_id"`
	VillageID      *int64         `json:"village_id"`
	Latitude       string         `json:"latitude"`
	Longitude      string         `json:"longitude"`
	CoverImagePath string         `json:"cover_image_path"`
	Status         *int64         `json:"status"`
}

type PhoneNumber struct {
	Type        string `json:"type"`
	PhoneNumber string `json:"phone_number"`
}

func Validation(in *CreatePhonebook) error {
	return validation.ValidateStruct(in,
		validation.Field(in.Name, validation.Required),
		validation.Field(in.PhoneNumbers, validation.Required),
		validation.Field(in.CategoryID, validation.Required),
		validation.Field(&in.Status, validation.Required, validation.In(helper.ACTIVED, helper.DELETED, helper.INACTIVED)),
	)
}
