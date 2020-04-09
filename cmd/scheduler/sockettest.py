import socket
import os
import json


dict_shellcmd = {
    "DestIP": "127.0.0.1", 
    "SrcID": "UMSP_Scheduler", 
    "ObjType": "shellcmd",
    "obj": {
        "script": "df -m"
        }
}

dict_transfile = {
    "DestIP": "127.0.0.1", 
    "SrcID": "UMSP_Scheduler", 
    "ObjType": "transfile",
    "obj": {
	    "SrcFilePath": "/tmp/transfile1/test.sh",
	    "DestFilePath": "/tmp/transfile2/test.sh",
	    "FileMode": "0744"
        }
}

dict_notsupported = {
    "DestIP": "127.0.0.1", 
    "SrcID": "UMSP_Scheduler", 
    "ObjType": "notsupported",
    "obj": {
	    "x": "/tmp/transfile1/test.sh",
	    "y": "/tmp/transfile2/test.sh",
	    "FileMode": "0744"
        }
}


SCHUNIXSOCKET   = '/var/run/ire_clshsheduler.sock'

print("Connecting...")
if os.path.exists(SCHUNIXSOCKET):
    client = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    client.connect(SCHUNIXSOCKET)
    print("Ready.")
    print("Ctrl-C to quit.")

    client.send(json.dumps(dict_shellcmd).encode('utf-8'))
    client.send(json.dumps(dict_notsupported).encode('utf-8'))
    client.send(json.dumps(dict_transfile).encode('utf-8'))


    """
    print("Sending 'DONE' shuts down the server and quits.")
    while True:
        try:
            x = input("> ")
            if "" != x:
                print("SEND:", x)
                client.send(x.encode('utf-8'))
                if "DONE" == x:
                    print("Shutting down.")
                    break
        except KeyboardInterrupt as k:
            print("Shutting down.")
            client.close()
            break
    """
else:
    print("Couldn't Connect!")
    print("Done")