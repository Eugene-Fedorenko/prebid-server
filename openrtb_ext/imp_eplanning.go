package openrtb_ext

// ExtImpEPlanning defines the contract for bidrequest.imp[i].ext.eplanning
type ExtImpEPlanning struct {
	ExtImpBase
	ClientID   string `json:"ci"`
	AdUnitCode string `json:"adunit_code"`
	SizeString string
}
