/// Mock for external Ravencoin API
/// See:
/// curl "http://{Ravencoin rpc}/address/RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS?details=txs"
/// curl "http://localhost:3000/ravencoin-api/address/RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS?details=txs"
/// curl "http://localhost:8420/v1/ravencoin/address/RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS"

module.exports = {
    path: '/ravencoin-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "addrStr": "RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS",
                        "balance": "0",
                        "totalReceived": "10.48",
                        "totalSent": "10.48",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 5,
                        "txs": [
                            {
                                "txid": "fc6c9e9c6cd0253f05e43a12d64154a689115f0103dc0b64957a8cb92b67f306",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "3717b528eb4925461d9de5a596d2eefe175985740b4fda153255e10135f236a6",
                                        "vout": 1,
                                        "sequence": 4294967293,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100b085c0f7d09937eed588f4c18e8ea600382926bcd0467a90436c7d3446ae2aa202206261e0b98c5d4852b17e32bcc04c2d6bc01dad327be0c16776ce9b6f57a44142012102138724e702d25b0fdce73372ccea9734f9349442d5a9681a5f4d831036cd9429"
                                        },
                                        "addresses": [
                                            "RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS"
                                        ],
                                        "value": "0.48"
                                    },
                                    {
                                        "txid": "128793fa1635c630a3463225515f9aabfa6160ae32028d5ae6f09b04802abebe",
                                        "vout": 0,
                                        "sequence": 4294967294,
                                        "n": 1,
                                        "scriptSig": {
                                            "hex": "47304402202f7f64346ddf51e1304e46407e1a285625b1c8e743b2a8ac919520a924be6ca102206a452426f9a1e105b0ba422b321cf358fe8d357cf985f43caf581ec7275f03ce0121030b22ea9d87e08432c18606849e801462189a45aaa9e158306cd5f26bcea3c61e"
                                        },
                                        "addresses": [
                                            "RVtWJKBbG7JN3GQBkgwz4n4GYVPeViNer7"
                                        ],
                                        "value": "5.988166"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "6.4679603",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a9148272c538003476df668581ad1336eecc774f274c88ac",
                                            "addresses": [
                                                "RMAwQqgkYur4un2JbE6vPnZWeyzunXkuWo"
                                            ]
                                        },
                                        "spent": true
                                    }
                                ],
                                "blockhash": "0000000000005cbc77e1a4b921e5de4fc5388fb661065e825877911d2364472a",
                                "blockheight": 753909,
                                "confirmations": 405243,
                                "time": 1560658623,
                                "blocktime": 1560658623,
                                "valueOut": "6.4679603",
                                "valueIn": "6.468166",
                                "fees": "0.0002057",
                                "hex": "0100000002a636f23501e1553215da4f0b74855917feeed296a5e59d1d462549eb28b51737010000006b483045022100b085c0f7d09937eed588f4c18e8ea600382926bcd0467a90436c7d3446ae2aa202206261e0b98c5d4852b17e32bcc04c2d6bc01dad327be0c16776ce9b6f57a44142012102138724e702d25b0fdce73372ccea9734f9349442d5a9681a5f4d831036cd9429fdffffffbebe2a80049bf0e65a8d0232ae6061faab9a5f51253246a330c63516fa938712000000006a47304402202f7f64346ddf51e1304e46407e1a285625b1c8e743b2a8ac919520a924be6ca102206a452426f9a1e105b0ba422b321cf358fe8d357cf985f43caf581ec7275f03ce0121030b22ea9d87e08432c18606849e801462189a45aaa9e158306cd5f26bcea3c61efeffffff01fe528d26000000001976a9148272c538003476df668581ad1336eecc774f274c88ac00000000"
                            },
                            {
                                "txid": "3717b528eb4925461d9de5a596d2eefe175985740b4fda153255e10135f236a6",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "0c7e82b44eec71d634c013e2db3cb4fa26f87fbc90eb8734da93807d23605544",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100d790bdaa3c44eb5e3a422365ca5fc009c4512625222e3378f2f16e7e6ef1732a0220688c1bb995b7ff2f12729e101d7c24b6314430317e7717911fdc35c0d84f2f0d012102138724e702d25b0fdce73372ccea9734f9349442d5a9681a5f4d831036cd9429"
                                        },
                                        "addresses": [
                                            "RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS"
                                        ],
                                        "value": "1"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.5",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a9149451f4546e09fc2e49ef9b5303924712ec2b038e88ac",
                                            "addresses": [
                                                "RNoSGCX8SPFscj8epDaJjqEpuZa2B5in88"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "0.48",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a9145d6e33f3a108bbcc586cbbe90994d5baf5a9cce488ac",
                                            "addresses": [
                                                "RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS"
                                            ]
                                        },
                                        "spent": true
                                    }
                                ],
                                "blockhash": "0000000000004b942df6b9736c0a4b1a320cadfe46ca10351e43b35cccee7be1",
                                "blockheight": 734142,
                                "confirmations": 425010,
                                "time": 1559464388,
                                "blocktime": 1559464388,
                                "valueOut": "0.98",
                                "valueIn": "1",
                                "fees": "0.02",
                                "hex": "0100000001445560237d8093da3487eb90bc7ff826fab43cdbe213c034d671ec4eb4827e0c000000006b483045022100d790bdaa3c44eb5e3a422365ca5fc009c4512625222e3378f2f16e7e6ef1732a0220688c1bb995b7ff2f12729e101d7c24b6314430317e7717911fdc35c0d84f2f0d012102138724e702d25b0fdce73372ccea9734f9349442d5a9681a5f4d831036cd9429ffffffff0280f0fa02000000001976a9149451f4546e09fc2e49ef9b5303924712ec2b038e88ac006cdc02000000001976a9145d6e33f3a108bbcc586cbbe90994d5baf5a9cce488ac00000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
