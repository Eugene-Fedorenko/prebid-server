package colossus

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

// NewColossusSyncer returns colossus syncer
func NewColossusSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("colossus", 0, temp, adapters.SyncTypeRedirect)
}
