package controllers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spur-dev/api/models"
)

/*Cache Utils*/
func encodeToBytes(v interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		log.Fatalln("Could not encode to bytes")
		log.Fatal(err)
	}
	return buf.Bytes()
}

func decodeToBytes(s []byte) (models.Entry, error) {
	e := models.Entry{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&e)

	return e, err
}

func getPreviewLink(vid string) string {
	return fmt.Sprintf("www.github.com/spur-dev/v/%s", vid)
}

func newVideErrorResponse(err error) []byte {
	res, _ := json.Marshal(models.MsgResponse{Msg: err.Error()})
	return res

}

func getInvalidVideoResponse(vid string) []byte {
	res, _ := json.Marshal(models.MsgResponse{Msg: fmt.Sprintf("Invalid Video Id %s", vid)})
	return res
}

func getCurrentTimestamp() int {
	return int(time.Now().UnixMilli())
}
