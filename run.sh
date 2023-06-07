mkdir logs

docker stop servicio-app-movil && docker rm servicio-app-movil

docker image rm servicio-app-movil

docker build -t servicio-app-movil .

docker run -d \
--restart always \
--name servicio-app-movil \
--net=upla \
-p 8892:80 \
-v $(pwd)/logs:/etc/app-movil \
servicio-app-movil