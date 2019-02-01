package objects

import (
	"awesomeProject/go/go-implement-your-object-storage/ch02/apiServer/heartbeat"
	"awesomeProject/go/go-implement-your-object-storage/lib/objectstream"
	"fmt"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectstream.NewPutStream(server, object), nil
}
