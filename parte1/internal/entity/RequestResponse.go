package entity

type ResponseBuy struct {
	Total     float32        `json:"total,omitempty"`
	BuyByTdc  map[string]int `json:"comprasPorTDC,omitempty"`
	NotBuy    int            `json:"nocompraron,omitempty"`
	BuyHigher float32        `json:"compraMasAlta,omitempty"`
}

type RequestBuy struct {
	Compro bool    `json:"compro"`
	Tdc    string  `json:"tdc"`
	Monto  float32 `json:"monto"`
}
