# design

restful user authentication and authorization micro-service, provides the following functionality:

* register user (email/password, google and facebook in the future)
* user authentication
* password reset
* delete user
* authorize action

## storage

uses redis to store user information.


### key space

* users (set)
* user:[NAME] (hashmap)

### user credential information

* email
* mobile
* password hash
* salt
* token
* creation date
* support google/facebook in the future

## authentication

essentially an opaque token will be used which carries the user credential information