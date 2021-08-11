## Marvel Comics API Server

### Installing on Ubuntu
#### Prerequisites
* A running Redis server
#### Installing and running the web API
* Clone the Repo `git clone https://github.com/rhperera/marvel-comic-api.git`
* `sudo apt install golang-go`  
* Add .env file with following properties 
  ```
  PRIVATE_KEY=MARVEL_COMICS_PRIVATE_KEY 
  PUBLIC_KEY=MARVEL_COMICS_PUBLIC_KEY
  CACHE_DOMAIN=REDIS_SERVER_DOMAIN:PORT
  ```
* Run `go get` to install dependencies
* `go run app.go`
