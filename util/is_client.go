package util

import (
	"strings"

	"goau-biat/util/winapi"
)

func IsClient() bool {
	title := winapi.GetTitle()
	return strings.Contains(title, "Tibia - ") && !strings.Contains(title, "Tibia - Free")
}
