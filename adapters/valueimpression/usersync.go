package valueimpression

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewValueImpressionSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("valueimpression", 0, temp, adapters.SyncTypeRedirect)
}
