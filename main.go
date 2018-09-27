package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"loginserver"
	"net/http"
	"os"
	"time"
)

func handlePrice(w http.ResponseWriter, r *http.Request) {
	ct := loginserver.Contract{ContractAddress: "0xdeadbeef", Price: 1, PublicKey: sr.Get()}
	output, err := json.MarshalIndent(ct, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Fprintf(w, string(output))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//input := r.Form["input"][0]

	// test data
	var testdata loginserver.ClearContent = loginserver.ClearContent{
		Content: loginserver.UserContent{
			TimeStamp: string(time.Now().Unix()),
			Action:    "login",
		},
		PublicKey: sr.Get(),
		Sig:       "0x1239ae258763",
	}

	var err error
	input, err := json.Marshal(&testdata)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(string(input))

	test_result := loginserver.SafeEncoder(sr, string(input))

	cc := loginserver.SafeDecoder(sr, test_result)
	fmt.Println(cc)

	var cc1 loginserver.ClearContent
	json.Unmarshal([]byte(cc), &cc1)

	//result := userManager.getUserInfo("")
	// 使用用户给与的key
	fmt.Fprintf(w, cc)
	//fmt.Fprintf(w, SafeEncoder(result))
}

var sr *loginserver.SafeRsa

func main() {

	var err error
	var publicKey, privateKey []byte

	publicKey, err = ioutil.ReadFile("/Users/wenbinsong/lilin/rsa/rsa_public_key.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("/Users/wenbinsong/lilin/rsa/rsa_private_key.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	sr = loginserver.NewSafeRsa(publicKey, privateKey)

	/*
		var theMsg = "the message you want to encode 你好 世界"

		enc, _ := sr.RsaEncrypt([]byte(theMsg))
		fmt.Println("Encrypted:", string(enc))

		decstr, _ := sr.RsaDecrypt(enc)
		fmt.Println("Decrypted:", string(decstr))
	*/

	fmt.Println("这是一个服务器")
	server := http.Server{
		Addr: "127.0.0.1:8877",
	}

	http.HandleFunc("/price.json", handlePrice)
	http.HandleFunc("/cert.info", handleLogin)
	server.ListenAndServe()

}
