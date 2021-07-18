package main

import (
	"fmt"
	goversion "github.com/hashicorp/go-version"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// app binary is in the same folder as the launcher
const appFilename = "app.exe"

var version = "unknown"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Print(version)
		os.Exit(0)
	}
	// get current version of app
	// get most recent version from update server
	// if server version is newer than app version, download the update
	// apply the update, if any

	currentAppVersion := getCurrentAppVersion(appFilename)
	fmt.Printf("Current app version: %s\n", currentAppVersion)

	cl := &http.Client{}
	res, err := cl.Get("http://localhost:7000/app-version")
	handleError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	handleError(err)
	serverVersion := string(body)

	localVer, err := goversion.NewVersion(currentAppVersion)
	handleError(err)
	remoteVer, err := goversion.NewVersion(serverVersion)
	handleError(err)

	if localVer.LessThan(remoteVer) {
		// initiate update process
		resp, err := cl.Get("http://localhost:7000/download-app")
		handleError(err)
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Fatal("could not download: status 500: ", err.Error())
		}

		fh, err := os.OpenFile(appFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0744)
		handleError(err)
		defer fh.Close()

		_, err = io.Copy(fh, resp.Body)
		handleError(err)
	}

	// start app
	//go func() {
	fmt.Println("starting app.exe...")
	cmd := exec.Command(appFilename)
	//bd, _ := filepath.Abs(".")
	//cmd.Dir = bd
	err = cmd.Run()
	if err == nil {
		cmd.Process.Release()
	} else {
		log.Fatal("could not start process:", err.Error())
	}

	//}()
	//time.Sleep(200 * time.Millisecond)

	// shutdown self

}

func getCurrentAppVersion(filename string) string {
	cmd := exec.Command(filename, "version")
	out, _ := cmd.Output()
	return string(out)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
