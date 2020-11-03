package openrtb_ext

type ExtImpEmxDigital struct {
	ExtImpBase
	TagID    string `json:"tagid"`
	BidFloor string `json:"bidfloor,omitempty"`
}
