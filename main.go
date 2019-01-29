package main

import (
	"fmt"

	"github.com/fwiedmann/airgab/pkg/opts"
)

func main() {
	o := opts.New()
	fmt.Print(o.Source)
}
