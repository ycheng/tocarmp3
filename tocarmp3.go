
package main

import (
    "strings"
    "flag"
    "fmt"
    "os"
    "io/ioutil"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: tocarmp3 input_dir output_dir\n")
    flag.PrintDefaults()
    os.Exit(2)
}

var audioExts = map[string]bool {
    "aac": true,
    "m4a": true,
    "mp3": true,
}

var targetExt = "mp3"

func main() {
    flag.Usage = usage
    flag.Parse()

    args := flag.Args()
    if len(args) < 2 {
        usage()
        os.Exit(1);
    }
    // fmt.Println("args len is", len(args))

    inputFiles, error := ioutil.ReadDir(args[0])
    if error != nil {
        fmt.Fprintf(os.Stderr, "Can't open directory", args[0])
        os.Exit(1);
    }

    for _, file := range inputFiles {
        if file.Size() == 0  || strings.HasPrefix(file.Name(), ".") || file.IsDir() {
            continue
        }
        filePart := strings.Split(file.Name(), ".")
        if len(filePart) == 1 { continue }
        // fmt.Println("# of file parts:", len(filePart))
    
        fileExt := strings.ToLower(filePart[len(filePart) - 1])

        if ! audioExts[fileExt] { continue }
        fmt.Println("audio file name:", file.Name(), "ext:", fileExt)
    }
    // fmt.Printf("opening %s\n", args[0]);
    // ...
}
