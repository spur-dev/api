package controllers

import (
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/spur-dev/api/models"
)

func (uc SessionController) GetCacheEntry(vid string) (models.Entry, error) {
	var res models.Entry
	err := uc.cache.Update(func(txn *badger.Txn) error {
		item, berr := txn.Get([]byte(vid))

		if berr != nil {
			return berr
		}

		return item.Value(func(be []byte) error {
			var err error
			res, err = decodeToBytes(be)
			return err
		})

	})
	return res, err
}

func (uc SessionController) CreateCacheEntry(vid string, uid string, ts int) error {
	e := models.Entry{
		VID:       vid,
		UID:       uid,
		State:     "",
		Timestamp: ts,
	}

	eb := encodeToBytes(e)

	return uc.cache.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(vid), eb)
	})

}

func (uc SessionController) UpdateCacheEntryState(vid string, newState string) error {
	err := uc.cache.Update(func(txn *badger.Txn) error {
		item, berr := txn.Get([]byte(vid))
		if berr != nil {
			return berr
		}
		var oe, ne models.Entry
		err := item.Value(func(be []byte) error {
			var err error
			oe, err = decodeToBytes(be)
			return err
		})

		if err != nil {
			fmt.Printf("Error fetching key %s  for update \n", vid)
			return err
		}

		ne = oe
		ne.State = newState
		neb := encodeToBytes(ne)

		err = txn.Set([]byte(vid), neb)

		if err != nil {
			fmt.Printf("Error when updating status key %s \n", vid)
			return err
		}
		return nil
	})
	return err
}

func (uc SessionController) DeleteCacheEntry(vid string) error {
	return uc.cache.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(vid))
	})
}
