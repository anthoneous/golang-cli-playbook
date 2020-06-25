package module7

import (
	"fmt"
	"runtime"
)

func content() {
	fmt.Println(runtime.GOOS)
}

//go:generate -command goimports fmt
//go:generate -command goimports runtime
