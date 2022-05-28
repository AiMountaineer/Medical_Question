package models

import (
	"MedicalQuestionnaireService/libs"
	"time"

	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"gorm.io/gorm"
)

type Doctors struct {
	gorm.Model
	FirstName    string `gorm:"type:varchar(100)" json:"first_name"`
	LastName     string `gorm:"type:varchar(100)" json:"last_name"`
	Mail         string `gorm:"type:varchar(100)" json:"mail"`
	PhoneNumber  string `gorm:"type:varchar(100)" json:"phone_number"`
	Birthday     string `gorm:"type:datetime" json:"birthday"`
	Password     string `gorm:"type:varchar(100)" json:"password"`
	Sex          string `gorm:"type:varchar(10);not null" json:"sex"`
	Affiliation  string `gorm:"type:varchar(100)" json:"affiliation"`
	Speciality   string `gorm:"type:varchar(100)" json:"speciality"`
	RefreshToken string `gorm:"type:varchar(100);not null" json:"refresh_token"`
	AccessToken  string `gorm:"type:varchar(100);not null" json:"access_token"`
}

func CreateDoctor(doctor *Doctors) error {
	doctor.Password = libs.HashPassword(doctor.Password)
	if err := libs.Db.Create(doctor).Error; err != nil {
		return err
	}
	return nil
}

func GetAllDoctors() (*gorm.DB, []*Doctors) {
	var doctors []*Doctors
	db := libs.Db.Find(&doctors)
	return db, doctors
}

func GetSelectedDoctor(doctorId uint) (*gorm.DB, []*Doctors) {
	var doctor []*Doctors
	db := libs.Db.Find(&doctor, doctorId)
	return db, doctor
}

func GetDoctor(search *Search) (*Doctors, error) {
	t := NewDoctor()
	err := Found(search).First(t).Error
	return t, err
}

func NewDoctor() *Doctors {
	return &Doctors{
		Model: gorm.Model{
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
}

func (u *Doctors) CheckLogin(ctx iris.Context, password string) (*jwt.TokenPair, int64, string) {
	if u.ID == 0 {
		return nil, 5, "Doctorが存在しません"
	} else {
		if bcrypt.Match(password, u.Password) {
			tokenPair := u.GenerateTokenPair(ctx)
			return tokenPair, libs.CodeSuc, "ログインに成功しました"
		} else {
			return nil, 6, "ログイン情報が正しくありません"
		}
	}
}

func (u *Doctors) UpdateToken(ctx iris.Context) (*jwt.TokenPair, int64, string) {
	tokenPair := u.refreshToken(ctx)
	if tokenPair != nil {
		return tokenPair, 200, "token更新成功"
	} else {
		return tokenPair, 400, "token更新失敗"
	}
}
