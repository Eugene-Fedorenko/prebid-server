package openrtb_ext

// ExtImpLockerDome defines the contract for bidrequest.imp[i].ext.lockerdome
type ExtImpLockerDome struct {
	ExtImpBase
	// LockerDome params
	AdUnitId string `json:"adUnitId"`
}
