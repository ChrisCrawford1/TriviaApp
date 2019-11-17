package helpers

import (
	"fmt"
	"os"
)

// Exit will display a message terminate the process
func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
