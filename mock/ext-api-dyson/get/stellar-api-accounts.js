/// Stellar API Mock, accounts
/// See:
/// curl "http://localhost:3000/stellar-api/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX/payments?order=desc&limit=25"
/// curl "https://horizon.stellar.org/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX/payments?order=desc&limit=25"
/// curl http://localhost:8420/v1/stellar/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX
module.exports = {
    path: "/stellar-api/accounts/:address/:operation?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.operation === 'payments') {
            if (params.address === 'GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX') {
                return JSON.parse(`
                    {
                        "_links": {
                            "self": {
                                "href": "https://horizon.stellar.org/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX/payments?cursor=\u0026limit=25\u0026order=desc"
                            },
                            "next": {
                                "href": "https://horizon.stellar.org/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX/payments?cursor=68651109446586369\u0026limit=25\u0026order=desc"
                            },
                            "prev": {
                                "href": "https://horizon.stellar.org/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX/payments?cursor=118409941953355777\u0026limit=25\u0026order=asc"
                            }
                        },
                        "_embedded": {
                            "records": [
                                {
                                    "_links": {
                                        "self": {
                                            "href": "https://horizon.stellar.org/operations/118409941953355777"
                                        },
                                        "transaction": {
                                            "href": "https://horizon.stellar.org/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029"
                                        },
                                        "effects": {
                                            "href": "https://horizon.stellar.org/operations/118409941953355777/effects"
                                        },
                                        "succeeds": {
                                            "href": "https://horizon.stellar.org/effects?order=desc\u0026cursor=118409941953355777"
                                        },
                                        "precedes": {
                                            "href": "https://horizon.stellar.org/effects?order=asc\u0026cursor=118409941953355777"
                                        }
                                    },
                                    "id": "118409941953355777",
                                    "paging_token": "118409941953355777",
                                    "transaction_successful": true,
                                    "source_account": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
                                    "type": "payment",
                                    "type_i": 1,
                                    "created_at": "2020-01-03T00:26:37Z",
                                    "transaction_hash": "2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029",
                                    "asset_type": "native",
                                    "from": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
                                    "to": "GDBX63ONLLI372D7FHJYEPKP3KOCH7HZPKJWZPYJMC5RN7L5HB4VFXLM",
                                    "amount": "500000.0000000"
                                },
                                {
                                    "_links": {
                                        "self": {
                                            "href": "https://horizon.stellar.org/operations/117528387031003137"
                                        },
                                        "transaction": {
                                            "href": "https://horizon.stellar.org/transactions/23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c"
                                        },
                                        "effects": {
                                            "href": "https://horizon.stellar.org/operations/117528387031003137/effects"
                                        },
                                        "succeeds": {
                                            "href": "https://horizon.stellar.org/effects?order=desc\u0026cursor=117528387031003137"
                                        },
                                        "precedes": {
                                            "href": "https://horizon.stellar.org/effects?order=asc\u0026cursor=117528387031003137"
                                        }
                                    },
                                    "id": "117528387031003137",
                                    "paging_token": "117528387031003137",
                                    "transaction_successful": true,
                                    "source_account": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
                                    "type": "payment",
                                    "type_i": 1,
                                    "created_at": "2019-12-20T23:06:36Z",
                                    "transaction_hash": "23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c",
                                    "asset_type": "native",
                                    "from": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
                                    "to": "GCGVWRW5MHZW5OB2IR6RRFYBG5NV4BXUT5GYBRH2E3FD3TVYVINA72KM",
                                    "amount": "3976053.0000000"
                                }
                            ]
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
