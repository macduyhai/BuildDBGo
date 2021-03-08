package dtos

type TypeProduct struct {
	Name   string `json:"name"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Weight int32  `json:"weight"`
	Length int32  `json:"length"`
	Note   string `json:"note"`
}

type AddtypeRequest struct {
	TypeProducts []TypeProduct `json:"typeproducts"`
}

type AddtypeResponse struct {
	ID             int64    `json:"id,omitempty"`
	RecordsSuccess []string `json:"recordsSuccess,omitempty"`
	RecordsFailes  []string `json:"recordsFailes,omitempty"`
}
type DeleteTypeRequest struct {
	Name string `json:"name"`
}

type SearchTypeProductQuery struct {
	Key   string
	Value interface{}
}
