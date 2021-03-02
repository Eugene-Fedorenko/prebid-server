package avocet

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAvocetSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("avocet", 63, temp, adapters.SyncTypeRedirect)
}
