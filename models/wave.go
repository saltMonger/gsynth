package models

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

const CD_SAMPLE_RATE = 44100

type WaveHeader struct {
	RiffId        [4]uint8
	RiffSize      uint32
	WaveType      [4]uint8
	FormatId      [4]uint8
	FormatSize    uint32
	FormatCode    uint16
	Channels      uint16
	SampleRate    uint32
	AverageBps    uint32
	Align         uint16
	BitsPerSample uint16
	WaveId        [4]uint8
	WaveSize      uint32
}

func (wh WaveHeader) ToBytes() []byte {
	ret := []byte{}
	ret = append(ret, wh.RiffId[:]...)
	ret = binary.LittleEndian.AppendUint32(ret, wh.RiffSize)
	ret = append(ret, wh.WaveType[:]...)
	ret = append(ret, wh.FormatId[:]...)
	ret = binary.LittleEndian.AppendUint32(ret, wh.FormatSize)
	ret = binary.LittleEndian.AppendUint16(ret, wh.FormatCode)
	ret = binary.LittleEndian.AppendUint16(ret, wh.Channels)
	ret = binary.LittleEndian.AppendUint32(ret, wh.SampleRate)
	ret = binary.LittleEndian.AppendUint32(ret, wh.AverageBps)
	ret = binary.LittleEndian.AppendUint16(ret, wh.Align)
	ret = binary.LittleEndian.AppendUint16(ret, wh.BitsPerSample)
	ret = append(ret, wh.WaveId[:]...)
	ret = binary.LittleEndian.AppendUint32(ret, wh.WaveSize)
	return ret
}

func SamplesToBytes(input []int16) []byte {
	bytes := []byte{}
	for _, b := range input {
		bytes = binary.LittleEndian.AppendUint16(bytes, uint16(b))
	}
	fmt.Println("sample len: ", len(bytes))
	return bytes
}

func WriteWaveFile(sampleBuffer []int16, sampleTotal uint32, channels uint16, path string) {
	byteTotal := sampleTotal * 2 * uint32(channels)

	align := (channels + 16) / 8
	fmt.Println("align", align)

	header := WaveHeader{
		RiffId:        [4]uint8{'R', 'I', 'F', 'F'},
		WaveType:      [4]uint8{'W', 'A', 'V', 'E'},
		FormatId:      [4]uint8{'f', 'm', 't', ' '},
		FormatSize:    16,
		FormatCode:    1, // PCM
		Channels:      channels,
		SampleRate:    CD_SAMPLE_RATE,
		BitsPerSample: 16,
		Align:         align,
		AverageBps:    CD_SAMPLE_RATE * uint32(align),
		WaveId:        [4]uint8{'d', 'a', 't', 'a'},
		WaveSize:      byteTotal,
	}
	header.RiffSize = byteTotal + uint32(binary.Size(header)) - 8

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("sample bytes", len(sampleBuffer))
	outBytes := header.ToBytes()
	fmt.Println("header bytes", len(outBytes))
	outBytes = append(outBytes, SamplesToBytes(sampleBuffer)...)

	outBuffer := bufio.NewWriter(file)
	n, err := outBuffer.Write(outBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote bytes: ", n)
}
