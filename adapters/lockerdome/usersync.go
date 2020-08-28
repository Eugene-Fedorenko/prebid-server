package lockerdome

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewLockerDomeSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("lockerdome", 0, temp, adapters.SyncTypeRedirect)
}
