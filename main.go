package main

import (
    "log"
    "net/http"
    "os"
    "os/exec"
    "fmt"
    "strings"
)

func main() {

    router := NewRouter()

    loadDevices()

    log.Fatal(http.ListenAndServe(":2488", router))
}

func loadDevices() {
    cmd := exec.Command("irsend", "list \"\" \"\"")
    printCommand(cmd)
    output, err := cmd.CombinedOutput()

    if err != nil {
        printError(err)
    } else {
        createDevices(output)
    }
}

func createDevices(args []byte) {
    printOutput(args)
}

func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
  if err != nil {
    os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  }
}

func printOutput(outs []byte) {
  if len(outs) > 0 {
    fmt.Printf("==> Output: %s\n", string(outs))
  }
}
