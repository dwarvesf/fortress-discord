package model

type AdapterBank struct {
	Data []Bank `json:"data"`
}

type Bank struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Bin       string `json:"bin"`
	ShortName string `json:"shortName"`
	Logo      string `json:"logo"`
	SwiftCode string `json:"swiftCode"`
}
