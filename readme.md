initial crud from https://www.youtube.com/watch?v=Yk5ZjKq4qDQ


para correr en local(comandos de consola)
```
docker compose up
```

y
este segundo solo funciona si se tiene instalado go(sino toca dentro del kubernetes con los que siguen)
```
DB_PATH=postgres://pg:pass@localhost:5432/crud go run cmd/main.go
```

el crea una tabla con la siguiente estructura
```
CREATE TABLE public.points (
	"IdUser" bigserial NOT NULL,
	points int8 NULL,
	CONSTRAINT points_pkey PRIMARY KEY ("IdUser")
);
```

aun no se como o si se puede setear los otros tipos.

para correr en kubernetes
```
kubectl apply -f kubernetes-resources/database.yml
kubectl apply -f kubernetes-resources/deployment.yml
```

para conectarse y mandarle peticiones en local 
```
kubectl port-forward service/servicio-puntos 4000:4000
```


y responde en localhost:4000
ya despues de esto se puede lanzar/probar las peticiones del postman

se puede cambiar las variables de entorno en kubernetes-resources/deployment.yaml linea 23 la variable de entorno para que se conecte a otra base de datos

---
---otras notas --

locally run
 go run cmd/main.go

como generar la imagen local 
https://docs.docker.com/language/golang/build-images/

https://blog.logrocket.com/dockerizing-go-application/

https://gobyexample.com/environment-variables
run locally DB_PATH=postgres://pg:pass@localhost:5432/crud go run cmd/main.go


run docker with enviroment variable docker run -e DB_PATH='postgres://pg:pass@localhost:5432/crud' servicio-punto:v0.0.1

