Simple User Agent grabber written in Go.

# Docker Usage

```
docker build . -t useragentverifier

docker run -p 3000:3000 -d useragentverifier
```

# Helm Deployment

Edit values.yaml to set nodePort.

```
cd user-agent-verifier
helm install . --values values.yaml
```

# Usage without Expect

## Create a User Agent grabbing session:

GET /create

Response :

```json
{
  "id": "RandomIDString",
  "agent": ""
}
```

## Caputure User Agent from Target Browser

GET /agent/{id}

Response :

```json
{
  "id": "id",
  "agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:79.0) Gecko/20100101 Firefox/79.0"
}
```

## Get Saved User Agent

GET /verify/{id}

Response :

```json
{
  "id": "id",
  "agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:79.0) Gecko/20100101 Firefox/79.0"
}
```

# Usage with Expect

## Create Expected User Agent

POST /expect/create

Body :

```json
{
  "expected_os": "Intel Mac OS X 10.15",
  "expected_browser": "Firefox",
  "expected_version": "79.0"
}
```

Response :

```json
{
  "id": "id",
  "agent": "",
  "os": "",
  "browser": "",
  "version": "",
  "expected_os": "Intel Mac OS X 10.15",
  "expected_browser": "Firefox",
  "expected_version": "79.0",
  "result": false
}
```

## Capture Agent from Target as before.

GET /agent/{id}
...

## Verify Expected User Agent

GET /expect/verify/{id}

Response :

```json
{
  "id": "id",
  "agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:79.0) Gecko/20100101 Firefox/79.0",
  "os": "Intel Mac OS X 10.15",
  "browser": "Firefox",
  "version": "79.0",
  "expected_os": "Intel Mac OS X 10.15",
  "expected_browser": "Firefox",
  "expected_version": "79.0",
  "result": true
}
```
