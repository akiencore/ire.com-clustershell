package communicate

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	//MAXPAYLOADLEN --
	MAXPAYLOADLEN = 32 * 1024 //32K Bytes

	//HOSTIDLEN - the string length of the identity of source host
	HOSTIDLEN = 20

	//HEADERLEN -  the length of packet header (other fields except of Data)
	HEADERLEN = HOSTIDLEN + 8 + 4*4
)

//Packet -- means udp packet
type Packet struct {
	SrcID      [HOSTIDLEN]byte
	TaskID     TaskIDType
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

	err := binary.Write(buf, binary.BigEndian, p.SrcID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, p.TaskID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, p.PacketNum)
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

	err := binary.Read(reader, binary.BigEndian, &p.SrcID)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.BigEndian, &p.TaskID)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.BigEndian, &p.PacketNum)
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
