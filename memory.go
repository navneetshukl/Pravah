package main

import (
	"log"
	"runtime"
)

func PrintMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("=== Memory Usage ===")
	log.Printf("Alloc:      %v KB", m.Alloc/1024)
	log.Printf("TotalAlloc: %v KB", m.TotalAlloc/1024)
	log.Printf("Sys:        %v KB", m.Sys/1024)
	log.Printf("NumGC:      %v", m.NumGC)
	log.Printf("====================")
}

func PrintMemoryUsageMB() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("=== Memory Usage (MB) ===")
	log.Printf("Alloc:      %.2f MB", float64(m.Alloc)/1024/1024)
	log.Printf("TotalAlloc: %.2f MB", float64(m.TotalAlloc)/1024/1024)
	log.Printf("Sys:        %.2f MB", float64(m.Sys)/1024/1024)
	log.Printf("NumGC:      %v", m.NumGC)
	log.Printf("===========================")
}