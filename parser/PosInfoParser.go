package parser

import (
	"github.com/arnobroekhof/rd2wgs84"
	xj "github.com/basgys/goxml2json"
	"github.com/nleeper/goment"
	"github.com/tidwall/gjson"
	"go/types"
	"io"
	"ov-geosearch/models"
)

func PosInfoParser(message io.Reader) (positionUpdates []models.PositionUpdate, err error) {
	jsonMessage, err := xj.Convert(message)

	VVTMPUSH := gjson.Get(jsonMessage.String(), "VV_TM_PUSH")
	VVTMPUSH.Get("KV6posinfo").ForEach(func(eventName, event gjson.Result) bool {
		for _, info := range event.Array() {
			kvEvent := KVEvent{
				lineplanningnumber:    info.Get("lineplanningnumber").String(),
				reinforcementnumber:   info.Get("reinforcementnumber").String(),
				journeynumber:         info.Get("journeynumber").String(),
				passagesequencenumber: info.Get("passagesequencenumber").String(),
				source:                info.Get("source").String(),
				vehiclenumber:         info.Get("vehiclenumber").String(),
				dataownercode:         info.Get("dataownercode").String(),
				rdX:                   info.Get("rd-x").Float(),
				rdY:                   info.Get("rd-y").Float(),
				operatingday:          info.Get("operatingday").String(),
			}

			rd := rd2wgs84.NewRDCoordinates(kvEvent.rdX, kvEvent.rdY)
			wgs84 := rd.ToWGS84()

			//log.Println(info.Get("rd-x").Raw, "-", info.Get("rd-x").String(), "-", info.Get("rd-x").Float(), "-", wgs84.Lon)

			g, _ := goment.New()

			positionUpdate := models.PositionUpdate{
				Vehicle:        kvEvent.vehiclenumber,
				Operator:       kvEvent.dataownercode,
				Lat:            wgs84.Lat,
				Lon:            wgs84.Lon,
				RealtimeTripId: kvEvent.dataownercode + ":" + kvEvent.lineplanningnumber + ":" + kvEvent.journeynumber,
				Date:           kvEvent.operatingday,
				Timestamp:		g.ToUnix(),
			}

			if info.Get("rd-x").Float() > 1 { // check if RD coords are useful
				positionUpdates = append(positionUpdates, positionUpdate)
			}

		}

		// ToDo: Parse single object events
		//log.Println(event.Map())

		return true
	})

	return positionUpdates, err
}

type VTMMPUSH struct {
	Version     string
	DossierName string
	Timestamp  string
	KV6posinfo KV6posinfo
}

type KV6posinfo struct {
	ONSTOP types.Array
}

type KVEvent struct {
	lineplanningnumber    string
	operatingday          string
	reinforcementnumber   string
	journeynumber         string
	passagesequencenumber string
	source                string
	vehiclenumber         string
	rdX                   float64
	rdY                   float64
	dataownercode         string
}
