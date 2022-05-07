package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"os"
)
func intToBytes(num int32) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.LittleEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
func createBitmap(data []byte, width int, height int){
	rowSize := int(math.Floor((24 * float64(width) + 31) / 32) * 4)
	pixelArraySize := rowSize * height
	fileSize := 54 + pixelArraySize
	offset := rowSize - width * 3
	header := []byte{0x42,0x4d}											//BMP header
	header = append(header,intToBytes(int32(fileSize))...)				//filesize
	header = append(header,intToBytes(int32(0))...)						//reserved
	header = append(header,intToBytes(int32(0x36))...)					//offset
	header = append(header,intToBytes(int32(0x28))...)					//bitmapInfoHeader size
	header = append(header,intToBytes(int32(width))...)					//width of image
	header = append(header,intToBytes(int32(height))...)				//height of image
	header = append(header,01,00)								//colour planes
	header = append(header,0x18,00)								//colour depth
	header = append(header,intToBytes(int32(0))...)						//compression
	header = append(header,intToBytes(int32(pixelArraySize))...)		//image size
	header = append(header,intToBytes(int32(0xc5))...)					//horizontal resolution
	header = append(header,intToBytes(int32(0xc5))...)					//vertical resolution
	header = append(header,intToBytes(int32(0))...)						//number of colours
	header = append(header,intToBytes(int32(0))...)						//number of important colours

	var pixelArray []byte

	for i := 0; i < len(data); i++ {
		if i%width == 0 {
			for j := 0; j < offset; j++ {
				pixelArray = append(pixelArray, 0x00)
			}
		}
		switch data[i] {
		case 0:
			pixelArray = append(pixelArray, 0x00, 0x00, 0x00)
		case 1:
			pixelArray = append(pixelArray, 0x00, 0x00, 0xff)
		case 2:
			pixelArray = append(pixelArray, 0x00, 0xff, 0x00)
		case 3:
			pixelArray = append(pixelArray, 0x00, 0xff, 0xff)
		case 4:
			pixelArray = append(pixelArray, 0xff, 0x00, 0x00)
		case 5:
			pixelArray = append(pixelArray, 0xff, 0x00, 0xff)
		case 6:
			pixelArray = append(pixelArray, 0xff, 0xff, 0x00)
		case 7:
			pixelArray = append(pixelArray, 0xff, 0xff, 0xff)
		}
	}
	bitmapData := append(header,pixelArray...)

	err := ioutil.WriteFile(os.Args[2], bitmapData,0644)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var data2 []byte
	for i:=0;i<len(data);i++{
		data2 = append(data2,data[i]-48)
	}
	createBitmap(data2,10,10)
}