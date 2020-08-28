package vrtcal

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewVrtcalSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("vrtcal", 0, temp, adapters.SyncTypeRedirect)
}
