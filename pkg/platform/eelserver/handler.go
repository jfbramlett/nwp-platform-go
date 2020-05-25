package eelserver

import (
	"context"
	"encoding/json"
	"net/http"

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

	htmlResp, err := http.Get("http://localhost:8090/accountlist")
	if err != nil {
		return nil, err
	}
	defer htmlResp.Body.Close()
	resp = &eelmodel.AccountListResponse{}
	decoder := json.NewDecoder(htmlResp.Body)
	err = decoder.Decode(&resp)

	return resp, err
}
