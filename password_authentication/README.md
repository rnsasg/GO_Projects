# Password Authentication 

Any application that stores passwords has to ensure that its passwords are stored securely. 
You can’t just store passwords as plain text, and it should ideally be impossible to guess your users passwords, 
even if you have access to their data.

## The Proper Way To Store Passwords

We cannot store our users passwords as we receive them.
We need to transform each password in such a way that it is easy to verify, but not easy to guess. 
We achieve this using a one-way hashing function (which would be the bcrypt algorithm in this case).


During login, we can check if a password is correct for a given username by retrieving the password hash for that username, and using the bcrypt compare function to verifying the given password with the hash:


## HTTP Server Implementation

We are going to build an HTTP server with two routes: /signup and /signin, and use a Postgres DB to store user credentials.
- [x] /signup will accept user credentials, and securely store them in our database.-
- [X] /signin will accept user credentials, and authenticate them by comparing them with the entries in the database.


## MySQL Table 

```
create database pwdauth;
use pwdauth;
create table users( username varchar(256) PRIMARY KEY , password varchar(256));
```