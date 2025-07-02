package simplefin

type URLConfig struct {
	// may be able to get away with int32 here, not sure if the API
	// actually supports a finer resolution thant 32bit numbers --
	// which should be done to minute resolution
	startdate 			int64
	enddate				int64
	pending				bool			// defaults to false
	accounts			[]string
	balancesOnly		bool			// defaults to false
}

type Organization struct {
	Domain 				string `json:"Domain"`
	SfinURL 			string `json:"sfin-url"`
	Name 				string `json:"name"`
	URL 				string `json:"url"`
	ID 					string `json:"id"`
} 

type Transaction struct {
	ID					string `json:"id"`
	Posted				int `json:"posted"`
	Amount				string `json:"amount"`
	Description			string `json:"description"`
	TransactedAt		int `json:"transacted_at"`
	Pending				bool `json:"pending"`
	Extra				[]any `json:"extra"`		
}

type Holding struct {
	ID					string `json:"id"`
	Created				int `json:"created"`
	Currency			string `json:"currency"`
	CostBasis			string `json:"cost_basis"`
	Description			string `json:"description"`
	MarketValue			string `json:"market_value"`
	PurchasePrice		string `json:"purchase_price"`
	Shares				string `json:"shares"`
	Symbol				string `json:"symbol"`
}

type Account struct {
	Org 				Organization `json:"org"`
	ID 					string `json:"id"`
	Name 				string `json:"name"`
	Currency 			string `json:"currency"`
	Balance 			string `json:"balance"`
	AvailableBalance 	string `json:"available-balance"`
	BalanceDate 		int `json:"balance-date"`
	Transactions 		[]Transaction `json:"transactions"`
	Holdings			[]Holding `json:"holdings"`
	Extra				[]any `json:"extra"`
}

type Accounts struct {
	Errors 				[]string `json:"errors"`
	Accounts 			[]Account `json:"accounts"`
}

