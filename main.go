package main

import (
	"flag"
	"fmt"
	"noauth/lib"
	"os"
	"runtime"
	"strings"
)

var (
	u     string
	n     string
	t     int
	h     bool
	a     string
	debug int
	list  bool
	proxy string
)

func init() {
	flag.BoolVar(&h, "h", false, "This help")
	flag.StringVar(&u, "u", "", "A target url(Please add http or https)")
	flag.StringVar(&n, "n", "", "An interface without authentication, such as /login")
	flag.StringVar(&a, "a", "", "An interface that requires authentication, such as /admin/adduser")
	flag.IntVar(&t, "t", runtime.NumCPU(), "Thread Num")
	flag.IntVar(&debug, "debug", 0, "choose start debug, such -debug 1")
	flag.BoolVar(&list, "list", false, "List mode, used to enumerate targets")
	flag.StringVar(&proxy, "proxy", "", "Set an HTTP proxy (e.g., http://127.0.0.1:8080)") // 新增 proxy 参数
	flag.Usage = usage
}

func checkFlags() {
	if list && u != "" {
		fmt.Println("Error: -list and -u cannot be used together. Please choose one.")
		os.Exit(0)
	}

	if n == "" || a == "" {
		fmt.Println("Error: Missing parameter. Please use the -h  to view the required parameters.")
		os.Exit(0)
	}

	if !list && (u == "") {
		fmt.Println("Error: Missing parameter. Please use the -h  to view the required parameters.")
		os.Exit(0)
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, `noauth version: 1.0.0
Usage:  [-unat] [-u url] [-n interface without authentication] [-a interface An interface that requires authentication] [-t thread] [-debug choose start debug] [-h help]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	lib.Logo()
	flag.Parse()

	if h {
		flag.Usage()
		os.Exit(0)
	}

	checkFlags()

	if list {
		lib.Dict(n, a)
		os.Exit(0)
	}

	res1 := strings.Contains(u, "http://")
	res2 := strings.Contains(u, "https://")

	if !res1 && !res2 {
		fmt.Println(lib.Red("[-] Please add http or https for url !!!"))
		os.Exit(0)

	}

	lib.GetStart(u, n, a, t, debug)
	lib.PostStart(u, n, a, t, debug)

}
