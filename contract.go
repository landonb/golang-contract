// File: contract.go
// Author: Landon Bouma (landonb &#x40; retrosoft &#x2E; com)
// Last Modified: 2016.10.29
// Project Page: https://github.com/landonb/golang_contract
// Summary: A design by contract `assert` mechanism for devs.
// License: GPLv3

package contract

import (
	"log"
	"os"
)

var LOG = log.New(os.Stdout, "[Contract] ", log.Ldate|log.Lmicroseconds|log.LUTC)

func Contract(condition bool) {
	if !condition {
		LOG.Printf("Contract failure: %+v", condition)
		// FIXME: Introspect and print file/line of caller.
		// LATER: dlv connect --init="set_breakpoints.dlv" localhost:3001
		//            where set_breakpoints.dlv is
		//            break contract.go:19
		//        doesn't work. Perhaps one day we can make it work?
	}
}

// MAYBE: Implement SetLogger so app can specify an alternative logger.
//        E.g., maybe the rest of the app uses another logger,
//              like the juju logger (https://github.com/juju/loggo)
//func SetLogger() {
//}

