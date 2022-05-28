package models

import (
	"MedicalQuestionnaireService/libs"
	"encoding/json"
	"time"

	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"

	"gorm.io/gorm"
)

type Clients struct {
	gorm.Model
	FirstName    string `gorm:"type:varchar(100)" json:"first_name"`
	LastName     string `gorm:"type:varchar(100)" json:"last_name"`
	Mail         string `gorm:"type:varchar(100)" json:"mail"`
	PhoneNumber  string `gorm:"type:varchar(100)" json:"phone_number"`
	Password     string `gorm:"type:varchar(100)" json:"password"`
	Sex          string `gorm:"type:varchar(100)" json:"sex"`
	Birthday     string `gorm:"type:datetime" json:"birthday"`
	RefreshToken string `gorm:"type:varchar(100)" json:"refresh_token"`
	AccessToken  string `gorm:"type:varchar(100)" json:"access_token"`
}

func CreatClients(client *Clients) error {
	client.Password = libs.HashPassword(client.Password)
	if err := libs.Db.Create(client).Error; err != nil {
		return err
	}
	return nil
}

//	ログイン後に返送する情報
type ReturnUserInfo struct {
	UserId       uint            `json:"userId" comment:"user id"`
	FirstName    string          `json:"firstName" comment:"firstname"`
	LastName     string          `json:"lastName" comment:"lastname"`
	Sex          string          `json:"sex" comment:"sex"`
	AccessToken  json.RawMessage `json:"access_token"`
	RefreshToken json.RawMessage `json:"refresh_token"`
}

func (u *Clients) CheckLogin(ctx iris.Context, password string) (*jwt.TokenPair, int64, string) {
	if u.ID == 0 {
		return nil, 5, "ユーザが存在しません"
	} else {
		if bcrypt.Match(password, u.Password) {
			tokenPair := u.GenerateTokenPair(ctx)
			return tokenPair, libs.CodeSuc, "ログインに成功しました"
		} else {
			return nil, 6, "ログイン情報が正しくありません"
		}
	}
}

func GetUser(search *Search) (*Clients, error) {
	t := NewUser()
	err := Found(search).First(t).Error
	return t, err
}

func NewUser() *Clients {
	return &Clients{
		Model: gorm.Model{
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
}

//generate new access token and refresh token for user
func (u *Clients) UpdateToken(ctx iris.Context) (*jwt.TokenPair, int64, string) {
	tokenPair := u.refreshToken(ctx)
	if tokenPair != nil {
		return tokenPair, 200, "token更新成功"
	} else {
		return tokenPair, 400, "token更新失敗"
	}
}
