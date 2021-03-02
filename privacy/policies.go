package privacy

import (
	"github.com/eugene-fedorenko/prebid-server/privacy/ccpa"
	"github.com/eugene-fedorenko/prebid-server/privacy/gdpr"
	"github.com/eugene-fedorenko/prebid-server/privacy/lmt"
)

// Policies represents the privacy regulations for an OpenRTB bid request.
type Policies struct {
	CCPA ccpa.Policy
	GDPR gdpr.Policy
	LMT  lmt.Policy
}
