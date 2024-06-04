all: windows linux

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/recipetomd-amd64.exe .

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/recipetomd-linux64 .
