package iteration

import (
	"strconv"
	"strings"
)

func Repeat(value string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated += value
	}
	count := strings.Count(repeated, value)
	return repeated + ", " + strconv.Itoa(count) + " times"

}
