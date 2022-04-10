package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	self := os.Getpid()

	go func() {

		args := []string{
			"attach", fmt.Sprint(self),
			"--listen=127.0.0.1:0",
			"--headless",
			"--api-version=2",
			//"--backend=native",
			//"--continue=true",
			"--accept-multiclient=true",
			//"--check-go-version=false",
			"--log",
			//"--log-output=rpc",
			//"--log-dest=dlv.log",
			//"--only-same-user=false",
		}

		cmd := exec.Command("dlv", args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		log.Printf("dlv will trace current process-%d in 5s", self)
		time.Sleep(time.Second * 5)

		if err := cmd.Start(); err != nil {
			log.Printf("dlv start fail: %v", err)
			return
		}
		log.Printf("dlv started: pid-%d", cmd.Process.Pid)

		if err := cmd.Wait(); err != nil {
			log.Printf("dlv status fail: %v", err)
			return
		}
		log.Printf("dlv status: %v", cmd.ProcessState)

		//log.Printf("dlv will stop in another 5s")
		//time.Sleep(time.Second * 5)
		//if err := cmd.Process.Kill(); err != nil {
		//log.Printf("dlv exit fail: %v", err)
		//return
		//}
		//log.Printf("dlv exited")
	}()

	var counter int
	for {
		log.Printf("count = %d", counter)
		counter++
		time.Sleep(time.Second)
	}
}
