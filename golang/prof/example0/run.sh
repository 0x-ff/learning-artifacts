#!/bin/sh

#go test -bench=. -cpuprofile cpu.out
GOGC=off go test -bench=BenchmarkRegex -cpuprofile cpu.out
go tool pprof cpu.out