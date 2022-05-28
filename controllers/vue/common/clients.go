package controllers

import (
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/models"

	"github.com/kataras/iris/v12"
)

func CreateClient(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	client := new(models.Clients)

	if err := ctx.ReadJSON(&client); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr3, nil, err.Error()))
		return
	}

	if !libs.VerifyMailFormat(client.Mail) {
		ctx.JSON(libs.ApiResource(libs.CodeErr, nil, "メール形式じゃないよ！"))
		return
	}

	s := &models.Search{
		Fields: []*models.Field{
			{
				Key:       "mail",
				Condition: "=",
				Value:     client.Mail,
			},
		},
	}

	_, err := models.GetUser(s)
	if err == nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr, nil, "メールアドレスは既に使用されています。"))
		return
	}
	
	models.CreatClients(client)
	ctx.JSON(libs.ApiResource(libs.CodeSuc, client, libs.MegSuc))
}
