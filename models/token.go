package models

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

const (
	accessTokenMaxAge  = 7 * 24 * time.Hour
	refreshTokenMaxAge = 14 * 24 * time.Hour
	leewayAge          = 8 * 24 * time.Hour
)

var (
	privateKey, publicKey = jwt.MustLoadRSA("key/rsa_private_key_pkcs8.pem", "key/rsa_public_key.pem")

	signer   = jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)
	verifier = jwt.NewVerifier(jwt.RS256, publicKey)
)

type UserClaims struct {
	ID string `json:"user_id"`
	// Do: `json:"username,required"` to have this field required
	// or see the Validate method below instead.
	Mail string `json:"mail"`
	exp  int64
	iat  int64
}

//generate token to user and store token
func (u *Clients) GenerateTokenPair(ctx iris.Context) *jwt.TokenPair {
	if u.ID == 0 {
		return nil
	} else {
		var tempID string
		fmt.Println("generate token pair")
		tempID = strconv.Itoa(int(u.ID))
		refreshClaims := jwt.Claims{Subject: tempID}

		accessClaims := UserClaims{
			ID:   tempID,
			Mail: u.Mail,
			exp:  time.Now().UTC().Add(time.Hour * time.Duration(1)).Unix(),
			iat:  time.Now().UTC().Unix(),
		}
		tokenPair, err := signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
		if err != nil {
			ctx.Application().Logger().Errorf("token pair: %v", err)
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return nil
		}
		u.AccessToken, _ = strconv.Unquote(string(tokenPair.AccessToken))
		u.RefreshToken, _ = strconv.Unquote(string(tokenPair.RefreshToken))

		ups := map[string]interface{}{"access_token": u.AccessToken, "refresh_token": u.RefreshToken}
		if err := Update(&Clients{}, ups, u.ID); err != nil {
			log.Println(fmt.Sprintf("Token保存失敗：%+v\n", u.ID))
			log.Println(fmt.Sprintf("Token保存失敗：%+v\n", err))
		}
		return &tokenPair
	}
}

func (u *Clients) refreshToken(ctx iris.Context) *jwt.TokenPair {
	var tempID string
	tempID = strconv.Itoa(int(u.ID))

	refreshToken := []byte(ctx.URLParam("refresh_token"))
	if len(refreshToken) == 0 {
		var tokenPair jwt.TokenPair
		if err := ctx.ReadJSON(&tokenPair); err != nil {
			ctx.StopWithError(iris.StatusBadRequest, err)
			return nil
		}
		refreshToken = tokenPair.RefreshToken
	}

	_, err := verifier.VerifyToken(refreshToken, jwt.Expected{Subject: tempID})
	if err != nil {
		ctx.Application().Logger().Errorf("verify refresh token: %v", err)
		ctx.StatusCode(iris.StatusUnauthorized)
		return nil
	}
	return u.GenerateTokenPair(ctx)

}

// func (u *Clients) generateAccessToken(ctx iris.Context) []byte {
// 	if u.ID == 0 {
// 		return nil
// 	} else {
// 		var tempID string
// 		tempID = strconv.Itoa(int(u.ID))

// 		accessClaims := UserClaims{
// 			ID:   tempID,
// 			Name: u.FirstName,
// 			exp:  time.Now().UTC().Add(time.Hour * time.Duration(1)).Unix(),
// 			iat:  time.Now().UTC().Unix(),
// 		}
// 		accessToken, err := signer.Sign(accessClaims)
// 		if err != nil {
// 			ctx.Application().Logger().Errorf("access token: %v", err)
// 			ctx.StopWithStatus(iris.StatusInternalServerError)
// 			return nil
// 		}
// 		u.AccessToken = string(accessToken)
// 		ups := map[string]interface{}{"access_token": u.AccessToken}
// 		if err := Update(&Clients{}, ups, u.ID); err != nil {
// 			log.Println(fmt.Sprintf("Token保存失敗：%+v\n", u.ID))
// 			log.Println(fmt.Sprintf("Token保存失敗：%+v\n", err))
// 		}
// 		return BytesQuote(accessToken)
// 	}
// }

// BytesQuote returns a double-quoted []byte slice representing "b".
func BytesQuote(b []byte) []byte {
	dst := make([]byte, len(b)+2)
	dst[0] = '"'
	copy(dst[1:], b)
	dst[len(dst)-1] = '"'
	return dst
}

func (u *Doctors) GenerateTokenPair(ctx iris.Context) *jwt.TokenPair {
	if u.ID == 0 {
		return nil
	} else {
		var tempID string
		fmt.Println("generate token pair")
		tempID = strconv.Itoa(int(u.ID))
		refreshClaims := jwt.Claims{Subject: tempID}

		accessClaims := UserClaims{
			ID:   tempID,
			Mail: u.Mail,
			exp:  time.Now().UTC().Add(time.Hour * time.Duration(1)).Unix(),
			iat:  time.Now().UTC().Unix(),
		}
		tokenPair, err := signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
		if err != nil {
			ctx.Application().Logger().Errorf("token pair: %v", err)
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return nil
		}
		u.AccessToken, _ = strconv.Unquote(string(tokenPair.AccessToken))
		u.RefreshToken, _ = strconv.Unquote(string(tokenPair.RefreshToken))

		ups := map[string]interface{}{"access_token": u.AccessToken, "refresh_token": u.RefreshToken}
		if err := Update(&Doctors{}, ups, u.ID); err != nil {
			log.Println(fmt.Sprintf("Token保存失敗：%+v\n", u.ID))
			log.Println(fmt.Sprintf("Token保存失敗：%+v\n", err))
		}
		return &tokenPair
	}
}

func (u *Doctors) refreshToken(ctx iris.Context) *jwt.TokenPair {
	var tempID string
	tempID = strconv.Itoa(int(u.ID))

	refreshToken := []byte(ctx.URLParam("refresh_token"))
	if len(refreshToken) == 0 {
		var tokenPair jwt.TokenPair
		if err := ctx.ReadJSON(&tokenPair); err != nil {
			ctx.StopWithError(iris.StatusBadRequest, err)
			return nil
		}
		refreshToken = tokenPair.RefreshToken
	}

	_, err := verifier.VerifyToken(refreshToken, jwt.Expected{Subject: tempID})
	if err != nil {
		ctx.Application().Logger().Errorf("verify refresh token: %v", err)
		ctx.StatusCode(iris.StatusUnauthorized)
		return nil
	}
	return u.GenerateTokenPair(ctx)

}
