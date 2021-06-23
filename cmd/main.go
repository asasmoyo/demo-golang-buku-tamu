package main

import (
	"fmt"
	"os"

	"github.com/asasmoyo/demo-golang-buku-tamu/httpsrv"
)

func main() {
	srv := httpsrv.Server{
		ListenIP:   "127.0.0.1",
		ListenPort: "8080",
		DBConnStr:  "postgres://postgres:password@127.0.0.1:6432/bukutamu",
	}

	fmt.Println("starting http server...")

	if err := srv.Init(); err != nil {
		fmt.Println("error initializing server")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := srv.Start(); err != nil {
		fmt.Println(err)
	}
}
