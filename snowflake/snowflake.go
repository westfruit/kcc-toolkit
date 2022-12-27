package snowflake

import (
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
)

var (
	sf *sonyflake.Sonyflake
)

func init() {
	st := sonyflake.Settings{
		MachineID: func() (uint16, error) {
			machineID := viper.GetInt32("snowflake.machineId")
			return uint16(machineID), nil
		},
	}

	sf = sonyflake.NewSonyflake(st)
}

func GenId() int64 {
	id, err := sf.NextID()
	if err != nil {
		return 0
	}

	return int64(id)
}
