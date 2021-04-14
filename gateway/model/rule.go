package model

type (
	// {
	//    "rule_name":"新增用餐规则1",
	//    "once_limit":30.50,
	//    "day_limit":88.88
	//}

	RuleFoodCreateRequest struct {
		HttpModel
		Title     string `json:"title"`
		OnceLimit string `json:"once_limit"`
		DayLimit  string `json:"day_limit"`
	}
	RuleFoodCreateResponse struct {
		BaseModel
	}
)
