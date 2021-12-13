# WSL:
#Go swagger can be installed with the following command:

#go get -u github.com/go-swagger/go-swagger/cmd/swagger
#You can generate the documentation using the command:

#swagger generate spec -o ./swagger.yaml --scan-models
#After running the application:

#go run main.go
#Swagger documentation can be viewed using the ReDoc UI in your browser at http://localhost:9090/docs.
check_install:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

