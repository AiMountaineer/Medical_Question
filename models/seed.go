package models

import (
	"MedicalQuestionnaireService/libs"
	"strconv"
)

func Seed() {
	MasterQuestionSeed() 
	ClientSeed()
	DoctorSeed()
	QuestionFormSeed()
	QuestionSeed()
	AnswerSeed()
}

func MasterQuestionSeed() {
	var mquestion MasterQuestions

	if recordSize := libs.Db.Find(&mquestion).RowsAffected; recordSize == 0 {
		var mquestions []MasterQuestions

		for i := 1; i < 5; i++ {
			for j := 1; j < 5; j++ {
				mquestions = append(mquestions, MasterQuestions{
					MasterQuestionText: "日常的に吐き気を感じることはありますか？",
					Selection:    j % 2,
					Statement:    "xxx",
					Paragraph:    "hoge",
					Jump:         j % 3,
					Ng:           j % 2,
					Department:   "産婦人科",
				})
			}
		}
		libs.Db.Create(&mquestions)
	}
}

func ClientSeed() {
	var client Clients

	if recordSize := libs.Db.Find(&client).RowsAffected; recordSize == 0 {
		var clients []Clients

		// PWの平文は"xxxxxxx"
		for j := 1; j < 11; j++ {
			clients = append(clients, Clients{
				FirstName:    "骨折",
				LastName:     "太郎",
				Mail:         strconv.Itoa(j) + "@example.com",
				PhoneNumber:  "0120-888-888",
				Password:     "$2a$12$/nRZ5DvDZsfUybM7WLphGe/YopgamPORNv.na18KRRPtR.emfIqcW",
				Sex:          "男性",
				Birthday:     "2022-04-27",
				RefreshToken: "xxxxxxxxxxxxxxxxxxx",
				AccessToken:  "xxxxxxxxxxxxxxxxxxxxx",
			})
		}
		libs.Db.Create(&clients)
	}
}

func DoctorSeed() {
	var doctor Doctors

	if recordSize := libs.Db.Find(&doctor).RowsAffected; recordSize == 0 {
		var doctors []Doctors

		for i := 1; i < 11; i++ {
			doctors = append(doctors, Doctors{
				FirstName:    "北条",
				LastName:     "雅子",
				Mail:         strconv.Itoa(i) + "@pexample.com",
				PhoneNumber:  "0120-555-555",
				Birthday:     "2022-11-11",
				Password:     "xxxxxxx",
				Sex:          "女性",
				Affiliation:  "方正病院",
				Speciality:   "内科",
				RefreshToken: "xxxxxxxxxxxxxxxxxxx",
				AccessToken:  "xxxxxxxxxxxxxxxxxxxxx",
			})
		}
		libs.Db.Create(&doctors)
	}
}

func QuestionFormSeed() {
	var form QuestionForms

	if recordSize := libs.Db.Find(&form).RowsAffected; recordSize == 0 {
		var question_forms []QuestionForms

		for j := 1; j < 11; j++ {
			question_forms = append(question_forms, QuestionForms{
				DoctorId:       j,
				QuestionIdList: "[1, 5, 4, 3, 2]",
			})
		}
		libs.Db.Create(&question_forms)
	}
}

func QuestionSeed() {
	var question Questions

	if recordSize := libs.Db.Find(&question).RowsAffected; recordSize == 0 {
		var questions []Questions

		for i := 1; i < 5; i++ {
			for j := 1; j < 11; j++ {
				questions = append(questions, Questions{
					QuestionText: "症状は何ですか？",
					Selection:    j % 2,
					Statement:    "xxx",
					DoctorId:     1,
					Paragraph:    "hoge",
					Jump:         j % 3,
					Ng:           j % 2,
					Department:   "内科",
				})
			}
		}
		libs.Db.Create(&questions)
	}
}

func AnswerSeed() {
	var answer Answers

	if recordSize := libs.Db.Find(&answer).RowsAffected; recordSize == 0 {
		var answers []Answers

		for i := 1; i < 5; i++ {
			for j := 1; j < 11; j++ {
				answers = append(answers, Answers{
					DoctorId:  i,
					ClientId:  j,
					Paragraph: "aaa",
					Statement: "bbb",
					Selection: "ccc",
				})
			}
		}
		libs.Db.Create(&answers)
	}
}
