#!/bin/bash

# From https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-deploy-eck.html

# Install custom resource definitions:
kubectl create -f https://download.elastic.co/downloads/eck/2.9.0/crds.yaml

# Install the operator with its RBAC rules:
kubectl apply -f https://download.elastic.co/downloads/eck/2.9.0/operator.yaml

# Monitor the operator logs:
kubectl -n elastic-system logs -f statefulset.apps/elastic-operator