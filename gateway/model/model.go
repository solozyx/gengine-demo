package model

type (
	HttpModel struct {
		Path   string `json:"path"`
		Method string `json:"method"`
		IP     string `json:"ip"`
	}

	BaseModel struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)
