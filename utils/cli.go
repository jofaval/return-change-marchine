package utils

import (
	"fmt"
	"return-change-machine/config"
)

func DebugPrint(a ...any) {
	if !config.DEBUG_MODE {
		return
	}

	fmt.Println(a...)
}
