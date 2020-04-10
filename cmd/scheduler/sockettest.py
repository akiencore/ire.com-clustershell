import socket
import os
import json
import time

dict_shellcmd1 = {
    "DestIPPort": "127.0.0.1:33225", 
    "SrcID": "UMSP_Scheduler", 
    "ObjType": "shellcmd",
    "obj": {
        "script": "/tmp/test.sh"
        }
}

dict_shellcmd2 = {
    "DestIPPort": "127.0.0.1:33225", 
    "SrcID": "UMSP_Scheduler", 
    "ObjType": "shellcmd",
    "obj": {
        "script": "df"
        }
}

dict_transfile = {
    "DestIPPort": "127.0.0.1:33225", 
    "SrcID": "UMSP_Scheduler", 
    "ObjType": "transfile",
    "obj": {
	    "SrcFilePath": "/tmp/transfile1/test.sh",
	    "DestFilePath": "/tmp/transfile2/test.sh",
	    "FileMode": "0744"
        }
}

dict_notsupported = {
    "DestIPPort": "127.0.0.1:33225", 
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

    for _ in range(5):
        client.send(json.dumps(dict_shellcmd1).encode('utf-8'))
        #client.send(json.dumps(dict_shellcmd2).encode('utf-8'))
        #client.send(json.dumps(dict_notsupported).encode('utf-8'))
        #client.send(json.dumps(dict_transfile).encode('utf-8'))
        #time.sleep(1)

    client.close()

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