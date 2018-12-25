package models

type Cargo struct {
	Product string `json:"product"`
	SrcCity string `json:"origin_city"`
	SrcState string `json:"origin_state"`
	SrcLat float64 `json:"origin_lat"`
	SrcLng float64 `json:"origin_lng"`
	DstCity string `json:"destination_city"`
	DstState string `json:"destination_state"`
	DstLat float64 `json:"destination_lat"`
	DstLng float64 `json:"destination_lng"`
}
