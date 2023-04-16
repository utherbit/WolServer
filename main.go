package main

import (
	"WolServer/util"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var tokensFileDir string
var tokenMap = make(map[string]string)

func main() {
	util.CheckEnvFile()
	util.LookupEnv(&tokensFileDir, "TOKENS_DIR", "input.txt")
	initTokenMap()

	var host, port string
	util.LookupEnv(&host, "LISTEN_HOST", "0.0.0.0")
	util.LookupEnv(&port, "LISTEN_PORT", "8087")

	http.HandleFunc("/wakeup/", handlerWol)

	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Listen %s", addr)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	util.PanicIfErr(err)
}

func handlerWol(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.URL.Path, "/wakeup/")

	s, find := tokenMap[token]
	if !find {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := util.SendMagicPacket(s)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Package sent successfully."))
	w.WriteHeader(http.StatusOK)
}

func initTokenMap() {
	file, err := os.Open(tokensFileDir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) > 1 {
			tokenMap[args[0]] = args[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
