package data_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stuarforrest-infinity/websocket-lambda/data"
)

func PutConnectionID(t *testing.T) {
	sess, table, drop, err := CreateTestTable()
	assert.NoError(t, err, "create table")
	defer drop()
	datastore := data.NewDataStore(table, sess)

	// Create a connection identifier.
	ci := data.ConnectionItemData{
		ID: "test-id",
	}

	d := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	err = datastore.PutConnectionID(ci, d)
	assert.NoError(t, err, "insert connection")

	// // Get a single count.
	// // No count daily count has yet been recorded for this sku
	// _, ok, err := datastore.GetDailySalesCount(d, "nonExistantSKU")
	// assert.NoError(t, err, "get single sku count - should be no error")
	// assert.Equal(t, false, ok, "should return ok = false when no count has been recorded")

	// // Count record exists
	// count, ok, err := datastore.GetDailySalesCount(d, "testSKU")
	// assert.NoError(t, err, "should be no error getting a single SKU count")
	// assert.Equal(t, true, ok, "should return ok = true when an record exists")
	// assert.Equal(t, int64(1), count, "expected count to be 1")

	// // Update a single SKU count.
	// err = datastore.UpdateCount(d, store.CountRecord{
	// 	SKU:   "testSKU",
	// 	Count: 2,
	// })
	// assert.NoError(t, err, "should be no error updating a count")
	// count, ok, err = datastore.GetDailySalesCount(d, "testSKU")
	// assert.NoError(t, err, "should be no error getting a single SKU count")
	// assert.Equal(t, true, ok, "should return ok = true when an record exists")
	// assert.Equal(t, int64(3), count, "should return a count of 3")
}