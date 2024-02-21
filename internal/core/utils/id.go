package utils

import (
	"strings"

	guuid "github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/rs/xid"
)

func NewCode() string {
	return xid.New().String()
}

func NewCodeWithPrefix(prefix string, upper bool) string {
	if upper {
		return strings.ToUpper(prefix + NewCode())
	}
	return prefix + NewCode()
}

func NewID() string {
	return guuid.New().String()
}

func NewCodeWithPrefixAndLength(prefix string, length int) string {
	return prefix + NewCodeWithLength(length)
}

func NewCodeWithLength(length int) string {
	id, _ := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length)
	return id
}
