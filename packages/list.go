package packages

const (
	URL_PACKAGES_LIST = "/ark/open_api/v0/packages"
)

type TimeType string

const (
	CREATED   TimeType = "created_at"
	CONFIRMED TimeType = "confirmed_at"
	UPDATED   TimeType = "updated_at"
)

func (t TimeType) String() string {
	return string(t)
}

type List struct {
	StartTime int64         `json:"start_time" structs:",omitempty"`
	EndTime   int64         `json:"end_time" structs:",omitempty"`
	TimeType  TimeType      `json:"time_type" structs:",omitempty"`
	Logistics LogisticsMode `json:"logistics" structs:",omitempty"`
	Status    OrderStatus   `json:"status" structs:",omitempty"`
	PageNo    int64         `json:"page_no" structs:",omitempty"`
	PageSize  int64         `json:"page_size" structs:",omitempty"`
}

type ListRsp struct {
	TotalNumber int `json:"total_number" mapstructure:"total_number"`
	CurrentPage int `json:"current_page" mapstructure:"current_page"`
	TotalPage   int `json:"total_page" mapstructure:"total_page"`
	PageSize    int `json:"page_size" mapstructure:"page_size"`
	PackageList []struct {
		PackageId          string        `json:"package_id" mapstructure:"package_id"`
		Time               int           `json:"time" mapstructure:"time"`
		ConfirmTime        int           `json:"confirm_time" mapstructure:"confirm_time"`
		ExpressCompanyCode string        `json:"express_company_code" mapstructure:"express_company_code"`
		ExpressCompanyName string        `json:"express_company_name" mapstructure:"express_company_name"`
		ExpressNo          string        `json:"express_no" mapstructure:"express_no"`
		Logistics          LogisticsMode `json:"logistics" mapstructure:"logistics"`
		Status             OrderStatus   `json:"status" mapstructure:"status"`
		ItemNumber         int           `json:"item_number" mapstructure:"item_number"`
	} `json:"package_list" mapstructure:"package_list"`
}
