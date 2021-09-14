package main

type dataType int
type DocType dataType

type ModelB struct {
	SourceCode      byte  `json:"source_code" bson:"source_code"`
	DestinationCode byte  `json:"destination_code" bson:"destination_code"`
	CommandCode     int16 `json:"command_code" bson:"command_code"`
	DataSize        int32 `json:"data_size" bson:"data_size"`
}

type IQData struct {
	I float32 `bson:"i"`
	Q float32 `bson:"q"`
}

type IQInfo struct {
	ID              DocType  `json:"doc_type" bson:"doc_type"`
	Header          ModelB   `json:"header" bson:"header"`
	SampleNumber    uint32   `json:"sample_number" bson:"sample_number"`
	CenterFrequency uint64   `json:"center_frequency" bson:"center_frequency"`
	AntennaNumber   uint32   `json:"antenna_number" bson:"antenna_number"`
	FFTSize         uint32   `json:"fft_size" bson:"fft_size"`
	Data            []IQData `json:"iq_data" bson:"iq_data"`
}
