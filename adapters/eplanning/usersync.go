package eplanning

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewEPlanningSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("eplanning", 90, temp, adapters.SyncTypeIframe)
}
