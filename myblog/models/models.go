package models

import (
	"strconv"
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

func RegisterDB() {
	//注册 model
	orm.RegisterModel(new(Category), new(Topic), new(Liuyan), new(Reply))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/default?charset=utf8") //密码为空格式

}
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
	return err

}
func AddLiuyan(allcontent string, Created time.Time) error {
	o := orm.NewOrm()
	o.Using("default")
	pdd := new(Liuyan)
	pdd.AllContent = allcontent
	pdd.Created = time.Now()
	_, err := o.Insert(pdd)
	return err
}
func AddCategory(name string, Created time.Time) error {
	o := orm.NewOrm()
	o.Using("default")
	pro := new(Category)
	pro.Title = name
	pro.Created = time.Now()
	pro.TopicTime = time.Now()
	_, err := o.Insert(pro)
	return err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	o.Using("default")

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func AddTopic(title, category, desc, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:     title,
		Category:  category,
		Desc:      desc,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-Created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}
func GetAllLiuyan(isDesc bool) ([]*Liuyan, error) {
	o := orm.NewOrm()

	liuyan := make([]*Liuyan, 0)
	qs := o.QueryTable("liuyan")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-Created").All(&liuyan)
	} else {
		_, err = qs.All(&liuyan)
	}
	return liuyan, err
}
func DelTopics(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	o.Using("default")

	topic := &Topic{Id: cid}
	_, err = o.Delete(topic)
	return err
}

func DelLiuyan(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	o.Using("default")

	liu := &Liuyan{Id: cid}
	_, err = o.Delete(liu)
	return err
}

func ModifyTopic(tid, title, category, desc, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Desc = desc
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return err
}
