package models

import (
	"MedicalQuestionnaireService/libs"

	"gorm.io/gorm"
)

type Questions struct {
	gorm.Model
	QuestionText   string `gorm:"type:varchar(100)" json:"question_text"`
	Selection      int    `gorm:"type:tinyint" json:"selection"`
	Statement      string `gorm:"type:varchar(100)" json:"statement"`
	DoctorId       int    `gorm:"type:int" json:"doctor_id"`
	MultipleChoice string `gorm:"type:varchar(100)" json:"multiple_choice"`
	SingleChoice   string `gorm:"type:varchar(100)" json:"single_choice"`
	Paragraph      string `gorm:"type:varchar(1000)" json:"paragraph"`
	Jump           int    `gorm:"type:tinyint" json:"jump"`
	Ng             int    `gorm:"type:tinyint" json:"ng"`
	Department     string `gorm:"type:varchar(100)" json:"department"`
}

func CreateQuestion(question *Questions) error {
	if err := libs.Db.Create(question).Error; err != nil {
		return err
	}
	return nil
}

func (q *Questions) UpdateQuestion(question *Questions) error {
	if err := libs.Db.Model(q).Updates(question).Error; err != nil {
		return err
	}
	return nil
}

func (q *Questions) DeleteQuestion(id uint) error {
	if err := libs.Db.Delete(q, id).Error; err != nil {
		return err
	}
	return nil
}
