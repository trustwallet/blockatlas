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

## Add a new file

To add a new test data file:

1. Start with the full external API URL, such as https://api.trongrid.io/v1/assets/1002798

2. If not exist yet, add mapping to a local mock URL prefix, such as mock/tron-api -- https://api.trongrid.io.
Add the mapping to urlmap.yaml.  Please observer the style of the mappings.

3. Derive the mock URL, such as mock/tron-api/v1/assets/1002798

4. Derive the filename, such as get/mock%2Ftron-api%2Fv1%2Fassets%2F1002798.json

5. Retrieve the response, and store it in the file.

6. Make sure the test uses this filem, and add the file to git repo.

For steps 3-5 you can use the testdatatool, like this:

```
$ go run testdatatool.go add https://api.trongrid.io/v1/assets/1002798 get
Filename:   mock%2Ftron-api%2Fv1%2Fassets%2F1002798.json
Mock URL:   mock/tron-api/v1/assets/1002798
Real URL:   https://api.trongrid.io/v1/assets/1002798
Response file written, 410 bytes, url https://api.trongrid.io/v1/assets/1002798, file ./get/mock%2Ftron-api%2Fv1%2Fassets%2F1002798.json
```

## Update data files

Use testdatatool to update a data file (or all):

`testdatatool update get/mock%2Ftron-api%2Fv1%2Fassets%2F1002798.json`

`testdatatool updateall .`
