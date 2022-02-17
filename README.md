# portunus

[![Build Status](https://travis-ci.com/xphyr/portunus.svg?branch=master)](https://travis-ci.com/xphyr/portunus)
[![Go Reportcard](https://goreportcard.com/badge/github.com/xphyr/portunus)](https://goreportcard.com/report/github.com/xphyr/portunus)
[![Quay.io Status](https://quay.io/repository/xphyr/portunus/status)](https://quay.io/repository/xphyr/portunus)

Portunus was the ancient Roman god of keys, doors, livestock and ports. It is also a small web application that will test for connectivity to host/port combinations.  

## Building

Portunus can be built locally, or be built inside a container image.

# Using in openshift/k8s as a testing tool
Deploy using the provided deploymentConfig.yaml in the k8s library. 
To move it to a different node use 
```
NODE=<node name> && \
oc rollout pause dc/portunus && \
oc patch dc portunus -p '{"spec":{"template":{"spec":{"nodeSelector":{"kubernetes.io/hostname": "'$NODE'"}}}}}' && \
oc set env dc/portunus MY_K8_NODE=$NODE  && \
oc rollout resume dc/portunus
```