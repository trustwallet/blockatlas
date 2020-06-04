# Lending API

## /lending/providers

Get Lending Providers Info.
Supported assets are also included (without full details).

Sample request: 
`"http://localhost:8420/v1/lending/providers"`

Sample Response:
```
{
    "docs": [
        {
            "id": "compound",
            "info": {
                "id": "compound",
                "description": "Compound Decentralized Finance Protocol",
                "image": "https://compound.finance/images/compound-logo.svg",
                "website": "https://compound.finance"
            },
            "type": "lending",
            "assets": [
                {
                    "symbol": "WBTC",
                    "description": "Compound Wrapped BTC",
                    "apy": 0.19609075835655262,
                    "yield_period": 0,
                    "yield_freq": 15,
                    "total_supply": "9561.87521399",
                    "minimum_amount": "0",
                    "meta_info": {}
                },
                {
                    "symbol": "USDC",
                    "description": "Compound USD Coin",
                    "apy": 1.8766239760801242,
                    "yield_period": 0,
                    "yield_freq": 15,
                    "total_supply": "1011833595.74210543",
                    "minimum_amount": "0",
                    "meta_info": {}
                },
                {
                    "symbol": "ETH",
                    "description": "Compound Ether",
                    "apy": 0.007986621792568961,
                    "yield_period": 0,
                    "yield_freq": 15,
                    "total_supply": "15500928.19070463",
                    "minimum_amount": "0",
                    "meta_info": {}
                },
                {
                    "symbol": "DAI",
                    "description": "Compound Dai",
                    "apy": 0.38386329729567875,
                    "yield_period": 0,
                    "yield_freq": 15,
                    "total_supply": "1257534130.63337750",
                    "minimum_amount": "0",
                    "meta_info": {}
                }
            ]
        }
    ]
}
```

## /lending/assets/:provider/:asset

Get Asset infos, with yield rates and full info.

Sample requests:
```
"http://localhost:8420/v1/lending/assets/compound"
"http://localhost:8420/v1/lending/assets/compound/USDC"
"http://localhost:8420/v1/lending/assets/compound/0x39aa39c021dfbae8fac545936693ac917d5e7563"
```

Provider: ID of the provider, e.g. "compound".  Mandatory.

Asset: Optional.  If missing, all assets are returned.  Can be symbol or token address.

Sample Response:
```
{
    "docs": [
        {
            "symbol": "USDC",
            "description": "Compound USD Coin",
            "apy": 1.8766419718848717,
            "yield_period": 0,
            "yield_freq": 15,
            "total_supply": "1011832110.29512977",
            "minimum_amount": "0",
            "meta_info": {
                "defi_info": {
                    "asset_token": {
                        "symbol": "USDC",
                        "chain": "ETH"
                    },
                    "technical_token": {
                        "symbol": "cUSDC",
                        "chain": "ETH",
                        "contract_address": "0x39aa39c021dfbae8fac545936693ac917d5e7563"
                    }
                }
            }
        }
    ]
}
```

## /lending/account/:provider

Get Account Contracts.
Alongside the current balance, full asset info is returned as well.

Sample request: 
`"http://localhost:8420/v1/lending/account/compound"`
```
{
    "addresses": [
        "0xf9C659D90663BC4e0F7a8766112fE806bae3b5aE"
    ],
    "assets": ["USDC"]
}
```

Provider: ID of the provider, e.g. "compound".  Mandatory.

Addresses: One or more wallet addresses.  Mandatory.

Assets: One or more assets to filter on.  Optional.

Sample Response:
```
{
    "docs": [
        {
            "address": "0xf9c659d90663bc4e0f7a8766112fe806bae3b5ae",
            "contracts": [
                {
                    "asset": {
                        "symbol": "DAI",
                        "description": "Compound Dai",
                        "apy": 0.38386276144011483,
                        "yield_period": 0,
                        "yield_freq": 15,
                        "total_supply": "1257535161.74798884",
                        "minimum_amount": "0",
                        "meta_info": {
                            "defi_info": {
                                "asset_token": {
                                    "symbol": "DAI",
                                    "chain": "ETH"
                                },
                                "technical_token": {
                                    "symbol": "cDAI",
                                    "chain": "ETH",
                                    "contract_address": "0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"
                                }
                            }
                        }
                    },
                    "current_amount": "4.0000527794"
                }
            ]
        }
    ]
}
```
