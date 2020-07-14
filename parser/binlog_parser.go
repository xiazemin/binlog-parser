package parser

import (
	"os"
	"github.com/xiazemin/binlog-parser/database"
	"github.com/xiazemin/binlog-parser/parser/parser"
)

func ParseBinlog(binlogFilename string, tableMap database.TableMap, consumerChain ConsumerChain) error {
	if _, err := os.Stat(binlogFilename); os.IsNotExist(err) {
		return err
	}

	return parser.ParseBinlogToMessages(binlogFilename, tableMap, consumerChain.consumeMessage)
}
