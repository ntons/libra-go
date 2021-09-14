package libra

type ErrorDetail struct {
	Err  int         `json:"err"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
