package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/visualfc/atk/tk"
)

func init() {
	if runtime.GOOS == "windows" {
		dir, _ := filepath.Split(os.Args[0])
		if !filepath.IsAbs(dir) {
			root, err := os.Getwd()
			if err == nil {
				dir = filepath.Join(root, dir)
			}
			os.Chdir(dir)
		}
		tcl_lib := filepath.Join(dir, "lib", "tcl8.6")
		if _, err := os.Lstat(tcl_lib); err == nil {
			tk.InitEx(true, tcl_lib, "")
		}
	}
}

func bestFont() string {
	if runtime.GOOS == "windows" {
		return "Times"
	}
	return ""
}
