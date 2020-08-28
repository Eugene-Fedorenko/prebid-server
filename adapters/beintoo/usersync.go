package beintoo

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewBeintooSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("Beintoo", 618, temp, adapters.SyncTypeIframe)
}
