/// Iotex API Mock
/// See:
/// curl "http://localhost:3347/mock/iotex-api/actions/addr/io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5?count=25&start=32"
/// curl "http://localhost:3347/mock/iotex-api/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "https://{iotex_rpc}/v1/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "http://localhost:8437/v2/iotex/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m?Authorization=Bearer"

module.exports = {
    path: "/mock/iotex-api/:command1/:command2/:arg3/:arg4?",
    template: function(params, query) {
        switch(params.command1) {
            case 'actions':
                switch(params.command2) {
                    case 'addr':
                        switch(params.arg3) {
                            case 'io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5':
                                var fn = "../../ext-api-data/get/" +
                                    "mock%2Fiotex-api%2Factions%2Faddr%2Fio1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5%3Fcount%3D25%26start%3D32.json";
                                var json = require(fn);
                                return json;
                        }
                }
                break;

            case 'staking':
                switch(params.command2) {
                    case 'delegations':
                        switch (params.arg3) {
                            case 'io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m':
                                var fn = "../../ext-api-data/get/" +
                                    "mock%2Fiotex-api%2Fstaking%2Fdelegations%2Fio1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m.json";
                                var json = require(fn);
                                return json;
                        }
                }
                break;
        }

        return {error: "Not implemented"};
    }
};

