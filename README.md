## Marvel Comics API Server

### Installing on Ubuntu
#### Prerequisites
* A running Redis server
#### Installing and running the web API
* Clone the Repo `git clone https://github.com/rhperera/marvel-comic-api.git`
* Run `sudo apt install golang-go`  to install golang
* Change path to cloned directory and add a .env file with following properties 
  ```
  PRIVATE_KEY=<MARVEL_COMICS_PRIVATE_KEY> 
  PUBLIC_KEY=<MARVEL_COMICS_PUBLIC_KEY>
  CACHE_DOMAIN=<REDIS_SERVER_DOMAIN>:<PORT>
  ```
* Run `go get` to install dependencies
* `go run app.go`

### Running with docker-compose 
This start up the two docker containers, redis and webapp
#### Prerequisites
* Docker engine and docker-compose should be installed
* Clone the Repo `git clone https://github.com/rhperera/marvel-comic-api.git`
* Change path to cloned directory and add a .env file with following properties
  ```
  PRIVATE_KEY=<MARVEL_COMICS_PRIVATE_KEY> 
  PUBLIC_KEY=<MARVEL_COMICS_PUBLIC_KEY>
  CACHE_DOMAIN=redis:6379
  ```
  
* Run `sudo docker-compose up --build`