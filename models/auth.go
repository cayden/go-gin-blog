// Package models -----------------------------
// @file      : auth.go
// @author    : cayden
// @contact   : cuiran2001@163.com
// @time      : 2022/6/17 13:03
// -------------------------------------------

package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
