package ninthdecimal

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewNinthDecimalSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("ninthdecimal", 0, temp, adapters.SyncTypeIframe)
}
