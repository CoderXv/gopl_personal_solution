package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"strings"

	//"io/ioutil"
	"log"
	"math"
	"math/rand"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

// 1.4
func solution4() {
	// maintain line - file info
	fileInfo := make(map[string]string)
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileInfo)
	} else {
		for _, arg :=  range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "1.4: %v\n", err)
				continue
			}
			countLines(f, counts, fileInfo)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("Appearance: %d\t Line: %s\tFileName: %s\n", n, line, fileInfo[line])
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int, fileInfo map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++
		fileInfo[input.Text()] = f.Name()
	}
}

// 1.5 & 1.6
var palette = []color.Color{color.Black, color.RGBA{
	R: 0,
	G: 0xFF,
	B: 0,
	A: 0xFF,
}}

func solution5 () {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器变化的个数
		res     = 0.001 //	角度分辨率
		size    = 100   //	图像画布包含 [-size..+size]
		nframes = 64    //	动画中的帧数
		delay   = 8     //	以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 1.6: 改变SetColorIndex()第三个参数
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// 1.7 1.8 1.9
func solution6() {
	for _, url := range os.Args[1:] {
		// 1.8
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// 1.9
		fmt.Println("Status Code: ", resp.StatusCode)

		// 1.7
		_, err = io.Copy(os.Stdout, resp.Body); if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}

// 1.10 1.11
func solution7() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// init an goroutine
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		// receive from an channel
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
