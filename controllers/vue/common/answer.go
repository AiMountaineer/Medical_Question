package controllers

import (
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/models"
	"fmt"

	"github.com/kataras/iris/v12"
)

func ShowAnswerLists(ctx iris.Context){
	ctx.StatusCode(iris.StatusOK)

	doctorId, err := ctx.Params().GetUint("doctorId")

	if err != nil {
		fmt.Printf("%v", "URLパラメータが受け取れていません。")
		return
	}

	_, answers := models.GetSelectedAnswer(doctorId)
	if len(answers) == 0 {
		ctx.JSON(libs.ApiResource(2, nil, "検索結果がありません"))
	} else {
		ctx.JSON(libs.ApiResource(libs.CodeSuc, answers, libs.MegSuc))
	}
}

func ShowClientAnswer(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	doctorId, err := ctx.Params().GetUint("doctorId")

	if err != nil {
		fmt.Printf("%v", "URLパラメータが受け取れていません。")
		return
	}

	clientId, err := ctx.Params().GetUint("clientId")


	if err != nil {
		fmt.Printf("%v", "URLパラメータが受け取れていません。")
		return
	}

	_, answers := models.GetAnswerByClientId(doctorId,clientId)
	if len(answers) == 0 {
		ctx.JSON(libs.ApiResource(2, nil, "検索結果がありません"))
	} else {
		ctx.JSON(libs.ApiResource(libs.CodeSuc, answers, libs.MegSuc))
	}
}
