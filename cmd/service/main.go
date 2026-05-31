// Package main запускає основний сервіс та сервер профілювання pprof.
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Підключаємо pprof для зняття метрик

	"lab3-detector/internal/processor"
)

func main() {
	// Запускаємо pprof-сервер на порту 6060 у фоновій горутині
	go func() {
		log.Println("Pprof server started on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Image Metadata Processor started...")

	// Запускаємо пул із 5 воркерів
	processor.RunWorkerPool(5)
}
