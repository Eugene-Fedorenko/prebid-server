package cpmstar

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

//NewCpmstarSyncer :
func NewCpmstarSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("cpmstar", 0, temp, adapters.SyncTypeRedirect)
}
