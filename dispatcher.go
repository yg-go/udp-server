package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
)

const (
	DataTypeFFTIQ = 1 + iota
)

var (
	COM_CODE_IQ_INFO int16 = 0x0012
)

type Payload []byte

func parseHeader(payload Payload) (*ModelB, int) {
	message := ModelB{}
	offset := 0

	message.SourceCode = payload[offset]
	offset += 1
	message.DestinationCode = payload[offset]
	offset += 1
	message.CommandCode = int16(binary.LittleEndian.Uint16(payload[offset : offset+2]))
	offset += 2
	message.DataSize = int32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4

	return &message, offset
}

func parseIQInfo(header *ModelB, payload Payload) *IQInfo {
	info := IQInfo{}

	fmt.Println("IQInfo : ", info)
	info.ID = DocType(DataTypeFFTIQ)
	info.Header = *header

	offset := 0
	info.SampleNumber = uint32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4
	info.CenterFrequency = uint64(binary.LittleEndian.Uint64(payload[offset : offset+8]))
	offset += 8
	info.AntennaNumber = uint32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4
	info.FFTSize = uint32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4

	info.Data = make([]IQData, 0)

	var i uint32
	for i = 0; i < info.FFTSize; i++ {
		data := IQData{}
		data.I = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		data.Q = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		info.Data = append(info.Data, data)
	}

	return &info
}

func DispatcherHeader(in <-chan Payload, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case payload := <-in:
			head, dataOffset := parseHeader(payload)
			switch head.CommandCode {
			case COM_CODE_IQ_INFO:
				plots := parseIQInfo(head, payload[dataOffset:])
				fmt.Println(plots)
				break
			}
		}
	}
}
