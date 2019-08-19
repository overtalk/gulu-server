package jwt_test

import (
	"testing"

	"gitlab.com/SausageShoot/admin-server/utils/gin-middleware/jwt"
)

func TestJWTAuth(t *testing.T) {
	j := jwt.NewJWT()

	token, err := j.CreateToken(jwt.CustomClaims{
		ID:    "aa",
		Name:  "qinhan",
		Phone: "18888888888",
	})

	if err != nil {
		t.Error("CreateToken Error, err = ", err)
		return
	}

	t.Log("token = ", token)

	CustomClaims, err := j.ParseToken(token)

	if err != nil {
		t.Error("ParseToken Error, err = ", err)
		return
	}

	t.Log(CustomClaims)
}
