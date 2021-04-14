package phonebook

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
