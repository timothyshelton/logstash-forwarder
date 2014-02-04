package main

// vim: tabstop=2:noexpandtab:shiftwidth=2

import (
	"os"
	"syscall"
)

func file_ids(info *os.FileInfo) (uint64, uint64) {
	fstat := (*info).Sys().(*syscall.Stat_t)
	return fstat.Ino, fstat.Dev
}
