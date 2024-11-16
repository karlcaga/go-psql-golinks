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

In your system's `hosts` file, add a line like `<SERVER DOMAIN> go`.
For example, if you're hosting this at `go.karlcaga.com` then add
```
go.karlcaga.com go
```

Then go to http://go/golinks to make your browser redirect to this repo.

## Deleting shortlinks

To delete `shortlinks` from your DB, run
```sql
DELETE FROM links WHERE shortlink='<shortlink>';
```

For example, to remove https://go/golinks, run
```sql
DELETE FROM links WHERE shortlink='golinks';
```

# Docker Instructions

Build the container with 
```
docker build -t golinks .
```

Run the container with
``` 
docker run --env-file=.env -d -p 8080:8080 golinks
```

# Kubernetes Instructions

If you want to deploy this app using Kubernetes, apply the following resources
```
kubectl apply -f service.yaml
kubectl apply -f deployment.yaml
kubectl apply -f ingress.yaml
```
An Ingress controller is needed to expose this service outside of the cluster. We use Contour's Ingress controller which can be applied with
```
k get -n projectcontour service envoy -o wide
```
The external IP or hostname of the Ingress Controller may then be used for the A or CNAME record to access the service.
```
kubectl get -n projectcontour service envoy -o wide
```

The pods will not enter a ready state without the environment variable set. To set the environment variable `CONN_STR`, we need to create a secret using Kubernetes
```
kubectl create secret generic conn-str --from-literal=conn-str='<YOUR_CONNECTION_STRING_HERE>'

```
To test this, you can port-forward to your local machine using
```
kubectl port-forward <PODNAME> 8080:8080
```
Check functionality by entering http://localhost:8080/pplox into your web browser.