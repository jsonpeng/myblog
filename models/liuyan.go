package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

func AddLiuyan(allcontent string, Created time.Time) error {
	o := orm.NewOrm()
	o.Using("default")
	pdd := new(Liuyan)
	pdd.AllContent = allcontent
	pdd.Created = time.Now()
	_, err := o.Insert(pdd)
	return err
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
