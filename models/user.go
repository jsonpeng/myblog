package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func RegisterUser(username, password, pass, email string) error {
	o := orm.NewOrm()
	o.Using("default")
	pro := new(User)
	pro.Username = username
	pro.Password = password
	pro.Pass = pass
	pro.Email = email
	pro.Created = time.Now()
	_, err := o.Insert(pro)
	return err
}

func FindUser(username string) (User, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Username: username}
	err := o.Read(&user, "username")

	return user, err
}
