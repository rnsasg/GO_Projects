
## prometheus.yml
```
global:
  scrape_interval: 30s 
scrape_configs:
  # Prometheus itself
  # This uses the static method to get metrics endpoints
  - job_name: "prometheus"
    honor_labels: true
    static_configs:
      - targets: ["prometheus:9090"]

```
## Prometheus Docker Installtion 

[Ref](https://www.linkedin.com/pulse/running-grafana-prometheus-docker-stephen-townshend/)
```
docker network create grafana-prometheus
docker pull prom/prometheus:latest

docker run --rm --name my-prometheus --network grafana-prometheus --network-alias prometheus --publish 9090:9090 --volume /Users/kumarro/go/src/GO_Projects/observability/promethues/prometheus.yml:/etc/prometheus/prometheus.yml --detach prom/prometheus

docker container ls 
```

## Grafana Docker Installation

```
docker pull grafana/grafana-oss:latest
docker run --rm --name grafana --network grafana-prometheus --network-alias grafana --publish 3000:3000 --detach grafana/grafana-oss:latest
```

## Custome Docker Image Build

```
FROM prom/prometheus
ADD prometheus.yml /etc/prometheus/
```

```
docker build -t my-prometheus .
docker run -p 9090:9090 my-prometheus
```


### Connecting Grafana to Prometheus
To display Prometheus metrics in Grafana:
* Log into Grafana (default username "admin" and password "admin").
* Go to Configuration > Data sources
* Click Add data source
* Pick Prometheus as the data type
* Set the URL as http://prometheus:9090 and click Save & test down the bottom.