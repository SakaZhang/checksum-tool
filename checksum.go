package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"hash/crc64"
	"io"
	"os"
	"time"
)

func HashFile(h hash.Hash, filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	h.Write([]byte(""))
	_, err = io.Copy(h, file)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func HashCRC(h hash.Hash64, filename string) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil
	}
	h.Write([]byte(""))
	_, err = io.Copy(h, file)
	if err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

func main() {
	startTime := time.Now()

	var filename string
	var (
		crc64Chan  = make(chan uint64)
		md5Chan    = make(chan []byte)
		sha1Chan   = make(chan []byte)
		sha256Chan = make(chan []byte)
	)
	flag.StringVar(&filename, "f", "", "checksum -f FILENAME")
	flag.Parse()

	crc64Hash := crc64.New(crc64.MakeTable(crc64.ECMA))
	md5Hash := md5.New()
	sha1Hash := sha1.New()
	sha256Hash := sha256.New()

	go func() {
		hashRes, _ := HashFile(md5Hash, filename)
		md5Chan <- hashRes
	}()

	go func() {
		hashRes, _ := HashFile(sha1Hash, filename)
		sha1Chan <- hashRes
	}()

	go func() {
		hashRes, _ := HashFile(sha256Hash, filename)
		sha256Chan <- hashRes
	}()

	go func() {
		hashRes, _ := HashCRC(crc64Hash, filename)
		crc64Chan <- hashRes
	}()

	crc64Result := <-crc64Chan
	md5Result := <-md5Chan
	sha1Result := <-sha1Chan
	sha256Result := <-sha256Chan

	endTime := time.Now()

	fmt.Printf("CRC64 checksum: %d\n", crc64Result)
	fmt.Printf("MD5 checksum: %x\n", md5Result)
	fmt.Printf("SHA1 checksum: %x\n", sha1Result)
	fmt.Printf("SHA256 checksum: %x\n", sha256Result)
	fmt.Printf("Spent time: %v\n", endTime.Sub(startTime))

}
