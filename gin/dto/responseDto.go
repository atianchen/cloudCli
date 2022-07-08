package dto

/**
 *
 * @author jensen.chen
 * @date 2022/7/7
 */
const (
	RESPONSEMSG_CODE_ERROR = iota
	RESPONSEMSG_CODE_SUCCESS
)

type ResponseMsg struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func BuildSuccessMsg(data interface{}) *ResponseMsg {
	return &ResponseMsg{Code: RESPONSEMSG_CODE_SUCCESS, Data: data}
}

func BuildEmptySuccessMsg() *ResponseMsg {
	return &ResponseMsg{Code: RESPONSEMSG_CODE_SUCCESS}
}

func BuildErrorMsg(error string) *ResponseMsg {
	return &ResponseMsg{Code: RESPONSEMSG_CODE_ERROR, Error: error}
}

type PageResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalCount int `json:"totalCount"`
	Items      any `json:"items"`
}
