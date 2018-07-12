package t2sdk

//#include <stdlib.h>
import (
	"C"
)

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	//"log"
	"syscall"
	"unsafe"
)

var (
	//dll
	t2sdkEx    uintptr
	t2sdkExDll *syscall.DLL
	// Functions

	getConfig                uintptr
	getNewConnection         uintptr
	loadT2Config             uintptr
	initCreat                uintptr
	t2Connect                uintptr
	syncSendBiz              uintptr
	getNewPacker             uintptr
	addPackerRef             uintptr
	getErrorMsg              uintptr
	beginPack                uintptr
	addField                 uintptr
	addStr                   uintptr
	addInt                   uintptr
	addDouble                uintptr
	addChar                  uintptr
	endPack                  uintptr
	recvBiz                  uintptr
	getDatasetCount          uintptr
	setCurrentDatasetByIndex uintptr
	getColCount              uintptr
	getColName               uintptr

	getRowCount        uintptr
	getColType         uintptr
	getIntByIndex      uintptr
	getCharByIndex     uintptr
	getDoubleByIndex   uintptr
	getStrByIndex      uintptr
	next               uintptr
	release            uintptr
	freeMem            uintptr
	asynSendBiz        uintptr
	registeredCallBack *syscall.Proc

	getNewPublisher             uintptr
	getMCLastError              uintptr
	pubMsgByPacker              uintptr
	getNewSubscriber            uintptr
	getNewSubscribeParam        uintptr
	setTopicName                uintptr
	setFromNow                  uintptr
	setReplace                  uintptr
	setAppData                  uintptr
	setFilter                   uintptr
	setSendInterval             uintptr
	subscribeTopic              uintptr
	cancelSubscribeTopic        uintptr
	getUnPack                   uintptr
	registeredSubscribeCallBack *syscall.Proc
)

func MustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}
func MustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}
	return uintptr(lib)
}

func init() {
	t2sdkEx = MustLoadLibrary("../lib/t2sdkEx.dll")
	t2sdkExDll = syscall.MustLoadDLL("../lib/t2sdkEx.dll")

	getConfig = MustGetProcAddress(t2sdkEx, "GetNewConfig")
	getNewConnection = MustGetProcAddress(t2sdkEx, "GetNewConnection")
	loadT2Config = MustGetProcAddress(t2sdkEx, "LoadT2Config")
	initCreat = MustGetProcAddress(t2sdkEx, "CreateConnection")
	t2Connect = MustGetProcAddress(t2sdkEx, "T2Connect")
	syncSendBiz = MustGetProcAddress(t2sdkEx, "SyncSendBiz")
	recvBiz = MustGetProcAddress(t2sdkEx, "SyncRecvBiz")
	getErrorMsg = MustGetProcAddress(t2sdkEx, "GetErrorMsg")
	getNewPacker = MustGetProcAddress(t2sdkEx, "GetNewPacker")
	addPackerRef = MustGetProcAddress(t2sdkEx, "AddPackerRef")
	beginPack = MustGetProcAddress(t2sdkEx, "BeginPack")
	endPack = MustGetProcAddress(t2sdkEx, "EndPack")
	addField = MustGetProcAddress(t2sdkEx, "AddField")
	addStr = MustGetProcAddress(t2sdkEx, "AddStr")
	addInt = MustGetProcAddress(t2sdkEx, "AddInt")
	addDouble = MustGetProcAddress(t2sdkEx, "AddDouble")
	addChar = MustGetProcAddress(t2sdkEx, "AddChar")
	getDatasetCount = MustGetProcAddress(t2sdkEx, "GetDatasetCount")
	setCurrentDatasetByIndex = MustGetProcAddress(t2sdkEx, "SetCurrentDatasetByIndex")
	getColCount = MustGetProcAddress(t2sdkEx, "GetColCount")
	getColName = MustGetProcAddress(t2sdkEx, "GetColName")
	getRowCount = MustGetProcAddress(t2sdkEx, "GetRowCount")
	getColType = MustGetProcAddress(t2sdkEx, "GetColType")
	getIntByIndex = MustGetProcAddress(t2sdkEx, "GetIntByIndex")
	getCharByIndex = MustGetProcAddress(t2sdkEx, "GetCharByIndex")
	getDoubleByIndex = MustGetProcAddress(t2sdkEx, "GetDoubleByIndex")
	getStrByIndex = MustGetProcAddress(t2sdkEx, "GetStrByIndex")
	next = MustGetProcAddress(t2sdkEx, "Next")
	release = MustGetProcAddress(t2sdkEx, "Release")
	freeMem = MustGetProcAddress(t2sdkEx, "FreeMem")
	asynSendBiz = MustGetProcAddress(t2sdkEx, "AsynSendBiz")
	registeredCallBack = t2sdkExDll.MustFindProc("RegisteredCallBack")
	getNewPublisher = MustGetProcAddress(t2sdkEx, "GetNewPublisher")
	getMCLastError = MustGetProcAddress(t2sdkEx, "GetMCLastError")
	pubMsgByPacker = MustGetProcAddress(t2sdkEx, "PubMsgByPacker")
	getNewSubscriber = MustGetProcAddress(t2sdkEx, "GetNewSubscriber")
	getNewSubscribeParam = MustGetProcAddress(t2sdkEx, "GetNewSubscribeParam")
	setTopicName = MustGetProcAddress(t2sdkEx, "SetTopicName")
	setFromNow = MustGetProcAddress(t2sdkEx, "SetFromNow")
	setReplace = MustGetProcAddress(t2sdkEx, "SetReplace")
	setAppData = MustGetProcAddress(t2sdkEx, "SetAppData")
	setFilter = MustGetProcAddress(t2sdkEx, "SetFilter")
	setSendInterval = MustGetProcAddress(t2sdkEx, "SetSendInterval")
	subscribeTopic = MustGetProcAddress(t2sdkEx, "SubscribeTopic")
	cancelSubscribeTopic = MustGetProcAddress(t2sdkEx, "CancelSubscribeTopic")
	registeredSubscribeCallBack = t2sdkExDll.MustFindProc("RegisteredSubscribeCallBack")
	getUnPack = MustGetProcAddress(t2sdkEx, "UnPack")
}

func GetNewConfig() uintptr {
	ret, _, _ := syscall.Syscall(getConfig, 0,
		0,
		0,
		0)

	return ret
}

func LoadT2Config(config uintptr, ini_File string) int32 {
	cs := C.CString(ini_File)
	defer C.free(unsafe.Pointer(cs))
	ret, _, _ := syscall.Syscall(loadT2Config, 2,
		config,
		uintptr(unsafe.Pointer(cs)),
		0)
	return int32(ret)
}

func GetNewConnection(config uintptr) uintptr {
	ret, _, _ := syscall.Syscall(getNewConnection, 1,
		config,
		0,
		0)

	return ret
}

func RegisteredCallBack(fun func(uintptr, int32, uintptr, int32) uintptr) int32 {
	cb := syscall.NewCallback(fun)
	r, _, _ := registeredCallBack.Call(cb)
	return int32(r)
}

func RegisteredSubscribeCallBack(fun func(uintptr, int32, uintptr, int32) uintptr) int32 {
	cb := syscall.NewCallback(fun)
	r, _, _ := registeredSubscribeCallBack.Call(cb)
	return int32(r)
}

func InitCreat(conn uintptr) int32 {
	ret, _, _ := syscall.Syscall(initCreat, 1,
		conn,
		0,
		0)
	return int32(ret)

}

func T2Connect(conn uintptr, timeout uint32) int32 {
	ret, _, _ := syscall.Syscall(t2Connect, 2,
		conn,
		uintptr(timeout),
		0)
	return int32(ret)
}

func GetNewPacker(version int32) uintptr {
	ret, _, _ := syscall.Syscall(getNewPacker, 1,
		uintptr(version),
		0,
		0)
	return ret
}

func AddPackerRef(packer uintptr) {
	syscall.Syscall(addPackerRef, 1,
		packer,
		0,
		0)
}

func BeginPack(packer uintptr) {
	syscall.Syscall(beginPack, 1,
		packer,
		0,
		0)
}

func AddField(packer uintptr, file_name string) {
	s := C.CString(file_name)
	defer C.free(unsafe.Pointer(s))
	syscall.Syscall(addField, 2,
		packer,
		uintptr(unsafe.Pointer(s)),
		0)
}

func AddStr(packer uintptr, value string) {
	s := C.CString(value)
	defer C.free(unsafe.Pointer(s))
	syscall.Syscall(addStr, 2,
		packer,
		uintptr(unsafe.Pointer(s)),
		0)
}

func AddInt(packer uintptr, value int32) {
	syscall.Syscall(addInt, 2,
		packer,
		uintptr(value),
		0)
}

func AddDouble(packer uintptr, value float64) {
	syscall.Syscall(addDouble, 2,
		packer,
		uintptr(value),
		0)
}

func AddChar(packer uintptr, value byte) {
	ch := C.char(value)
	syscall.Syscall(addChar, 2,
		packer,
		uintptr(ch),
		0)
}

func EndPack(packer uintptr) {
	syscall.Syscall(endPack, 1,
		packer,
		0,
		0)
}

func SyncSendMsg(conn uintptr, functionNo int32, request IRequest) int32 {
	ptr := request.ParseRequestContent()

	ret, _, _ := syscall.Syscall(syncSendBiz, 3,
		conn,
		uintptr(functionNo),
		ptr)

	return int32(ret)
}

func AsynSendBiz(conn uintptr, functionNo int32, request IRequest) int32 {
	ptr := request.ParseRequestContent()

	ret, _, _ := syscall.Syscall(asynSendBiz, 3,
		conn,
		uintptr(functionNo),
		ptr)

	return int32(ret)
}

func SyncRecvBiz(conn uintptr, retCode int32) (int32, uintptr) {
	var p = unsafe.Pointer(nil)

	ret, _, _ := syscall.Syscall(recvBiz, 3,
		conn,
		uintptr(retCode),
		uintptr(unsafe.Pointer(&p)))

	return int32(ret), uintptr(p)
}

func GetDatasetCount(unpakcer uintptr) int32 {
	ret, _, _ := syscall.Syscall(getDatasetCount, 1,
		unpakcer,
		0,
		0)

	return int32(ret)
}

func SetCurrentDatasetByIndex(unpakcer uintptr, nIndex int32) int32 {
	ret, _, _ := syscall.Syscall(setCurrentDatasetByIndex, 2,
		unpakcer,
		uintptr(nIndex),
		0)
	return int32(ret)
}

func GetColCount(unpakcer uintptr) int32 {
	ret, _, _ := syscall.Syscall(getColCount, 1,
		unpakcer,
		0,
		0)
	return int32(ret)
}

func GetErrorMsg(conn uintptr, recode int32) string {
	size := 0
	buf := make([]byte, 1024+size*1, 2*(1024+size*1))
	s := C.CString(string(buf[:]))
	defer C.free(unsafe.Pointer(s))
	syscall.Syscall(getErrorMsg, 3,
		conn,
		uintptr(recode),
		uintptr(unsafe.Pointer(s)))

	r, _ := simplifiedchinese.GBK.NewDecoder().String(C.GoString(s))
	return r
}

func GetColName(unpakcer uintptr, column int32) string {
	size := 0
	buf := make([]byte, 1024+size*1, 2*(1024+size*1))
	s := C.CString(string(buf[:]))
	defer C.free(unsafe.Pointer(s))

	syscall.Syscall(getColName, 3,
		unpakcer,
		uintptr(column),
		uintptr(unsafe.Pointer(s)))
	return C.GoString(s)
}

func GetRowCount(unpakcer uintptr) int32 {
	ret, _, _ := syscall.Syscall(getRowCount, 1,
		unpakcer,
		0,
		0)

	return int32(ret)
}

func GetColType(unpakcer uintptr, column int32) string {
	ret, _, _ := syscall.Syscall(getColType, 2,
		unpakcer,
		uintptr(column),
		0)

	return string(byte(ret))
}

func GetIntByIndex(unpakcer uintptr, column int32) int32 {
	ret, _, _ := syscall.Syscall(getIntByIndex, 2,
		unpakcer,
		uintptr(column),
		0)

	return int32(ret)
}

func GetCharByIndex(unpakcer uintptr, column int32) string {
	ret, _, _ := syscall.Syscall(getCharByIndex, 2,
		unpakcer,
		uintptr(column),
		0)

	return string(byte(ret))
}

func GetDoubleByIndex(unpakcer uintptr, column int32) float64 {
	ret, _, _ := syscall.Syscall(getDoubleByIndex, 2,
		unpakcer,
		uintptr(column),
		0)
	return float64(ret)
}

func GetStrByIndex(unpakcer uintptr, column int32) string {
	s := C.CString("")
	defer C.free(unsafe.Pointer(s))

	syscall.Syscall(getStrByIndex, 3,
		unpakcer,
		uintptr(column),
		uintptr(unsafe.Pointer(s)))
	r, _ := simplifiedchinese.GBK.NewDecoder().String(C.GoString(s))
	return r
}

func Next(unpakcer uintptr) {
	syscall.Syscall(next, 1,
		unpakcer,
		0,
		0)
}

func Release(known uintptr) int32 {
	r, _, _ := syscall.Syscall(release, 1,
		known,
		0,
		0)
	return int32(r)
}

func FreeMem(packer uintptr) {
	syscall.Syscall(freeMem, 1,
		packer,
		0,
		0)
}

func GetNewPublisher(conn uintptr, pubName string, msgCount int32, timeOut int32) uintptr {
	s := C.CString(pubName)
	defer C.free(unsafe.Pointer(s))
	r, _, _ := syscall.Syscall6(getNewPublisher, 4,
		conn,
		uintptr(unsafe.Pointer(s)),
		uintptr(msgCount),
		uintptr(timeOut), 0, 0)
	return r
}

func GetMCLastError(conn uintptr) string {
	size := 0
	buf := make([]byte, 1024+size*1, 2*(1024+size*1))
	s := C.CString(string(buf[:]))
	defer C.free(unsafe.Pointer(s))
	syscall.Syscall(getMCLastError, 2,
		conn,
		uintptr(unsafe.Pointer(s)), 0)
	r, _ := simplifiedchinese.GBK.NewDecoder().String(C.GoString(s))
	return r
}

func PubMsgByPacker(lpPublish uintptr, topicName string, unPack uintptr, timeOut int32) int32 {
	s := C.CString(topicName)
	defer C.free(unsafe.Pointer(s))
	r, _, _ := syscall.Syscall6(pubMsgByPacker, 4,
		lpPublish,
		uintptr(unsafe.Pointer(s)),
		unPack,
		uintptr(timeOut), 0, 0)
	return int32(r)
}

func GetNewSubscriber(conn uintptr, subName string, timeOut int32) uintptr {
	s := C.CString(subName)
	defer C.free(unsafe.Pointer(s))
	r, _, _ := syscall.Syscall(getNewSubscriber, 3,
		conn,
		uintptr(unsafe.Pointer(s)),
		uintptr(timeOut))
	return r
}

func GetNewSubscribeParam() uintptr {
	r, _, _ := syscall.Syscall(getNewSubscribeParam, 0,
		0,
		0,
		0)
	return r
}
func SetTopicName(subscribeParam uintptr, topicName string) {
	s := C.CString(topicName)

	defer C.free(unsafe.Pointer(s))

	syscall.Syscall(setTopicName, 2,
		subscribeParam,
		uintptr(unsafe.Pointer(s)),
		0)
}
func SetFromNow(subscribeParam uintptr, isFromNow bool) {
	v := 0
	if !isFromNow {
		v = 1
	}
	syscall.Syscall(setFromNow, 2,
		subscribeParam,
		uintptr(v),
		0)
}

func SetReplace(subscribeParam uintptr, isReplace bool) {
	v := 0
	if !isReplace {
		v = 1
	}
	syscall.Syscall(setReplace, 2,
		subscribeParam,
		uintptr(v),
		0)
}

func SetAppData(subscribeParam uintptr, appData string, nAppData int32) {
	s := C.CString(appData)

	defer C.free(unsafe.Pointer(s))

	syscall.Syscall(setAppData, 3,
		subscribeParam,
		uintptr(unsafe.Pointer(s)),
		uintptr(nAppData))
}

func SetFilter(subscribeParam uintptr, filterName string, filterValue string) {
	s_filterName := C.CString(filterName)
	s_filterValue := C.CString(filterValue)
	defer func() {
		C.free(unsafe.Pointer(s_filterName))
		C.free(unsafe.Pointer(s_filterValue))
	}()
	syscall.Syscall(setFilter, 3,
		subscribeParam,
		uintptr(unsafe.Pointer(s_filterName)),
		uintptr(unsafe.Pointer(s_filterValue)))
}

func SetSendInterval(subscribeParam uintptr, sendInterval int32) {
	syscall.Syscall(setSendInterval, 2,
		subscribeParam,
		uintptr(sendInterval),
		0)
}

func SubscribeTopic(lpSub, lpSubscribeParam uintptr, uTimeOut int32) int32 {
	r, _, _ := syscall.Syscall(subscribeTopic, 3,
		lpSub,
		lpSubscribeParam,
		uintptr(uTimeOut))
	return int32(r)
}

func CancelSubscribeTopic(lpSub uintptr, subIndex int32) int32 {
	r, _, _ := syscall.Syscall(cancelSubscribeTopic, 2,
		lpSub,
		uintptr(subIndex),
		0)
	return int32(r)
}

func GetUnPack(packer uintptr) uintptr {
	r, _, _ := syscall.Syscall(getUnPack, 1,
		packer,
		0,
		0)
	return r
}
