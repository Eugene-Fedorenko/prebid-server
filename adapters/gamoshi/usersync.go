package gamoshi

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewGamoshiSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("gamoshi", 644, temp, adapters.SyncTypeRedirect)
}
