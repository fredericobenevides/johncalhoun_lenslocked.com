create() {
docker run --name postgres_dev -v "$PWD/code/pgdata":/var/lib/postgresql/data -d postgres:11
}

connect() {
docker exec -it postgres_dev psql -U postgres
}

$1
