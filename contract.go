// File: contract.go
// Author: Landon Bouma (landonb &#x40; retrosoft &#x2E; com)
// Last Modified: 2017.03.13
// Project Page: https://github.com/landonb/golang_contract
// Summary: A design by contract `assert` mechanism for devs.
// License: Apache 2.0 (See file: LICENSE).

package contract

import (
	"log"
	"os"
	"runtime"
)

var LOG = log.New(os.Stdout, "[Contract] ", log.Ldate|log.Lmicroseconds|log.LUTC)

func Contract(condition bool) {
	if !condition {
		// Send 1 to Caller, not 0, to get caller's info, not this fcn's.
		pc, fn, line, _ := runtime.Caller(1)
		fcn := runtime.FuncForPC(pc).Name()
		LOG.Printf("Contract failure in %s[%s:%d]: %+v", fcn, fn, line, condition)
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

