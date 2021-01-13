package bounce

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/client"
)

const (
	httpScheme = "http"
	ipfsScheme = "ipfs"
)

type Client struct {
	client.Request
}

func InitClient(url string) *Client {
	c := Client{internal.InitClient(url)}
	return &c
}

func (c Client) getCollections(owner string, chainId int) ([]Collection, error) {
	query := url.Values{
		"address":  {owner},
		"chain_id": {strconv.Itoa(chainId)},
	}
	var resp CollectionResponse
	err := c.Get(&resp, "nft", query)
	if err != nil {
		return nil, err
	}
	return resp.Data.Collections, nil
}

func (c Client) getCollectibles(owner string, collectionID string, chainId int) ([]Collectible, error) {
	query := url.Values{
		"user_addr":     {owner},
		"contract_addr": {collectionID},
		"chain_id":      {strconv.Itoa(chainId)},
	}

	var resp CollectibleResponse
	err := c.Get(&resp, "erc721", query)
	if err != nil {
		return nil, err
	}
	return resp.Data.Collectibles, err
}

func fetchTokenURI(uri string) (info CollectionInfo, err error) {
	url, err := url.Parse(uri)
	if err != nil {
		return
	}

	var c client.Request
	if strings.HasPrefix(url.Scheme, httpScheme) {
		c = client.InitClient(uri)
	} else if strings.HasPrefix(url.Scheme, ipfsScheme) {
		c = client.InitClient(ipfsGatewayUrl(url))
	} else {
		return info, errors.New("not supported url scheme: " + url.Scheme)
	}

	err = c.Get(&info, "", nil)
	return
}

func normalizeImageUrl(uri string) string {
	url, err := url.Parse(uri)
	if err != nil {
		return uri
	}
	if url.Scheme != ipfsScheme {
		return uri
	}
	return ipfsGatewayUrl(url)
}

func ipfsGatewayUrl(url *url.URL) string {
	return fmt.Sprintf("https://ipfs.io/ipfs/%s%s", url.Host, url.Path)
}
