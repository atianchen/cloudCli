package dto

/**
 *
 * @author jensen.chen
 * @date 2022/7/8
 */

/**
分页
*/
type PageRequestDto struct {
	Page    int
	Limit   int
	Keyword string
}
