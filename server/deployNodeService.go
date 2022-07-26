package server

import (
	"cloudCli/db"
	"cloudCli/domain"
	"github.com/google/uuid"
)

/**
 * 客户端信息
 * @author jensen.chen
 * @date 2022/7/25
 */

const (
	NODE_STATUS_ERR = iota
	NODE_STATUS_VALID
	NODE_STATUS_OFFLINE
)

type DeployNodeService struct {
}

func (c *DeployNodeService) Save(node *domain.DeployNode) error {
	nodeNum, _ := c.CountByIpAndPort(node.Ip, node.Port)
	if nodeNum < 1 {
		_, err := db.DbInst.Execute("insert into deploy_node (id,name,ip,port,ts,status) values (?,?,?,?,?,?)",
			uuid.New(), node.Name, node.Ip, node.Port, node.Ts, node.Status)
		return err
	}
	return nil
}

func (c *DeployNodeService) UpdateTs(node *domain.DeployNode) error {
	_, err := db.DbInst.Execute("update deploy_node set ts=?,status=? and id=?", node.Ts, node.Status, node.Id)
	return err
}

func (r *DeployNodeService) LoadAll(dest *[]domain.DeployNode) error {
	return db.DbInst.Query(dest, "select  * from deploy_node")
}

func (c *DeployNodeService) Remove(node *domain.DeployNode) error {
	_, err := db.DbInst.Execute("delete from deploy_node where ip=? and port=?", node.Ip, node.Port)
	return err
}

func (r *DeployNodeService) GetByPrimary(priKey string) (*domain.DeployNode, error) {
	node := domain.DeployNode{}
	err := db.DbInst.Get(&node, "select * from deploy_node where id=?", priKey)
	return &node, err
}

func (r *DeployNodeService) GetByIpAndPort(ip string, port int) (*domain.DeployNode, error) {
	node := domain.DeployNode{}
	err := db.DbInst.Get(&node, "select * from deploy_node where ip=? and port=?", ip, port)
	return &node, err
}

func (r *DeployNodeService) CountByIpAndPort(ip string, port int) (int, error) {
	var num int
	err := db.DbInst.Get(&num, "select count(*) from deploy_node where ip=? and port=? ", ip, port)
	return num, err
}
