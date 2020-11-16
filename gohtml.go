package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/PuerkitoBio/goquery"
)

type AppInfo struct {
	Name         string `json:"name"`
	NameNC       string `json:"name_nc"`
	OweApp       string `json:"owe_app"`
	Type         string `json:"type"`
	Docker       string `json:"docker"`
	BU           string `json:"bu"`
	Station      string `json:"station"`
	Architecture string `json:"architecture"`
	DevDomain    string `json:"dev_domain"`
	Zone         string `json:"zone"`
	OnlineT      string `json:"online_t"`
	Owner        string `json:"owner"`
}

func appInfo(appName string) {

	url := "https://test/app/showApp.htm?type=basic&appName=%s"
	url = fmt.Sprintf(url, appName)
	cookie := "test"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("req err ", err)
	}
	req.Header.Add("cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp err ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("new query  err ", err)
	}

	s1 := doc.Find("td[class=table-label]").Next().Text()
	//s1 := doc.Find("tr").Text()

	s2, err := decodeToGBK(s1)
	if err != nil {
		fmt.Println("get gbk err ", err)
	}

	//删除 tab 空行
	s2 = strings.ReplaceAll(s2, " ", "\t")
	s2 = strings.ReplaceAll(s2, "\t", "")
	//s2 = strings.ReplaceAll(s2, "   ", "")
	//	s2 = strings.ReplaceAll(s2, " ", "nihao")

	//	s2 = strings.ReplaceAll(s2, " ", "")
	//fmt.Println(s2)

	//https://my.oschina.net/u/4321737/blog/3530170

	//r := regexp.MustCompile(``)
	//s3 := r.FindAllString(s2, -1)
	//s3 := r.ReplaceAllString(s2, ".")

	//同时 按照 \n和,进行分割
	//s3 := strings.FieldsFunc(s2, func(r rune) bool {
	//	return strings.ContainsRune("\n \t", r)
	//})

	s3 := strings.Split(s2, "\n")
	for i, j := range s3 {
		fmt.Printf("i==%d,s==%s\n", i, j)
	}

	//fmt.Println(s3)

	//var app AppInfo
	////app.Name = s3[]
	//app.NameNC = s3[1]
	//app.OweApp = s3[3]
	//app.Docker = s3[5]
	//app.BU = s3[13]
	//app.Station = s3[25]
	//app.Architecture = s3[35]
	//app.DevDomain = s3[36]
	//app.Zone = s3[37]
	//app.OnlineT = s3[42]
	//app.Owner = s3[37]
	//
	//fmt.Printf("%v+\n", s3)

}

//copy
func decodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}
