package logicad

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewLogicadSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("logicad", 0, temp, adapters.SyncTypeRedirect)
}
