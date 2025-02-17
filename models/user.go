package models

import (
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/golang-jwt/jwt/v5"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int64
	Username string
	Password string
}

// Создание секретного ключа
var SecretKey = []byte("your-secret-key")

func CreateToken(u User) (string, error) {
	// создаем заявку
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   u.Id,
		"name": u.Username,
	})
	// генерируем токен
	tokenString, err := claims.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	// Проверка на ошибки
	if err != nil {
		return nil, err
	}

	// Проверка валидности токена
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Возврат данных токена
	return token.Claims.(jwt.MapClaims), nil
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
	qb.Select("id", "username", "password").From("user").Where("id > ?").OrderBy("id").Desc().Limit(10)
	o.Raw(qb.String(), 0).QueryRows(&users)
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

func Login(u User) string {
	o := orm.NewOrmUsingDB("test")
	err := o.Read(&u, "username", "password")
	if err != orm.ErrNoRows {
		tokenString, _ := CreateToken(u)
		return tokenString
	}
	return ""
}

func DeleteUser(uid int64) bool {
	o := orm.NewOrmUsingDB("test")
	user := User{Id: uid}
	_, err := o.Delete(&user)
	return err == nil
}
