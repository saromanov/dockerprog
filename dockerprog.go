package main

import
(
    "fmt"
    "os/exec"
    "os"
    "sync"
    "flag"
)

func goget(target string) {
    _, err := exec.Command("go", "get", target).Output()
    if err != nil {
        fmt.Printf("Error to getting %s", target)
    }
}

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

func build() {
    _, err := exec.Command("go", "build").Output()
    if err != nil {
        fmt.Printf("Found error %v\n", err)
    }

    fmt.Println("Successful build")
}

func godep() {
    goget("github.com/tools/godep")
     _, err := exec.Command("godep", "save", "-r").Output()
    if err != nil {
        fmt.Printf("Found error %v\n", err)
    }
}

func gotest() {
    fmt.Println("TESTS: ")
    result, err := exec.Command("go", "test").Output()
    if err != nil {
        fmt.Printf("Found error %v\n", err)
    }

    fmt.Println(string(result))
}

func main() {
    os.Chdir("/app")
    gotest()
    build()
    var param = flag.String("value", "default", "try")
    flag.Parse()
    fmt.Println(*param)
}