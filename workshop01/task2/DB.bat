docker container run  -d  -e DB_PASSWORD=changeit --mount type=volume,src=myvol,dst=/var/lib/mysql  --name mydb stackupiss/northwind-db:v1