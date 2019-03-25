# Kafka K8s Topic Controller

This Kubernetes controller listen for a custom resource definition and create corresponding Kafka topics.

## Run

```bash
go run main.go \
-kubeconfig=$HOME/.kube/config \
-kafka-brokers=kafka:9092 \
-v=2
```

```
kubectl apply -f example/topic-a.yaml
kubectl apply -f example/topic-b.yaml
```

## Links

https://github.com/kubernetes/sample-controller

https://github.com/openshift-evangelists/crd-code-generation

https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/
