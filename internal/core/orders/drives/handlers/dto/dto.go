package dto

type (
	PriceRequest struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}

	AddressRequest struct {
		Country      string  `json:"country"`
		State        string  `json:"state"`
		Municipality string  `json:"municipality"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
	}

	// OrderRequest contains the body of http request
	OrderRequest struct {
		Id               string `json:"id"`
		CreateAt         string `json:"createAt"`
		Status           string `json:"status"`
		PriceRequest     `json:"price"`
		AddressRequest   `json:"address"`
		RequestedTime    string `json:"requestedTime"`
		IsProduct        bool   `json:"isProduct"`
		IsSubscription   bool   `json:"isSubscription"`
		TypeSubscription string `json:"typeSubs"`
		UserId           string `json:"userId"`
	}
)
