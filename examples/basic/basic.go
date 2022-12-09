package main

import (
	"bufio"
	"fmt"
	"github.com/gen2brain/x264-go"
	"image"
	_ "image/jpeg"
	"os"
	"os/exec"
)

//ffmpeg -re -i video.mp4 -c:a copy -c:v copy -f flv rtsp://127.0.0.1:1935/live/test

func main() {
	opts := &x264.Options{
		Width:     640,
		Height:    480,
		FrameRate: 25,
		Preset:    "veryfast",
		Tune:      "stillimage",
		Profile:   "baseline",
		LogLevel:  x264.LogDebug,
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	//w, err := os.Create("example.264")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//defer w.Close()

	enc, err := x264.NewEncoder(w, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer enc.Close()

	r, err := os.Open("example.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img, _, err := image.Decode(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i <= 100; i++ {
		err = enc.Encode(img)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	err = enc.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd := exec.Command("ffmpeg -i pipe:0 -c:a copy -c:v copy -f flv rtsp://127.0.0.1:1935/live/test")
	err = cmd.Run()
	fmt.Fprintf(os.Stderr, "stderr: %s", err)
}
