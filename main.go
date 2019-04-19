// Todo allows you to create and delete items on a to-do list
package main

import (
	"bufio"
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

}

func list() {
	fmt.Println("LISTING ITEMS HERE") // TODO: implement list method
}

func writeData(items []string) {
	s := strings.Join(items, "\n")

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Fprint(w, s)

	err = w.Flush()
	if err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		os.Exit(1)
	}
}
