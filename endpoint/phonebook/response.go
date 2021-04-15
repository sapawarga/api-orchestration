package phonebook

import "time"

type GetListResponse struct {
	Data     []*Phone  `json:"data"`
	Metadata *Metadata `json:"metadata"`
}

type Phone struct {
	ID           int64                    `json:"id"`
	PhoneNumbers []map[string]interface{} `json:"phone_numbers"`
	Description  string                   `json:"description"`
	Name         string                   `json:"name"`
	Address      string                   `json:"address"`
	Latitude     string                   `json:"latitude"`
	Longitude    string                   `json:"longitude"`
	Status       int64                    `json:"status"`
	Category     string                   `json:"category"`
	Distance     string                   `json:"distance"`
}

type Metadata struct {
	Page      int64 `json:"page"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"total_page"`
}

type PhoneDetail struct {
	ID           int64                    `json:"id"`
	PhoneNumbers []map[string]interface{} `json:"phone_numbers"`
	Description  string                   `json:"description"`
	Name         string                   `json:"name"`
	Address      string                   `json:"address"`
	Latitude     string                   `json:"latitude"`
	Longitude    string                   `json:"longitude"`
	Status       int64                    `json:"status"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
	CategoryID   int64                    `json:"category_id"`
	CategoryName string                   `json:"category_name"`
	Distance     string                   `json:"distance"`
	RegencyID    int64                    `json:"regency_id"`
	RegencyName  string                   `json:"regency_name"`
	DistrictID   int64                    `json:"district_id"`
	DistrictName string                   `json:"district_name"`
	VillageID    int64                    `json:"village_id"`
	VillageName  string                   `json:"village_name"`
}
