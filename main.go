// Todo allows you to create and delete items on a to-do list
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

const fileName = ".items.tdo"

var filePath string
var l = flag.Bool("l", false, "display items on the list")

func main() {
	flag.Parse()
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	filePath = fmt.Sprintf("%s/%s", usr.HomeDir, fileName)
	fmt.Println(filePath)

	add(flag.Args())
	if *l {
		list()
	}
}

func add(items []string) {
	s := strings.Join(flag.Args(), "\n")
	if len(s) < 1 {
		return
	}

	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write: %v", err)
		os.Exit(1)
	}
}

func list() {
	fmt.Println("LISTING ITEMS HERE") // TODO: implement list method
}
