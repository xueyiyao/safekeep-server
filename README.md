# safekeep-server

## Installing Requirements

```console
foo@bar:~$ go build
```

## Running the server

```console
foo@bar:~$ CompileDaemon -command="./safekeep"
```

## Migrating the database

```console
foo@bar:~$ go run migrate/migrate.go
```

## Admin Login (Deprecated)

To build and debug APIs, logging in as the admin user with the url below will return a cookie that allows access to the rest of Safekeep's API. If you are using Postman, the cookie should be automatically saved so that you may start accessing the other APIs immediately.

```
curl --location --request POST '{url}/admin/login' \
--header 'id: {admin-id}' \
--header 'Authorization: {auth}'
```

## Testing Locally

To run tests, simply execute the following command, which will use a throwaway postgres docker container to run all unit tests.

```
./test.sh
```
