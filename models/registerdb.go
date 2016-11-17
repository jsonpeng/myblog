package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Desc            string `orm:"size(800)"`
	Content         string `orm:"size(8000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}
type Reply struct {
	Id       int64
	Tid      int64
	Nickname string
	Content  string    `orm:"size(1000)"`
	Created  time.Time `orm:"index"`
	Rcount   int64
}
type Liuyan struct {
	Id         int64
	AllContent string    `orm:"size(1000)"`
	Created    time.Time `orm:"index"`
	Lcount     int64
}

type User struct {
	Id       int64
	Username string
	Password string
	Pass     string
	Email    string
	Created  time.Time `orm:"index"`
}

func RegisterDB() {
	//注册 model
	orm.RegisterModel(new(Category), new(Topic), new(Liuyan), new(Reply), new(User))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/default?charset=utf8") //密码为空格式

}
