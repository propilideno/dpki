# Descentralized Public Key Infrastructure (DPKI)

## Execution

The follow command will run frontend, backend, blockchain and database containers:

```bash
> docker-compose up -d
```

To stop them, just run:

```bash
> docker-compose down
```

## Endpoints Documentation

### PHP Helper
- Server Port: 9000

#### GET /certificate/new

response 200

response 200
```json
{
    "private_key": "example_base_64_private_key",
    "public_key": "example_base_64_public_key",
    "certificate": "example_base_64_certificate"
}
```

#### GET /sign?message=example_message&private_key=example_base_64_private_key

response 200
```json
{
    "sign": "example_base_64_sign"
}
```

#### POST /verify

body
```json
{
    "public_key": "example_base_64_public_key",
    "message": "example_message",
    "sign": "example_base_64_sign",
}
```

response 200
```json
{
    "verify": true,
}
```

#### GET /hash?message=example_message&key=example_key

response 200
```json
{
    "hash": "example_base_64_sha_256_hash",
}
```

### PHP DNS Server
- Server Port: 9000

#### GET /dns/{domain}

response 200
```json
{
    "id": 1,
    "domain": "domain",
    "txt": "example_txt",
    "created_at": "2024-01-01T00:00:00.000000Z",
    "updated_at": "2024-01-01T00:00:00.000000Z"
}
```

#### PATCH /dns/{domain}/clear-txt

header
```json
{
    "hash": "example_base_64_hash" // HMAC Authorization
}
```

response 200

```json
{
    "id": 1,
    "domain": "domain",
    "txt": null,
    "created_at": "2024-01-01T00:00:00.000000Z",
    "updated_at": "2024-01-01T00:00:00.000000Z"
}
```

#### POST /dns

body
```json
{
    "domain": "example_domain",
    "txt": "example_txt"
}
```

header
```json
{
    "hash": "example_base_64_hash", // Required if domain already exists.
}
```

### Go Blockchain
- Server Port: 7000

#### POST /mine/block

body
```json
{
    "wallet": "example_wallet"
}
```

response 200
```json
{
    "block": {
        "data": {
            "contract_execution_history": null,
            "contracts": null,
            "transactions": [
                {
                    "from": "Block Reward",
                    "to": "wallet",
                    "amount": 10,
                    "miner": ""
                }
            ]
        },
        "previous_hash": "",
        "hash": "00d69dad76a96aaeb05a9653762bbb5ad1d81431fcc8d1c8c3063e0c8627d210",
        "timestamp": "2024-08-19T20:15:14.167207882Z",
        "nonce": 313
    },
    "index": 1,
    "message": "New Block Forged"
}
```
#### POST /mine/transaction

body
```json
{
    "wallet": "example_wallet"
}
```

response 404
```json
{
    "message": "No transactions to mine"
}
```

response 200
```json
{
    "message": "Transaction mined successfully"
}
```

#### POST /mine/contract

body
```json
{
    "wallet": "example_wallet"
}
```

response 404
```json
{
    "message": "No transactions to mine"
}
```

response 200
```json
{
    "message": "Contract executed successfully",
    "gas": 0.1
}
```

#### GET /chain

response 200
```json
{
    "chain": [
        {
            "data": {
                "contract_execution_history": null,
                "contracts": null,
                "transactions": [
                    {
                        "from": "Block Reward",
                        "to": "wallet",
                        "amount": 10,
                        "miner": ""
                    }
                ]
            },
            "previous_hash": "",
            "hash": "00d69dad76a96aaeb05a9653762bbb5ad1d81431fcc8d1c8c3063e0c8627d210",
            "timestamp": "2024-08-19T20:15:14.167207882Z",
            "nonce": 313
        },
        {
            "data": {
                "contract_execution_history": null,
                "contracts": null,
                "transactions": [
                    {
                        "from": "wallet",
                        "to": "b",
                        "amount": 1,
                        "miner": "teste"
                    }
                ]
            },
            "previous_hash": "00d69dad76a96aaeb05a9653762bbb5ad1d81431fcc8d1c8c3063e0c8627d210",
            "hash": "8fff5098c68b889b3e3f8a0ead2be7677c79ad19a390584e8bcbfba283192984",
            "timestamp": "2024-08-19T20:35:41.314961941Z",
            "nonce": 0
        }
    ],
    "isValid": false,
    "length": 2,
    "minedCoins": 10
}
```

#### GET /memorypool

response 200
```json
{
    "contractexecutionpool": [
        {
            "contract_id": "example_contract_id",
            "result": "",
            "timestamp": "2024-01-01T00:00:00.218002954Z",
            "miner": ""
        }
    ],
    "transactionpool": [
        {
            "from": "example_wallet_1",
            "to": "example_wallet_2",
            "amount": 1,
            "miner": ""
        }
    ]
}
```

#### GET /info/{wallet}

response 200
```json
{
    "balance": 8.95
}
```

#### GET /certificate/status/{certificate}

response 200
```json
{
    "certificate": "certificate",
    "status": true
}
```

#### POST /transaction/new

body
```json
{
    "from": "example_wallet_1",
    "to": "example_wallet_2",
    "amount": 1
}
```

response 201
```json
{
    "message": "Transaction added to the pool"
}
```

#### POST /contract/execute

body
```json
{
    "contract_id": "example_contract_id"
}
```

header
```json
{
    "Authorization": "example_base_64_sign"
}
```

response 201
```json
{
    "message": "Contract execution added to the pool"
}
```

#### POST /certificate/request

body
```json
{
    "certificate": "example_base_64_certificate",
    "domain": "example_domain"
}
```

response 201
```json
{
    "certificate": "example_base_64_certificate",
    "domain": "example_domain",
    "message": "Certificate successfully added to the current block, complete acme challenge to turn it valid.",
    "wallet": "example_wallet"
}
```

### Notes
- RFC 8855 is a subdomain _acme-challenge.domain.com we'll mimic as a route /contract/execute
> https://datatracker.ietf.org/doc/html/rfc8555/#section-8.4
