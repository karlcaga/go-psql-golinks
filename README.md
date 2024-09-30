`go-psql-golinks` is an implementation of the popular `golinks` functionality for sharing URLs as user-defined string using Go for the HTTP server and PostgreSQL for the link database.

# Installation
Clone this repo
```sh
git clone https://github.com/karlcaga/go-psql-golinks.git
```

With your database, run all of the SQL files in `migrations/` sequentially starting from `01_create_table_links.sql`.

Create a `.env` file in the repo root with contents
```
CONN_STR=postgresql://<USER>:<PASSWORD>@<DB_URL>/<DB_NAME>?<OPTIONS>
```
You may be able to get this from your DB provider.

With Go v1.23 installed, run
```bash
go build .
./go-psql-golinks
```

# Usage

Adding and removing `shortlinks` is done directly on the DB with SQL.

## Adding shortlinks

To add shortlinks to your DB, run
```sql
INSERT INTO links (shortlink, url) VALUES ('<shortlink>', '<URL>');
```

For example, to add https://go/golinks to this repo, run
```sql
INSERT INTO links (shortlink, url) VALUES ('golinks', 'https://github.com/karlcaga/go-psql-golinks');
``` 

## Using shortlinks

In your OS' `hosts` file, add a line with your server's domain name and go.
For example, if you're hosting this at `go.karlcaga.com` then add
```
go.karlcaga.com go
```

Then go to http://go/golinks to make your browser redirect to this repo.

## Deleting shortlinks

To delete `shortlinks` from your db, run
```sql
DELETE FROM links WHERE shortlink='<shortlink>';
```

For example, to remove https://go/golinks, run
```sql
DELETE FROM links WHERE shortlink='golinks';
```
