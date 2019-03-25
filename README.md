# Kafka K8s Topic Controller

This Kubernetes controller listen for a custom resource definition and create corresponding Kafka topics.

## Run

```bash
go run main.go \
-kubeconfig=$HOME/.kube/config \
-v=2
```

## Links

https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/
