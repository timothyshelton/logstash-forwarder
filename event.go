package main

// vim: tabstop=2:noexpandtab:shiftwidth=2

import "os"

type FileEvent struct {
	Source *string `json:"source,omitempty"`
	Offset int64   `json:"offset,omitempty"`
	Line   uint64  `json:"line,omitempty"`
	Text   *string `json:"text,omitempty"`
	Fields *map[string]string

	fileinfo *os.FileInfo
}
