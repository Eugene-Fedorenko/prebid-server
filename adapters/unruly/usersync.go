package unruly

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewUnrulySyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("unruly", 162, temp, adapters.SyncTypeIframe)
}
