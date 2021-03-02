package connectad

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewConnectAdSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("connectad", 138, temp, adapters.SyncTypeIframe)
}
