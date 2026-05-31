package processor

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"lab3-detector/internal/stats"
)

var (
	mu        sync.Mutex
	LeakCache = make(map[string][]byte)
)

// ОПТИМІЗАЦІЯ: Компілюємо регулярний вираз один єдиний раз при старті сервісу
var imageRegex = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)

func RunWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				ProcessImage(id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	select {}
}

func ProcessImage(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	matched := imageRegex.MatchString(data)
	if matched {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())

		mu.Lock()
		LeakCache[key] = make([]byte, 1024*10)
		mu.Unlock()

		stats.IncrementProcessed("jpeg")

		// ВИПРАВЛЕНО ЗА ЗАУВАЖЕННЯМ: структурований JSON лог
		fmt.Printf("{\"level\":\"info\",\"ts\":%d,\"msg\":\"image successfully processed\",\"worker_id\":%d,\"type\":\"jpeg\"}\n", time.Now().Unix(), workerID)
	}
}
