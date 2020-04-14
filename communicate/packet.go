package communicate

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (

	//SCHRECVPPORT --  on the udp port scheduler listen and recerive
	SCHRECVPPORT = ":33225"
	//SCHSENDPORT -- on the udp port scheduler send to any host you desire
	SCHSENDPORT = ":52233"
	//XCTRECVPPORT -- on the udp port executor listen and recerive
	XCTRECVPPORT = ":33235"
	//XCTSENDPORT -- on the udp port executor send to any host you desire
	XCTSENDPORT = ":52243"

	//MAXPAYLOADLEN --
	MAXPAYLOADLEN = 32 * 1024 //32K Bytes
	//MAXUDPPACKET --
	MAXUDPPACKET = 40 * 1024

	//HEADERLEN -  the length of packet header (other fields except of Data)
	HEADERLEN = 4 * 4

	//EOF -- end signal of session communication
	EOF = "<E<O>F>"
	//STDOUTSTR --
	STDOUTSTR = "STDOUT"
	//STDERRSTR --
	STDERRSTR = "STDOUT"
	
	//STDOUTEOF -- end signal of stdout line
	STDOUTEOF = STDOUTSTR + EOF
	//STDERREOF -- end signal of stderr line
	STDERREOF = STDERRSTR + EOF

	//DEFAULTTIMEOUT --
	DEFAULTTIMEOUT = 3
)

//Packet -- means udp packet
type Packet struct {
	PacketNum  uint32
	Index      uint32
	PayloadLen uint32
	Checksum   uint32
	Payload    []byte
}

//MarshalPacket -
func (p *Packet) MarshalPacket() ([]byte, error) {
	p.PayloadLen = uint32(len(p.Payload))

	if p.PayloadLen > MAXPAYLOADLEN {
		return nil, fmt.Errorf("payload size of packet should not be greater than %d bytes", MAXPAYLOADLEN)
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, p.PacketNum)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, p.Index)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, p.PayloadLen)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, p.Checksum)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, p.Payload)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

//UnMarshalPacket -
func UnMarshalPacket(data []byte) (*Packet, error) {
	var p Packet

	reader := bytes.NewReader(data)

	err := binary.Read(reader, binary.BigEndian, &p.PacketNum)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.BigEndian, &p.Index)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.BigEndian, &p.PayloadLen)
	if err != nil {
		return nil, err
	}
	if int(p.PayloadLen) != (len(data) - HEADERLEN) {
		return nil, fmt.Errorf("failed recerived a packet")
	}

	err = binary.Read(reader, binary.BigEndian, &p.Checksum)
	if err != nil {
		return nil, err
	}

	payload := make([]byte, p.PayloadLen)

	err = binary.Read(reader, binary.BigEndian, &payload)
	if err != nil {
		return nil, err
	}

	p.Payload = payload

	return &p, nil
}
