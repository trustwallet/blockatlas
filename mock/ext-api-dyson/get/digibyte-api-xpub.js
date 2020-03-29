/// Mock for external Digibyte API
/// See:
/// curl "http://{digibyte rpc}/v2/xpub/zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow?details=txs"
/// curl "http://localhost:3000/digibyte-api/v2/xpub/zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow?details=txs"
/// curl "http://localhost:8420/v1/digibyte/xpub/zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow"

module.exports = {
    path: '/digibyte-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow",
                        "balance": "2599896040",
                        "totalReceived": "6895667780",
                        "totalSent": "4295771740",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 7,
                        "transactions": [
                            {
                                "txid": "e13c81f4450e43ef8ef10dea8f4f9d53f357266563462df32a7fd61eb71bd190",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "34fac508c701939ac0fcb62bc436281cc28588d7986c5aa51de99e9835b9217b",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "dgb1qxxdszsfkv4uvw8kzzl2wfatts5r9zex69x6076"
                                        ],
                                        "isAddress": true,
                                        "value": "598908696"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "10000000",
                                        "n": 0,
                                        "hex": "0014428a3010eb79c4d32478404200dbd41042e61fce",
                                        "addresses": [
                                            "dgb1qg29rqy8t08zdxfrcgppqpk75zppwv87wknuder"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "588896040",
                                        "n": 1,
                                        "hex": "00146542b8b7859994b18df4b71c72c4a72ed113662e",
                                        "addresses": [
                                            "dgb1qv4pt3du9nx2trr05kuw893989mg3xe3w322vy7"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "000000000000000733a44ab2efd51c9485b05d905d2ba9092f716dcc0083116c",
                                "blockHeight": 9763292,
                                "confirmations": 777936,
                                "blockTime": 1573501957,
                                "value": "598896040",
                                "valueIn": "598908696",
                                "fees": "12656",
                                "hex": "010000000001017b21b935989ee91da55a6c98d78885c21c2836c42bb6fcc09a9301c708c5fa34010000000000000000028096980000000000160014428a3010eb79c4d32478404200dbd41042e61fce28d71923000000001600146542b8b7859994b18df4b71c72c4a72ed113662e02473044022023b050ed97c4c7e334fec8efa3f27754e65c1e66a60ce142b2f8e2c16a1c847b02204a2e7d8a470a8978735a0b6d0f3649da6489612256f7cacd52e95034cc05365e012102cc9d7c2383e8f7c18e5af1852ddc35c5c4355a262a68b9d499ba173f6200ff9500000000"
                            },
                            {
                                "txid": "303aa433530cb32762f8068ae14022e9a759b6c40bb69cee7e77abca333db317",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "d8d05cee8623257853b90cb1b0ff6c8e4bf5509818e812510f6715015fac1599",
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "addresses": [
                                            "DTVYvQYvoXQjsiZQdvT9AaBjyQJr8TjAWu"
                                        ],
                                        "isAddress": true,
                                        "value": "48587940518800",
                                        "hex": "483045022100de758903643db73bc56bcd6b3e1e4bdafb3bd9b7a844286f66317de7c90b868702207eba0eeb32c817df9e79df2aad5f192808941ccb5d18430f9ee7ea18b8b1abad0121038f47ad4221aa24dbefe2b439f644b3af71428a7eb9e837517a162710fe631cf7"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "48585940500900",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a914c30c2dd0e7b7f6506259f11e956d0f1085ca437588ac",
                                        "addresses": [
                                            "DNvQtshdnrehs8MNRx6Mh9XDPz4UZucfEv"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "2000000000",
                                        "n": 1,
                                        "hex": "00147ad8c91046724785f78b48d4ae606a3556ddf9b4",
                                        "addresses": [
                                            "dgb1q0tvvjyzxwfrctautfr22ucr2x4tdm7d5va4lr6"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "311aa13d66a4cf6be602209601d7b54bc715320d9255d1abc822ad20a7e01a20",
                                "blockHeight": 9756845,
                                "confirmations": 784383,
                                "blockTime": 1573405513,
                                "value": "48587940500900",
                                "valueIn": "48587940518800",
                                "fees": "17900",
                                "hex": "01000000019915ac5f0115670f5112e8189850f54b8e6cffb0b10cb95378252386ee5cd0d8000000006b483045022100de758903643db73bc56bcd6b3e1e4bdafb3bd9b7a844286f66317de7c90b868702207eba0eeb32c817df9e79df2aad5f192808941ccb5d18430f9ee7ea18b8b1abad0121038f47ad4221aa24dbefe2b439f644b3af71428a7eb9e837517a162710fe631cf7ffffffff02a481b94b302c00001976a914c30c2dd0e7b7f6506259f11e956d0f1085ca437588ac00943577000000001600147ad8c91046724785f78b48d4ae606a3556ddf9b400000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
