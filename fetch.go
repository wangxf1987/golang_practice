package main
import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
       if strings.HasPrefix(url, "http://") != true {
            url = "http://"+url
       }
       resp, err := http.Get(url)
       if err != nil {
           fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
           os.Exit(1)
       }
       if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
           fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
           os.Exit(1)
       }
       fmt.Printf("http code: %s", resp.Status)
       resp.Body.Close()
       fmt.Printf("%s", os.Stdout)
    }
}
