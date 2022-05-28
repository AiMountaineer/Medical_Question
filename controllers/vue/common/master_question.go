package controllers

import (
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/models"
	"fmt"

	"github.com/kataras/iris/v12"
)

func ShowAllMasterQuestions(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_, masterQuestions := models.GetAllMasterQuestions()
	ctx.JSON(libs.ApiResource(libs.CodeSuc, masterQuestions, libs.MegSuc))
}

func ShowMasterQuestion(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)

	masterQuestionId, err := ctx.Params().GetUint("master_question_id")

	if err != nil {
		fmt.Printf("%v", "URLパラメータが受け取れていません。")
		return
	}

	_, masterQuestion := models.GetMasterQuestion(masterQuestionId)
	ctx.JSON(libs.ApiResource(libs.CodeSuc, masterQuestion, libs.MegSuc))
}

func CreateMasterQuestion(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	var masterQuestion models.MasterQuestions

	if err := ctx.ReadJSON(&masterQuestion); err != nil {
		fmt.Printf("%v", "JSONが受け取れていません。")
		return
	}
	err := masterQuestion.CreateMasterQuestion(&masterQuestion)

	if err != nil {
		fmt.Printf("%v", "追加が失敗しました。")
		return
	}
	ctx.JSON(libs.ApiResource(1, nil, "テーブルに追加しました。"))
}

func UpdateMasterQuestion(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	var masterQuestion models.MasterQuestions

	if err := ctx.ReadJSON(&masterQuestion); err != nil {
		fmt.Printf("%v", "JSONが受け取れていません。")
		return
	}
	err := masterQuestion.UpdateMasterQuestion(&masterQuestion)
	if err != nil {
		fmt.Printf("%v", "JSONが受け取れていません。")
		return
	}
	ctx.JSON(libs.ApiResource(1, nil, "テーブルがアップデートされました。"))
}

func DeleteMasterQuestionById(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	var masterQuestion models.MasterQuestions
	if err := ctx.ReadJSON(&masterQuestion); err != nil {
		fmt.Printf("%v", "JSONが受け取れていません。")
		return
	}

	err := masterQuestion.DeleteMasterQuestion(masterQuestion.ID)

	if err != nil {
		fmt.Printf("%v", "JSONが受け取れていません。")
		return
	}

	ctx.JSON(libs.ApiResource(1, nil, "テーブルから削除されました。"))
}