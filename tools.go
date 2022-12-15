//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golang/mock/mockgen/model"
	_ "golang.org/x/tools/cmd/stringer"
)
