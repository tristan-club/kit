package customid

import (
	"fmt"
	"github.com/tristan-club/kit/uid"
	"strconv"
	"strings"
)

type CustomId struct {
	ct  int32
	id  string
	cbt int32
}

func NewCustomId(customType int32, id string, callbackType int32) *CustomId {
	if id == "" {
		id = uid.GenerateUuid(true)
	}
	return &CustomId{
		ct:  customType,
		id:  id,
		cbt: callbackType,
	}
}

func (c *CustomId) GetCustomType() int32 {
	return c.ct
}

func (c *CustomId) GetId() string {
	return c.id
}

func (c *CustomId) GetCallbackType() int32 {
	return c.cbt
}

func (c *CustomId) String() string {
	return fmt.Sprintf("%d_%s_%d", c.ct, c.id, c.cbt)
}

func ParseCustomId(ciStr string) (*CustomId, bool) {
	a := strings.Split(ciStr, "_")
	if len(a) != 3 {
		a = strings.Split(ciStr, "/")
		if len(a) != 3 {
			return nil, false
		}
	}

	ct, _ := strconv.ParseInt(a[0], 10, 64)
	cbt, _ := strconv.ParseInt(a[2], 10, 64)

	ci := &CustomId{
		ct:  int32(ct),
		id:  a[1],
		cbt: int32(cbt),
	}
	return ci, true
}
