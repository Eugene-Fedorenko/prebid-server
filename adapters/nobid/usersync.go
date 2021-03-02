package nobid

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewNoBidSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("nobid", 816, temp, adapters.SyncTypeRedirect)
}
