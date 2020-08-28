package adman

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

// NewAdmanSyncer returns adman syncer
func NewAdmanSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("adman", 149, temp, adapters.SyncTypeRedirect)
}
