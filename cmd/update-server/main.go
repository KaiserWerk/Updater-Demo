package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	appVersion = "1.0.1"
	launcherVersion = "1.0.3"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/app-version", appVersionHandler)
	router.HandleFunc("/app-download", appDownloadHandler)
	router.HandleFunc("/launcher-version", launcherVersionHandler)
	router.HandleFunc("/launcher-download", launcherDownloadHandler)
	fmt.Println("Update server listening on :7000...")
	log.Fatal(http.ListenAndServe(":7000", router))
}

func launcherVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(launcherVersion))
}

func launcherDownloadHandler(w http.ResponseWriter, r *http.Request) {
	bp, _ := filepath.Abs(".")
	filename := filepath.Join(bp, "data", fmt.Sprintf("launcher_v%s.exe", launcherVersion))
	fmt.Println("launcher file:", filename)
	if !fileExists(filename) {
		fmt.Println("file does not exist")
		w.WriteHeader(500)
		w.Write([]byte("nope"))
		return
	}

	cont, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file file: " + err.Error())
		w.WriteHeader(500)
		w.Write([]byte("nope"))
		return
	}

	w.Write(cont)
}

func appVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(appVersion))
}

func appDownloadHandler(w http.ResponseWriter, r *http.Request) {
	bp, _ := filepath.Abs(".")
	filename := filepath.Join(bp, "data", fmt.Sprintf("app_v%s.exe", appVersion))
	fmt.Println("app file:", filename)
	if !fileExists(filename) {
		fmt.Println("file does not exist")
		w.WriteHeader(500)
		w.Write([]byte("nope"))
		return
	}

	cont, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file file: " + err.Error())
		w.WriteHeader(500)
		w.Write([]byte("nope"))
		return
	}

	w.Write(cont)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
