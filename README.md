# igridmq
MQTT Message Bus for iGrid Nodes Network. It uses eclipse mosquitto as the underlying broker

## run
to run igridmq make sure you have git, docker and docker compose installed, then run the command 

```bash
git clone https://github.com/dataleodev/igridmq.git
cd igridmq
docker-compose up --build
```

to stop the running containers just go to igrid root directory where the file docker-compose.yaml is found
and then run the command
```bash
docker-compose down

```