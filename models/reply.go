package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

func GetAllReply(tid string) (reply []*Reply, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	reply = make([]*Reply, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("reply")
	_, err = qs.Filter("tid", tidNum).All(&reply)
	return reply, err
}
func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Reply{
		Tid:      tidNum,
		Nickname: nickname,
		Content:  content,
		Created:  time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err

}
