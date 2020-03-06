bug:
	cd js && npm i
	cd js && node server.js &
	cd go && go run main.go &
	cd go2 && go run main.go