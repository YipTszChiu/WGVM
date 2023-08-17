package main

import (
	"WGVM/src/common"
	"WGVM/src/instance"
	"WGVM/src/interpreter"
	"flag"
	"fmt"
	"os"
)

func main() {
	dumpFlag := flag.Bool("d", false, "dump")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: wasmgo [-d] filename")
		os.Exit(1)
	}

	module, err := common.DecodeFile(flag.Args()[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if *dumpFlag {
		dump(module)
	} else {
		instantiateAndExecMainFunc(module)
	}
}

func instantiateAndExecMainFunc(module common.Module) {
	mm := map[string]instance.Module{"env": newEnv()}
	m, err := interpreter.New(module, mm)
	if err == nil {
		_, err = m.InvokeFunc("main")
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
