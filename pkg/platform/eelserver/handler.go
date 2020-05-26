package eelserver

import (
	"context"

	"github.com/jfbramlett/go-aop/pkg/rest"
	"github.com/jfbramlett/nwp-platform-go/pkg/eelmodel"

	"github.com/jfbramlett/go-aop/pkg/aop"
)

func NewEELHandler() *EELHandler {
	return &EELHandler{}
}

type EELHandler struct {
}

func (e *EELHandler) GetAccountList(ctx context.Context, req *eelmodel.AccountListRequest) (resp *eelmodel.AccountListResponse, err error) {
	aspectCtx := aop.Before(ctx)
	defer func() { aop.After(aspectCtx, err) }()

	resp = &eelmodel.AccountListResponse{}
	err = rest.GetRequest(aspectCtx, "http://localhost:8090/accountlist", &resp)
	return resp, err
}
