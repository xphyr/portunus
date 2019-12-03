# portunus
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fxphyr%2Fportunus.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fxphyr%2Fportunus?ref=badge_shield)

Portunus was the ancient Roman god of keys, doors, livestock and ports. It is also a small web application that will test for connectivity to host/port combinations.  


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

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fxphyr%2Fportunus.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fxphyr%2Fportunus?ref=badge_large)