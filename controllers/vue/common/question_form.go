package controllers

import (
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/models"

	"github.com/kataras/iris/v12"
)

func SaveForm(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	form := new(models.QuestionForms)
	if err := ctx.ReadJSON(&form); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr3, nil, err.Error()))
	}
	models.SaveForm(form)
	ctx.JSON(libs.ApiResource(libs.CodeSuc, form, libs.MegSuc))	
}
