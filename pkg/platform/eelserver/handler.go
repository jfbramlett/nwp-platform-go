package eelserver

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jfbramlett/go-aop/pkg/tracing"
	"github.com/jfbramlett/go-aop/pkg/web"

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

	traceId := tracing.GetTraceFromContext(aspectCtx)

	httpReq, _ := http.NewRequest("GET", "http://localhost:8090/accountlist", nil)
	httpReq.Header.Set(web.HeaderRequestId, traceId)

	client := &http.Client{Timeout: time.Second * 10}

	htmlResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer htmlResp.Body.Close()
	resp = &eelmodel.AccountListResponse{}
	decoder := json.NewDecoder(htmlResp.Body)
	err = decoder.Decode(&resp)

	return resp, err
}
