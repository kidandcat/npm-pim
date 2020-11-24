package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"

	"github.com/qeesung/asciiplayer/pkg/player"
)

func audio() {
	d, err := mp3.NewDecoder(bytes.NewReader(MustAsset("media/aa.mp3")))
	if err != nil {
		fmt.Println(err)
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	for {
		p := c.NewPlayer()
		defer p.Close()

		if _, err := io.Copy(p, d); err != nil {
			fmt.Println(err)
		}
	}
}

func video() {
	p := player.NewGifTerminalPlayer()
	p.Play("media/aa.gif", &player.DefaultPlayOptions)
}

func npm() string {
	cmd := exec.Command("npm", "install")
	stdout, _ := cmd.CombinedOutput()
	return string(stdout)
}

func main() {
	go video()
	go audio()
	res := npm()
	defer fmt.Println(res)
}
