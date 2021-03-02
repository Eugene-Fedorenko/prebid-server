package openrtb_ext

// ExtImpVerizonMedia defines the contract for bidrequest.imp[i].ext.verizonmedia
type ExtImpVerizonMedia struct {
	ExtImpBase
	Dcn string `json:"dcn"`
	Pos string `json:"pos"`
}
