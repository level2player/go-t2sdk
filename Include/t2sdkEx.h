#include "t2sdk_interface.h"

extern "C"
{
   CConfigInterface*  GetNewConfig(){
        return NewConfig();
   }

   int LoadT2Config(CConfigInterface* config,const char *szFileName){
      return config->Load(szFileName);
   }

   CConnectionInterface* GetNewConnection(CConfigInterface* config){
       return NewConnection(config);
   }

   int CreateConnection(CConnectionInterface* conn){
      return conn->Create(NULL);
   }

   int T2Connect(CConnectionInterface* conn, int timeOut){
       return conn->Connect(timeOut);
   }

    const char* GetErrorMsg(CConnectionInterface* conn, int retCode){
        return conn->GetErrorMsg(retCode);
    }

   IF2Packer* GetNewPacker(int version){
       return NewPacker(version);
   }

   int AddPackerRef(IF2Packer* packer){
       return packer->AddRef();
   }

   void BeginPack(IF2Packer* packer){
       packer->BeginPack();
   }

   void EndPack(IF2Packer* packer){
        packer->EndPack();
   }

   int AddField(IF2Packer* packer,const char *szFieldName){
      return packer->AddField(szFieldName);
   }

   int AddStr(IF2Packer* packer,const char *szValue ){
        return packer->AddStr(szValue);
   }

   int AddInt(IF2Packer* packer,int szValue ){
        return packer->AddInt(szValue);
   }

   int AddDouble(IF2Packer* packer,double szValue ){
        return packer->AddDouble(szValue);
   }

   int AddChar(IF2Packer* packer,char szValue ){
        return packer->AddChar(szValue);
   }

   int SyncSendBiz(CConnectionInterface* conn,int functionNo,IF2Packer* packer){
      return conn->SendBiz(functionNo,packer);
   }

   int SyncRecvBiz(CConnectionInterface* conn,int retCode,void** pointer){
      return conn->RecvBiz(retCode,pointer);
   }

   int GetDatasetCount(void** pointer){
      return ((IF2UnPacker *)pointer)->GetDatasetCount();
   }
   int SetCurrentDatasetByIndex(void** pointer,int nIndex){
      return ((IF2UnPacker *)pointer)->SetCurrentDatasetByIndex(nIndex);
   }


   int GetColCount(void** pointer){
      return ((IF2UnPacker *)pointer)->GetColCount();
   }

    const char* GetColName(void** pointer,int column){
      return ((IF2UnPacker *)pointer)->GetColName(column);
   }

}