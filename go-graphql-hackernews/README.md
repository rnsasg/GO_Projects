
## what is GraphQL ?
GraphQL is a query language for APIs and runtime for fulfilling those queries with your existing data. GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.

## What is GraphQL Server ?
A GraphQL server is able to receive requests in GraphQL Query Language format and return response in desired form. GraphQL is a query language for API so you can send queries and ask for what you need and exactly get that piece of data. In this sample query we are looking for address, title of the links and name of the user who added it:

## Schema-Driven Development

In GraphQL your API starts with a schema that defines all your types, queries and mutations, It helps others to understand your API. So it’s like a contract between server and the client. Whenever you need to add a new capability to a GraphQL API you must redefine schema file and then implement that part in your code. GraphQL has its Schema Definition Language for this purpose. gqlgen is a Go library for building GraphQL servers and has a nice feature that generates code based on your schema definition.

### Request 
```
query {
	links{
    	title
    	address,
    	user{
      		name
    	}
  	} 
}
```

### Response 

```
{
  "data": {
    "links": [
      {
        "title": "our dummy link",
        "address": "https://address.org",
        "user": {
          "name": "admin"
        }
      }
    ]
  }
}
```

## What Is A Query ?
A query in graphql is asking for data, you use a query and specify what you want and graphql will return it back to you.

```
query {
	links{
    title
    address,
    user{
      name
    }
  }
}
```

## What Is A Mutation ?

Simply mutations are just like queries but they can cause a data write, Technically Queries can be used to write data too however it’s not suggested to use it. So mutations are like queries, they have names, parameters and they can return data.

```
mutation {
  createLink(input: {title: "new link", address:"http://address.org"}){
    title,
    user{
      name
    }
    address
  }
}
```

## Models and migrations
```
go-graphql-hackernews
--internal
----pkg
------db
--------migrations
----------mysql
```
```
go get -u github.com/go-sql-driver/mysql
go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
cd internal/pkg/db/migrations/
migrate create -ext sql -dir mysql -seq create_users_table
migrate create -ext sql -dir mysql -seq create_links_table
```

```
CREATE TABLE IF NOT EXISTS Links(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
)
```

```
migrate -database mysql://root:dbpass@/hackernews -path internal/pkg/db/migrations/mysql up
```


## References :

1. [graphql-go](https://github.com/graph-gophers/graphql-go)
2. [graphql](https://github.com/graphql-go/graphql)
3. [thunder](https://github.com/samsarahq/thunder)
4. [gqlgen](https://github.com/99designs/gqlgen)
5. [Go Graphql Documentation example](https://www.howtographql.com/graphql-go/0-introduction/)
6. [Github Repo](https://github.com/howtographql/graphql-golang)
7. [Schema-Driven Development](https://graphql.org/learn/schema/)

## go mod download

```go get github.com/99designs/gqlgen/codegen/config@v0.17.47
go get github.com/99designs/gqlgen/internal/imports@v0.17.47
go get github.com/99designs/gqlgen/graphql/handler/transport@v0.17.47
go get github.com/99designs/gqlgen/codegen@v0.17.47
go get github.com/99designs/gqlgen@v0.17.47
go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.47
go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.47

go mod download github.com/google/uuid```
