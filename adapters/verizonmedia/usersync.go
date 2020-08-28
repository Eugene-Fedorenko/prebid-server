package verizonmedia

import (
	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
	"text/template"
)

func NewVerizonMediaSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("verizonmedia", 25, temp, adapters.SyncTypeRedirect)
}
