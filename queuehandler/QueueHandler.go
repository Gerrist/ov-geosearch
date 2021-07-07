package queuehandler

import (
	"bytes"
	"compress/gzip"
	"github.com/pebbe/zmq4"
	"io"
	"log"
	"ov-geosearch/parser"
	"ov-geosearch/processor"
	"time"
)

// /CXX/KV6posinfo
func QueueHandler() {
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)

	defer subscriber.Close()

	subscriber.SetLinger(0)
	subscriber.SetRcvtimeo(1 * time.Second)

	subscriber.Connect("tcp://pubsub.besteffort.ndovloket.nl:7658")

	subscriber.SetSubscribe("/ARR/KV6posinfo")
	subscriber.SetSubscribe("/CXX/KV6posinfo")
	subscriber.SetSubscribe("/DITP/KV6posinfo")
	subscriber.SetSubscribe("/CXX/KV6posinfo")
	subscriber.SetSubscribe("/EBS/KV6posinfo")
	subscriber.SetSubscribe("/GVB/KV6posinfo")
	subscriber.SetSubscribe("/OPENOV/KV6posinfo")
	subscriber.SetSubscribe("/QBUZZ/KV6posinfo")
	subscriber.SetSubscribe("/RIG/KV6posinfo")
	subscriber.SetSubscribe("/SYNTUS/KV6posinfo")

	log.Println("Handling KV6 queue")

	for {
		msg, err := subscriber.RecvMessageBytes(0)

		if err != nil {
			continue
		}

		//envelope := string(msg[0])
		message, _ := gunzip(msg[1])

		positionUpdates, err := parser.PosInfoParser(message)

		for _, positionUpdate := range positionUpdates {
			processor.ProcessPosition(positionUpdate)
		}
	}

}

func gunzip(data []byte) (io.Reader, error) {
	buf := bytes.NewBuffer(data)
	reader, err := gzip.NewReader(buf)
	defer reader.Close()

	if err != nil {
		return nil, err
	}

	buf3 := new(bytes.Buffer)
	buf3.ReadFrom(reader)

	return buf3, nil
}
