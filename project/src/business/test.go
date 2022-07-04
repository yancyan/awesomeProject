package business

import (
	"io/ioutil"
	"net/http"
	. "project/src/log"
)

//var wg sync.WaitGroup
//var lock sync.Mutex
//var once sync.Once

//var x int
//
//func add() {
//	for i := 0; i < 100000; i++ {
//		lock.Lock()
//		x = x + 1
//		lock.Unlock()
//	}
//	wg.Done()
//}

func Test() {

	//test()

	//testHttp()

}
func test() {
	//fmt.Println("", time.Now())
	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println("x = " + strconv.Itoa(x))

	//cCount := make(chan string, 1000)
	//go func(c chan string) {
	//	for k := range c {
	//		fmt.Println(k)
	//	}
	//}(cCount)
	//
	//for i := 0; i < 1000; i++ {
	//	cCount <- strconv.Itoa(i) + " count"
	//}
	//fmt.Println("send chan end.")
	//close(cCount)
	//time.Sleep(5 * time.Second)
	//c_count := make(chan string, 2)
	//
	//close(c_count)
	//
	//ct := <-c_count
	//fmt.Println("chan <= content is " + ct)

	//runtime.GOMAXPROCS(8)
	//wg.Add(5)
	//for i := 0; i < 5; i++ {
	//	go func(i int) {
	//		fmt.Println("exec " + string(i))
	//		wg.Done()
	//	}(i)
	//	fmt.Println(runtime.NumGoroutine())
	//}
	//wg.Wait()
	//fmt.Println("end.")

	//countChan := make(chan int)
	//
	//go func() {
	//	select {
	//	case a := <-countChan:
	//		fmt.Println(a)
	//		countChan <- a * 10
	//
	//	}
	//}()
	//for i := 0; i < 5; i++ {
	//	countChan <- i
	//}
	//time.Sleep(5 * time.Second)
	//fmt.Println("end...")

	// 只读信道
	//type Receiver <-chan int
	//var receiver Receiver = countChan

	//	只写信道
	//type Sender chan <- int
	//var sender Sender = countChan
}
func testHttp() {
	//
	//addrs, _ := net.InterfaceAddrs()
	//Log.Infoln(addrs)
	//
	//interfs, _ := net.Interfaces()
	//Log.Infoln(interfs)

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bdy, _ := ioutil.ReadAll(resp.Body)
	Log.Infoln(string(bdy))
}
