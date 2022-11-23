//go:build tools
// +build tools

package main

import (
	_ "github.com/cespare/reflex"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/securego/gosec/v2/cmd/gosec"
)
