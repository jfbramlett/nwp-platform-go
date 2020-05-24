package eel

import (
	"context"

	"github.com/jfbramlett/go-aop/pkg/aop"
)

func NewEELHandler() *EELHandler {
	return &EELHandler{}
}

type EELHandler struct {
}

func (e *EELHandler) GetAccountList(ctx context.Context, req *AccountListRequest) (resp *AccountListResponse, err error) {
	aspectCtx := aop.Before(ctx)
	defer func() { aop.After(aspectCtx, err) }()

	return NewAccountListResponse([]*AccountResponse{NewAccountResponse("10",
		"123", "CHECKING", 55.7, "CHECK")}), nil
}
