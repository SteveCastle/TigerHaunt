package grifts

import (
	"github.com/SteveCastle/tigerhaunt/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
