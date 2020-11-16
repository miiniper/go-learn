package le1

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//  ,channel,lock

func IfDemo1() {
	if s := 5; s > 5 {
		fmt.Println("a")
	} else if s < 5 {
		fmt.Println("b")
	} else {
		fmt.Println(s)
	}
}

func IfDemo2() {
	switch n := 7; n {
	case 5:
		fmt.Println("a")
	case 7:
		fmt.Println("b")

	}
}

func MapDemo1() {
	scoreMap := make(map[string]int, 8)
	scoreMap["aa"] = 90
	scoreMap["bb"] = 100
	fmt.Println(scoreMap)
	if v, ok := scoreMap["aa"]; ok {
		fmt.Println(v)
	}
}

func Sdemo() {
	str := "a a a a n n n l o p j k k k k"
	var sm = make(map[string]int)
	sl := strings.Split(str, " ")
	for _, k := range sl {
		sm[k] = sm[k] + 1
	}
	fmt.Println(sm)

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func ChannelDemo1() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func RWLockDemo1() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
