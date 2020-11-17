package filecoin

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlatform_CurrentBlockNumber(t *testing.T) {
	chainHead, err := mock.JsonFromFilePathToString("mocks/ChainHead.json")
	assert.Nil(t, err)

	data := make(map[string]func(http.ResponseWriter, *http.Request))
	data["/"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, chainHead); err != nil {
			panic(err)
		}
	}

	server := httptest.NewServer(mock.CreateMockedAPI(data))
	defer server.Close()

	p := Init(server.URL)
	block, err := p.CurrentBlockNumber()
	assert.Nil(t, err)
	assert.Equal(t, int64(243590), block)
}

func TestPlatform_GetBlockByNumber(t *testing.T) {
	data := make(map[string]func(http.ResponseWriter, *http.Request))
	data["/"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		type Request map[string]interface{}
		var p Request

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			panic(err)
		}

		resp, ok := p["method"]
		if !ok {
			panic("bad json request")
		}
		var d string

		switch resp {
		case "Filecoin.ChainGetTipSetByHeight":
			chainHead, err := mock.JsonFromFilePathToString("mocks/ChainGetTipSetByHeight.json")
			if err != nil {
				panic(err)
			}
			d = chainHead
		case "Filecoin.ChainGetBlockMessages":
			blockMsg, err := mock.JsonFromFilePathToString("mocks/ChainGetBlockMessages.json")
			if err != nil {
				panic(err)
			}
			d = blockMsg
		}

		if _, err := fmt.Fprint(w, d); err != nil {
			panic(err)
		}
	}

	server := httptest.NewServer(mock.CreateMockedAPI(data))
	defer server.Close()

	p := Init(server.URL)
	block, err := p.GetBlockByNumber(243590)
	assert.Nil(t, err)
	raw, err := json.Marshal(block)
	assert.Nil(t, err)
	assert.Equal(t, wantedResponse, string(raw))
}

const wantedResponse = `{"number":243590,"txs":[{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f029223","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246187,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f026582","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246188,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f029084","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246189,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f026721","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246190,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f027694","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246191,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f028389","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246192,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f027138","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246193,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f028945","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246194,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f028250","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246195,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f027555","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246196,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f026860","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246197,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f028806","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246198,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f026999","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246199,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f028528","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246200,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f13sb4pa34qzf35txnan4fqjfkwwqgldz6ekh5trq","to":"f028667","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":246201,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}},{"id":"","coin":461,"from":"f1i5kvaeurfv27jncddyvcuzgegm4j46y5u7okcza","to":"f2mgv6khl6s6oukeyi3nm3ja67s7fl3cecy2stvny","fee":"0","date":1605614100,"block":243590,"status":"completed","sequence":2,"type":"transfer","memo":"","metadata":{"value":"0","symbol":"FIL","decimals":18}}]}`
