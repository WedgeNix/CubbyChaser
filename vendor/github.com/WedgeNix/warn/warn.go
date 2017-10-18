package warn

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	warnings = map[string]*sync.Once{}
	debugr   = bufio.NewReader(os.Stdin)
)

func Do(warning string) {
	w, found := warnings[warning]
	if !found {
		w = new(sync.Once)
		warnings[warning] = w
	}
	w.Do(func() {
		for {
			fmt.Print("WARNING: " + warning + "  (continue? y/n)  ")
			input, _ := debugr.ReadString('\n')
			if strings.HasPrefix(input, "n") {
				os.Exit(1)
			} else if strings.HasPrefix(input, "y") {
				return
			}
		}
	})
}
