package snowflake

import (
	"testing"
)

func init() {

}

// 生成一个雪花Id
func TestGenId(t *testing.T) {
	newId, err := sf.NextID()

	if err != nil {
		t.Fatalf("GenId error, %s", err)
	}
	t.Logf("new genid = %d", newId)
}
