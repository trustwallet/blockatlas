package bounce

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/network/middleware"
)

const (
	httpScheme = "http"
	ipfsScheme = "ipfs"
)

type Client struct {
	client.Request
}

func InitClient(url string) *Client {
	c := Client{client.InitClient(url, middleware.SentryErrorHandler)}
	return &c
}

func (c Client) getCollections(owner string) ([]Collection, error) {
	query := url.Values{
		"user_address": {owner},
	}
	var resp CollectionResponse
	err := c.Get(&resp, "/v2/bsc/nft", query)
	if err != nil {
		return nil, err
	}
	return resp.Data.Collections, nil
}

func (c Client) getCollectibles(owner string, collectionID string) ([]Collectible, error) {
	query := url.Values{
		"user_address":     {owner},
		"contract_address": {collectionID},
	}

	var resp CollectibleResponse
	err := c.Get(&resp, "/v2/bsc/erc721", query)
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
		c = client.InitClient(uri, middleware.SentryErrorHandler)
	} else if strings.HasPrefix(url.Scheme, ipfsScheme) {
		c = client.InitClient(ipfsGatewayUrl(url), middleware.SentryErrorHandler)
	} else {
		return info, errors.New("not supported url scheme: " + url.Scheme)
	}

	err = c.Get(&info, "", nil)
	return
}

func normalizeUrl(uri string) string {
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
	components := strings.TrimPrefix(url.String(), ipfsScheme+"://")
	components = strings.TrimPrefix(components, "/")
	components = strings.TrimPrefix(components, "ipfs/")
	return fmt.Sprintf("https://ipfs.io/ipfs/%s", components)
}
