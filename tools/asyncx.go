package tools

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
)

type Async struct {
}

// WithGoroutine 并发执行f函数
func (receive *Async) WithGoroutine(f func(), count int) {
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

// ScheduleTask 定义一个函数,接受cron格式的字符串和要执行的函数
//
//	分钟 (0-59)
//	小时 (0-23)
//	日期 (1-31)
//	月份 (1-12)
//	星期几 (0-7,其中0和7都表示星期日)
//	* * * * *
func (receive *Async) ScheduleTask(cronStr string, task func()) error {
	// 创建一个新的cron调度器
	c := cron.New()

	// 添加任务到调度器
	_, err := c.AddFunc(cronStr, task)
	if err != nil {
		return fmt.Errorf("添加任务失败: %v", err)
	}

	// 启动调度器
	c.Start()

	// 保持程序运行
	select {}
}
