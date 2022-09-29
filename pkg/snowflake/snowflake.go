package snowflake

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var tParse time.Time
	tParse, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return fmt.Errorf("time.Parse fail err:%s", err.Error())
	}
	snowflake.Epoch = tParse.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return err
}

func GenID() int64 {
	return node.Generate().Int64()
}
