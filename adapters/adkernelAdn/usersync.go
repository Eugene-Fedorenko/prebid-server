package adkernelAdn

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

const adkernelGDPRVendorID = uint16(14)

func NewAdkernelAdnSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("adkernelAdn", 14, temp, adapters.SyncTypeRedirect)
}
