/// Mock for external coin API, transactions
/// See
/// curl https://vethor-pubnode.digonchain.com/blocks/best
/// curl https://vethor-pubnode.digonchain.com/transactions/0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7
/// curl http://localhost:3000/vechain-api/blocks/best
/// curl http://localhost:3000/vechain-api/transactions/0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7
/// curl "http://localhost:8420/v1/vechain/0xB5e883349e68aB59307d1604555AC890fAC47128"

module.exports = {
    path: '/vechain-api/:entity/:id?',
    template: function(params, query, body) {
        //console.log(params)
        if (params.entity === 'blocks') {
            if (params.id === 'best') {
                return JSON.parse(`
                    {
                        "number": 5466405,
                        "id": "0x00536925d12b64746d8ca1d75db659b0688da13191f709304c88d8372de53519",
                        "size": 1021,
                        "parentID": "0x00536924049d49e0f4605b72f313ccac653ee54e2523f515113db3998ac30cc5",
                        "timestamp": 1585149040,
                        "gasLimit": 35201963,
                        "beneficiary": "0x5643143716537c9c86c558091d2a30710f71fec7",
                        "gasUsed": 184188,
                        "totalScore": 526861845,
                        "txsRoot": "0x81c1361fff2a3cb206f6ecd94d644220d11d91dde4295a0a7b9b0f0e549b9008",
                        "txsFeatures": 1,
                        "stateRoot": "0xcbf49e247ca8bb1826cf101480c05c1470a54c4fdc10702932ac7613ea297d81",
                        "receiptsRoot": "0x0e7dfcee0f0e1fcc1b6888e6eb9f767b6f37aaa5614efcee0faf94e6a62bcb6d",
                        "signer": "0x2da258cae01aac5cd0e4bd876f708081f78b327d",
                        "isTrunk": true,
                        "transactions": [
                            "0x3985abfc8963bc7adfc54d534f0a89498aa986bb49ff87b4eb3d0aa47aeefa87",
                            "0xeedebfba7af9d6a99561f1c02509ce792eb6a2b582f0e35ea4dc571f9ae2306c",
                            "0xbb386e125bfb9bd81e32af6a4b2ef30ff8f89461a94f7be2b3a42135b490562a",
                            "0x49c295e28b017cf337325ea7d6e112ab11afde0433c65c8e71778e21b5acc7d9"
                        ]
                    }
                `);
            }
        }
        if (params.entity === 'transactions') {
            switch (params.id) {
                case '0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7':
                    return JSON.parse(`
                        {
                            "id": "0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7",
                            "chainTag": 74,
                            "blockRef": "0x004313a393a18efb",
                            "expiration": 720,
                            "clauses": [
                                {
                                    "to": "0x2c7a8d5cce0d5e6a8a31233b7dc3dae9aae4b405",
                                    "value": "0x12b1815d00738000",
                                    "data": "0x"
                                }
                            ],
                            "gasPriceCoef": 0,
                            "gas": 21000,
                            "origin": "0xb5e883349e68ab59307d1604555ac890fac47128",
                            "delegator": null,
                            "nonce": "0x8cff29df64a414f8",
                            "dependsOn": null,
                            "size": 129,
                            "meta": {
                                "blockID": "0x004313a4bd4286e821b684cc1749deb3df12fa2a8114435fbd35baa155e82016",
                                "blockNumber": 4395940,
                                "blockTimestamp": 1574410670
                            }
                        }
                    `);

                case '0x004aa0448e458105b098aea2a764a1d54ab95451bee488869f417b351857c3c5':
                    return JSON.parse(`
                        {
                            "id": "0x004aa0448e458105b098aea2a764a1d54ab95451bee488869f417b351857c3c5",
                            "chainTag": 74,
                            "blockRef": "0x0042249a647f63e7",
                            "expiration": 720,
                            "clauses": [
                                {
                                    "to": "0x00bae5ed35736e4ef17af1be0c6f50e0fb73d685",
                                    "value": "0x38f6ea18e810b6f00",
                                    "data": "0x"
                                }
                            ],
                            "gasPriceCoef": 0,
                            "gas": 21000,
                            "origin": "0xb5e883349e68ab59307d1604555ac890fac47128",
                            "delegator": null,
                            "nonce": "0x6fa76caac4c18bd",
                            "dependsOn": null,
                            "size": 130,
                            "meta": {
                                "blockID": "0x0042249bee56223e0ed7a9c7fcfffe8e61b9fd95d29d24843c558ff2c46ea094",
                                "blockNumber": 4334747,
                                "blockTimestamp": 1573795570
                            }
                        }
                    `);    
            }
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
