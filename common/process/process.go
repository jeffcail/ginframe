package _process

import "runtime"

func GroRuntimeMaxCpu() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
