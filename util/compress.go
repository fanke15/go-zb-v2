package util

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"

	. "github.com/fanke15/go-zb-v2/log"
)

func GzipCompress(msg []byte) []byte {
	var gzipBuf bytes.Buffer
	var gzipWrite *gzip.Writer
	gzipWrite = gzip.NewWriter(&gzipBuf)

	_, err := gzipWrite.Write(msg)
	if err != nil {
		Log.Warn().Msg("gzipWrite error:" + err.Error())
		return nil
	}
	gzipWrite.Close()

	return gzipBuf.Bytes()
}

func GzipUnCompress(buf []byte) []byte {
	gzipRead, err := gzip.NewReader(bytes.NewReader(buf))
	if err != nil {
		Log.Warn().Msg("gzip.NewReader error:" + err.Error())
		return nil
	}

	decodeBuf, err := ioutil.ReadAll(gzipRead)
	if err != nil {
		Log.Warn().Msg("ReadAll(gzipRead) error:" + err.Error())
		return nil
	}
	return decodeBuf
}

func UnCompress(message []byte) []byte {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	n, err := base64.StdEncoding.Decode(buf, message)
	if err != nil {
		Log.Warn().Msg("base64 Decode error:" + err.Error())
		Log.Warn().Msg(string(message))
		return nil
	}
	return GzipUnCompress(buf[:n])
}
