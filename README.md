### PHP Mimic DNS Server
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

## Notes
- RFC 8855 is a subdomain _acme-challenge.domain.com we'll mimic as a route /_acme-challenge
> https://datatracker.ietf.org/doc/html/rfc8555/#section-8.4
