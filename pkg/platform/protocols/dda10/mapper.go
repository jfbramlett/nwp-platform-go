package dda10

import (
	"context"

	"github.com/jfbramlett/go-aop/pkg/aop"
	"github.com/jfbramlett/nwp-platform-go/pkg/eelmodel"
)

func AccountListRequestToEEL(ctx context.Context, req *DDA10AccountListRequest) (eelReq *eelmodel.AccountListRequest, err error) {
	aopCtx := aop.Before(ctx)
	defer func() { aop.After(aopCtx, err) }()

	return eelmodel.NewAccountListRequest(req.CustomerId, req.FID), nil
}

func AccountListResponseFromEEL(ctx context.Context, eelResp *eelmodel.AccountListResponse) (resp *DDA10AccountListResponse, err error) {
	aopCtx := aop.Before(ctx)
	defer func() { aop.After(aopCtx, err) }()

	accounts := make([]*DDA10AccountResponse, 0)
	for _, a := range eelResp.Accounts {
		accounts = append(accounts, NewDDA10AccountResponse(a.AccountID, a.AccountNumber, a.AccountName, a.Balance, a.AccountType))
	}
	return NewDDA10AccountListResponse(accounts), nil
}
