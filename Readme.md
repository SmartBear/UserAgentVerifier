Simple User Agent grabber written in Go.

# Usage

## Create a User Agent grabbing session:

GET /create

Response :

```json
{
  "id": "RandomIDString",
  "agent": ""
}
```

## Save User Agent

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
