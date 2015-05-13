package errors

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func stack(skip int) string {

	buffer := &bytes.Buffer{}

	for i := skip; ; i++ {

		pc, file, line, ok := runtime.Caller(i)

		if !ok {

			break
		}

		var function string

		if fn := runtime.FuncForPC(pc); fn != nil {

			parts := strings.Split(fn.Name(), "/")

			function = parts[len(parts)-1]
		}

		fmt.Fprintf(buffer, "%s:%d %s(0x%x)\n", file, line, function, pc)
	}

	return buffer.String()
}
