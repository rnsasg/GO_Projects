# SQL Transaction 

Transactions are very useful when you want to perform multiple operations on a database, but still treat them as a single unit. We will go into detail on why this is useful, and how you can use it in your Go applications.

> [!NOTE]
> This is n


## References
* [A Guide On SQL Database Transactions In Go](https://www.sohamkamani.com/golang/sql-transactions/)
* [Github](https://github.com/sohamkamani/golang-sql-transactions)


## Postgress Database 

```
psql -U postgres 
admin123

\list 
\c postgres
\dt 

TRUNCATE TABLE food,pets;
```