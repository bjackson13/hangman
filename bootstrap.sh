container_id=$(docker build . | tail -1 | awk '{print $3}')

docker run --net=host -p 8080:8080 -d $container_id