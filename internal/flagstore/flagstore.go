package flagstore

import (
	"sync"

	flags "github.com/jessevdk/go-flags"
)

type FlagStore struct {
	MountLoc    string `short:"m" description:"The mount path on the host"`
	Selection   string `short:"s" description:"The selection of the scenarios to run, default=all" optional:"true" default:"all"`
	ScreenLines int    `short:"l" description:"The screenheight in lines to project information in the lower box." default:"5"`
	Verbose     bool   `short:"v" description:"Enable verbose, default=false."`
}

var flagstore FlagStore
var once sync.Once

// getFlags() obtains the called flags when executing the program
// and parses those into a FlagStore struct, containing definitions for possible flags.
// For now, mountlocation, ScreenLines (used for Terminal UI) and verbose are the only implemented flags.
func getFlags() FlagStore {
	once.Do(func() {
		_, err := flags.Parse(&flagstore)
		if err != nil {
			panic(err)
		}
	})
	return flagstore
}
