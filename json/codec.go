package json

import (
	"github.com/json-iterator/go"
	"time"
	"unsafe"
)

const (
	date_partten = "2006-01-02"
	time_partten = "2006-01-02 15:04:05"
)

func init() {
	RegisterDateAsStringCodec()
	RegisterTimeAsStringCodec()
}

type timeAsStringCodec struct {
}

func RegisterTimeAsStringCodec() {
	jsoniter.RegisterTypeEncoder("time.Time", &timeAsStringCodec{})
	jsoniter.RegisterTypeDecoder("time.Time", &timeAsStringCodec{})
}

func (codec *timeAsStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	time_str := iter.ReadString()
	*((*time.Time)(ptr)), _ = time.Parse(time_partten, time_str)
}

func (codec *timeAsStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.UnixNano() == 0
}
func (codec *timeAsStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	stream.WriteString(ts.Format(time_partten))
}

type dateAsStringCodec struct {
	precision time.Duration
}

func RegisterDateAsStringCodec() {
	jsoniter.RegisterTypeEncoder("time.Time", &dateAsStringCodec{})
	jsoniter.RegisterTypeDecoder("time.Time", &dateAsStringCodec{})
}

func (codec *dateAsStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	time_str := iter.ReadString()
	*((*time.Time)(ptr)), _ = time.Parse(date_partten, time_str)
}

func (codec *dateAsStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.UnixNano() == 0
}
func (codec *dateAsStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	stream.WriteString(ts.Format(date_partten))
}
