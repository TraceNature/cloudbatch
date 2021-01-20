/**
  @author: xinyulu
  @date: 2021/1/8 14:19
  @note: 
**/
package main

import (
    _ "cloudbatch/conf"
    "cloudbatch/interact"
    "fmt"
    "io/ioutil"
    "os"
    "os/signal"
    "strings"
    "syscall"
    //"github.com/c-bata/go-prompt"
)

func main() {

    //cmd.Execute()

    pdAddr := os.Getenv("PD_ADDR")
    if pdAddr != "" {
        os.Args = append(os.Args, "-u", pdAddr)
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT)

    go func() {
        sig := <-sc
        fmt.Printf("\nGot signal [%v] to exit.\n", sig)
        switch sig {
        case syscall.SIGTERM:
            os.Exit(0)
        default:
            os.Exit(1)
        }
    }()

    var input []string
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        b, err := ioutil.ReadAll(os.Stdin)
        if err != nil {
            fmt.Println(err)
            return
        }
        input = strings.Split(strings.TrimSpace(string(b[:])), " ")
    }

    interact.MainStart(append(os.Args[1:], input...))
}
