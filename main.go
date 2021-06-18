package main

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"golang.org/x/sys/windows"
)

func main() {
	var free, total, avail uint64

	path := "c:\\"
	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		fmt.Println("failed to convert the path: ", err)
		os.Exit(1)
	}

	err = windows.GetDiskFreeSpaceEx(pathPtr, &free, &total, &avail)
	if err != nil {
		fmt.Println("failed to get free space: ", err)
		os.Exit(1)
	}

	fmt.Println("## IEC size (like Windows)")
	fmt.Printf("Total: %v, Free: %v, Available: %v\n", humanize.IBytes(total), humanize.IBytes(free), humanize.IBytes(avail))

	fmt.Println("## SI size")
	fmt.Printf("Total: %v, Free: %v, Available: %v\n", humanize.Bytes(total), humanize.Bytes(free), humanize.Bytes(avail))
}
