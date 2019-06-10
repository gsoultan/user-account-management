package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type SignUpRequest struct {
	firstName string `json:"first_name"`
	lastName  string `json:"last_name""`
	company   string `json:"company"`
	email     string `json:"email"`
	mobile    int    `json:"mobile"`
}

type SignUpResponse struct {
	ID  int64 `json:"id"`
	Err error `json:"error, omitempty"`
}

func makeSignUpEndpoint(c CommandService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SignUpRequest)

		a := Account{}
		a.FirstName = req.firstName
		a.LastName = req.lastName
		a.Company = req.company
		a.Email = req.email
		a.Mobile = req.mobile

		err = c.SignUp(&a)
		return a.ID, err
	}
}

type UpdateRequest struct {
	ID        int64  `json:"id"`
	firstName string `json:"first_name"`
	lastName  string `json:"last_name""`
	company   string `json:"company"`
	email     string `json:"email"`
	mobile    int    `json:"mobile"`
}
