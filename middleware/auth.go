package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/BigListRyRy/harbourlivingapi/token"
	"github.com/BigListRyRy/harbourlivingapi/util"
)

var tokenMaker token.TokenService

func init() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot find config ", err)
	}
	tokenMaker, err = token.NewJWTService(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalln("cannot create a token maker ", err)
	}

}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := tokenMaker.VerifyToken(auth)
		if err != nil {
			fmt.Printf("Could not authenticate token ....%v", err)
			http.Error(w, "Could not authenticate token", http.StatusForbidden)
			return
		}
		if validate.Valid() != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		userinfo := token.UserInfo{
			UserID:   validate.UserID,
			Username: validate.Username,
			Email:    validate.Email,
			UserType: validate.UserType,
		}

		ctx := context.WithValue(r.Context(), "userinfo", userinfo)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) token.UserInfo {
	raw, _ := ctx.Value("userinfo").(token.UserInfo)
	return raw
}
