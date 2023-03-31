//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package inverted

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sync"

	"github.com/pkg/errors"
)

type PropLenData struct {
	BucketedData map[string][]int
}

type JsonPropertyLengthTracker struct {
	path string
	data PropLenData
	sync.Mutex
}

func NewJsonPropertyLengthTracker(path string) (*JsonPropertyLengthTracker, error) {

	t := &JsonPropertyLengthTracker{
		data: PropLenData{make(map[string][]int)},
		path: path,
	}

	
		// read the file into memory
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				return t, nil
			}
			return nil, err
		}

		var data PropLenData
		if err := json.Unmarshal(bytes, &data); err != nil {
			return nil, err
		}
		t.data = data


	return t, nil
}

func (t *JsonPropertyLengthTracker) FileName() string {
	return t.path
}

func (t *JsonPropertyLengthTracker) TrackProperty(propName string, value float32) error {
	if value == 0 {
		return nil
	}
	t.Lock()
	defer t.Unlock()

	

	bucketId := t.bucketFromValue(value)
	if bucket, ok := t.data.BucketedData[propName]; ok {
		bucket[bucketId] = bucket[bucketId] + 1
	} else {
		bucket := make([]int, 64)
		bucket[bucketId] = 1
		t.data.BucketedData[propName] = bucket
	}
	return nil
}


func (t *JsonPropertyLengthTracker) bucketFromValue(value float32) uint16 {
	if value <= 5.00 {
		return uint16(value) - 1
	}

	bucket := int(math.Log(float64(value)/4.0)/math.Log(1.25) + 4)
	if bucket > 63 {
		return 64
	}
	if bucket == 63 {
		panic("bucket 63 is too big, you shouldn't be writing to this")
	}
	return uint16(bucket)
}

func (t *JsonPropertyLengthTracker) valueFromBucket(bucket uint16) float32 {
	if bucket <= 5 {
		return float32(bucket + 1)
	}

	return float32(4 * math.Pow(1.25, float64(bucket)-3.5))
}

func (t *JsonPropertyLengthTracker) PropertyMean(propName string) (float32, error) {
	t.Lock()
	defer t.Unlock()

	bucket, ok := t.data.BucketedData[propName]
	if !ok {
		return 0, nil
	}

	sum := float32(0)
	totalCount := float32(0)

	for i := 0; i < len(bucket); i++ {
		sum += t.valueFromBucket(uint16(i)) * float32(bucket[i])
		totalCount += float32(bucket[i])
	}

	if totalCount == 0 {
		return 0, nil
	}

	return sum / totalCount, nil
}

func (t *JsonPropertyLengthTracker) PropertyTally(propName string) (uint64, uint64, float64, uint64, uint64, error) {
	t.Lock()
	defer t.Unlock()

	bucket, ok := t.data.BucketedData[propName]
	if !ok {
		return 0, 0, 0, 0, 0, nil
	}

	sum := uint64(0)
	totalCount := uint64(0)
	countTally := uint64(0)
	proplenTally := uint64(0)

	for i := 0; i < len(bucket); i++ {
		count := bucket[i]
		value := t.valueFromBucket(uint16(i))
		countTally += uint64(count)
		proplenTally += uint64(value)
		sum += uint64(value * float32(count))
		totalCount += uint64(count)
		fmt.Printf("Bucket: %v, Count: %v, Proplen: %v, Sum: %v, TotalCount: %v\n", i,  count, value, sum, totalCount)

		
	}

	if totalCount == 0 {
		return 0, 0, 0, 0, 0, nil
	}

	return sum, totalCount, float64(sum) / float64(totalCount), countTally, proplenTally, nil
}



func (t *JsonPropertyLengthTracker) Flush() error {
	t.Lock()
	defer t.Unlock()

	bytes, err := json.Marshal(t.data)
	if err != nil {
		return err
	}

	ioutil.WriteFile(t.path, bytes, 0o666)
	return nil
}

func (t *JsonPropertyLengthTracker) Close() error {
	if err := t.Flush(); err != nil {
		return errors.Wrap(err, "flush before closing")
	}

	t.Lock()
	defer t.Unlock()


	t.data.BucketedData = nil

	return nil
}

func (t *JsonPropertyLengthTracker) Drop() error {
	t.Lock()
	defer t.Unlock()


	t.data.BucketedData = nil

	if err := os.Remove(t.path); err != nil {
		return errors.Wrap(err, "remove prop length tracker state from disk")
	}

	return nil
}
