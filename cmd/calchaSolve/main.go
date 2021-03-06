/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: qbCalchaSolve
 * File Name: main.go
 * Last Modified: 21/01/2021, 08:53
 */

package main

import (
	"calchaSolve/pkg/solv"
	"flag"
	"fmt"
	"github.com/Bytesimal/goutils/pkg/fileio"
	"github.com/Bytesimal/goutils/pkg/httputil"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const quizURL = "https://quantbet.com/quiz/dev"
const submitURL = "https://quantbet.com/submit"

var repeat bool
var saveSolutions bool
var rateLimit time.Duration

var debugProxy, _ = url.Parse("http://localhost:9090")
var cli = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyURL(debugProxy),
	},
}

func init() {
	// flag parse
	flag.BoolVar(&repeat, "r", false, "controls if the program should repeatedly solve new calcas.")
	flag.BoolVar(&saveSolutions, "s", false, "if true, saves all returned solution pages in a temp dir")
	flag.DurationVar(&rateLimit, "l", 1, "time between each calcha solve")
	flag.Parse()

	// init tmp dir
	if saveSolutions {
		fileio.Init("", "calchaSolve_*")
		log.Printf("Saving solutions in %s", fileio.TmpDir)
	}

	// Init cookiejar
	cli.Jar, _ = cookiejar.New(nil)
}

func main() {
	t := time.NewTicker(rateLimit)

	for {
		// rq and parse html
		rq, _ := http.NewRequest("GET", quizURL, nil)
		rq.Header.Set("User-Agent", httputil.RandUA())
		rsp, err := cli.Do(rq)
		if err != nil {
			log.Fatalf("can't request calcha page: %s", err)
		}
		page, err := goquery.NewDocumentFromReader(rsp.Body)
		if err != nil {
			log.Fatalf("can't parse html: %s", err)
		}
		rsp.Body.Close()

		// Parse n1 and n2
		n1, err := strconv.ParseInt(page.Find("form#quiz > p > strong").Get(0).FirstChild.Data, 10, 64)
		if err != nil {
			log.Fatalf("can't parse int64 from HTML: %s", err)
		}
		n2, err := strconv.ParseInt(page.Find("form#quiz > p > strong").Get(1).FirstChild.Data, 10, 64)
		if err != nil {
			log.Fatalf("can't parse int64 from HTML: %s", err)
		}

		// solve
		solution := solv.GCD(n1, n2)
		log.Printf("n1: %10d and n2: %10d | Solution: %d\n", n1, n2, solution)

		// POST solutions
		rq, _ = http.NewRequest("POST", submitURL, strings.NewReader(
			fmt.Sprintf("divisor=%d", solution)))
		rq.Header.Set("User-Agent", httputil.RandUA())
		rq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rq.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		rsp, err = cli.Do(rq)
		if err != nil {
			log.Fatalf("can't submit solution: %s", err)
		}

		// save solutions
		if saveSolutions {
			tmpFPath, err := fileio.TmpPath(fmt.Sprintf("%dx%d.html", n1, n2))
			if err != nil {
				log.Fatalf("can't generate tmp path: %s", err)
			}
			f, err := os.Create(tmpFPath)
			if err != nil {
				log.Fatalf("can't create file in tmp path %s: %s", tmpFPath, err)
			}
			// Copy rsp contents into file
			_, err = io.Copy(f, rsp.Body)
			if err != nil {
				log.Fatalf("can't copy rsp body into f at %s: %s", tmpFPath, err)
			}
			f.Close()
		}

		rsp.Body.Close()
		// Next iter if repeat
		if !repeat {
			break
		}

		// Rate limit
		<-t.C
	}
}
