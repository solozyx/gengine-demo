package model

type (
	// {
	//    "rule_name":"新增用餐规则1",
	//    "once_limit":30.50,
	//    "day_limit":88.88
	//}

	RuleFoodCreateRequest struct {
		HttpModel
		Title     string  `json:"title"`
		OnceLimit float32 `json:"once_limit"` //OnceLimit string `json:"once_limit"`
		DayLimit  float32 `json:"day_limit"`  //DayLimit  string `json:"day_limit"`
	}
	RuleFoodCreateResponse struct {
		BaseModel
	}

	RuleFoodCheckRequest struct {
		HttpModel
		Once float32 `json:"once"`
		Day  float32 `json:"day"`
	}
	RuleFoodCheckData struct {
		Permitted bool `json:"permitted"`
	}
	RuleFoodCheckResponse struct {
		BaseModel
		Data RuleFoodCheckData `json:"data"`
	}
)
