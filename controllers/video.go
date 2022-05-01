package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spur-dev/api/models"
)

func (sc SessionController) NewVideo(vid string, uid string) []byte {
	// add entry to cache
	ts := getCurrentTimestamp()
	err := sc.CreateCacheEntry(vid, uid, ts)

	if err != nil {
		return newVideErrorResponse(err)
	}
	r := models.NewVideoResponse{
		VID:       vid,
		UID:       uid,
		Timestamp: ts,
	}

	rj, _ := json.Marshal(r)

	return rj
}

func (sc SessionController) GetVideo(vid string) ([]byte, bool) {
	// Get from cache
	ve, err := sc.GetCacheEntry(vid)
	r := models.GetVideoResponse{VID: vid}

	found := true
	if err != nil {
		// TODO: check specifically that error shows video not present in cache
		// Fetch response from db
		ve, err := sc.GetVideoMetadata(vid)

		if err != nil {
			found = false
			return getInvalidVideoResponse(vid), found
		}
		r.UID = ve.UID
		r.State = ve.State

	} else { // Video present in cache
		r.UID = ve.UID
		r.State = models.StatusProcessing
		if ve.State == "" { // update cache state to fetched
			sc.UpdateCacheEntryState(vid, models.StatusFetched)
		}
		//TODO: create DB record with status processing
		v := models.MetaData{
			VID:       r.VID,
			UID:       r.UID,
			Timestamp: ve.Timestamp,
			State:     models.StatusProcessing,
		}
		err = sc.CreateVideoMetadata(v)

		if err != nil {
			fmt.Println("Error when creating entry in dynamo")
			log.Fatalln(err)
		}

	}

	r.Preview = getPreviewLink(r.VID)
	r.Src = getPreviewLink(r.VID)

	if ve.State == models.StatusCancelled { // If video fetched after cancellation
		r.State = models.StatusDelete
		r.Preview = ""
		r.Src = ""
		found = false
	}

	rj, _ := json.Marshal(r)
	return rj, found
}

// TODO: use go routines
func (sc SessionController) UpdateVideoStatus(vid string, status string) []byte {
	// check status in cache
	ve, err := sc.GetCacheEntry(vid)
	if err != nil {
		log.Fatalln("Unexpected Error when updating status on lambda callback")
		log.Fatal(err)
		// TDOO: should return a 503
	}

	if ve.State == models.StatusCancelled {
		// delete video from final bucket
		//debug sc.DeleteVideoFromBucket(vid, "final-videos-bucket") // TODO: find a way to import this from main
	} else {
		/* cache state is either fetched or empty */

		// update db record with lambda status
		sc.UpdateVideoMetadataStatus(vid, status)
	}

	//debug sc.DeleteVideoFromBucket(vid, "raw-videos-bucket") // TODO: find a way to import this from main
	sc.DeleteCacheEntry(vid) // TODO: Error handling

	r := models.MsgResponse{Msg: "OK"}
	rj, _ := json.Marshal(r)
	return rj
}

/*Called when recording is cancelled. Hence only updating cache*/
func (sc SessionController) DeleteVideo(vid string) error {
	return sc.UpdateCacheEntryState(vid, models.StatusCancelled)
}
