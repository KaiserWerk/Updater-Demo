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

const launcherFilename = "launcher.exe"

// set at compile time
var version = "unknown"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Print(version)
		os.Exit(0)
	}

	currentLauncherVersion := getCurrentLauncherVersion(launcherFilename)
	cl := &http.Client{}
	res, err := cl.Get("http://localhost:7000/launcher-version")
	handleError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	handleError(err)
	serverVersion := string(body)

	localVer, err := goversion.NewVersion(currentLauncherVersion)
	handleError(err)
	remoteVer, err := goversion.NewVersion(serverVersion)
	handleError(err)

	if localVer.LessThan(remoteVer) {
		fmt.Println("local launcher version is lower, initiating update process...")
		// initiate update process
		resp, err := cl.Get("http://localhost:7000/launcher-download")
		handleError(err)

		if resp.StatusCode != 200 {
			log.Fatalf("expected status code 200, got %d", resp.StatusCode)
		}

		fh, err := os.OpenFile(launcherFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0744)
		handleError(err)

		_, err = io.Copy(fh, resp.Body)
		handleError(err)

		_ = resp.Body.Close()
		_ = fh.Close()
	} else {
		fmt.Println("launcher is at the most recent version")
	}

	fmt.Printf("Hello! This is version '%s'\n", version)

	fmt.Scanln()
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getCurrentLauncherVersion(filename string) string {
	cmd := exec.Command(filename, "version")
	out, _ := cmd.Output()
	return string(out)
}