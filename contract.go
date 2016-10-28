// File: contract.go
// Author: Landon Bouma (landonb &#x40; retrosoft &#x2E; com)
// Last Modified: 2016.10.28
// Project Page: https://github.com/landonb/golang_contract
// Summary: A design by contract `assert` mechanism for devs.
// License: GPLv3

package contract

import (
	"log"
)

func contract(condition) {
	if condition {
		LOG.Printf("contract failure: %+v", condition)
// FIXME: Introspect and print file/line of cbdfail.
// FIXME: dlv connect --init="set_breakpoints.dlv" localhost:3001
//          where set_breakpoints.dlv is
//          break contract.go:16
	}
}

func setLogger() {
	// FIXME: Implement.
}

