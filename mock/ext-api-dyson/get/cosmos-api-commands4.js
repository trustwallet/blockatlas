/// Cosmos API Mock
/// See:
/// curl "http://localhost:3000/cosmos-api/staking/delegators/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq/delegations"
/// curl "https://{cosmos_rpc}/staking/delegators/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq/delegations"
/// curl "http://localhost:8420/v2/cosmos/staking/delegations/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq?Authorization=Bearer"

module.exports = {
    path: "/cosmos-api/:command1/:command2/:arg3/:arg4?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'staking':
                switch(params.command2) {
                    case 'delegators':
                        switch (params.arg3) {
                            case 'cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq':
                                switch (params.arg4) {
                                    case 'delegations':
                                    case 'unbonding_delegations':
                                        return JSON.parse(`
                                            {
                                                "height": "1419065",
                                                "result": [
                                                    {
                                                        "delegator_address": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq",
                                                        "validator_address": "cosmosvaloper17h2x3j7u44qkrq0sk8ul0r2qr440rwgjkfg0gh",
                                                        "shares": "2211271.000000000000000000",
                                                        "balance": "2211271"
                                                    }
                                                ]
                                            }
                                        `);
                                }
                        }
                }
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};

