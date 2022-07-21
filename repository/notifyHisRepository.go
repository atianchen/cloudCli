package repository

import (
	"cloudCli/db"
	"cloudCli/domain"
	"cloudCli/utils/timeUtils"
	"github.com/google/uuid"
)

/**
 * 通知历史
 * @author jensen.chen
 * @date 2022/7/20
 */
type NotifyHistoryRepository struct {
}

func (*NotifyHistoryRepository) Save(nh *domain.NotifyHistory) error {
	nh.SendTime = timeUtils.NowUnixTime()
	_, err := db.DbInst.Execute("insert into notify_history (id,receiver,content,send_time) values (?,?,?,?)",
		uuid.New(), nh.Receiver, nh.Content, nh.SendTime)
	return err

}

/**
 * 分页查询
 */
func (r *NotifyHistoryRepository) PageQuery(dest *[]domain.NotifyHistory, startIndex int, limit int) error {
	sql := "select * from notify_history"
	sql += " limit ? offset  ?"
	return db.DbInst.Query(dest, sql, limit, startIndex)
}

/**
 * 根据主键查询
 */
func (r *NotifyHistoryRepository) GetByPrimary(priKey string) (*domain.NotifyHistory, error) {
	doc := domain.NotifyHistory{}
	err := db.DbInst.Get(&doc, "select * from notify_history where id=?", priKey)
	return &doc, err
}
