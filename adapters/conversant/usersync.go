package conversant

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewConversantSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("conversant", 24, temp, adapters.SyncTypeRedirect)
}
