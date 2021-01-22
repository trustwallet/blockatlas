package blockatlas

import "fmt"

func GenCollectibleId(contract, tokenId string) string {
	return fmt.Sprintf("%s-%s", contract, tokenId)
}
