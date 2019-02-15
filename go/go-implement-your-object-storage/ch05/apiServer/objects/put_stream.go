package objects

import (
	"../heartbeat"
	"awesomeProject/go/go-implement-your-object-storage/lib/rs"
	"fmt"
)

func putStream(hash string, size int64) (*rs.RSPutStream, error) {
	server := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(server) != rs.ALL_SHARDS {
		return nil, fmt.Errorf("cannot find enough dataServer")
	}
	return rs.NewRSPutStream(server, hash, size)
}
