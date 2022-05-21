# Stringsvc1

Microservicio para realizar operaciones con cadenas enviadas como `JSON` a través de peticiones `POST`.

Este código está basado en los tutoriales de [Go-kit](https://gokit.io/examples/stringsvc.html), publicados bajo la licencia `MIT`.

## Formato de parámetros
El microservicio acepta peticiones con un `JSON` que incluya una cadena con clave `s`
```json
{
    "s": "Hello, World!"
}
```


## Operaciones permitidas
`/uppercase`: Devuelve una cadena en mayúsculas a partir de una cadena enviada como entrada.
`/count`: Devuelve el número de caracteres de una cadena enviada como entrada.


## Uso con Dockerfile
`$ docker build -t stringsvc1 .`

`$ docker run -p 8080:8080 stringsvc1`
