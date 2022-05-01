package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spur-dev/api/models"

	"github.com/julienschmidt/httprouter"
)

func (sc SessionController) NewVideoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	vid := GenerateUniqueVideoId(uid)
	res := sc.NewVideo(vid, uid)

	// TODO: Add tests to justify avoding handling this error

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	w.Write(res)
}

func (sc SessionController) GetVideoMetadaHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vid := ps.ByName("vid")

	res, found := sc.GetVideo(vid)

	w.Header().Set("Content-Type", "application/json")
	if found {
		w.WriteHeader(http.StatusOK) // 200
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
	}
	fmt.Fprintf(w, "%s\n", res)
}

func (sc SessionController) UpdateVideoStateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vid := ps.ByName("vid")
	var body models.LambdaResponse
	json.NewDecoder(r.Body).Decode(&body)
	sc.UpdateVideoStatus(vid, body.Status)
	w.WriteHeader(http.StatusOK) // 200
}

func (sc SessionController) CancelRecordingHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vid := ps.ByName("vid")

	// Delete user, TODO: set status deleted
	err := sc.DeleteVideo(vid)

	if err != nil {
		log.Fatalln(err) // TODO: return here with a valid response
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted video", vid, "\n")
}
