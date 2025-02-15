package models

import (
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)



func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int64
	Username string
	Password string
}

func AddUser(u User) int64 {
	o := orm.NewOrmUsingDB("test")
	id, err := o.Insert(&u)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func GetUser(uid int64) (u *User, err error) {
	o := orm.NewOrmUsingDB("test")
	user := User{Id: uid}
	err = o.Read(&user)
	if err == orm.ErrNoRows {
		return nil, errors.New("Пользователь с таким id не найден")
	}
	return &user, nil
}

func GetAllUsers() *[]User {
	var users []User
	o := orm.NewOrmUsingDB("test")
	qb, _ := orm.NewQueryBuilder("postgres")
	qb.Select("id,username").From("user").Where("id > 0").OrderBy("id").Desc().Limit(10)
	o.Raw(qb.String()).QueryRows(&users)
	return &users
}

func UpdateUser(uu *User) (err error) {
	o := orm.NewOrmUsingDB("test")
	_, err = o.Update(uu, "username", "password")
	if err != nil {
		return errors.New("Пользователь не найден")
	}
	return nil
}

func Login(username, password string) bool {
	o := orm.NewOrmUsingDB("test")
	user := User{Username: username, Password: password}
	err := o.Read(&user, "username", "password")
	fmt.Println(err, &user)
	return err != orm.ErrNoRows
}

func DeleteUser(uid int64) bool {
	o := orm.NewOrmUsingDB("test")
	user := User{Id: uid}
	_, err := o.Delete(&user)
	return err == nil
}
