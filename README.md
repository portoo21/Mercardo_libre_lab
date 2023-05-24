# Pre-requisites

Install docker and golang

# Run the App

`go build -o main main.go`

# How to get token

`go run scripts/create-token.go <role>`

Where role can be `regular` or `sensitive`

# How to load clients

`go run scripts/add-clients.go <token>`

Where token is generated from the step above, worth mention you need to have **sensitive** role to make it work


# Out of scope

Add audit logs when accessing

# endpoints

- `localhost/clients/`
- - Require regular role
- `localhost/clients/sensitive?id=<client_id>`
- - Require sensitive role

# docker

It should be enough to run docker-composer up