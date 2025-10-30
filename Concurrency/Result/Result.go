package Result

// A struct it to hold the result of a concurrent operation
// Similar to a model in c#
import (
	"time"
)

type Result struct {
	URL      string
	Duration time.Duration
	Error    error
}
