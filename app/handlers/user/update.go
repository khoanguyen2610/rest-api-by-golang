package user

import (
	"fmt"
	"user-service/handlers"
	"user-service/handlers/response"
)

type UserPutForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func Update(ctx handlers.Context) response.ApiResponse {
	payload := UserPutForm{}
	if err := ctx.DecodePayload(&payload); err != nil {
		return response.BadRequest(
			fmt.Errorf("cannot parse JSON to object: %s", err.Error()),
		)
	}

	return response.Ok(payload)
	//tx := ctx.DB().Begin()
	//userRepo := mysql.NewUserRepo(tx)
}
