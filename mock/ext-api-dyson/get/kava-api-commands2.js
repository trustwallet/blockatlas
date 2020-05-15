/// Kava API Mock
/// See:
/// curl "http://localhost:3347/mock/kava-api/staking/validators?status=bonded"
/// curl "http://localhost:3347/mock/kava-api/staking/pool"
/// curl "http://localhost:3347/mock/kava-api/minting/inflation"
/// curl "https://{kava_rpc}/staking/validators?status=bonded"
/// curl "https://{kava_rpc}/staking/pool"
/// curl "https://{kava_rpc}/minting/inflation"
/// curl "http://localhost:8437/v2/kava/staking/validators"
/// curl "http://localhost:8437/v2/kava/staking/delegations/kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m?Authorization=Bearer"

module.exports = {
    path: "/mock/kava-api/:command1/:command2?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'minting':
                switch(params.command2) {
                    case 'inflation':
                        // status=bonded
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fkava-api%2Fminting%2Finflation.json";
                        var json = require(fn);
                        return json;
                    }

            case 'txs': {
                if (query["transfer.recipient"] === 'kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fkava-api%2Ftxs%3Flimit%3D25%26page%3D1%26transfer.recipient%3Dkava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m.json";
                    var json = require(fn);
                    return json;
                }
        
                if (query["message.sender"] === 'kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fkava-api%2Ftxs%3Flimit%3D25%26message.sender%3Dkava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m%26page%3D1.json";
                    var json = require(fn);
                    return json;
                }        
            }

            case 'staking':
                switch(params.command2) {
                    case 'pool':
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fkava-api%2Fstaking%2Fpool.json";
                        var json = require(fn);
                        return json;

                    case 'validators':
                        // status=bonded
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fkava-api%2Fstaking%2Fvalidators%3Fstatus%3Dbonded.json";
                        var json = require(fn);
                        return json;
                }
        }

        return {error: "Not implemented"};
    }
};

