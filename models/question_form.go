package models

import (
	"MedicalQuestionnaireService/libs"

	"gorm.io/gorm"
)

type QuestionForms struct {
	gorm.Model
	DoctorId       int    `gorm:"type:uint" json:"doctor_id"`
	QuestionIdList string `gorm:"type:varchar(100)" json:"question_id_list"`
}

func SaveForm(form *QuestionForms) error {
	if err := libs.Db.Create(form).Error; err != nil {
		return err
	}
	return nil
}
