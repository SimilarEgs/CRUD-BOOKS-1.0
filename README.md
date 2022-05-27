
## How to use

- Clone this repository
- Execute next in your PostgreSQL client:
```SQL
CREATE TABLE books(
    id SERIAL PRIMARY KEY,
    bookname TEXT,
    author TEXT,
    date INT
);
```
- In a root folder create `.env` file and fill the blank with your postgres url:
``` GOLANG
POSTGRES_URL="postgress_url"
```
- Run `go build` to create the executable file
- Execute generated program
 

