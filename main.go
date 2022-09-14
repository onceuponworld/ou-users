package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/onceuponworld/ouw-sdk"
)

var (
	conf		= flag.String("conf", DEFAULT_CONF_FILE, "conf file path")
)

var app ouwsdk.AppConfig


func main() {

	flag.Parse()

	ouwsdk.ParseConfig(*conf, &app)

	log.Println(app)
	ouwsdk.InitRedis(ouwsdk.Addr(app.Store.Host, app.Store.Port))

	http.HandleFunc("/auth", authHandler)

	address := ouwsdk.Addr(app.Host, app.Port)

	fmt.Printf("Starting %s on %s...", APP_NAME, address)

	log.Fatal(http.ListenAndServe(address, nil))

} // main
