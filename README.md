# ftpgo
A minimal ftp client command (download only) written in Go

# Usage

```
$ ./ftp -h
Usage of ftp:
  -o string
        output filename(or path/filename)
  -v    show version.

Example: 
   ./ftp ftp://ip/pub/test
   ./ftp ftp://user:pass@ip/pub/test
   ./ftp -o file ftp://ip/pub/test
   ./ftp -o dir1/file ftp://ip/pub/test
``` 

# Auth

Add user and password in the url to set user and password

For example

```
./ftp ftp://user:pass@ip/pub/test
```

end.