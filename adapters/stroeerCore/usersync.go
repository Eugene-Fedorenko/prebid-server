package stroeerCore

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewStroeerCoreSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("stroeerCore", 136, temp, adapters.SyncTypeIframe)
}
