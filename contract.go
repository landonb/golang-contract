// File: contract.go
// Author: Landon Bouma (landonb &#x40; retrosoft &#x2E; com)
// Project Page: https://github.com/landonb/golang_contract
// Summary: A design by contract `assert` mechanism for devs.
// License: Apache 2.0 (See file: LICENSE).

package contract

import (
	"fmt"
	"log"
	"os"
	"runtime"

	// 2018-05-03: Whelp, so much for not having any dependencies.
	"github.com/fatih/color"
)

var LOG = log.New(os.Stdout, "[Contract] ", log.Ldate|log.Lmicroseconds|log.LUTC)

type ContractOptions struct {
	Color bool
	Split bool
	Loggr func(format string, v ...interface{})
}

var contractOpts = &ContractOptions{
	Color: false,
	Split: false,
	Loggr: LOG.Printf,
}

func Contract(condition bool, args ...interface{}) {
	if !condition {
		// Send 1 to Caller, not 0, to get caller's info, not this fcn's.
		// Gets pc (program counter addy), file name, line number, and ok.
		pc, file, lnum, _ := runtime.Caller(1)
		fcn := runtime.FuncForPC(pc).Name()
		line := fmt.Sprintf("%d", lnum)
		if contractOpts.Color {
			colFcn := color.New(color.FgHiBlue).Add(color.Underline).Add(color.Bold).SprintFunc()
			colFnL := color.New(color.FgHiGreen).Add(color.Underline).Add(color.Bold).SprintFunc()
			fcn = colFcn(fcn)
			file = colFnL(file)
			line = colFnL(line)
		}
		callerSays := ""
		if len(args) > 0 {
			callerSays = fmt.Sprintf(": %s", fmt.Sprintf(args[0].(string), args[1:]...))
		}
		firstPart := fmt.Sprintf("Contract failure in %s", fcn)
		if !contractOpts.Split {
			contractOpts.Loggr("%s [%s:%s]%s", firstPart, file, line, callerSays)
		} else {
			contractOpts.Loggr("%s", firstPart)
			contractOpts.Loggr("at %s:%s", file, line)
			if callerSays != "" {
				contractOpts.Loggr("Note%s", callerSays)
			}
		}
		// Golang/Delve does not let code force a break, like JS's debugger,
		// Ruby's byebug, or Python's pdb.set_trace. But you can tell dlv to
		// run commands when it loads, e.g., to tell it to break herein. E.g.,
		//
		//    echo "break ${HOME}/.gopath/src/github.com/landonb/golang-contract/contract.go:69" > bps.dlv
		//    dlv exec --init="bps.dlv" -- foo bar --baz --bat
		//
		_ = "BREAK HERE"
		if false {
			_ = "All this just for a break!"
		}
	}
}

func SetColor(enable bool) {
	contractOpts.Color = true
}

func SetSplit(enable bool) {
	contractOpts.Split = true
}

// Let app specify an alternative logger func.
func SetLogger(logFcn func(format string, v ...interface{})) {
	contractOpts.Loggr = logFcn
}

