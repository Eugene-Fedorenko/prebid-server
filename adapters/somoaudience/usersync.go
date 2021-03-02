package somoaudience

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewSomoaudienceSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("somoaudience", 341, temp, adapters.SyncTypeRedirect)
}
