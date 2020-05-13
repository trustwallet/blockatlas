/// Cosmos API Mock
/// See:
/// curl "http://localhost:3347/mock/cosmos-api/staking/delegators/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq/delegations"
/// curl "https://{cosmos_rpc}/staking/delegators/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq/delegations"
/// curl "http://localhost:8437/v2/cosmos/staking/delegations/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq?Authorization=Bearer"

module.exports = {
    path: "/mock/cosmos-api/:command1/:command2/:arg3/:arg4?",
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
                                        var fn = "../../ext-api-data/get/" +
                                            "mock%2Fcosmos-api%2Fstaking%2Fdelegators%2Fcosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq%2Fdelegations.json";
                                        var json = require(fn);
                                        return json;
                                }
                        }
                }
        }

        return {error: "Not implemented"};
    }
};

