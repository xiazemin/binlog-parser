package parser

import (
	"github.com/xiazemin/binlog-parser/parser/conversion"
)

type RowsEventBuffer struct {
	buffered []conversion.RowsEventData
}

func NewRowsEventBuffer() RowsEventBuffer {
	return RowsEventBuffer{}
}

func (mb *RowsEventBuffer) BufferRowsEventData(d conversion.RowsEventData) {
	mb.buffered = append(mb.buffered, d)
}

func (mb *RowsEventBuffer) Drain() []conversion.RowsEventData {
	ret := mb.buffered
	mb.buffered = nil

	return ret
}
