## Create Kind Cluster with 3 Node( 1 Master Node & 2 Worker Node)

```
kind create cluster --config multinode_kind_cluster.yaml
kubectl get nodes
docker ps
```

```
$ kubectl get nodes
NAME                 STATUS   ROLES           AGE     VERSION
kind-control-plane   Ready    control-plane   3m17s   v1.27.3
kind-worker          Ready    <none>          2m55s   v1.27.3
kind-worker2         Ready    <none>          2m57s   v1.27.3
$  docker ps                            
CONTAINER ID   IMAGE                  COMMAND                  CREATED         STATUS         PORTS                       NAMES
25e9536d4d05   kindest/node:v1.27.3   "/usr/local/bin/entr…"   4 minutes ago   Up 4 minutes                               kind-worker
45a732c8587e   kindest/node:v1.27.3   "/usr/local/bin/entr…"   4 minutes ago   Up 4 minutes   127.0.0.1:65235->6443/tcp   kind-control-plane
8122ddcc7970   kindest/node:v1.27.3   "/usr/local/bin/entr…"   4 minutes ago   Up 4 minutes                               kind-worker2

```
## Install Nginx pods and services

```
kubectl apply -f foo-service.yaml
kubectl apply -f bar-service.yaml
kubectl get pods
kubectl get svc
```

```
$ docker exec -it kind-control-plane bash
$$> root@kind-control-plane:/# curl 10.96.112.166
bar
$$> root@kind-control-plane:/# curl 10.96.162.119
foo
```


## MetalLB 

```
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml 
kubectl wait --namespace metallb-system \
    --for=condition=ready pod \
    --selector=app=metallb \
    --timeout=180s
kubectl apply -f metallb-config.yaml
```