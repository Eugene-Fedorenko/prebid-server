package smartrtb

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewSmartRTBSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("smartrtb", 0, temp, adapters.SyncTypeRedirect)
}
