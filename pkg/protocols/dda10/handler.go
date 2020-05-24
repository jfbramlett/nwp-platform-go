package dda10

import (
	"context"

	"github.com/jfbramlett/go-aop/pkg/aop"
	"github.com/jfbramlett/nwp-platform-go/pkg/eel"
)

func NewDDA10Handler(eelHandler *eel.EELHandler) *DDA10Handler {
	return &DDA10Handler{eelHandler: eelHandler}
}

type DDA10Handler struct {
	eelHandler *eel.EELHandler
}

func (d *DDA10Handler) GetAccountList(ctx context.Context, request *DDA10AccountListRequest) (resp *DDA10AccountListResponse, err error) {
	aopCtx := aop.Before(ctx)
	defer func() { aop.After(aopCtx, err) }()

	eelReq, err := AccountListRequestToEEL(aopCtx, request)
	if err != nil {
		return nil, err
	}

	eelResp, err := d.eelHandler.GetAccountList(aopCtx, eelReq)
	if err != nil {
		return nil, err
	}

	return AccountListResponseFromEEL(aopCtx, eelResp)
}
