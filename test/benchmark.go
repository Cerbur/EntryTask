package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	c        = flag.Int("c", 200, "Plz input client quantity")
	t        = flag.Int("t", 5, "Plz input times quantity")
	u        = flag.String("u", "http://127.0.0.1:8888/api/login", "Plz input url")
	isRandom = flag.Bool("r", true, "Plz input true or false")
)

var client = &http.Client{}

func init() {
	tr := &http.Transport{
		MaxIdleConns:        *c,
		MaxIdleConnsPerHost: *c,
	}
	client = &http.Client{Transport: tr}
}

var (
	total   = 0.0
	success = 0.0
	failure = 0.0
)

var wg sync.WaitGroup

func run(num int, gnum int) {
	defer wg.Done()

	no := 0.0
	ok := 0.0

	for i := 0; i < num; i++ {
		var nums int
		if *isRandom {
			nums = rand.Intn(10000000 + 1)
		} else {
			nums = gnum + 1
		}
		param := fmt.Sprintf("name=cerbur%d&password=123456", nums)

		// req, err := http.NewRequest("POST", *u, strings.NewReader(param))
		// req.Header.Add("Connection", "close")
		// resp, err := client.Do(req)
		resp, err := client.Post(*u, "application/x-www-form-urlencoded", strings.NewReader(param))
		if err != nil {
			log.Println(err)
			no += 1
			continue
		}
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		if resp.StatusCode != 200 {
			no += 1
			continue
		}

		ok += 1
		continue
	}

	success += ok
	failure += no
	total += float64(num)
}

func main() {
	startTime := time.Now().UnixNano()

	flag.Parse()

	if *c == 0 || *t == 0 || *u == "" {
		flag.PrintDefaults()
		return
	}

	for i := 0; i < *c; i++ {
		wg.Add(1)
		go run(*t, i)
	}

	wg.Wait()
	endTime := time.Now().UnixNano()

	cost := float64(endTime-startTime) / 1e9
	avgCost := cost / float64(*t)
	qps := float64(*c) / avgCost

	fmt.Println("is random:", *isRandom)
	fmt.Println("PreTotal:", (*c)*(*t))
	fmt.Println("Total:", success)
	fmt.Println("Concurrent:", *c)
	fmt.Println("UseTime:", fmt.Sprintf("%.4f", float64(endTime-startTime)/1e9), "s")
	fmt.Println("QPS:", fmt.Sprintf("%.4f", qps))
}
