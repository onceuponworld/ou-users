package main

import (
	//"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/onceuponworld/ouw-sdk"
	//"github.com/go-redis/redis/v8"
)


var (
	conf = flag.String("conf", DEFAULT_CONF_FILE, "conf file path")
)


//var rds *redis.Client

//var app ouwsdk.AppConfig

//var ctx = context.Background()


func main() {

	flag.Parse()

	ouwsdk.ParseConfig(*conf, &ouwsdk.App, true)

	ouwsdk.Rds = ouwsdk.InitRedis(ouwsdk.App.Store.Host,
		ouwsdk.App.Store.Port)

	http.HandleFunc("/auth", authHandler)

	address := ouwsdk.Addr(ouwsdk.App.Host,
		ouwsdk.App.Port)

	fmt.Printf("Starting %s on %s...\n", APP_NAME, address)

	log.Fatal(http.ListenAndServe(address, nil))

} // main
