# GoLang-WebServer

# Install dependencies

`go mod tidy`

`go get github.com/gin-gonic/gin`

`go get go.mongodb.org/mongo-driver/mongo`

`go get github.com/gin-contrib/cors`

`go get golang.org/x/crypto/bcrypt`

`go mod download golang.org/x/sys`


## Request Flow on API

Request ==> app.go ==> apiRouter.go ==> \<matchingRouter>.go ==> \<matchingController>.go ==> \<matchingModel>.go ==> Response

## To use live-reload during devlopment using nodemon(npm),

Run the following command from repo's root directory

`nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run app.go`# Movie-Review-Web-App
# Movie-Review-Web-App
