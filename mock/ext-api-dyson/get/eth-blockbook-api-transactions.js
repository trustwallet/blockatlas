/// Ethereum Blockbook API Mock
/// See:
/// curl "http://localhost:3000/eth-blockbook-api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://{eth blockbook api}/api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7&details=txs"
/// curl http://localhost:8420/v1/ethereum/0x0875BCab22dE3d02402bc38aEe4104e1239374a7

module.exports = {
    path: '/eth-blockbook-api/v2/address/:address',
    template: function(params, query, body) {
        //console.log(query)
        if (params.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
            return JSON.parse(`{
                "page": 1,
                "totalPages": 11,
                "itemsOnPage": 25,
                "address": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                "balance": "185659745674589722",
                "unconfirmedBalance": "0",
                "unconfirmedTxs": 0,
                "txs": 259,
                "nonTokenTxs": 240,
                "transactions": [
                  {
                    "txid": "0x2a9fd94735e273526a2cde57c6a19b9d488e0c9a960565c3e19be5e12d4b4b47",
                    "vin": [
                      {
                        "n": 0,
                        "addresses": [
                          "0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
                        ],
                        "isAddress": true
                      }
                    ],
                    "vout": [
                      {
                        "value": "17635730000000000",
                        "n": 0,
                        "addresses": [
                          "0x1717f94202c126ef71d6C562de253Fe95eEbDD5f"
                        ],
                        "isAddress": true
                      }
                    ],
                    "blockHash": "0x5b8b39099d025a8feeed7575ebff6d0b2508e0542d6279c1b3b4f8b893e74296",
                    "blockHeight": 9551915,
                    "confirmations": 231227,
                    "blockTime": 1582624428,
                    "value": "17635730000000000",
                    "fees": "90720000000000",
                    "ethereumSpecific": {
                      "status": 1,
                      "nonce": 227,
                      "gasLimit": 21000,
                      "gasUsed": 21000,
                      "gasPrice": "4320000000"
                    }
                  },
                  {
                    "txid": "0xb669b69afee75c6ef073a603600041d3708d54da8d43cab7b35ee66baa7510d3",
                    "vin": [
                      {
                        "n": 0,
                        "addresses": [
                          "0xeCe114137b2e9Dbf29712BDC39639EB0B72B41b8"
                        ],
                        "isAddress": true
                      }
                    ],
                    "vout": [
                      {
                        "value": "0",
                        "n": 0,
                        "addresses": [
                          "0x0D8775F648430679A709E98d2b0Cb6250d2887EF"
                        ],
                        "isAddress": true
                      }
                    ],
                    "blockHash": "0x770332c9b4ea42e518f8c5a0178ca5732fd5edfe5e8e011e1362e6598de262d5",
                    "blockHeight": 9519169,
                    "confirmations": 263973,
                    "blockTime": 1582189159,
                    "value": "0",
                    "fees": "425822000000000",
                    "tokenTransfers": [
                      {
                        "type": "ERC20",
                        "from": "0xeCe114137b2e9Dbf29712BDC39639EB0B72B41b8",
                        "to": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                        "token": "0x0D8775F648430679A709E98d2b0Cb6250d2887EF",
                        "name": "Basic Attention Token",
                        "symbol": "BAT",
                        "decimals": 18,
                        "value": "400000000000000000"
                      }
                    ],
                    "ethereumSpecific": {
                      "status": 1,
                      "nonce": 16,
                      "gasLimit": 51839,
                      "gasUsed": 37028,
                      "gasPrice": "11500000000"
                    }
                  }
                ],
                "nonce": "228",
                "tokens": []
              }`);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
