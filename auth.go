package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	//"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/onceuponworld/ouw-sdk"
)


func generateSalt() string {

	now := time.Now().String()

	x := rand.Intn(DEFAULT_RANGE)

	return fmt.Sprintf("%s%d", now, x)

} // generateSalt


func hash(p string, s string, l int) string {

	digest := hmac.New(sha256.New, []byte(s + APP_SALT))

	digest.Write([]byte(p))

	h := hex.EncodeToString(digest.Sum(nil))

	return h[:l]

} // hash


func generateToken(p string) string {

  salt := generateSalt()

	//return hash(p, salt, DEFAULT_HASH_SIZE)

	return salt

} // generateToken


func encrypt(s string) string {
	return s
} // encrypt


func decrypt(s string) string {
	return s
} // decrypt


func authHandler(w http.ResponseWriter, r *http.Request) {

	user 		:= r.FormValue(API_PARAM_USER)
	pass 		:= r.FormValue(API_PARAM_PASS)
		
	log.Println(user)
	log.Println(pass)

	if !ouwsdk.CheckStr(user, DEFAULT_USER_LENGTH) || !ouwsdk.CheckStr(
		pass, DEFAULT_PASS_LENGTH) {

			w.WriteHeader(http.StatusBadRequest)

	} else {

		switch r.Method {
		case http.MethodPost:
	
			email 	:= r.FormValue(API_PARAM_EMAIL)
			
			if !ouwsdk.CheckStr(email, DEFAULT_EMAIL_LENGTH) {
				w.WriteHeader(http.StatusBadRequest)
			} else {
	
				if ouwsdk.CheckExists(ouwsdk.KEY_USERS, user) {
					
					ouwsdk.SendErr(ERR_USER_EXISTS, w)
					return

				}

				if ouwsdk.CheckExists(ouwsdk.KEY_EMAILS, email) {
					
					ouwsdk.SendErr(ERR_EMAIL_EXISTS, w)
					return

				}

				salt := generateSalt()
	
				log.Println(hash(pass, salt, DEFAULT_HASH_SIZE))
	
			}
	
		default:
			log.Println("http method not supported ", r.Method)
		}
	
	}	

} // authHandler