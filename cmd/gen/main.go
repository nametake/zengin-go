package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/nametake/zengin-go/internal"
)

func main() {
	var (
		datapath  string
		outputdir string
	)

	flag.StringVar(&datapath, "data", "", "data directory path")
	flag.StringVar(&outputdir, "output", "", "putput directory path")

	flag.Parse()

	if datapath == "" || outputdir == "" {
		flag.Usage()
		os.Exit(1)
	}

	banks, err := internal.Read(path.Join(datapath))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := os.Create(path.Join(outputdir, "banks.go"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := internal.Output(f, banks); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
