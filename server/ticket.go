package server

import "github.com/golang-jwt/jwt"

/**
 * 节点控制的票据
 * @author jensen.chen
 * @date 2022/7/26
 */
type NcTicket struct {
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Ticket string `json:"ticket"` //票据号
}
type NcTicketToken struct {
	NcTicket
	jwt.StandardClaims
}
