### Run app
```
    $ make run
```

### Url
GET: http://localhost:3000/account/{id}
POST: http://localhost:3000/account/

### Setup postgres databese docker container
```
    // set up a postgres database named pgdb with a password 123456(user default is postgres)
    $ docker run --name pgdb -p 5432:5432 -e  POSTGRES_PASSWORD=123456 -d postgres
```

### How to start and stop postgres database
```
    $ docker start pgdb
    $ docker stop pgdb
```