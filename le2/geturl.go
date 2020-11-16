package le2

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetUrl() {
	urls := GetUrls()
	for _, j := range urls {
		res, err := http.Get(j)
		if err != nil {
			fmt.Println("get err :", err)
		}
		if res.StatusCode != 200 {
			fmt.Println("get not 200 is ===", j)
		}
		fmt.Println("ok", j)
	}
}

func GetUrls() []string {
	b, err := ioutil.ReadFile("./ip")
	if err != nil {
		fmt.Println("1111111111", err)
	}
	UrlList := strings.Split(string(b), "\n")
	//	fmt.Println(UrlList)
	return UrlList
}
