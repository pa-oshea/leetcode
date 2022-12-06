package main

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"time"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
