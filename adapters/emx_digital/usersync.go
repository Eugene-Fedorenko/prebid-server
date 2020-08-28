package emx_digital

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewEMXDigitalSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("emx_digital", 183, temp, adapters.SyncTypeIframe)
}
