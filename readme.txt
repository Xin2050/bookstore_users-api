nodemon --exec go run src/main.go --signal SIGTERM
lsof -i tcp:3000
kill -9 xxxx
