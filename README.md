# Get KV Consul
Version  v0.0.1

The kv get data value via API consul use Authentication

## Feature
- Basic Auth
- Token
- Store Data
- Parsing Base64
- Loop Update

## How Usage

This can be used by all realtime automation key update values ​​from consul.

### Pull Images Docker
```bash
docker pull fajarhide/getconsul
```
### Kubernetes

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: get-consul
spec:
  replicas: 1
  selector:
    matchLabels:
      app: get-consul
  template:
    metadata:
      labels:
        app: get-consul
    spec:
      containers:
      - name: get-consul
        image: fajarhide/getconsul
        imagePullPolicy: Always
        volumeMounts:
          - name: env-consul
            mountPath: "/app/.env"
            subPath: .env
          - name: shared-data
            mountPath: /directory/store/data
      volumes:
      - name: env-consul
        secret:
          secretName: consul-secret
          items:
          - key: .env
            path: .env
```

secret env-consul is a `.env` get-consul which is saved in kubernetes secrets.


## TODO
- Web GUI
- External Service
- Record response status for each executed