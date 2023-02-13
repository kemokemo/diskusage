package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"golang.org/x/sys/windows"
)

var (
	ver bool
)

func init() {
	flag.BoolVar(&ver, "v", false, "display version")
	flag.Parse()
}

func main() {
	os.Exit(run())
}

func run() int {
	if ver {
		fmt.Fprintf(os.Stdout, "diskusage version %s.%s\n", Version, Revision)
		return 0
	}

	var free, total, avail uint64

	path := "c:\\"
	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		fmt.Println("failed to convert the path: ", err)
		return 1
	}

	err = windows.GetDiskFreeSpaceEx(pathPtr, &free, &total, &avail)
	if err != nil {
		fmt.Println("failed to get free space: ", err)
		return 1
	}

	fmt.Println("## IEC size (like Windows)")
	fmt.Printf("Total: %v, Free: %v, Available: %v\n", humanize.IBytes(total), humanize.IBytes(free), humanize.IBytes(avail))

	fmt.Println("## SI size")
	fmt.Printf("Total: %v, Free: %v, Available: %v\n", humanize.Bytes(total), humanize.Bytes(free), humanize.Bytes(avail))

	return 0
}
