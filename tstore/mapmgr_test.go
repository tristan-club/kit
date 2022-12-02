package tstore

import (
	"github.com/stretchr/testify/assert"
	uid2 "github.com/tristan-club/kit/uid"
	"os"
	"testing"
)

func TestMapMgr(t *testing.T) {
	err := InitTStore(os.Getenv("TSTORE_SVC"))
	if err != nil {
		t.Fatal(err)
	}

	mm := GetMapMgr()
	uid := uid2.GenerateUuid(true)
	path := "PathA"

	err = mm.Save(uid, path, "key1", "value1")
	if err != nil {
		t.Fatal(err)
	}

	err = mm.Save(uid, path, "key2", "value2")
	if err != nil {
		t.Fatal(err)
	}

	arr := make(map[string]string, 0)
	err = mm.Fetch(uid, path, "").Scan(&arr).Error()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(arr)
	}

	err = mm.BatchSave(uid, path, map[string]interface{}{
		"key1": "value1New",
		"key3": "value3New",
	})
	if err != nil {
		t.Fatal(err)
	}

	k1, err := mm.Fetch(uid, path, "key1").GetStrValue()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(k1)
		assert.Equal(t, k1, "value1New")
	}
}
