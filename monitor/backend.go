package monitor

import (
	"encoding/json"
	"log"
	"time"

	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type LastSeenEvent struct {
	Address   string `binding:"required" json:"address"`
	Timestamp int64  `binding:"required" json:"timestamp"`
}

func (mms *MqttMonitor) AddParticipant(address string, lastSeenTS int64) (err error) {
	// store receive address - planetmint address pair
	var lastSeen LastSeenEvent
	lastSeen.Address = address
	lastSeen.Timestamp = lastSeenTS

	lastSeenBytes, err := json.Marshal(lastSeen)
	if err != nil {
		mms.Log("[Monitor] Error serializing ConversionRequest: " + err.Error())
		return
	}
	increaseCounter := false
	// returns an error if the entry does not exist (we have to increase the counter in this case)
	_, err = mms.db.Get([]byte(address), nil)
	if err != nil {
		increaseCounter = true
	}
	mms.dbMutex.Lock()
	if increaseCounter {
		mms.numberOfElements++
	}
	err = mms.db.Put([]byte(address), lastSeenBytes, nil)
	mms.dbMutex.Unlock()
	if err != nil {
		log.Println("[Monitor] storing addresses in DB: " + err.Error())
		return
	}
	return
}

func (mms *MqttMonitor) deleteEntry(key []byte) (err error) {
	mms.dbMutex.Lock()
	err = mms.db.Delete(key, nil)
	mms.numberOfElements--
	mms.dbMutex.Unlock()
	return
}

func (mms *MqttMonitor) getAmountOfElements() (amount int64, err error) {
	iter := mms.db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		amount++
	}

	// Check for any errors encountered during iteration
	if err := iter.Error(); err != nil {
		log.Println("[Monitor] " + err.Error())
	}
	return
}
func (mms *MqttMonitor) getDataFromIter(iter iterator.Iterator) (lastSeen LastSeenEvent, err error) {
	key := iter.Key()
	value := iter.Value()
	err = json.Unmarshal(value, &lastSeen)
	if err != nil {
		mms.Log("[Monitor] Failed to unmarshal entry: " + string(key) + " - " + err.Error())
	}
	return
}

func (mms *MqttMonitor) CleanupDB() {
	// Create an iterator for the database
	iter := mms.db.NewIterator(nil, nil)
	defer iter.Release() // Make sure to release the iterator at the end

	// Iterate over all elements in the database
	for iter.Next() {
		// Use iter.Key() and iter.Value() to access the key and value
		lastSeen, err := mms.getDataFromIter(iter)
		if err != nil {
			mms.Log("[Monitor] Failed to unmarshal entry: " + string(iter.Key()) + " - " + err.Error())
			continue
		}
		timeThreshold := time.Now().Add(-1 * mms.CleanupPeriodicityInMinutes * time.Minute).Unix()
		if lastSeen.Timestamp <= timeThreshold {
			// If the entry is older than 12 hours, delete it
			err := mms.deleteEntry(iter.Key())
			if err != nil {
				mms.Log("[Monitor] Failed to delete entry: " + err.Error())
			}
		}
	}

	// Check for any errors encountered during iteration
	if err := iter.Error(); err != nil {
		mms.Log(err.Error())
	}
}
