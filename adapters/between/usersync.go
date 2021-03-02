package between

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

// NewBetweenSyncer returns "between" syncer
func NewBetweenSyncer(template *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("between", 724, template, adapters.SyncTypeRedirect)
}
