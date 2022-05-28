package routes

import (
	controllers "MedicalQuestionnaireService/controllers/vue/common"
	_ "MedicalQuestionnaireService/docs"
	"MedicalQuestionnaireService/libs"
	"MedicalQuestionnaireService/middleware"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	privateKey, publicKey = jwt.MustLoadRSA("key/rsa_private_key_pkcs8.pem", "key/rsa_public_key.pem")
	verifier              = jwt.NewVerifier(jwt.RS256, publicKey)
)

func App(api *iris.Application) {
	//CORSクロスドメインで安全に通信
	api.UseRouter(middleware.CrsAuth())
	app := api.Party("/").AllowMethods(iris.MethodOptions)
	{
		// バイナリモード、プロジェクトエントリを有効にする
		if libs.Config.Bindata {
			app.Get("/", func(ctx iris.Context) { // 首页模块
				_ = ctx.View("index.html")
			})
		}

		// swagger配置
		app.Get("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

		vue := app.Party("/vue")
		{
			/****verify token ****/
			// verifyMiddleware := verifier.Verify(func() interface{} {
			// 	return new(models.UserClaims)
			// })
			// vue.Use(verifyMiddleware, middleware.VerifierTokenWithDb)

			vue.Post("/clients/create", controllers.CreateClient)
			// vue.Get("/clients/{client_id:uint}", controllers.ShowClient)
			// vue.Post("/clients/update", controllers.UpdateClient)
			// vue.Post("/clients/delete", controllers.DeleteClient)
			// vue.Post("/clients/updatePassword", controllers.UpdateClientPassword)
			vue.Post("/clients/login", controllers.ClientLogin)
			// vue.Post("/clients/logout", controllers.ClientLogout)

			vue.Get("/doctors/all", controllers.ShowAllDoctors)
			vue.Get("/doctors/{doctorId:uint}", controllers.ShowSelectedDoctor)
			vue.Post("/doctors/create", controllers.CreateDoctor)
			// vue.Post("/doctors/update", controllers.UpdateDoctor)
			// vue.Post("/doctors/delete", controllers.DeleteDoctor)
			// vue.Post("/doctors/updatePassword", controllers.UpdateDoctorPassword)
			vue.Post("/doctors/login", controllers.DoctorLogin)
			// vue.Post("/doctors/logout", controllers.DoctorLogout)

			// vue.Post("/forms", controllers.Createform)
			// vue.Get("/forms/{doctor_id:uint}", controllers.ShowDoctorForm)
			vue.Post("/forms/create", controllers.CreateQuestion)
			vue.Post("/forms/update", controllers.UpdateQuestion)
			vue.Post("/forms/delete", controllers.DeleteQuestion)
			vue.Post("/forms/save", controllers.SaveForm)

			vue.Get("/masterQuestions/all", controllers.ShowAllMasterQuestions)
			vue.Get("/masterQuestions/{master_question_id:uint}", controllers.ShowMasterQuestion)
			vue.Post("/masterQuestions/create", controllers.CreateMasterQuestion)
			vue.Post("/masterQuestions/update", controllers.UpdateMasterQuestion)
			vue.Post("/masterQuestions/delete", controllers.DeleteMasterQuestionById)

			vue.Get("/answers/{doctorId:uint}", controllers.ShowAnswerLists)
			vue.Get("/answers/{doctorId:uint}/{clientId:uint}", controllers.ShowClientAnswer)
			// vue.Post("/answers", controllers.CreateAnswer)

			vue.Post("/genToken", controllers.GenTokenForTest)
			vue.Get("/refreshToken", controllers.RefreshUserToken)

			/****verify token****/
			// vue.PartyFunc("/admin", func(admin iris.Party) {

			// 	/* casbin 使用*/
			// 	// casbinMiddleware := middleware.New(libs.Enforcer)       //casbin for gorm
			// 	// admin.Use(verifyMiddleware, casbinMiddleware.AuthToken) //登录验证
			// 	/* casbin 未使用*/
			// 	//
			// 	adminRoutes(admin)

			// })
			// vue.PartyFunc("/users", func(users iris.Party) {
			// 	usersRoutes(users)
			// })

			//クエリパラメータの部分
			// app.Get("/welcome", func(ctx iris.Context) {
			// 	firstname := ctx.URLParamDefault("firstname", "Guest")
			// 	lastname := ctx.URLParam("lastname")

			// 	ctx.Writef("Hello %s %s", firstname, lastname)
			// })
		}
	}
}
