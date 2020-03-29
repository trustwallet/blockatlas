/// Mock for external Ravencoin API
/// See:
/// curl "http://{Ravencoin rpc}/v2/xpub/xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B?details=txs"
/// curl "http://localhost:3000/ravencoin-api/v2/xpub/xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B?details=txs"
/// curl "http://localhost:8420/v1/ravencoin/xpub/xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B"

module.exports = {
    path: '/ravencoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B",
                        "balance": "299445646",
                        "totalReceived": "2599356720",
                        "totalSent": "2299911074",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 8,
                        "transactions": [
                            {
                                "txid": "5d22e5a84b0eb054b623fdc3caf88a60eaa646dc97922f06d325aa8d6f18c564",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "baa2ee45ad8b9922a9db3fdcc22f8c3b5257dc81fef609800e7bd138748b6977",
                                        "sequence": 4294967294,
                                        "n": 0,
                                        "addresses": [
                                            "RHhtrB8eUZuTp8m8HVEyGidEiSkSz5CRdM"
                                        ],
                                        "isAddress": true,
                                        "value": "299456974",
                                        "hex": "47304402204d44f9873e894314746557e4e7ba74e9436d7fa81986d0386684179d8acb77fe02206ff207862b3badbd41cc2ab584ff334410ab2bd089c6c4a1f6ffa59ce2b5d7da012102d34aee6afc67a1b86789beadf0ef2ed7dd5665776b734c0a4dba727bf532bd59"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "299445646",
                                        "n": 0,
                                        "hex": "76a91469a82a3f4a8441d7eb0805e9c1ca421761eb456588ac",
                                        "addresses": [
                                            "RJurUyAqde3GLkDgkTzAuHHMbdMPMqfBsz"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000000013c420937088dd9b69e5ef38c2c1e544828d222c1e8998d1358a",
                                "blockHeight": 843530,
                                "confirmations": 315623,
                                "blockTime": 1566071064,
                                "value": "299445646",
                                "valueIn": "299456974",
                                "fees": "11328",
                                "hex": "010000000177698b7438d17b0e8009f6fe81dc57523b8c2fc2dc3fdba922998bad45eea2ba000000006a47304402204d44f9873e894314746557e4e7ba74e9436d7fa81986d0386684179d8acb77fe02206ff207862b3badbd41cc2ab584ff334410ab2bd089c6c4a1f6ffa59ce2b5d7da012102d34aee6afc67a1b86789beadf0ef2ed7dd5665776b734c0a4dba727bf532bd59feffffff018e2dd911000000001976a91469a82a3f4a8441d7eb0805e9c1ca421761eb456588ac00000000"
                            },
                            {
                                "txid": "baa2ee45ad8b9922a9db3fdcc22f8c3b5257dc81fef609800e7bd138748b6977",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "0f517fd0d4e86070fdff40e95f5202a1b343f75d3168842feb2e8fd6f1be0ef9",
                                        "n": 0,
                                        "addresses": [
                                            "RJurUyAqde3GLkDgkTzAuHHMbdMPMqfBsz"
                                        ],
                                        "isAddress": true,
                                        "value": "9990400",
                                        "hex": "483045022100d8e3598c4d8829b39433c43d4a4c3c438754e8fa77284e1de606233b954770f2022034eb8b6ab6ca103f613d14ccc38fc6aea7bc8661dbff87b0c576b4712d202b3a012102ee0d81b0e40adde804683388953c9cef6276703a3d26e65f9f3433dff0b82dba"
                                    },
                                    {
                                        "txid": "b9db035506eaeb0de0dbc98a277c29f41c08317515538c79cdc542d759918004",
                                        "vout": 1,
                                        "n": 1,
                                        "addresses": [
                                            "RGG4nwkVgFSpy5wdZVKkKFySUCz8N35B5L"
                                        ],
                                        "isAddress": true,
                                        "value": "89988700",
                                        "hex": "473044022075c8f6c3ce3de19c86bbaa845871e25af82ea40bb3793896e657dccbd350209a022033b983ef8733a74ca8e977e08c4ded3a027a800dedbe8aa65a06bca0dc87d09b012103d66703c2448791bcb6828c060de9c568853aa43a753ed227a1f14d70612d3144"
                                    },
                                    {
                                        "txid": "44172def2b5f799963adf28f0a3129e6f7910c8bc5ad09ad0e8155da24748d25",
                                        "vout": 1,
                                        "n": 2,
                                        "addresses": [
                                            "RHj2wiQ8tDhJZKsNKsf48WzxAj2pVnhEnX"
                                        ],
                                        "isAddress": true,
                                        "value": "199965874",
                                        "hex": "4730440220601d4e0d8bc03cf5064dad0316cf61d996ec507aa92105711a0d480a561d9bbb02200acaf5ff2deed01fca255caeaeb63925097bad740adbfd952ab0df66fbfa8552012102eff42496547e74666a392bf8399072dc289ae1d1e835a1c7068650a76be44709"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "299456974",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a9145c6d05bdca691ab3acfd59ad1a70e55aea1cece388ac",
                                        "addresses": [
                                            "RHhtrB8eUZuTp8m8HVEyGidEiSkSz5CRdM"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "0000000000006c4af5eab913f98e03bb3ff45c8b580b88cbeba39b0748ec6f8c",
                                "blockHeight": 843527,
                                "confirmations": 315626,
                                "blockTime": 1566070929,
                                "value": "299456974",
                                "valueIn": "299944974",
                                "fees": "488000",
                                "hex": "0100000003f90ebef1d68f2eeb2f8468315df743b3a102525fe940fffd7060e8d4d07f510f000000006b483045022100d8e3598c4d8829b39433c43d4a4c3c438754e8fa77284e1de606233b954770f2022034eb8b6ab6ca103f613d14ccc38fc6aea7bc8661dbff87b0c576b4712d202b3a012102ee0d81b0e40adde804683388953c9cef6276703a3d26e65f9f3433dff0b82dba0000000004809159d742c5cd798c53157531081cf4297c278ac9dbe00debea065503dbb9010000006a473044022075c8f6c3ce3de19c86bbaa845871e25af82ea40bb3793896e657dccbd350209a022033b983ef8733a74ca8e977e08c4ded3a027a800dedbe8aa65a06bca0dc87d09b012103d66703c2448791bcb6828c060de9c568853aa43a753ed227a1f14d70612d314400000000258d7424da55810ead09adc58b0c91f7e629310a8ff2ad6399795f2bef2d1744010000006a4730440220601d4e0d8bc03cf5064dad0316cf61d996ec507aa92105711a0d480a561d9bbb02200acaf5ff2deed01fca255caeaeb63925097bad740adbfd952ab0df66fbfa8552012102eff42496547e74666a392bf8399072dc289ae1d1e835a1c7068650a76be447090000000001ce59d911000000001976a9145c6d05bdca691ab3acfd59ad1a70e55aea1cece388ac00000000"
                            },
                            {
                                "txid": "0f517fd0d4e86070fdff40e95f5202a1b343f75d3168842feb2e8fd6f1be0ef9",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "b9db035506eaeb0de0dbc98a277c29f41c08317515538c79cdc542d759918004",
                                        "n": 0,
                                        "addresses": [
                                            "RArA49JQFoQZPZRUJqbnE6tp8fVRmAp8v7"
                                        ],
                                        "isAddress": true,
                                        "value": "10000000",
                                        "hex": "4830450221009370c852500b17d02ad9cc2f539a7b56fa237ccd108ec39b2e914fa9fdcde95e02203a730d6431a5e4ada7580644f0fc68ec57fc4e3519d8656e70790239bf12e6b501210214632a2120dc88ae7624ccb3e63534ae3c1d84aefe9f1b72ee23ba9876e3cbc6"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "9990400",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a91469a82a3f4a8441d7eb0805e9c1ca421761eb456588ac",
                                        "addresses": [
                                            "RJurUyAqde3GLkDgkTzAuHHMbdMPMqfBsz"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000000060b0a773ca5f98d4dfdfab1610d62c85f367aee85768852ba5f6",
                                "blockHeight": 741026,
                                "confirmations": 418127,
                                "blockTime": 1559880464,
                                "value": "9990400",
                                "valueIn": "10000000",
                                "fees": "9600",
                                "hex": "010000000104809159d742c5cd798c53157531081cf4297c278ac9dbe00debea065503dbb9000000006b4830450221009370c852500b17d02ad9cc2f539a7b56fa237ccd108ec39b2e914fa9fdcde95e02203a730d6431a5e4ada7580644f0fc68ec57fc4e3519d8656e70790239bf12e6b501210214632a2120dc88ae7624ccb3e63534ae3c1d84aefe9f1b72ee23ba9876e3cbc6000000000100719800000000001976a91469a82a3f4a8441d7eb0805e9c1ca421761eb456588ac00000000"
                            }
                        ],
                        "usedTokens": 7,
                        "tokens": [
                            {
                                "type": "XPUBAddress",
                                "name": "RJurUyAqde3GLkDgkTzAuHHMbdMPMqfBsz",
                                "path": "m/44'/175'/0'/0/2",
                                "transfers": 3,
                                "decimals": 8,
                                "balance": "299445646",
                                "totalReceived": "309436046",
                                "totalSent": "9990400"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
