package opts

import (
	"github.com/gobwas/glob"
	"sync/atomic"
)

type Opts struct {
	Base       string `short:"b" long:"base" description:"The root folder from where to look after files" required:"false" default:"./"`
	Stop       bool   `short:"s" long:"stop" description:"Stops search on first match" required:"false"`
	StopFlag   bool
	Max        uint64 `short:"m" long:"max" description:"Maximum amount of matches" required:"false" default:"10000"`
	MaxCounter uint64
	Search     glob.Glob
}

// Check if a match was found on the stop flag or
// the requested matches exceeded limit
func (opts *Opts) Check() bool {
	if opts.Stop && opts.StopFlag {
		return false
	}
	if opts.MaxCounter < opts.Max {
		return true
	}
	return false
}

func (opts *Opts) MatchFound() {
	if opts.Stop && !opts.StopFlag {
		opts.StopFlag = true
	}
	atomic.AddUint64(&opts.MaxCounter, 1)
}
