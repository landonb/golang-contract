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
		// Golang does not let code force a break, like JavaScript's debugger,
		// or Ruby byebug's byebug, or Python's pdb.set_trace(), etc. However,
		// you can tell dlv to run commands when it loads, e.g., tell it to break
		// herein. E.g.,
		//
		//    echo "break ${HOME}/.gopath/src/github.com/landonb/golang-contract/contract.go:23" > bps.dlv
		//    dlv exec --init="bps.dlv" -- foo bar --baz --bat
	}
}

// MAYBE: Implement SetLogger so app can specify an alternative logger.
//        E.g., maybe the rest of the app uses another logger,
//              like the juju logger (https://github.com/juju/loggo)
//func SetLogger() {
//}

