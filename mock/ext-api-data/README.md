# Test data files

Responses for mock tests are stored in data files.
The mock uses these files to return canned responses.

The data files can be updated, by reading and storing responses from the real external APIs.

## File structure

**Subfolder**
The subfolder is either **get** or **post** that corresponds to the HTTP method.

**Filename**
The filename must be the URL path of the call, in the mock.  The path is URL encoded (special characters would casue trouble in the filename).

Example:

Filename:  `mock%2Ftron-api%2Fv1%2Fassets%2F1002798.json`

Corresponding unescaped path: `mock/tron-api/v1/assets/1002798`

Corresponding real path: `https://api.trongrid.io/v1/assets/1002798`

provided the mapping exists in urlmap.yaml `mock/tron-api` to `https://api.trongrid.io`.

## List files

Use testdatatool to list data files and info about them:  `testdatatool list .`

## Update data files

Use testdatatool to update a data file (or all):

`testdatatool update get/mock%2Ftron-api%2Fv1%2Fassets%2F1002798.json`

`testdatatool updateall .`
