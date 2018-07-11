package main

import (
	"log"
	"go-t2sdk/t2sdk"
)

type Request_105010 struct {
	FunctionNo int32
}

func (req Request_105010) ParseRequestContent() uintptr {
	packer := t2sdk.GetNewPacker(2)
	t2sdk.AddPackerRef(packer)
	t2sdk.BeginPack(packer)
	t2sdk.AddField(packer, "function_id")
	t2sdk.AddInt(packer, req.FunctionNo)
	t2sdk.EndPack(packer)
	return packer
}






func main() {
	config := t2sdk.GetNewConfig()
	log.Printf("T2 Config Pointer:%x", config)
	iniFile := "t2sdk.ini"
	ret := t2sdk.LoadT2Config(config, iniFile)
	if ret != 0 {
		log.Printf("T2 Config Read Error,ErrorNo:%d", ret)
	}
	log.Printf("T2 Config Read Reulst:%d", ret)

	conn := t2sdk.GetNewConnection(config)
	log.Printf("T2 NewConnection Pointer:%x", conn)

	ret = t2sdk.InitCreat(conn)
	if ret != 0 {
		log.Printf("InitCreat Error,ErrorNo:%d", ret)
	}

	ret = t2sdk.T2Connect(conn, 100)
	if ret != 0 {
		err_info := t2sdk.GetErrorMsg(conn, ret)
		log.Printf("T2Connect Error,ErrorNo:%d,ErrorInfo:%s", ret, err_info)
	} else {
		log.Printf("T2 Connection Ok...")

		var iq t2sdk.IRequest = Request_105010{105010}
		ret = t2sdk.SyncSendMsg(conn, 105010, iq)

		if ret < 0 {
			log.Printf("syncSendMsg Error,FunctionNo:%d,Error:%s", 105010, t2sdk.GetErrorMsg(conn, ret))
		} else {
			log.Printf("SyncSendMsg Ok,ret:%d", ret)

			ret1, p := t2sdk.SyncRecvBiz(conn, ret)

			log.Printf("SyncRecvBiz Ok,ret1:%d,unpacker:%x", ret1, p)

			c1 := t2sdk.GetDatasetCount(p)
			log.Printf("DatasetCount:%d", c1)
			for x := int32(0); x < c1; x++ {
				c2 := t2sdk.GetColCount(p)
				for y := int32(0); y < c2; y++ {
					c_name := t2sdk.GetColName(p, y)
					log.Printf(" %s", c_name)
				}
				log.Println("")
				row_count := t2sdk.GetRowCount(p)
				for y := int32(0); y < row_count; y++ {
					for z := int32(0); z < c2; z++ {
						c_type := t2sdk.GetColType(p, z)
						switch c_type {
						case "S":
							//log.Printf(" %s", t2sdk.GetStrByIndex(p, z))
						case "C":
							//log.Printf(" %s", t2sdk.GetCharByIndex(p, z))
						case "I":
							//log.Printf(" %d", t2sdk.GetIntByIndex(p, z))
						case "F":
							//log.Printf(" %f", t2sdk.GetDoubleByIndex(p, z))
						default:
							log.Printf(" %s", c_type)
						}
					}
					t2sdk.Next(p)
					//log.Println("")
				}
				log.Printf("row_count:%d", row_count)
			}
		}
	}

	for{

	}
}
