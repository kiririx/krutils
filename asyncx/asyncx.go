package asyncx

import "sync"

// WithGoroutine 并发执行f函数
func WithGoroutine(f func(), count int) {
	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}
	wg.Wait()
}
