/// Iotex API Mock
/// See:
/// curl "http://localhost:3000/iotex-api/actions/addr/io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5?start=0&count=25"
/// curl "https://pharos.iotex.io/v1/actions/addr/io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5?start=0&count=25"
/// curl http://localhost:8420/v1/iotex/io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5
module.exports = {
    path: "/iotex-api/actions/:action/:address?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.action === 'addr') {
            if (params.address === 'io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5') {
                return JSON.parse(`
                    {
                        "total": "2",
                        "actionInfo": [
                            {
                                "action": {
                                    "core": {
                                        "version": 1,
                                        "nonce": "40",
                                        "gasLimit": "10000",
                                        "gasPrice": "1000000000000",
                                        "transfer": {
                                            "amount": "5010000000000000000000",
                                            "recipient": "io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5"
                                        }
                                    },
                                    "senderPubKey": "BHd42w46LIO7qw5+TsTo+QjAUZBhVc3NPwEVh3IHePcFQQeb9rFSPmEV6LhFS7cqRzDBpgSHuQbelY6wnPvMYmw=",
                                    "signature": "CiJdvAO3jNUX7agpE4AvtrBdEzgFsm1MANhXTXeVtdkXxNrDV8xbSiFz6FuFUTd1L5uaFXoe1mxUPIrbTabh4QE="
                                },
                                "actHash": "7aceeda86535b8dd345e1ab2176a7f12e1907aac3591cf9422a5393feadb6bb6",
                                "blkHash": "57ca5160cb9aac7a16b4e5006295d93857220de966335a2a453992c824037ebe",
                                "blkHeight": "3291297",
                                "sender": "io1r8cwhdec6y3fpv42hv7ak0aqh0cfyc279ladxl",
                                "gasFee": "10000000000000000",
                                "timestamp": "2020-02-14T01:44:30Z"
                            },
                            {
                                "action": {
                                    "core": {
                                        "version": 1,
                                        "nonce": "1",
                                        "gasLimit": "500000",
                                        "gasPrice": "1000000000000",
                                        "execution": {
                                            "amount": "5000000000000000000000",
                                            "contract": "io1zn9mn4v63jg3047ylqx9nqaqz0ev659777q3xc",
                                            "data": "B8NfwAAAAAAAY29yZWRldgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
                                        }
                                    },
                                    "senderPubKey": "BKarIa2N7C/tz6r93fnnHXHEXyGC4aG9fiTYSjwSKpUOWl2efDuLb9Fk+NysVn1hEFgFLjRONu/KlM4TnHaIQLQ=",
                                    "signature": "0Pn96CJXbUmNvXsnvx0RgqPYuac7YD699liT/WqmE9Ak4EHYZXhTozu+t+JrM4cHmt37z7M3R4lbzC77gggHdgA="
                                },
                                "actHash": "7644e425163920f1d5658e12c96d6dfb4961ee35bc9886347fbb339d6b939dd8",
                                "blkHash": "1f7ab328160d0643b7a6af287a1559a88de70cd8a2068d40dcc814fdcdfba996",
                                "blkHeight": "3292065",
                                "sender": "io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5",
                                "gasFee": "264146000000000000",
                                "timestamp": "2020-02-14T02:48:30Z"
                            }
                        ]
                    }
                `);
            }
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
