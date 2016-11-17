package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

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
