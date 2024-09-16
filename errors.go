package main

import "log"

// Colors for logging
var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"

func NonFatal(err interface{}, f string, m string) {
	if err != nil {
		log.Printf(red+"[%s] - [ERROR] - %s - Trace: %v"+reset, f, m, err)
	} else if Debug {
		log.Printf(green+"[%s] - [OK] - %s - Trace: %v"+reset, f, m, err)
	}
}

func Fatal(err interface{}, m string, f string) {
	if err != nil {
		log.Fatalf(red+"[%s] - [ERROR] - %s - Trace: %v"+reset, f, m, err)
	} else if Debug {
		log.Fatalf(green+"[%s] - [OK] - %s - Trace: %v"+reset, f, m, err)
	}
}
