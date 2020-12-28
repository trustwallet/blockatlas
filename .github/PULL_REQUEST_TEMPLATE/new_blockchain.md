# Adding a new blockchain

Please follow this checklist:

 - `platform/mycoin`
   - [ ] `model.go`: Platform-specific models
   - [ ] `client.go`: Platform getter methods (API, RPC)
   - [ ] `api.go`:
     - Gin route: _GET /mycoin/txs_
     - Getting blockchain info
     - Normalizing platform-specific
       models to BlockAtlas model
   - [ ] `api_test.go`
     - Unit tests
     - Check if normalization from
       response body matches expected model
 - [ ] `test/main.go`:
   Add example address for integration test
 - [ ] `loaders.go`: Add route
 - [ ] `config.yml`: Add default config
   (comment out if no public endpoint)
 - [ ] `config.go`: Add default config

If you have any questions, contact us at
https://t.me/walletcore or @hewigovens / @vikmeup on Telegram.
