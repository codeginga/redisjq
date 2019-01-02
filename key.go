package redisjq

import (
	"github.com/codeginga/redisjq/cnst"
)

func keyQName(name string) string {
	return cnst.App + cnst.Separator + name
}

func keyMessage(id string) string {
	return cnst.App + cnst.Separator + id
}
