package synacormedia

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewSynacorMediaSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("synacormedia", 0, temp, adapters.SyncTypeIframe)
}
