package types

type Quote struct {
	Data Data `json:"USDBRL"`
}

type Data struct {
	Bid string `json:"bid"`
}
