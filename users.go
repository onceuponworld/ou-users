package main

const (
	APP_NAME								= "ouw-users"
	APP_SALT                = "the world is your oyster"
	APP_VERSION							= "1.0"
)


const (
	DEFAULT_CONF_FILE				= "./config.json"
	DEFAULT_HASH_SIZE				= 48
	DEFAULT_RANGE						= 5000
)


const (
	DEFAULT_EMAIL_LENGTH    = 8
	DEFAULT_PASS_LENGTH			= 8
	DEFAULT_USER_LENGTH			= 2
)


const (
	API_PARAM_EMAIL					= "email"
	API_PARAM_PASS          = "pass"
	API_PARAM_USER					= "user"
)


const (
	ERR_EMAIL_EXISTS				= "Email already exists"
	ERR_USER_EXISTS					= "Username already exists"
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
