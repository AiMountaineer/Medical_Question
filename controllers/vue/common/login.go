package controllers

import (
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/models"
	"fmt"

	"github.com/kataras/iris/v12"
)

// UserLogin ユーザログイン
// @Summary ユーザログイン
// @Description ユーザログイン
// @Tags ログイン
// @Accept  application/json
// @Produce application/json
// @Param EdiGo-AppKey header string true "EdiGo-AppKey"
// @Param tUserAccount body models.TUserAccount true "アカウント"
// @Success 200 {object} libs.Response "code=1 成功, code=2 データ通信失敗,code=3 ログイン情報が正しくありません,code=4 無効アカウント,code=5 ユーザが存在しません,code=6 ログイン情報が正しくありません"
// @Failure 400 {object} libs.Response
// @Router /login [post]
func ClientLogin(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	client := new(models.Clients)

	if err := ctx.ReadJSON(&client); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr, nil, err.Error()))
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

	user, err := models.GetUser(s)
	if err != nil || client.Mail == "" || client.Password == "" {
		ctx.JSON(libs.ApiResource(3, nil, "ログイン情報が正しくありません"))
		return
	}

	// if client.Password == cl.Password {
	// 	ctx.JSON(libs.ApiResource(1, nil, "ログイン成功しました"))
	// 	return
	// }

	// if user.LicenseId != 1 {
	// 	ctx.JSON(libs.ApiResource(4, nil, "無効アカウント"))
	// 	return
	// }

	pairtoken, code, msg := user.CheckLogin(ctx, client.Password)
	if code != libs.CodeSuc {
		ctx.JSON(libs.ApiResource(code, nil, msg))
		return
	}

	var respo models.ReturnUserInfo
	respo.UserId = user.ID
	respo.FirstName = user.FirstName
	respo.LastName = user.LastName
	respo.Sex = user.Sex
	respo.AccessToken = pairtoken.AccessToken
	respo.RefreshToken = pairtoken.RefreshToken

	ctx.JSON(libs.ApiResource(code, respo, msg))
}

func DoctorLogin(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	doctor := new(models.Doctors)

	if err := ctx.ReadJSON(&doctor); err != nil {
		ctx.JSON(libs.ApiResource(libs.CodeErr, nil, err.Error()))
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

	user, err := models.GetDoctor(s)
	if err != nil || doctor.Mail == "" || doctor.Password == "" {
		ctx.JSON(libs.ApiResource(3, nil, "ログイン情報が正しくありません"))
		return
	}

	// if client.Password == cl.Password {
	// 	ctx.JSON(libs.ApiResource(1, nil, "ログイン成功しました"))
	// 	return
	// }

	// if user.LicenseId != 1 {
	// 	ctx.JSON(libs.ApiResource(4, nil, "無効アカウント"))
	// 	return
	// }

	pairtoken, code, msg := user.CheckLogin(ctx, doctor.Password)
	if code != libs.CodeSuc {
		ctx.JSON(libs.ApiResource(code, nil, msg))
		return
	}

	var respo models.ReturnUserInfo
	respo.UserId = user.ID
	respo.FirstName = user.FirstName
	respo.LastName = user.LastName
	respo.Sex = user.Sex
	respo.AccessToken = pairtoken.AccessToken
	respo.RefreshToken = pairtoken.RefreshToken

	ctx.JSON(libs.ApiResource(code, respo, msg))
}

// RefreshUserToken token更新
// @Summary token更新
// @Description token更新
// @Tags ログイン
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param EdiGo-AppKey header string true "EdiGo-AppKey"
// @Param refreshToken body string true "refresh token"
// @Success 200 {object} libs.Response "code=1 成功,code=2 ユーザ見つかりませんでした,code=3 token更新失敗"
// @Failure 400 {object} libs.Response
// @Router /refreshtoken [post]
func RefreshUserToken(ctx iris.Context) {
	client := new(models.Clients)
	if !ctx.URLParamExists("refresh_token") {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(libs.ApiResource(400, nil, "refreshtoken not found"))
		return
	}
	client.RefreshToken = ctx.URLParam("refresh_token")
	s := &models.Search{
		Fields: []*models.Field{
			{
				Key:       "refresh_token",
				Condition: "=",
				Value:     client.RefreshToken,
			},
		},
	}

	user, err := models.GetUser(s)
	if err != nil {
		fmt.Println("i am here")
		ctx.JSON(libs.ApiResource(400, nil, "token更新に失敗しました"))
		return
	}
	ctx.StatusCode(iris.StatusOK)
	_, code, msg := user.UpdateToken(ctx)
	if code == 400 {

		ctx.JSON(libs.ApiResource(401, nil, msg))
		return
	}
	ctx.JSON(libs.ApiResource(libs.CodeSuc, user, "操作成功"))
	// refreshToken := ctx.URLParam("refresh_token")
	// s := &models.Search{
	// 	Fields: []*models.Field{
	// 		{
	// 			Key:       "refresh_token",
	// 			Condition: "=",
	// 			Value:     refreshToken,
	// 		},
	// 	},
	// }
	// user, err := models.GetUser(s)
	// if err != nil {
	// 	ctx.JSON(libs.ApiResource(libs.CodeErr, nil, "token更新に失敗しました"))
	// 	return
	// }
	ctx.StatusCode(iris.StatusOK)
	// response, code, msg := user.UpdateToken(ctx)
	// ctx.JSON(libs.ApiResource(code, response, msg))
	// return
}

func RefreshDoctorToken(ctx iris.Context) {
	doctor := new(models.Doctors)
	if !ctx.URLParamExists("refresh_token") {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(libs.ApiResource(400, nil, "refreshtoken not found"))
		return
	}
	doctor.RefreshToken = ctx.URLParam("refresh_token")
	s := &models.Search{
		Fields: []*models.Field{
			{
				Key:       "refresh_token",
				Condition: "=",
				Value:     doctor.RefreshToken,
			},
		},
	}

	user, err := models.GetDoctor(s)
	if err != nil {
		fmt.Println("i am here")
		ctx.JSON(libs.ApiResource(400, nil, "token更新に失敗しました"))
		return
	}
	ctx.StatusCode(iris.StatusOK)
	_, code, msg := user.UpdateToken(ctx)
	if code == 400 {

		ctx.JSON(libs.ApiResource(401, nil, msg))
		return
	}
	ctx.JSON(libs.ApiResource(libs.CodeSuc, user, "操作成功"))
	// refreshToken := ctx.URLParam("refresh_token")
	// s := &models.Search{
	// 	Fields: []*models.Field{
	// 		{
	// 			Key:       "refresh_token",
	// 			Condition: "=",
	// 			Value:     refreshToken,
	// 		},
	// 	},
	// }
	// user, err := models.GetUser(s)
	// if err != nil {
	// 	ctx.JSON(libs.ApiResource(libs.CodeErr, nil, "token更新に失敗しました"))
	// 	return
	// }
	ctx.StatusCode(iris.StatusOK)
	// response, code, msg := user.UpdateToken(ctx)
	// ctx.JSON(libs.ApiResource(code, response, msg))
	// return
}

func GenTokenForTest(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	client := new(models.Clients)
	if err := ctx.ReadJSON(&client); err != nil {
		ctx.JSON(libs.ApiResource(400, nil, err.Error()))
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

	user, err := models.GetUser(s)
	if err != nil {
		ctx.JSON(libs.ApiResource(400, nil, err.Error()))
		return
	}
	tokenPair := user.GenerateTokenPair(ctx)
	// response, code, msg := user.CheckLogin(ctx, client.Password)
	ctx.JSON(libs.ApiResource(libs.CodeSuc, tokenPair, "操作成功"))
}
