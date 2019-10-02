# portunus
Portunus was the ancient Roman god of keys, doors, livestock and ports. It is also a small web application that will test for connectivity to host/port combinations.  


# Using in openshift/k8s as a testing tool
Deploy using the provided deploymentConfig.yaml in the k8s library. 
To move it to a different node use 
```
NODE=node-1.kbutton.lab.upshift.rdu2.redhat.com && \
oc rollout pause dc/portunus && \
oc patch dc portunus -p '{"spec":{"template":{"spec":{"nodeSelector":{"kubernetes.io/hostname": "'$NODE'"}}}}}' && \
oc set env dc/portunus MY_K8_NODE=$NODE  && \
oc rollout resume dc/portunus
```