package main

import (
    "fmt"
    "os/exec"
    "os"
)

func main() {
    // os.start processes
    env := os.Environ()
    procAttr := &os.ProcAttr{
        Env: env,
        Files: []*os.File{
            os.Stdin,
            os.Stdout,
            os.Stderr,
        },
    }
    pid, err := os.StartProcess("/bin/ls", []string{"-a", "-l"}, procAttr)
    if err != nil {
        fmt.Printf("Error %v starting process!", err)
        os.Exit(1)
    }
    fmt.Printf("The process id is %v\n", pid)

    pid, err = os.StartProcess("/bin/ps", []string{"-e", "opid,ppid,comm"}, procAttr)
    if err != nil {
        fmt.Printf("Error %v starting process!", err)
        os.Exit(1)
    }
    fmt.Printf("The process id is %v\n", pid)

    cmd := exec.Command("vim")
    cmd.Stdin  = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        fmt.Printf("Error %v executing command!", err)
        os.Exit(1)
    }
    fmt.Printf("The command is %v\n", cmd)
}