package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch() {
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w,"<h1>Hello,this is my deploy server!</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8080",nil)
}

// env GOOS=linux GOARCH=amd64 go build
// cp newweb/deployserver/deployserver deployserver
// cp newweb/deploy.sh deploy.sh

