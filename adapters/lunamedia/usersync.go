package lunamedia

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewLunaMediaSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("lunamedia", 0, temp, adapters.SyncTypeIframe)
}
