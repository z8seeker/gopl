// Echo4 prints its command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()                               // must call before the flags are used, to update the flag variables from their default values
	fmt.Print(strings.Join(flag.Args(), *sep)) // get the non-flag arguments
	if !*n {
		fmt.Println()
	}
}
