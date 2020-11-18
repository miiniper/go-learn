package le2

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func GetUrl(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("get err :", err)
		return
	}
	if res.StatusCode != 200 {
		//fmt.Printf("get not 200 is ===%d ,,url=%s\n", res.StatusCode, url)
		fmt.Printf("1")
	}
}

//直接读取文件到内存，大文件不可取
func GetUrls(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("1111111111", err)
		return nil
	}
	UrlList := strings.Split(string(b), "\n")

	fmt.Println("urlList len ==", len(UrlList))

	return UrlList
}

//read big file,line read
func ReadLineReader(filePath string, ff func(string2 string)) {
	start := time.Now()
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer f.Close()
	lineReader := bufio.NewReader(f)
	for {
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		//业务逻辑
		ff(string(line))
		//fmt.Println(string(line))
	}
	fmt.Println("use readline time ==", time.Now().Sub(start))
}

//限制并发数
func GetUrls3() {
	start := time.Now()
	jobs := make(chan string, 10)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			for j := range jobs {
				GetUrl(j)
				wg.Done()
			}
		}()
	}
	urls := GetUrls("./ip")
	for _, j := range urls {
		jobs <- j
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("\ndoUrl use time ==", time.Now().Sub(start))
}

//限制并发数和行读
func ReadLineReader2(filePath string) {
	start := time.Now()
	var Wg sync.WaitGroup
	j := make(chan string)

	for i := 0; i < 100; i++ {
		go func() {
			for k := range j {
				GetUrl(k)
				Wg.Done()
			}
		}()
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer f.Close()
	lineReader := bufio.NewReader(f)
	for {
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		//业务逻辑
		Wg.Add(1)
		j <- string(line)
	}

	Wg.Wait()
	fmt.Println("\ndoUrl use time ==", time.Now().Sub(start))

}

//func TG() {
//	ReadLineReader2("./ip") // 输出结果2m25.077326452s
//	time.Sleep(time.Second)
//	GetUrls3() // 输出结果 2m23.229934631s
//	time.Sleep(time.Second)
//	ReadLineReader("./ip", GetUrl) // 输出结果 4m42.65718002s
//}

//func GetR() {
//	//reg, err := regexp.Compile("(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]")
//	//	reg, err := regexp.Compile("(ht|f)tp(s?)\\:\\/\\/[0-9a-zA-Z]([-.\\w]*[0-9a-zA-Z])*(:(0-9)*)*(\\/?)([a-zA-Z0-9\\-\\.\\?\\,\\'\\/\\\\\\+&amp;%\\$#_]*)?")
//	reg, err := regexp.Compile("http://([\\w-]+\\.)+[\\w-]+(/[\\w- ./?%&=]*)?")
//	if err != nil {
//		fmt.Println("create reg err = ", err)
//	}
//	s1 := reg.FindAllString("update pdc_rsc_rel set rsc_url = 'https://
//	fmt.Println(s1)
//}
