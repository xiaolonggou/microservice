# WSL:
#Go swagger can be installed with the following command:

#go get -u github.com/go-swagger/go-swagger/cmd/swagger
#You can generate the documentation using the command:

#swagger generate spec -o ./swagger.yaml --scan-models
#After running the application:

#go run main.go
#Swagger documentation can be viewed using the ReDoc UI in your browser at http://localhost:9090/docs.




# trouble-shooting / dubug:
# PS C:\Users\D058009\lawrence\microservice> make swagger
# swagger generate spec -o ./swagger.yaml --scan-models
 #PS C:\Users\D058009\lawrence\microservice> swagger validate swagger.yaml


# The swagger spec at "swagger.yaml" is invalid against swagger specification 2.0.
# See errors below:
# - path param "{id}" has no parameter definition

check_install:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

