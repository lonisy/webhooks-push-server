package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	flag.Parse()
	defer glog.Flush()
	query := request.URL.Query()
	gitAct := query.Get("git")
	if gitAct == "" {
		glog.Errorf("Time: %s \nIP: %s \n action is empty by %s", time.Now(), request.RemoteAddr, request.URL.Path)
		fmt.Fprintf(writer, "action is empty by %s!\n", request.URL.Path)
		return
	}
	command := "/data/" + gitAct + "-update.sh"
	if IsExist(command) == false {
		fmt.Fprintf(writer, "Executable script not found !\n") // request.URL.Path = URI
		return
	}
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		glog.Errorf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	glog.Infof("Time: %s \nIP: %s \nExecute Shell:%s \nfinished with output:\n%s", time.Now(), request.RemoteAddr, command, string(output))
	fmt.Fprintf(writer, "succeed ! \n")
	return
}

func IsExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4433", nil)
}
