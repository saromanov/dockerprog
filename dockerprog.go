package main

import
(
    "fmt"
    "os/exec"
    "os"
    "sync"
    "flag"
    "strings"
)

//run go build on the app directory
func build() {
    _, err := exec.Command("go", "build").Output()
    if err != nil {
        fmt.Printf("Found error %v\n", err)
    }

    fmt.Println("Successful build")
}

//run godep tool
func godep() {
    goget("github.com/tools/godep")
     _, err := exec.Command("godep", "save", "-r").Output()
    if err != nil {
        fmt.Printf("Found error %v\n", err)
    }
}

//run go test on the app directory
func gotest() {
    fmt.Println("TESTS: ")
    result, err := exec.Command("go", "test").Output()
    if err != nil {
        fmt.Printf("Found error %v\n", err)
    }

    fmt.Println(string(result))
}

//run cover tool
func gotestcover() {
    fmt.Println("TESTS AND COVER")
    goget("golang.org/x/tools/cover")
    result, err := exec.Command("go", "test", "-cover").Output()
    if err != nil {
        fmt.Printf("Found error %v\n", err)
    }
    fmt.Println(string(result))
}


//parsing of arguments. Parsing arguments starts if config file is not found
func parsing() {
    var param = flag.String("process", "test,build", "build project")
    flag.Parse()
    items := strings.Split(*param, ",")
    if len(items) == 0 {
        fmt.Printf("Error in the parsing of arguments")
        return
    }

    for _, command := range items {
        switch command {
           case "build":
             build()
           case "test":
             gotest()
           case "cover":
             gotestcover()
           case "godep":
             godep()
        }
    }
}

//helpful method
func goget(target string) {
    _, err := exec.Command("go", "get", target).Output()
    if err != nil {
        fmt.Printf("Error to getting %s", target)
    }
}

//helful method
func gogetall(targets []string) {
    var wg *sync.WaitGroup
    wg.Add(len(targets))
    for _, target := range targets {
        go func(target string) {
            goget(target)
            wg.Done()
        }(target)
    }
    wg.Wait()
}

func main() {
    os.Chdir("/app")
    parsing()
}