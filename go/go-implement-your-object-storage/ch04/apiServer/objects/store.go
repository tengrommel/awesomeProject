package objects

import (
	"../locate"
	"awesomeProject/go/go-implement-your-object-storage/lib/utils"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}
	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusInternalServerError, e
	}
	reader := io.TeeReader(r, stream)
	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, request=%s", d, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}
