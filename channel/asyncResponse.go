package channel

/**
 *	异步执行的响应
 * @author jensen.chen
 * @date 2022/6/15
 */
type AsyncResponse struct {
	Payload interface{} //载体
	Err     error       //错误信息
}

/**
空Response
*/
func BuildEmptyResponse() *AsyncResponse {
	return &AsyncResponse{}
}

/**
构建Response
*/
func BuildResponse(result interface{}) *AsyncResponse {
	return &AsyncResponse{Payload: result}
}

/**
构建错误的Response
*/
func BuildErrorResponse(errInfo error) *AsyncResponse {
	return &AsyncResponse{Err: errInfo}
}
