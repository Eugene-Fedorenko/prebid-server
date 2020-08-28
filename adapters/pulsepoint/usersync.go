package pulsepoint

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewPulsepointSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("pulsepoint", 81, temp, adapters.SyncTypeRedirect)
}
