package ix

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewIxSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("ix", 10, temp, adapters.SyncTypeRedirect)
}
