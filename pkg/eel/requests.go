package eel

type AccountListRequest struct {
	CustomerId 	string	`json:"customerId"`
	FID 		string	`json:"fid,omitempty"`
}

func NewAccountListRequest(customerId string, fid string) *AccountListRequest {
	return &AccountListRequest{CustomerId:  customerId, FID:  fid}
}


type AccountResponse struct {
	AccountID		string		`json:"accountId"`
	AccountNumber	string		`json:"accountNumber"`
	AccountName		string		`json:"accountName"`
	Balance			float64		`json:"balance"`
	AccountType		string		`json:"accountType"`
}

func NewAccountResponse(accountId string, accountNumber string, accountName string, balance float64, accountType string) *AccountResponse {
	return &AccountResponse{AccountID: accountId,
		AccountNumber: accountNumber,
		AccountName:  accountName,
		Balance: balance,
		AccountType: accountType,
	}
}

type AccountListResponse struct {
	Accounts []*AccountResponse 	`json:"accounts"`
}

func NewAccountListResponse(accounts []*AccountResponse) *AccountListResponse{
	return &AccountListResponse{Accounts: accounts}
}