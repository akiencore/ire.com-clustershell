package communicate

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"ire.com/clustershell/logger"
)

const (
	//ObjTypeTransFile -- type of transfer file
	ObjTypeTransFile = "transfile"
	//ObjTypeShellCmd -- type of shell command
	ObjTypeShellCmd = "shellcmd"
)

//TransFile -- for transferring file between two nodes.
type TransFile struct {
	SrcFilePath  string
	DestFilePath string
	FileMode     string
}

//ShellCMD -- call remote node to do some commands
type ShellCMD struct {
	Script string
}

//MSGObj -- for transferring obj between two nodes.
//receiver will do consequent action according to ObjType
type MSGObj struct {
	TaskID     TaskIDType
	DestIPPort string
	SrcID  string

	//2 types: "shellcmd", "transfile"
	ObjType string
	//for example with ObjType = "shellcmd":
	//obj = {"script": "df -m"}
	//for example with ObjType = "transfile":
	/*obj = {
	"SrcFilePath": "/myapp/scripts/dosomething.sh",
	"DestFilePath": "/tmp/dosomething.sh",
	"FileMode": "0744"}
	*/
	Obj interface{}
}

//TaskIDType --
type TaskIDType int64

//TASKID -- the global task id for caller
var (
	TASKID = TaskIDType(time.Now().UnixNano())
)

//SocketTask --
type SocketTask struct {
	TaskID TaskIDType
	SocketConn   *net.UnixConn
	Msgobj *MSGObj
}

//MarshalMSG --
func (m *MSGObj) MarshalMSG() ([]byte, error) {

	msgBytes, err := json.Marshal(m)

	return msgBytes, err
}

//UnMarshalMSG --
func (m *MSGObj) UnMarshalMSG(data []byte) error {
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}

	if m.ObjType == ObjTypeShellCmd {
		var sc ShellCMD

		objstr, _ := json.Marshal(m.Obj)
		err = json.Unmarshal(objstr, &sc)
		if err != nil {
			return fmt.Errorf("%s Unmarshall:%v", ObjTypeShellCmd, err)
		}
		m.Obj = sc

	} else if m.ObjType == ObjTypeTransFile {
		var tf TransFile

		objBytes, _ := json.Marshal(m.Obj)
		err = json.Unmarshal(objBytes, &tf)
		if err != nil {
			return fmt.Errorf("%s Unmarshall:%v", ObjTypeTransFile, err)
		}
		m.Obj = tf

	} else {
		if m.ObjType == "" {
			return fmt.Errorf("UnMarshalMSG -- ObjType is empty")
		}
		return fmt.Errorf("UnMarshalMSG -- not support %s", m.ObjType)
	}

	return nil
}

//GetMSGListFromSocketBuf --
func GetMSGListFromSocketBuf(buf []byte, conn *net.UnixConn) (*map[TaskIDType]SocketTask, error) {
	var (
		err            error
		socketTasksMap = make(map[TaskIDType]SocketTask)
	)

	defer func() {
		logger.Debug("GetMSGListFromSocketBuf exit -- generated socket tasks:", len(socketTasksMap))
	}()

	logger.Debug("GetMSGListFromSocketBuf enter to analysize buf:", string(buf))

	brace := 0
	jsonStarted := false
	startIdx := 0

	for i := 0; i < len(buf); i++ {
		if buf[i] == byte(123) { // '{' is 123
			logger.Debug("got a {")
			brace++
			if !jsonStarted {
				jsonStarted = true
				startIdx = i
			}
		} else if buf[i] == byte(125) { // '}' is 125
			logger.Debug("got a }")
			brace--
		}

		if i > 0 && brace < 0 {
			return nil, fmt.Errorf("there is json format error of brace pairing")
		}

		if jsonStarted && brace == 0 {
			var m MSGObj
			err = m.UnMarshalMSG(buf[startIdx:(i + 1)])
			if err != nil {
				logger.Error("UnMarshalMSG --", err) //todo: return err to caller
			} else {

				logger.Info("got a task message:", m)

				TASKID++
				m.TaskID = TASKID
				socketTasksMap[TASKID] = SocketTask{Msgobj: &m, TaskID: TASKID, SocketConn: conn}
			}

			jsonStarted = false
		}
	}

	return &socketTasksMap, nil
}
