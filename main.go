package rofigo

import (
	"fmt"
	"log"
	"os/exec"
	// "os"
	"strings"
)

func Hello() string {
	return "hello world"
}

type Page struct {
	Options   []string
	Selection string
}

func New(options ...string) *Page {
	return &Page{
		Options: options,
	}
}

func (p *Page) Show() {
	command := fmt.Sprintf("echo -e '%s' | rofi -dmenu", strings.Join(p.Options, "\n"))
	cmd := exec.Command("/bin/sh", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	p.Selection = strings.TrimSuffix(string(out), "\n")
	// fmt.Println(p.Selection)
}

// func main() {
// 	page1 := New("option1", "option2", "option2")
// 	page1.Show()
// 	log.Println(page1.Selection)
// 	// var options = "1\n2\n3\n"

// 	// options := []string{"one", "two", "three", "four"}
// 	// command := fmt.Sprintf("echo -e '%s' | rofi -dmenu", strings.Join(options, "\n"))
// 	// log.Println(command)
// 	// // os.Exit(0)
// 	// othercommand := "echo -e 'hello\nhi\nello' | rofi -dmenu"
// 	// log.Print(othercommand)
// 	// // cmd := exec.Command("/bin/sh", "-c", "'echo -e \"1\n2\n3\n\" | rofi -dmenu'")

// 	// cmd := exec.Command("/bin/sh", "-c", command)

// 	// // cmd := exec.Command("/bin/sh", "-c", command)
// 	// // cmd := exec.Command("rofi", "-show", "drun")

// 	// // uncomment to pipe to stdout of go program
// 	// // cmd.Stdout = os.Stdout
// 	// // cmd.Stderr = os.Stderr
// 	// out, err := cmd.Output()
// 	// if err != nil {
// 	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
// 	// }
// 	// log.Print(string(out))

// 	// fmt.Println(Hello())
// }
