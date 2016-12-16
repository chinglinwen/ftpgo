package main

import (
        "flag"
        "fmt"
        urlpkg "net/url"
        "os"
        "path"

        "github.com/secsy/goftp"
)

var example = `
Example: 
   ./ftp ftp://ip/pub/test
   ./ftp ftp://user:pass@ip/pub/test
   ./ftp -o file ftp://ip/pub/test
   ./ftp -o dir1/file ftp://ip/pub/test
`

func main() {
        flag.Usage = func() {
                fmt.Println("Usage of ftp:")
                flag.PrintDefaults()
                fmt.Printf(example)
        }
        outfile := flag.String("o", "", "output filename(or path/filename)")
        version := flag.Bool("v", false, "show version.")
        flag.Parse()

        if *version {
                fmt.Println("version=1.0.1, 2016-12-16")
                os.Exit(1)
        }

        args := flag.Args()
        fmt.Println(args)
        if len(args) < 1 {
                fmt.Println("url not provided")
                os.Exit(1)
        }
        url := args[0]

        u, err := urlpkg.Parse(url)
        checkErr(err)

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

        client, err := goftp.DialConfig(config, u.Host)
        checkErr(err)

        filename := path.Base(u.Path)
        if *outfile != "" {
                outDir := path.Dir(*outfile)
                if outDir != "." {
                        checkErr(os.MkdirAll(outDir, 0777))
                }
                filename = *outfile
        }
        tmpfilename := filename + ".tmp"
        file, err := os.Create(tmpfilename)
        checkErr(err)
        defer file.Close()

        // Download a file to disk
        err = client.Retrieve(u.Path, file)
        checkErr(err)

        err = os.Rename(tmpfilename, *outfile)
        checkErr(err)
}

func checkErr(err error) {
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}