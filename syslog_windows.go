package main

// vim: tabstop=2:noexpandtab:shiftwidth=2

import "log"

func configureSyslog() {
	log.Printf("Logging to syslog not supported on this platform\n")
}
