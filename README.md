# ccap.live-api

## Encryption process
Password => PBKDF2 with salt saved on server => AES encryption

## Dev
For ease of use, docker is used to dev with hot reloading

http://localhost:8080/playground

### Schema
```shell script
go run github.com/99designs/gqlgen init
# Update
go run github.com/99designs/gqlgen
```

## Ressources
- [Go modules](https://levelup.gitconnected.com/switch-to-go-modules-from-go-dep-fcdd4aa41bd5)