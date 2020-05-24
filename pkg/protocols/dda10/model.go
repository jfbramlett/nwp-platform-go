package dda10

type DDA10AccountListRequest struct {
	CustomerId 	string	`json:"customerId"`
	FID 		string	`json:"fid,omitempty"`
}

func NewDDA10AccountListRequest(customerId string, fid string) *DDA10AccountListRequest {
	return &DDA10AccountListRequest{CustomerId:  customerId, FID:  fid}
}


type DDA10AccountResponse struct {
	AccountID		string		`json:"accountId"`
	AccountNumber	string		`json:"accountNumber"`
	AccountName		string		`json:"accountName"`
	Balance			float64		`json:"balance"`
	AccountType		string		`json:"accountType"`
}

func NewDDA10AccountResponse(accountId string, accountNumber string, accountName string, balance float64, accountType string) *DDA10AccountResponse {
	return &DDA10AccountResponse{AccountID: accountId,
		AccountNumber: accountNumber,
		AccountName:  accountName,
		Balance: balance,
		AccountType: accountType,
	}
}

type DDA10AccountListResponse struct {
	Accounts []*DDA10AccountResponse 	`json:"accounts"`
}

func NewDDA10AccountListResponse(accounts []*DDA10AccountResponse) *DDA10AccountListResponse{
	return &DDA10AccountListResponse{Accounts: accounts}
}
