package phonebook

type GetListRequest struct {
	Search     string
	RegencyID  int64
	DistrictID int64
	VillageID  int64
	Status     *int64
	Limit      int64
	Page       int64
	Latitude   string
	Longitude  string
}

type AddPhonebookRequest struct {
	Name           string `json:"name"`
	PhoneNumbers   string `json:"phone_numbers"`
	Address        string `json:"address"`
	Description    string `json:"description"`
	RegencyID      int64  `json:"regency_id"`
	DistrictID     int64  `json:"district_id"`
	VillageID      int64  `json:"village_id"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	CoverImagePath string `json:"cover_image_path"`
	Status         int64  `json:"status"`
	CategoryID     int64  `json:"category_id"`
}

type UpdatePhonebookRequest struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	PhoneNumbers   string `json:"phone_numbers"`
	Address        string `json:"address"`
	Description    string `json:"description"`
	RegencyID      int64  `json:"regency_id"`
	DistrictID     int64  `json:"district_id"`
	VillageID      int64  `json:"village_id"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	CoverImagePath string `json:"cover_image_path"`
	Status         *int64 `json:"status"`
	CategoryID     int64  `json:"category_id"`
}
