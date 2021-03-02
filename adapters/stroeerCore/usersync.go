package stroeerCore

import (
	"text/template"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/usersync"
)

func NewStroeerCoreSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("stroeerCore", 136, temp, adapters.SyncTypeIframe)
}
