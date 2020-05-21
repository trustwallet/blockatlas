# Test data files

Responses for mock tests are stored in data files.
The mock uses these files to return canned responses.

The data files can be updated, by reading and storing responses from the real external APIs.

## File list

There is a file `mock/datafiles.yaml` containing API paths and corresponding data files.

Example:

```
- file: mock/ext-api-data/tron-api_v1_assets_1002814.json
  mockURL: /mock/tron-api/v1/assets/1002814
  method: GET
  extURL: https://api.trongrid.io/v1/assets/1002814
```

Fields:

* file: name of response data json file (relative to repository root).
* mockURL: the API path for this call, in the mock server.  Usually starts with '/mock/<service>'.
* method: GET or POST
* extURL: Optional.  The full URL of the external reap API.  Used to refresh the data file, manually during development, or automatically.
* reqFile: In case of POST requests, the json file containing POST request.  Used to select which resposne to return, and when invoking external API.
* reqField: Optional. Some POST requests cannot be matched by full request json matching, because they contain a changing field, typically call id.  In this case one field can be selected (.e.g 'address'), and input is matched by the field only.

