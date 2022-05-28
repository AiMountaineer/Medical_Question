package controllers

import (
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/models"
	"fmt"

	"github.com/kataras/iris/v12"
)

func CreateQuestion(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	question := new(models.Questions)
	if err := ctx.ReadJSON(&question); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr3, nil, err.Error()))
		return
	}
	models.CreateQuestion(question)
	ctx.JSON(libs.ApiResource(libs.CodeSuc, question, libs.MegSuc))
}

func UpdateQuestion(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	question := new(models.Questions)
	if err := ctx.ReadJSON(&question); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr3, nil, err.Error()))
		return
	}
	if err := question.UpdateQuestion(question); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr18, nil, err.Error()))
		return
	}
	ctx.JSON(libs.ApiResource(libs.CodeSuc, question, libs.MegSuc))
}

func DeleteQuestion(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	question := new(models.Questions)
	if err := ctx.ReadJSON(&question); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr3, nil, err.Error()))
		return
	}
	if err := question.DeleteQuestion(question.ID); err != nil {
		fmt.Printf("%v", "JSONが受け取れていません。")
		return
	}
	ctx.JSON(libs.ApiResource(libs.CodeSuc, question, libs.MegSuc))
}
