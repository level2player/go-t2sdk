package t2sdk

import (
	"fmt"
	"testing"
)

type Request_105010 struct {
	FunctionNo int32
}

func (req Request_105010) ParseRequestContent() uintptr {
	packer := GetNewPacker(2)
	AddPackerRef(packer)
	BeginPack(packer)
	AddField(packer, "function_id")
	//AddField(packer, "config_no")
	//AddField(packer, "branch_no")
	AddInt(packer, req.FunctionNo)
	//AddInt(packer, 683501)
	//AddInt(packer, 100)
	EndPack(packer)
	return packer
}

// func TestSyncSendMsg(t *testing.T) {
// 	config := GetNewConfig()
// 	iniFile := "t2sdk.ini"
// 	ret := LoadT2Config(config, iniFile)
// 	if ret != 0 {
// 		t.Errorf("T2 Config Read Error,ErrorNo:%d", ret)
// 	}
// 	conn := GetNewConnection(config)
// 	ret = InitCreat(conn)
// 	if ret != 0 {
// 		t.Errorf("InitCreat Error,ErrorNo:%d", ret)
// 	}
// 	ret_conn := T2Connect(conn, 1000)
// 	if ret_conn != 0 {
// 		err_info := GetErrorMsg(conn, ret_conn)
// 		t.Errorf("T2Connect Error,ErrorNo:%d,ErrorInfo:%s", ret_conn, err_info)
// 	} else {
// 		t.Log("T2 Connection Ok...")

// 		var iq IRequest = Request_105010{105010}
// 		ret = SyncSendMsg(conn, 105010, iq)
// 		if ret < 0 {
// 			t.Errorf("syncSendMsg Error,FunctionNo:%d,Error:%s", 105010, GetErrorMsg(conn, ret))
// 		} else {
// 			_, p := SyncRecvBiz(conn, ret)
// 			c1 := GetDatasetCount(p)
// 			for x := int32(0); x < c1; x++ {
// 				c2 := GetColCount(p)
// 				for y := int32(0); y < c2; y++ {
// 					c_name := GetColName(p, y)
// 					fmt.Printf(" %s", c_name)
// 				}
// 				fmt.Println("")
// 				row_count := GetRowCount(p)
// 				for y := int32(0); y < row_count; y++ {
// 					for z := int32(0); z < c2; z++ {
// 						c_type := GetColType(p, z)
// 						switch c_type {
// 						case "S":
// 							fmt.Printf(" %s", GetStrByIndex(p, z))
// 						case "C":
// 							fmt.Printf(" %s", GetCharByIndex(p, z))
// 						case "I":
// 							fmt.Printf(" %d", GetIntByIndex(p, z))
// 						case "F":
// 							fmt.Printf(" %f", GetDoubleByIndex(p, z))
// 						default:
// 							fmt.Printf(" %s", c_type)
// 						}
// 					}
// 					Next(p)
// 					fmt.Println("")
// 				}
// 				t.Logf("dataTable,row_count:%d", row_count)
// 			}
// 			FreeMem(p)
// 			Release(p)
// 		}
// 	}
// 	Release(conn)
// 	Release(config)

// }

// func TestXXX(t *testing.T){
// 	AskForCallback()
// }

func TestAsynSendMsg(t *testing.T) {
	config := GetNewConfig()
	iniFile := "t2sdk.ini"
	ret := LoadT2Config(config, iniFile)
	if ret != 0 {
		t.Errorf("T2 Config Read Error,ErrorNo:%d", ret)
	}
	conn := GetNewConnection(config)

	ch := make(chan int32)

	v := func(CConnectionInterface uintptr, hSend int32, lpUnPackerOrStr uintptr, nResult int32) uintptr {
		fmt.Printf("revice msg,%x,", CConnectionInterface)
		c1 := GetDatasetCount(lpUnPackerOrStr)
		for x := int32(0); x < c1; x++ {
			c2 := GetColCount(lpUnPackerOrStr)
			for y := int32(0); y < c2; y++ {
				c_name := GetColName(lpUnPackerOrStr, y)
				fmt.Printf(" %s", c_name)
			}
			fmt.Println("")
			row_count := GetRowCount(lpUnPackerOrStr)
			for y := int32(0); y < row_count; y++ {
				for z := int32(0); z < c2; z++ {
					c_type := GetColType(lpUnPackerOrStr, z)
					switch c_type {
					case "S":
						fmt.Printf(" %s", GetStrByIndex(lpUnPackerOrStr, z))
					case "C":
						fmt.Printf(" %s", GetCharByIndex(lpUnPackerOrStr, z))
					case "I":
						fmt.Printf(" %d", GetIntByIndex(lpUnPackerOrStr, z))
					case "F":
						fmt.Printf(" %f", GetDoubleByIndex(lpUnPackerOrStr, z))
					default:
						fmt.Printf(" %s", c_type)
					}
				}
				Next(lpUnPackerOrStr)
				fmt.Println("")
			}
			t.Logf("dataTable,row_count:%d", row_count)
		}
		ch <- nResult
		return 0
	}

	RegisteredCallBack(v)

	ret = InitCreat(conn)
	if ret != 0 {
		t.Errorf("InitCreat Error,ErrorNo:%d", ret)
	}
	ret_conn := T2Connect(conn, 1000)
	if ret_conn != 0 {
		err_info := GetErrorMsg(conn, ret_conn)
		t.Errorf("T2Connect Error,ErrorNo:%d,ErrorInfo:%s", ret_conn, err_info)
	} else {
		t.Log("T2 Connection Ok...")

		var iq IRequest = Request_105010{105010}

		ret = AsynSendBiz(conn, 105010, iq)
		if ret < 0 {
			t.Errorf("AsynSendBiz Error,FunctionNo:%d,Error:%s", 105010, GetErrorMsg(conn, ret))
			return
		}
	}
	//await asyn msg
	for {
		<-ch
		break
	}
	Release(conn)
	Release(config)

}
