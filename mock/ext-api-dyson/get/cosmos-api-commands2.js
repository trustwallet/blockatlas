/// Cosmos API Mock
/// See:
/// curl "http://localhost:3347/mock/cosmos-api/txs?limit=25&page=1&transfer.recipient=cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
/// curl "http://localhost:3347/mock/cosmos-api/txs?limit=25&message.sender=cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq&page=1"
/// curl "http://localhost:3347/mock/cosmos-api/staking/validators?status=bonded"
/// curl "http://localhost:3347/mock/cosmos-api/staking/pool"
/// curl "http://localhost:3347/mock/cosmos-api/minting/inflation"
/// curl "https://{cosmos_rpc}/staking/validators?status=bonded"
/// curl "https://{cosmos_rpc}/staking/pool"
/// curl "https://{cosmos_rpc}/minting/inflation"
/// curl "http://localhost:8437/v1/cosmos/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
/// curl "http://localhost:8437/v2/cosmos/staking/validators"
/// curl "http://localhost:8437/v2/cosmos/staking/delegations/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq?Authorization=Bearer"

module.exports = {
    path: "/mock/cosmos-api/:command1/:command2?",
    template: function(params, query) {
        switch(params.command1) {
            case 'minting':
                switch(params.command2) {
                    case 'inflation':
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fcosmos-api%2Fminting%2Finflation.json";
                        var json = require(fn);
                        return json;
                }

            case 'txs': {
                if (query["transfer.recipient"] === 'cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fcosmos-api%2Ftxs%3Flimit%3D25%26page%3D1%26transfer.recipient%3Dcosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq.json";
                    var json = require(fn);
                    return json;
                }
        
                if (query["message.sender"] === 'cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fcosmos-api%2Ftxs%3Flimit%3D25%26message.sender%3Dcosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq%26page%3D1.json";
                    var json = require(fn);
                    return json;
                }        
            }
            
            case 'staking':
                switch(params.command2) {
                    case 'pool':
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fcosmos-api%2Fstaking%2Fpool.json";
                        var json = require(fn);
                        return json;

                    case 'validators':
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fcosmos-api%2Fstaking%2Fvalidators%3Fstatus%3Dbonded.json";
                        var json = require(fn);
                        return json;
                }
        }

        return {error: "Not implemented"};
    }
};

