/// Kava API Mock
/// See:
/// curl "http://localhost:3000/kava-api/staking/delegators/kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m/delegations"
/// curl "https://{kava_rpc}/staking/delegators/kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m/delegations"
/// curl "http://localhost:8420/v2/kava/staking/delegations/kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m?Authorization=Bearer"

module.exports = {
    path: "/kava-api/:command1/:command2/:arg3/:arg4?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'staking':
                switch(params.command2) {
                    case 'delegators':
                        switch (params.arg3) {
                            case 'kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m':
                                switch (params.arg4) {
                                    case 'delegations':
                                    case 'unbonding_delegations':
                                        return {height: "1793219", result: []};
                                }
                        }
                }
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};

