goa + gorm sample application
====


- [goa](https://github.com/goadesign/goa) : for API design
- [gorma](https://github.com/goadesign/gorma) : for model design
- [qor admin](https://github.com/qor/admin) ; for admin UI
- [swagger ui](https://github.com/swagger-ui) : for API reference

Setup
----

```sh
$ dep ensure
```

Install swagger-ui

```sh
$ make swagger-ui
```

Modify db configuration
----

edit db.go


Start server
----

```sh
$ make run
```

Open in web browser
----

- Swagger UI: http://localhost:8000/swaggerui
- Admin: http://localhost:8000/admin



Development
----

generate files

```sh
make all
```