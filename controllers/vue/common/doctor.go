package controllers

import (
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/models"
	"fmt"

	"github.com/kataras/iris/v12"
)

func ShowAllDoctors(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_, doctors := models.GetAllDoctors()
	ctx.JSON(libs.ApiResource(libs.CodeSuc, doctors, libs.MegSuc))
}

func ShowSelectedDoctor(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	doctorId, err := ctx.Params().GetUint("doctorId")
	if err != nil {
		fmt.Printf("%v", "URLパラメータが受け取れていません。")
		return
	}
	_, doctor := models.GetSelectedDoctor(doctorId)

	if len(doctor) == 0 {
		ctx.JSON(libs.ApiResource(2, nil, "検索結果がありません"))
	} else {
		ctx.JSON(libs.ApiResource(libs.CodeSuc, doctor, libs.MegSuc))
	}	
}

func CreateDoctor(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	doctor := new(models.Doctors)

	if err := ctx.ReadJSON(&doctor); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr3, nil, err.Error()))
		return
	}

	if !libs.VerifyMailFormat(doctor.Mail) {
		ctx.JSON(libs.ApiResource(libs.CodeErr, nil, "メール形式じゃないよ！"))
		return
	}

	s := &models.Search{
		Fields: []*models.Field{
			{
				Key:       "mail",
				Condition: "=",
				Value:     doctor.Mail,
			},
		},
	}

	_, err := models.GetDoctor(s)
	if err == nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr, nil, "メールアドレスは既に使用されています。"))
		return
	}

	models.CreateDoctor(doctor)
	ctx.JSON(libs.ApiResource(libs.CodeSuc, doctor, libs.MegSuc))
}
