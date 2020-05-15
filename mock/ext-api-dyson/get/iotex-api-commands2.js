/// Iotex API Mock
/// See:
/// curl "http://localhost:3347/mock/iotex-api/staking/validators?status=bonded"
/// curl "http://localhost:3347/mock/iotex-api/accounts/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "https://{iotex_rpc}/v1/staking/validators?status=bonded"
/// curl "https://{iotex_rpc}/v1/accounts/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "http://localhost:8437/v2/iotex/staking/validators"
/// curl "http://localhost:8437/v2/iotex/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m?Authorization=Bearer"

module.exports = {
    path: "/mock/iotex-api/:command1/:command2?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'accounts':
                switch(params.command2) {
                    case 'io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m':
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fiotex-api%2Faccounts%2Fio1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m.json";
                        var json = require(fn);
                        return json;

                    case 'io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5':
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fiotex-api%2Faccounts%2Fio1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5.json";
                        var json = require(fn);
                        return json;
                }
                break;

            case 'staking':
                switch(params.command2) {
                    case 'validators':
                        // status=bonded
                        var fn = "../../ext-api-data/get/" +
                            "mock%2Fiotex-api%2Fstaking%2Fvalidators%3Fstatus%3Dbonded.json";
                        var json = require(fn);
                        return json;
                }
                break;
        }

        return {error: "Not implemented"};
    }
};

