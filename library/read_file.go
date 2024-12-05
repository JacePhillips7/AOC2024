package read_file

import (
	"os"
	"strings"
)

func ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	read := string(file)
	return read
}
func SplitOnLine(s string) []string {
	read := strings.ReplaceAll(s, "\r\n", "\n")
	return strings.Split(read, "\n")
}
