resource: https://gowebexamples.com/

```sh
go run main.go
```

Creates go.mod for dependency management; replace placeholders with your GitHub username/project name

> go.mod will display: `module github.com/rishi/golang-learn`

```sh
go mod init github.com/<yourname>/<yourproject>
```

For **installing the packages** we type `go get ...` in the same folder where the go.mod is created

using gomux:
https://gowebexamples.com/routes-using-gorilla-mux/
![subrouter gomux](image.png)

run this to start the mysql docker container (NO password needed here); The port maps host to container for access.

```sh
docker run -d --name mysql-container -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql
```

`db.Query` which can query multiple rows, for us to iterate over
`db.QueryRow` in case we only want to query a specific row.
