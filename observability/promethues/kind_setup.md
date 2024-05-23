

## Kind Create Cluster
```
kind create cluster --name kind-prometheus-cluster
```

### Helm 

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install kube-prometheus-stack \
  --create-namespace \
  --namespace kube-prometheus-stack \
  prometheus-community/kube-prometheus-stack
```
helm uninstall  --namespace kube-prometheus-stack kube-prometheus-stack



## Acme APP 

```
git clone -b dkalani-dev3  https://github.com/vmwarecloudadvocacy/acme_fitness_demo.git

kubectl create ns observability

docker pull gcr.io/vmwarecloudadvocacy/acmeshop-cart:1.0.0
docker pull gcr.io/vmwarecloudadvocacy/acmeshop-catalog:rel1
docker pull gcr.io/vmwarecloudadvocacy/acmeshop-front-end:rel1
docker pull gcr.io/vmwarecloudadvocacy/acmeshop-order:1.0.1
docker pull gcr.io/vmwarecloudadvocacy/acmeshop-payment:1.0.0
docker pull gcr.io/vmwarecloudadvocacy/acmeshop-user:1.0.0
``` 

## References

1. [Promethus-Kubernetes Installation](https://spacelift.io/blog/prometheus-kubernetes)
2. [Promethus-Kind Cluster Installation](https://medium.com/@giorgiodevops/kind-install-prometheus-operator-and-fix-missing-targets-b4e57bcbcb1f)
