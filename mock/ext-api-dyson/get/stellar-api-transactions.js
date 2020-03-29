/// Stellar API Mock, transactions
/// See:
/// curl "http://localhost:3000/stellar-api/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029?"
/// curl "https://horizon.stellar.org/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029?"
/// curl http://localhost:8420/v1/stellar/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX
module.exports = {
    path: "/stellar-api/transactions/:txid?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.txid === '2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029') {
            return JSON.parse(`
                {
                    "_links": {
                        "self": {
                            "href": "https://horizon.stellar.org/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029"
                        },
                        "account": {
                            "href": "https://horizon.stellar.org/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX"
                        },
                        "ledger": {
                            "href": "https://horizon.stellar.org/ledgers/27569463"
                        },
                        "operations": {
                            "href": "https://horizon.stellar.org/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029/operations{?cursor,limit,order}",
                            "templated": true
                        },
                        "effects": {
                            "href": "https://horizon.stellar.org/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029/effects{?cursor,limit,order}",
                            "templated": true
                        },
                        "precedes": {
                            "href": "https://horizon.stellar.org/transactions?order=asc\u0026cursor=118409941953355776"
                        },
                        "succeeds": {
                            "href": "https://horizon.stellar.org/transactions?order=desc\u0026cursor=118409941953355776"
                        }
                    },
                    "id": "2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029",
                    "paging_token": "118409941953355776",
                    "successful": true,
                    "hash": "2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029",
                    "ledger": 27569463,
                    "created_at": "2020-01-03T00:26:37Z",
                    "source_account": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
                    "source_account_sequence": "25002129911447568",
                    "fee_charged": 100,
                    "max_fee": 100,
                    "operation_count": 1,
                    "envelope_xdr": "AAAAANSEpQq63M02LHthmlMK1zL6lF7mvGyAj9qoryiJWAQrAAAAZABY004AAAAQAAAAAAAAAAAAAAABAAAAAQAAAADUhKUKutzNNix7YZpTCtcy+pRe5rxsgI/aqK8oiVgEKwAAAAEAAAAAw39tzVrRv+h/KdOCPU/anCP8+XqTbL8JYLsW/X04eVIAAAAAAAAEjCc5UAAAAAAAAAAAA4lYBCsAAABAM4j0FWLKI8qMUcJMO3tP4rMAoIDZdbRlFjiFNACVoAd9dCtdgBBpbi6ZnXtimE370ar19q/qrbvOKBJrnYXjC2b459kAAABAtIedy2ER3ep90HE5rEPwh03iyQRp644Sm8/wK0c9sk6qPSTG2VEFEj85Jk0TrXLxDmxvnktEhPyec67ZcdxfDrEQC4sAAABAoabKlFbeUmDTtAiGVOHthzrCZHK8mpyFkKIS7QzjL4a6ASQ6RweWNUKedUBj4uuO5z1vHDr3I5EiLUhK8xceAQ==",
                    "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAA=",
                    "result_meta_xdr": "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",
                    "fee_meta_xdr": "AAAAAgAAAAMBoYtyAAAAAAAAAADUhKUKutzNNix7YZpTCtcy+pRe5rxsgI/aqK8oiVgEKwCFi6si9J0GAFjTTgAAAA8AAAAEAAAAAQAAAADs3GnuWZXu1qpB+qhI/62WqM4dZ4A1pBgWqEOQJARFlAAAAAAAAAAFd29ybGQAAAABAwMDAAAABAAAAABYBYcJn4H9yxCzrVqI0n0el4b8AUafg9sxvcgJsRALiwAAAAEAAAAAc7yNUAU1YcAb84D5ufCOZyVNpEeCSNA5jaXDCRlL7xoAAAABAAAAALr2jV2dUnqsx1WUvs9TdYuzLargokVQGQ09R9I8gUuYAAAAAQAAAADVvMwB64q5vJdS6D9efKPUpiS/5xJS0JK8BzMjZvjn2QAAAAEAAAAAAAAAAAAAAAEBpK03AAAAAAAAAADUhKUKutzNNix7YZpTCtcy+pRe5rxsgI/aqK8oiVgEKwCFi6si9JyiAFjTTgAAAA8AAAAEAAAAAQAAAADs3GnuWZXu1qpB+qhI/62WqM4dZ4A1pBgWqEOQJARFlAAAAAAAAAAFd29ybGQAAAABAwMDAAAABAAAAABYBYcJn4H9yxCzrVqI0n0el4b8AUafg9sxvcgJsRALiwAAAAEAAAAAc7yNUAU1YcAb84D5ufCOZyVNpEeCSNA5jaXDCRlL7xoAAAABAAAAALr2jV2dUnqsx1WUvs9TdYuzLargokVQGQ09R9I8gUuYAAAAAQAAAADVvMwB64q5vJdS6D9efKPUpiS/5xJS0JK8BzMjZvjn2QAAAAEAAAAAAAAAAA==",
                    "memo_type": "none",
                    "signatures": [
                        "M4j0FWLKI8qMUcJMO3tP4rMAoIDZdbRlFjiFNACVoAd9dCtdgBBpbi6ZnXtimE370ar19q/qrbvOKBJrnYXjCw==",
                        "tIedy2ER3ep90HE5rEPwh03iyQRp644Sm8/wK0c9sk6qPSTG2VEFEj85Jk0TrXLxDmxvnktEhPyec67ZcdxfDg==",
                        "oabKlFbeUmDTtAiGVOHthzrCZHK8mpyFkKIS7QzjL4a6ASQ6RweWNUKedUBj4uuO5z1vHDr3I5EiLUhK8xceAQ=="
                    ]
                }
            `);
        }
        if (params.txid === '23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c') {
            return JSON.parse(`
                {
                    "_links": {
                        "self": {
                            "href": "https://horizon.stellar.org/transactions/23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c"
                        },
                        "account": {
                            "href": "https://horizon.stellar.org/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX"
                        },
                        "ledger": {
                            "href": "https://horizon.stellar.org/ledgers/27364210"
                        },
                        "operations": {
                            "href": "https://horizon.stellar.org/transactions/23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c/operations{?cursor,limit,order}",
                            "templated": true
                        },
                        "effects": {
                            "href": "https://horizon.stellar.org/transactions/23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c/effects{?cursor,limit,order}",
                            "templated": true
                        },
                        "precedes": {
                            "href": "https://horizon.stellar.org/transactions?order=asc\u0026cursor=117528387031003136"
                        },
                        "succeeds": {
                            "href": "https://horizon.stellar.org/transactions?order=desc\u0026cursor=117528387031003136"
                        }
                    },
                    "id": "23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c",
                    "paging_token": "117528387031003136",
                    "successful": true,
                    "hash": "23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c",
                    "ledger": 27364210,
                    "created_at": "2019-12-20T23:06:36Z",
                    "source_account": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
                    "source_account_sequence": "25002129911447567",
                    "fee_charged": 100,
                    "max_fee": 100,
                    "operation_count": 1,
                    "envelope_xdr": "AAAAANSEpQq63M02LHthmlMK1zL6lF7mvGyAj9qoryiJWAQrAAAAZABY004AAAAPAAAAAAAAAAAAAAABAAAAAQAAAADUhKUKutzNNix7YZpTCtcy+pRe5rxsgI/aqK8oiVgEKwAAAAEAAAAAjVtG3WHzbrg6RH0YlwE3W14G9J9NgMT6Jso9zriqGg8AAAAAAAAkKXhESIAAAAAAAAAAA2b459kAAABA5EcZFubvZDa2IytO64bCB85UpjzLqHIS3FBWshoigfice819SwpBYkk7xa3kIWseL/kiQBCnrIAsU7ujYAR3AIlYBCsAAABA1oF6s5NMUEJas97SMRqHgeuXxCdKetSEPBtFGlgzSPTqG6mzAGPuwibKs3HaBjdmfbwh7ffTtjAAesLHIApdCRlL7xoAAABAkADolN0l8llCFxrkqMLbbU4NNKwPsQDrYDoxERo7uvuVJgg18zMBk4t1eqXqQfiYZVeIdAN/q9zcL2SGGYJaBg==",
                    "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAA=",
                    "result_meta_xdr": "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",
                    "fee_meta_xdr": "AAAAAgAAAAMBn66qAAAAAAAAAADUhKUKutzNNix7YZpTCtcy+pRe5rxsgI/aqK8oiVgEKwCFr9SbOOXqAFjTTgAAAA4AAAAEAAAAAQAAAADs3GnuWZXu1qpB+qhI/62WqM4dZ4A1pBgWqEOQJARFlAAAAAAAAAAFd29ybGQAAAABAwMDAAAABAAAAABYBYcJn4H9yxCzrVqI0n0el4b8AUafg9sxvcgJsRALiwAAAAEAAAAAc7yNUAU1YcAb84D5ufCOZyVNpEeCSNA5jaXDCRlL7xoAAAABAAAAALr2jV2dUnqsx1WUvs9TdYuzLargokVQGQ09R9I8gUuYAAAAAQAAAADVvMwB64q5vJdS6D9efKPUpiS/5xJS0JK8BzMjZvjn2QAAAAEAAAAAAAAAAAAAAAEBoYtyAAAAAAAAAADUhKUKutzNNix7YZpTCtcy+pRe5rxsgI/aqK8oiVgEKwCFr9SbOOWGAFjTTgAAAA4AAAAEAAAAAQAAAADs3GnuWZXu1qpB+qhI/62WqM4dZ4A1pBgWqEOQJARFlAAAAAAAAAAFd29ybGQAAAABAwMDAAAABAAAAABYBYcJn4H9yxCzrVqI0n0el4b8AUafg9sxvcgJsRALiwAAAAEAAAAAc7yNUAU1YcAb84D5ufCOZyVNpEeCSNA5jaXDCRlL7xoAAAABAAAAALr2jV2dUnqsx1WUvs9TdYuzLargokVQGQ09R9I8gUuYAAAAAQAAAADVvMwB64q5vJdS6D9efKPUpiS/5xJS0JK8BzMjZvjn2QAAAAEAAAAAAAAAAA==",
                    "memo_type": "none",
                    "signatures": [
                        "5EcZFubvZDa2IytO64bCB85UpjzLqHIS3FBWshoigfice819SwpBYkk7xa3kIWseL/kiQBCnrIAsU7ujYAR3AA==",
                        "1oF6s5NMUEJas97SMRqHgeuXxCdKetSEPBtFGlgzSPTqG6mzAGPuwibKs3HaBjdmfbwh7ffTtjAAesLHIApdCQ==",
                        "kADolN0l8llCFxrkqMLbbU4NNKwPsQDrYDoxERo7uvuVJgg18zMBk4t1eqXqQfiYZVeIdAN/q9zcL2SGGYJaBg=="
                    ]
                }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
