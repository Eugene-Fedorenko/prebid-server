package beachfront

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

var VENDOR_ID uint16 = 335

func NewBeachfrontSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer(
		"beachfront",
		VENDOR_ID,
		temp,
		adapters.SyncTypeIframe)
}
