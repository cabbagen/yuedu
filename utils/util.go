package utils

import (
	"bytes"
	"fmt"
	"crypto/md5"
	"log"
)

func MakeMD5(s string) string {
	var buffer *bytes.Buffer = bytes.NewBufferString("")

	if _, error := fmt.Fprintf(buffer, "%x", md5.Sum([]byte(s))); error != nil {
		log.Println(error)
	}

	return buffer.String()
}
