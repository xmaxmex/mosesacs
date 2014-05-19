package gocwmp

import (
  "encoding/xml"
)

type SoapEnvelope struct {
	XMLName xml.Name
	Header  SoapHeader
	Body    SoapBody
}

type SoapHeader struct{}
type SoapBody struct {
	CWMPMessage CWMPMessage `xml:",any"`
}

type CWMPMessage struct {
	XMLName xml.Name
}

type Event struct {
}

type CWMPInform struct {
	DeviceId DeviceID `xml:"Body>Inform>DeviceId"`
	Events   []Event
}

type DeviceID struct {
	Manufacturer string
	OUI          string
	SerialNumber string
}

func InformResponse() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:cwmp="urn:dslforum-org:cwmp-1-0" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:schemaLocation="urn:dslforum-org:cwmp-1-0 ..\schemas\wt121.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <soap:Header/>
  <soap:Body soap:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <cwmp:InformResponse>
      <MaxEnvelopes>1</MaxEnvelopes>
    </cwmp:InformResponse>
  </soap:Body>
</soap:Envelope>`
}