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

const _version = "1.0.1"
const fileName = ".items.tdo"

var filePath string
var l = flag.Bool("l", false, "Display items on the list") // print item list
var d = flag.Int("d", 0, "Delete item number")             // delete item d
var v = flag.Bool("v", false, "Display version")

func main() {
	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	filePath = fmt.Sprintf("%s/%s", usr.HomeDir, fileName)

	if *v {
		fmt.Printf("todo v%s\n", _version)
	}

	addItem()
	if *d > 0 {
		deleteItem(*d - 1)
	}
	if *l || (*d == 0 && !*v && len(flag.Args()) <= 0) {
		listItems()
	}
}

func addItem() {
	if len(flag.Args()) <= 0 {
		return
	}

	item := strings.Join(flag.Args(), " ")
	items := append(readData(), item)
	writeData(items)
}

func deleteItem(i int) {
	items := readData()

	if i < 0 || i >= len(items) {
		fmt.Println("Item not found.")
		os.Exit(1)
	}

	item := items[i]
	items = append(items[:i], items[i+1:]...)
	writeData(items)
	fmt.Printf("Deleted: %s\n", item)
}

func listItems() {
	items := readData()
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

// readData returns the item list from the data file as a slice of string.
func readData() []string {
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var items []string

	input := bufio.NewScanner(f)
	for input.Scan() {
		items = append(items, input.Text())
	}
	if err := input.Err(); err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}

	return items
}

// writeData overwrites the data file with the item list of the items slice.
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
