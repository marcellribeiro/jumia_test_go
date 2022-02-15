package models

type Cargo struct {
	Cargo []Courier `json:"cargos"`
}

type Courier struct {
	OrderIds    []int   `json:"order_ids"`
	TotalWeight float32 `json:"total_weight"`
}
