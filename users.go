package main

const (
	APP_NAME								= "ouw-users"
	APP_SALT                = "the world is your oyster"
	APP_VERSION							= "1.0"
)


const (
	DEFAULT_CONF_FILE				= "./config.json"
)


type UserCredentials struct {
	Email						string				`redis:"email"`
	Name            string        `redis:"name"`
	Password				string				`redis:"password"`
	Mobile					string				`redis:"mobile"`
	Salt            string        `redis:"salt"`
	Token           string        `redis:"token"`
	Expiration      string        `redis:"expiration"`
	Creation        string        `redis:"creation"`
}
