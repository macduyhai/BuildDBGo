package dtos

type Meta struct {
	Code    int    `json:code`
	Message string `json:message`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
type AddUserRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Fullname     string `json:"fullname"`
	Phone_number string `json:"phone_number"`
}

//------------ THAM KHAO ------------------
// type Item struct {
// 	OrderNumber      string     `json:"orderNumber"`
// 	SkuNumber        string     `json:"skuNumber"`
// 	PackagedQuantity int32      `json:"packagedQuantity"`
// 	ItemDescription  string     `json:"itemDescription"`
// 	CreatedAt        *time.Time `json:"created_at"`
// }
// type SearchItemRequest struct {
// 	orderNumber string
// }
// type SearchItemResponse struct {
// 	Items []Item `json:"items"`
// }
// type AddOrderRequest struct {
// 	Orders []Order `json:"orders"`
// }
// type OrderNew struct {
// 	OrderNumber string `json:"orderNumber"`
// 	Name        string `json:"name"`
// 	Note        string `json:"note"`
// 	Address1    string `json:"address1"`
// 	Address2    string `json:"address2"`
// 	City        string `json:"city"`
// 	State       string `json:"state"`
// 	PostalCode  string `json:"postalCode"`
// 	Country     string `json:"country"`
// 	Phone       string `json:"phone"`
// 	BranchSell  string `json:"branchsell"`
// 	Seller      string `json:"seller"`
// }

// // type AddOrderNewRequest struct {
// // 	Orders []OrderNew `json:"orders"`
// // }
// type NumberOrderRequest struct {
// 	Steptime string `json:"steptime"`
// 	Value    string `json:"value"`
// }

// type NumberOrderResponse struct {
// 	Steptime     string             `json:"steptime"`
// 	Orders       []NumberOrderInfor `json:"orders"`
// 	BranchSells  []NumberOrderInfor `json:"branchsells"`
// 	TypeProducts []NumberOrderInfor `json:"typeproducts"`
// 	Sellers      []NumberOrderInfor `json:"sellers"`
// }
// type NumberOrderInfor struct {
// 	Key   string `json:"key"`
// 	Value int64  `json:"value"`
// }
// type AddorderResponse struct {
// 	ID             int64    `json:"id,omitempty"`
// 	RecordsSuccess []string `json:"recordsSuccess,omitempty"`
// 	RecordsFailes  []string `json:"recordsFailes,omitempty"`
// }
// type PrintersRequest struct {
// 	OrderNumber []string `json:"orderNumber"`
// 	PrintStatus int32    `json:"printstatus"`
// }
// type PrintersResponse struct {
// 	PrintsOrder    []NumberOrderInfor `json:"printsorder"`
// 	RecordsSuccess []string           `json:"recordsSuccess,omitempty"`
// 	RecordsFailes  []string           `json:"recordsFailes,omitempty"`
// }
// type LableDetails struct {
// 	TrackingNumber        string `json:"trackingNumber"`
// 	URL                   string `json:"url"`
// 	PartnerTrackingNumber string `json:"partnerTrackingNumber"`
// }

// type AddLabelRequest struct {
// 	OrderNumber  string       `json:"orderNumber"`
// 	Items        []Item       `json:"items"`
// 	LableDetails LableDetails `json:"labelDetails"`
// }

// type ShippingInfor struct {
// 	Status        int32  `json:"status"`
// 	BeginShipping string `json:"beginShipping,omitempty"`
// 	TimeCompleted string `json:"timeCompleted,omitempty"`
// }
// type AddShippingTimeRequest struct {
// 	OrderNumber       string     `json:"orderNumber"`
// 	BeginShipping     string     `json:"beginShipping"`
// 	TimeCompleted     string     `json:"timeCompleted"`
// 	BeginShippingReal *time.Time `json:"-"`
// 	TimeCompletedReal *time.Time `json:"-"`
// }

// type FullOrderInformation struct {
// 	Order
// 	ShippingInfor
// 	LableDetails LableDetails `json:"lableDetails"`
// }

// type SearchQuery struct {
// 	Key   string
// 	Value interface{}
// }
// type SearchItemsQuery struct {
// 	Key   string
// 	Value interface{}
// }

// type NumberOrderQuery struct {
// 	Key   string
// 	Value interface{}
// }

// type ChangeStatusToCompleted struct {
// 	OrderNumber string `json:"orderNumber"`
// }
// type ChangeStatusToDelay struct {
// 	OrderNumber string `json:"orderNumber"`
// }
// type CheckRequest struct {
// 	PartnerTrackingNumber string `json:"partner_tracking_number"`
// }
// type CheckResponse struct {
// 	Order        OrderNew
// 	LableDetails LableDetails `json:"lableDetails"`
// 	Items        []Item       `json:"items"`
// }

// // TICH HOP TOOL
// type ItemNew struct {
// 	SkuNumber        string `json:"skuNumber"`
// 	PackagedQuantity int32  `json:"packagedQuantity"`
// 	ItemDescription  string `json:"itemDescription"`
// }
// type OrderFull struct {
// 	OrderNumber string    `json:"orderNumber"`
// 	Name        string    `json:"name"`
// 	Note        string    `json:"note"`
// 	Address1    string    `json:"address1"`
// 	Address2    string    `json:"address2"`
// 	City        string    `json:"city"`
// 	State       string    `json:"state"`
// 	PostalCode  string    `json:"postalCode"`
// 	Country     string    `json:"country"`
// 	Phone       string    `json:"phone"`
// 	BranchSell  string    `json:"branchsell"`
// 	Seller      string    `json:"seller"`
// 	Items       []ItemNew `json:"items"`
// }
// type AddfullOrderRequest struct {
// 	Orders []OrderFull `json:"orders"`
// }
