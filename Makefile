BINARY_NAME=goer

mac:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} main.go

linux:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go

win:
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-win main.go

run:
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-win