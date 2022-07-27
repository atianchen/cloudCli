package server

import "github.com/golang-jwt/jwt"

/**
 * DeployNodeDto
 * @author jensen.chen
 * @date 2022/7/26
 */
type NodePayload struct {
	Content string
}

type DeployNodeDto struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port uint   `json:"port"`
}

type NcTicket struct {
	Ip     string `json:"ip"`
	Port   uint   `json:"port"`
	Ticket string `json:"ticket"` //票据号
	jwt.StandardClaims
}
