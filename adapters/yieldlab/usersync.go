package yieldlab

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewYieldlabSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("yieldlab", 70, temp, adapters.SyncTypeRedirect)
}
