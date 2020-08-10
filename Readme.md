Simple User Agent grabber written in Go.

# Usage
## Create a User Agent grabbing session:  
GET /create

Response :
```json
{
"id" : "RandomIDString",
"agent": ""
}
```
## Save User Agent
GET /agent/{id}

Response :
```json
{
"id" : "id",
"agent": "AgentString"
}
```

## Get Saved User Agent
GET /verify/{id}

Response :
```json
{
"id" : "id",
"agent": "AgentString"
}
```