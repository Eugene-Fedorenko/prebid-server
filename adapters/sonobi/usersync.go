package sonobi

import (
	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
	"text/template"
)

func NewSonobiSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("sonobi", 104, temp, adapters.SyncTypeRedirect)
}
