package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

import b64 "encoding/base64"

type Contract struct {
	ContractAddress string
	Price           int
	PublicKey       string
}

func getUserData(data string) (string, error) {
	sDec, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(sDec), nil
}

func handlePrice(w http.ResponseWriter, r *http.Request) {
	ct := Contract{ContractAddress: "0xdeadbeef", Price: 1, PublicKey: "alsdjflsjkdflusldfjasldjf"}
	output, err := json.MarshalIndent(ct, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Fprintf(w, string(output))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("====")
	fmt.Println(r.Form["id"][0])
	fmt.Println("====")
	fmt.Fprintf(w, "给个登录证书")
}

func main() {
	fmt.Println("这是一个服务器")
	server := http.Server{
		Addr: "127.0.0.1:8877",
	}

	http.HandleFunc("/price.json", handlePrice)
	http.HandleFunc("/cert.info", handleLogin)
	server.ListenAndServe()
}
