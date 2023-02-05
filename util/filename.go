package util

import (
	"fmt"
	"strings"
	"time"
)

func GenerateFilenameByUploadDatetime(filename string) (string, error) {
	label := time.Now().Format("20060102150404")
	var b strings.Builder
	b.Grow(len(label) + len(filename) + 1)
	if _, err := fmt.Fprintf(&b, "%s-%s", label, filename); err != nil {
		return "", err
	}
	return b.String(), nil
}
