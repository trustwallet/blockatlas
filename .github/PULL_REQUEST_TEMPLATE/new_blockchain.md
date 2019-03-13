# Adding a new blockchain

Please follow this checklist:

 - [ ] `config_sample.yml`
   - [ ] Add coin to the platforms list (commented out):
     `#  - mycoin`
   - Add your configuration block to the bottom of the file
 - [ ] `/platform/mycoin/source.go`
    - Get information from a JSON-RPC or HTTP API
 - [ ] `/platform/mycoin/api.go`
    - Public API routes
    - [ ] `GET /mycoin/tx/:id`:
      Detailed information about a transaction.
      Must be compatible with `models/api.go`
    - [ ] `GET /mycoin/account/:addr/txs`:
      List of transactions at an address.
      Must be compatible with `models/api.go`
    - Feel free to add more useful calls.
 - [ ] `/loaders.go`
    - Register your platform in the loaders map
