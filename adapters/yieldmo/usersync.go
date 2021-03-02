package yieldmo

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewYieldmoSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("yieldmo", 173, temp, adapters.SyncTypeRedirect)
}
