package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	goversion "github.com/hashicorp/go-version"
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
		fmt.Println("local app version is lower, initiating update process...")
		// initiate update process
		resp, err := cl.Get("http://localhost:7000/app-download")
		handleError(err)

		if resp.StatusCode != 200 {
			log.Fatal("could not download: status 500: ", err.Error())
		}

		fh, err := os.OpenFile(appFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0744)
		handleError(err)

		_, err = io.Copy(fh, resp.Body)
		handleError(err)

		_ = resp.Body.Close()
		_ = fh.Close()
	} else {
		fmt.Println("app is at the most recent version")
	}

	// start app
	fmt.Println("starting app.exe...")
	time.Sleep(2 * time.Second)

	bd, _ := filepath.Abs(".")
	cmd := exec.Command("cmd.exe", "/c", "start", filepath.Join(bd, appFilename))
	//cmd.Dir = bd
	_ = cmd.Run()
	_ = cmd.Process.Release()

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
