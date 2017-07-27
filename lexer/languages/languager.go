package languages

import (
	"github.com/shirou/golightan"
)

type Languager interface {
	GetModeMap() golightan.ModeMap
	GetCaseInsentive() bool
	GetDetect() golightan.Detect
}
