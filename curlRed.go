// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	if len(os.Args[1:]) >= 1 {
		for _, url := range os.Args[1:] {
			go fetch(url, ch) // start a goroutine
		}
		for range os.Args[1:] {
			fmt.Println(<-ch) // receive from channel ch
		}
		argCount := strconv.Itoa(len(os.Args[1:]))
		fmt.Printf("%.2fs elapsed %s %s", time.Since(start).Seconds(), argCount, "lookup(s)")
	} else {
		banner := figure.NewColorFigure("curlRED", "graffiti", "red", true)
		go fetch("", ch) // start a goroutine
		banner.Print()
		fmt.Println(<-ch) // receive from channel ch
		fmt.Println("\n\nUsage: curlRed.exe / curlRed\nExample (ipv4): curlRed.exe 1.1.1.1\nExample (ipv6): curlRed.exe 2a03:2880:2130:cf05:face:b00c::1.Faceb00c\nExample (domain): curlRed.exe google.com\nExample (batch mode): curlRed.exe 1.1.1.1 2a03:2880:2130:cf05:face:b00c::1.Faceb00c google.com")
	}
	fmt.Println()

}

func fetch(lookup string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get("https://curl.red/" + lookup)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", lookup, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s %s %s %s", secs, "|", lookup, "|", body)
}
