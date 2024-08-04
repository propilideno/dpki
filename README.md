## PHP Mimic DNS Server
- Server Port: 9000

#### GET /_acme-challenge/?domain_name=propi.dev

response 200
```javascript
{
    "TXT": "qjugIKmy1qcD+I/ZJbtKsog3/n1oqfP6hYMMUy3AgFnrSBt2B8Ms9Dzw4DHryxZ0UY64G9FQCwnlGxYppR55pWsBYcfIHgAOJZdJzndxO3v5633isgDyzZ7KQ6Ttw9otl0gqM1MgRdHHkGfYCff0oYblVLpoL+ESDqXP/2vfYgs=" //base64 encoded string
}
```

#### POST /_acme-challenge

body
```javascript
{
    "domain_name": "propi.dev" //string
    "signature": "qjugIKmy1qcD+I/ZJbtKsog3/n1oqfP6hYMMUy3AgFnrSBt2B8Ms9Dzw4DHryxZ0UY64G9FQCwnlGxYppR55pWsBYcfIHgAOJZdJzndxO3v5633isgDyzZ7KQ6Ttw9otl0gqM1MgRdHHkGfYCff0oYblVLpoL+ESDqXP/2vfYgs=" //base64 encoded string
}
```

## Go Blockchain
GET /certificate?name=base64-encoded-string
response
```javascript
{
    "status": bool //active or inactive
    "issued_at": datetime //iso 8601 format
    "expiration": datetime //iso 8601 format
    "revoke_token": string
}
```
POST /certificate/new
```javascript
{
    "certificate": string, //base64 encoded certificate
}
```
POST /certificate/revoke
```javascript
{
    "certificate": string, //base64encoded
    "encrypted_revoke_token": string //encrypt revoke_token with your private key
}
```

### Notes
- RFC 8855 is a subdomain _acme-challenge.domain.com we'll mimic as a route /_acme-challenge
> https://datatracker.ietf.org/doc/html/rfc8555/#section-8.4
