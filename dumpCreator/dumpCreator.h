#pragma once

#ifdef DUMPCREATOR_EXPORTS
#define DUMPCREATOR_API __declspec(dllexport)
#else
#define DUMPCREATOR_API __declspec(dllimport)
#endif

extern "C" DUMPCREATOR_API int DumpProcessImpl(DWORD processId);
extern "C" DUMPCREATOR_API BOOL SetPrivilege(HANDLE hToken, LPCTSTR lpszPrivilege, BOOL bEnablePrivilege);