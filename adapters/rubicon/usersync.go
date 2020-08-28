package rubicon

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewRubiconSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("rubicon", 52, temp, adapters.SyncTypeRedirect)
}
