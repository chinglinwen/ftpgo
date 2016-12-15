package main

import (
	"flag"
	"fmt"
	urlpkg "net/url"
	"os"
	"path"

	"github.com/secsy/goftp"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: ./ftp ftp://ip/pub/test")
	}
	flag.Parse()

	if len(os.Args) <= 1 {
		fmt.Println("url not provided")
		os.Exit(1)
	}
	url := os.Args[1]

	u, err := urlpkg.Parse(url)
	if err != nil {
		checkErr(err)
	}

	var config goftp.Config

	if u.User != nil {
		user := u.User.Username()
		if user != "" {
			config.User = user
		}
		if pass, ok := u.User.Password(); ok {
			config.Password = pass
		}
	}

	// Create client object with default config
	client, err := goftp.DialConfig(config, u.Host)
	if err != nil {
		checkErr(err)
	}

	// Download a file to disk
	filename := path.Base(u.Path)
	file, err := os.Create(filename)
	if err != nil {
		checkErr(err)
	}

	err = client.Retrieve(u.Path, file)
	if err != nil {
		checkErr(err)
	}
}

func checkErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}
