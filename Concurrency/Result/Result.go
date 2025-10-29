package Result

import (
	"time"
)

type Result struct {
	URL      string
	Duration time.Duration
	Error    error
}
