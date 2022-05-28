package models

import (
	"MedicalQuestionnaireService/libs"

	"gorm.io/gorm"
)

type MasterQuestions struct {
	gorm.Model
	MasterQuestionText string `gorm:"type:varchar(100)" json:"master_question_text"`
	Selection          int    `gorm:"type:tinyint" json:"selection"`
	Statement          string `gorm:"type:varchar(100)" json:"statement"`
	MultipleChoice     string `gorm:"type:varchar(100)" json:"multiple_choice"`
	SingleChoice       string `gorm:"type:varchar(100)" json:"single_choice"`
	Paragraph          string `gorm:"type:varchar(1000)" json:"paragraph"`
	Jump               int    `gorm:"type:tinyint" json:"jump"`
	Ng                 int    `gorm:"type:tinyint" json:"ng"`
	Department         string `gorm:"type:varchar(100)" json:"department"`
}

func GetAllMasterQuestions() (*gorm.DB, []*MasterQuestions) {
	var masterQuestions []*MasterQuestions
	db := libs.Db.Find(&masterQuestions)
	return db, masterQuestions
}

func GetMasterQuestion(MasterQuestionId uint) (*gorm.DB, []*MasterQuestions) {
	var masterQuestion []*MasterQuestions
	db := libs.Db.Find(&masterQuestion, MasterQuestionId)
	return db, masterQuestion
}

func (m *MasterQuestions) CreateMasterQuestion(question *MasterQuestions) error {
	if err := libs.Db.Model(m).Create(question).Error; err != nil {
		return err
	}
	return nil
}

func (m *MasterQuestions) UpdateMasterQuestion(question *MasterQuestions) error {
	if err := libs.Db.Model(m).Updates(question).Error; err != nil {
		return err
	}
	return nil
}

func (m *MasterQuestions) DeleteMasterQuestion(id uint) error {
	if err := libs.Db.Delete(m, id).Error; err != nil {
		return err
	}
	return nil
}