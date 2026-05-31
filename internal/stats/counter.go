package stats

import "sync"

var (
	mu          sync.RWMutex
	GlobalStats = make(map[string]int)
)

// IncrementProcessed безпечно збільшує лічильник завдяки м'ютексу
func IncrementProcessed(imageType string) {
	mu.Lock() // Блокуємо мапу перед записом
	GlobalStats[imageType]++
	mu.Unlock() // Обов'язково звільняємо після запису
}
