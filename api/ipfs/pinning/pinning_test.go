package pinning

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/jmtx1020/go_quicknode/client"
)

const (
	charset = "0123456789"
	length  = 6
)

func TestCreatePinnedObject(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/pinning")

	pinningAPI := &PinningAPI{API: apiWrapper}

	randomStr := randomString(length)

	_, err := pinningAPI.CreatePinnedObject(
		"QmWTqpfKyPJcGuWWg73beJJiL6FrCB5yX8qfcCF4bHanes",
		fmt.Sprintf("testing_%s.png", randomStr),
		[]string{
			"/ip4/123.12.113.142/tcp/4001/p2p/SourcePeerId",
			"/ip4/123.12.113.114/udp/4001/quic/p2p/SourcePeerId",
		},
		PinnedObjectPayloadMeta{
			Test:      "test_metadata",
			MoreValue: map[string]interface{}{"location": "/home"},
		},
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetAllPinnedObjects(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/pinning")

	pinningAPI := &PinningAPI{API: apiWrapper}

	_, err := pinningAPI.GetAllPinnedObjects(1, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
