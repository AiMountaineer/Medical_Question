package models

import (
	"MedicalQuestionnaireService/libs"

	"gorm.io/gorm"
)

type Answers struct {
	gorm.Model
	DoctorId  int    `gorm:"type:uint" json:"doctor_id"`
	ClientId  int    `gorm:"type:uint" json:"client_id"`
	Paragraph string `gorm:"type:varchar(255)" json:"paragraph"`
	Statement string `gorm:"type:varchar(100)" json:"statement"`
	Selection string `gorm:"type:varchar(100)" json:"selection"`
}

func GetSelectedAnswer(doctorId uint) (*gorm.DB, []*Answers) {
	var answer []*Answers
	db := libs.Db.Where("doctor_id = ?", doctorId).Find(&answer)
	return db, answer
}


func GetAnswerByClientId(doctorId, clientId uint) (*gorm.DB, []*Answers){
	var answer []*Answers
	db := libs.Db.Where("doctor_id = ? AND client_id = ?", doctorId, clientId).Find(&answer)
	return db, answer
}
