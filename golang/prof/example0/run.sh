#!/bin/sh

GOGC=off go test -bench=BenchmarkRegex -cpuprofile cpu.out
go tool pprof cpu.out
