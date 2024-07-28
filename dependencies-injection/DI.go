package dependenciesinjection

import (
	"fmt"
	"io"
)

// logMessage can write to any io.Writer without modification
func logMessage(w io.Writer, level, msg string) {
	fmt.Fprintf(w, "[%s] %s", level, msg)
}
