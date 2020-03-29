/// Waves API Mock
/// See:
/// curl "http://localhost:3000/waves-api/transactions/address/3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD/limit/25"
/// curl "https://nodes.wavesnodes.com/transactions/address/3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD/limit/25"
/// curl http://localhost:8420/v1/waves/3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD
module.exports = {
    path: "/waves-api/transactions/address/:address/limit/:limit",
    template: function (params, query, body) {
        //console.log(params)
        if (params.address === '3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD') {
            return [
                [
                    {
                        senderPublicKey: "2UstBx1nMYQ2mPeaJi6tv9oUFCUHLvU1G7nS8Leazsbw",
                        amount: parseFloat("369133368000"),
                        signature: "5Hw5SW7C1EK8hE2YKawFCAWcJoB1Z5NSZFkA65bS3PcDY9w2fi7etPCJDamK2WNb14RWa3BykdT5yFd64SxodjeQ",
                        fee: parseFloat("100000"),
                        type: parseFloat("4"),
                        version: parseFloat("1"),
                        attachment: "",
                        sender: "3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD",
                        feeAssetId: null,
                        proofs: [
                            "5Hw5SW7C1EK8hE2YKawFCAWcJoB1Z5NSZFkA65bS3PcDY9w2fi7etPCJDamK2WNb14RWa3BykdT5yFd64SxodjeQ"
                        ],
                        assetId: null,
                        recipient: "3PLrCnhKyX5iFbGDxbqqMvea5VAqxMcinPW",
                        feeAsset: null,
                        id: "23JGFzBh65fzZArK6KSeRqjjBG5WnQshJJkUv53hCE1E",
                        timestamp: parseFloat("1582527770493"),
                        height: parseFloat("1943922")
                    },
                    {
                        senderPublicKey: "3uT3a9ceebFf6vEa5DmedD5pK6xa1GY7PnvhWPMVyqxC",
                        amount: parseFloat("369133468000"),
                        signature: "3Nrz47KpD3U39Nf7Los23MDoukqXCrCJun1BqnUgjpy9iZLfXc163dTrz4wVvURC2yiULsNYDYA2pxTPGWpBotc5",
                        fee: parseFloat("100000"),
                        type: parseFloat("4"),
                        version: parseFloat("1"),
                        attachment: "Paribu",
                        sender: "3PHYYqBA6ZfBsoGsrXP8r7ZptLXYdGDt1Cm",
                        feeAssetId: null,
                        proofs: [
                            "3Nrz47KpD3U39Nf7Los23MDoukqXCrCJun1BqnUgjpy9iZLfXc163dTrz4wVvURC2yiULsNYDYA2pxTPGWpBotc5"
                        ],
                        assetId: null,
                        recipient: "3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD",
                        feeAsset: null,
                        id: "1456haw7zSTKDmVSbY1njrYdskX7xKbv241c5WJjWhCu",
                        timestamp: parseFloat("1582527244669"),
                        height: parseFloat("1943911")
                    }
                ]
            ]
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
}
;
