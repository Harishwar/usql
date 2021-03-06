package ql

import (
	// DRIVER: ql
	_ "github.com/cznic/ql/driver"

	"github.com/xo/usql/drivers"
)

func init() {
	drivers.Register("ql", drivers.Driver{
		AllowMultilineComments: true,
		AllowCComments:         true,
	})
}
