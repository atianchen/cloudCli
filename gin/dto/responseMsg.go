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
	Code  int
	Error string
	Data  interface{}
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
