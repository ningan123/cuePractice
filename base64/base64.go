package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueerror "cuelang.org/go/cue/errors"
)

func main() {

	ctx := cuecontext.New()

	clusterVal := ctx.NewList()

	confPath := "/root/cuePractice/base64"
	fileData, err := os.ReadFile(filepath.Join(confPath, "base64.cue"))
	if err != nil {
		fmt.Println(err.Error())
	}

	scheme := ctx.CompileBytes(fileData, cue.Scope(clusterVal))
	if scheme.Err() != nil {
		msg := cueerror.Details(scheme.Err(), nil)
		fmt.Println(errors.New(msg))
	}
	fmt.Printf("scheme: \n%s\n", scheme)
}
