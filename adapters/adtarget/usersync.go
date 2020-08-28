package adtarget

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAdtargetSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("adtarget", 0, temp, adapters.SyncTypeRedirect)
}
