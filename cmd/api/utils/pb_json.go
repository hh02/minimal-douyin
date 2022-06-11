package utils

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

var jsonpbMarshaler = &jsonpb.Marshaler{
	EmitDefaults: true,
	OrigName:     true,
}

type PbJSON struct {
	Data proto.Message
}

func (r PbJSON) Render(w http.ResponseWriter) (err error) {
	if err = WritePbJSON(w, r.Data); err != nil {
		panic(err)
	}
	return nil
}

func WritePbJSON(w http.ResponseWriter, m proto.Message) error {
	var bf bytes.Buffer
	if err := jsonpbMarshaler.Marshal(&bf, m); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(bf.Bytes())
	return err
}

func (r PbJSON) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func PbJSONResponse(c *gin.Context, code int, m proto.Message) {
	c.Render(code, PbJSON{
		Data: m,
	})
}
