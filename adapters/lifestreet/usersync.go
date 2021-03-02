package lifestreet

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewLifestreetSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("lifestreet", 67, temp, adapters.SyncTypeRedirect)
}
