package sharethrough

import (
	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
	"text/template"
)

func NewSharethroughSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("sharethrough", 80, temp, adapters.SyncTypeRedirect)
}
