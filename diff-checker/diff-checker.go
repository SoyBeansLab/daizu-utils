package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	diff "github.com/sergi/go-diff/diffmatchpatch"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("There is a problem with the given command line argument.")
		os.Exit(2)
	}
}

func isDiff(a, b string) bool {
	dmp := diff.New()
	diffs := dmp.DiffMain(a, b, false)

	if len(diffs) != 1 { // 差分が無い場合は1が帰ってくる
		return true
	}
	return false
}

func main() {
	flag.Parse()

	f1, err := os.Open(flag.Arg(0))
	defer f1.Close()
	checkErr(err)
	text1, err := ioutil.ReadAll(f1)
	checkErr(err)

	f2, err := os.Open(flag.Arg(1))
	defer f2.Close()
	checkErr(err)
	text2, err := ioutil.ReadAll(f2)
	checkErr(err)

	if isDiff(string(text1), string(text2)) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
