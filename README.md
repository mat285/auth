auth
======

`auth` is an auth manager for go! It provides management of user auth via various login methods and authentication methods. 

# Getting Started

See `example/main.go` for a basic example of how the package works using basic creds for login and JWT for authentication.

# Persistence

In order to configure persistence you just need to pass something satisfying the `Persist` interface. This will be used by the auth Manager to retrieve and store identities.

# Login Methods

Current login methods supported are:

- Basic (user/pass)


# Auth Methods

Current auth methods supported are:

- JWT