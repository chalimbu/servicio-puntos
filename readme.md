initial crud from https://www.youtube.com/watch?v=Yk5ZjKq4qDQ


insert into public.users ("id","username","password","role","subscribe","points") values(1,'sebastian',null,null,null,5);

locally run
 go run cmd/main.go

como generar la imagen local 
https://docs.docker.com/language/golang/build-images/

https://blog.logrocket.com/dockerizing-go-application/

https://gobyexample.com/environment-variables
run locally DB_PATH=postgres://pg:pass@localhost:5432/crud go run cmd/main.go
