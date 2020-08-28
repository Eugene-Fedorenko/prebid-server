package advangelists

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAdvangelistsSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("advangelists", 0, temp, adapters.SyncTypeIframe)
}
