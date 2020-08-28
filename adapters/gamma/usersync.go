package gamma

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewGammaSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("gamma", 0, temp, adapters.SyncTypeIframe)
}
