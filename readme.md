initial crud from https://www.youtube.com/watch?v=Yk5ZjKq4qDQ


para correr en local
docker compose up
y
DB_PATH=postgres://pg:pass@localhost:5432/crud go run cmd/main.go


para correr en kubernetes
kubectl apply -f kubernetes-resources/database.yml
kubectl apply -f kubernetes-resources/deployment.yml


locally run
 go run cmd/main.go

como generar la imagen local 
https://docs.docker.com/language/golang/build-images/

https://blog.logrocket.com/dockerizing-go-application/

https://gobyexample.com/environment-variables
run locally DB_PATH=postgres://pg:pass@localhost:5432/crud go run cmd/main.go


run docker with enviroment variable docker run -e DB_PATH='postgres://pg:pass@localhost:5432/crud' servicio-punto:v0.0.1

