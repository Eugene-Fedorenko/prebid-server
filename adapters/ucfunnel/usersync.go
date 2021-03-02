package ucfunnel

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewUcfunnelSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("ucfunnel", 607, temp, adapters.SyncTypeRedirect)
}
