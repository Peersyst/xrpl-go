package types

type IssuedCurrency struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"`
}

func (i *IssuedCurrency) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})
	flattened["currency"] = i.Currency
	flattened["issuer"] = i.Issuer
	return flattened
}
