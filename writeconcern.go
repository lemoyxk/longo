/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 21:07
**/

package longo

import (
	"time"
)

// WriteConcern

type WriteConcern struct {
	W        int
	J        bool
	Wtimeout time.Duration
}
