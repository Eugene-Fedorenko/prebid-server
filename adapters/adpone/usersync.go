package adpone

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

const adponeGDPRVendorID = uint16(16)
const adponeFamilyName = "adpone"

func NewadponeSyncer(urlTemplate *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer(
		adponeFamilyName,
		adponeGDPRVendorID,
		urlTemplate,
		adapters.SyncTypeRedirect,
	)
}
