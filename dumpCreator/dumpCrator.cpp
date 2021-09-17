#include "pch.h"
#include <iostream>
#include <windows.h>
#include <tlhelp32.h>
#include <process.h>
#include <psapi.h>
#include <minidumpapiset.h>
#include <atlstr.h>
#include <time.h>
#include "dumpCreator.h"

BOOL SetPrivilege(
    HANDLE hToken,          // access token handle
    LPCTSTR lpszPrivilege,  // name of privilege to enable/disable
    BOOL bEnablePrivilege   // to enable or disable privilege
)
{
    TOKEN_PRIVILEGES tp;
    LUID luid;

    if (!LookupPrivilegeValue(
        NULL,            // lookup privilege on local system
        lpszPrivilege,   // privilege to lookup 
        &luid))        // receives LUID of privilege
    {
        return FALSE;
    }

    tp.PrivilegeCount = 1;
    tp.Privileges[0].Luid = luid;
    if (bEnablePrivilege)
        tp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED;
    else
        tp.Privileges[0].Attributes = 0;

    // Enable the privilege or disable all privileges.

    if (!AdjustTokenPrivileges(
        hToken,
        FALSE,
        &tp,
        sizeof(TOKEN_PRIVILEGES),
        (PTOKEN_PRIVILEGES)NULL,
        (PDWORD)NULL))
    {
        return FALSE;
    }

    if (GetLastError() == ERROR_NOT_ALL_ASSIGNED)

    {
        return FALSE;
    }

    return TRUE;
}

int DumpProcessImpl(DWORD processId)
{
    // NOTE: The Windows API MiniDumpWriteDump() used below is not thread safe!
    // Ensure this method is only called by a single thread at a time.

    CString fileName = "postgre_dump.dmp";

    HANDLE hDumpFile = CreateFile(fileName,
        GENERIC_READ | GENERIC_WRITE,
        FILE_SHARE_DELETE | FILE_SHARE_READ | FILE_SHARE_WRITE,
        nullptr,
        CREATE_ALWAYS,
        FILE_ATTRIBUTE_NORMAL,
        nullptr);

    if (hDumpFile == INVALID_HANDLE_VALUE)
    {
        return 0; // error when create file
    }

    HANDLE hToken;
    HANDLE hCurrProcess = GetCurrentProcess();

    if (OpenProcessToken(hCurrProcess, TOKEN_ADJUST_PRIVILEGES, &hToken))
    {
        if (SetPrivilege(hToken, SE_DEBUG_NAME, TRUE))
        {
            CloseHandle(hToken);
        }
    }

    HANDLE hProc = OpenProcess(PROCESS_ALL_ACCESS, TRUE, processId);

    if (hProc == nullptr)
    {
        return 1; //cant't open process
    }

    //////////////////////////////////////////////////////////////////////////
    auto hDbgHelp = LoadLibraryA("dbghelp");
    if (hDbgHelp == nullptr)
        return 5; // can't load dbghelp
    auto pMiniDumpWriteDump = (decltype(&MiniDumpWriteDump))GetProcAddress(hDbgHelp, "MiniDumpWriteDump");
    if (pMiniDumpWriteDump == nullptr)
        return 6; // cant't getprocaddress 

    int result = 4; // assume the worst

    // Dump the sucker.
    if (!pMiniDumpWriteDump(
        hProc,
        processId,
        hDumpFile,
        static_cast<MINIDUMP_TYPE>(MiniDumpWithFullMemory |
            MiniDumpWithFullMemoryInfo |
            MiniDumpWithHandleData |
            MiniDumpWithThreadInfo |
            MiniDumpWithUnloadedModules),
        nullptr,
        nullptr,
        nullptr
    ))
    {
        HRESULT lastError = static_cast<HRESULT>(GetLastError());

        if (lastError != HRESULT_FROM_WIN32(ERROR_CANCELLED)) // Cancelled dump.
        {
            CString msg;
        }
        else
        {
            result = 2; // ERROR_CANCELED
        }
    }
    else
    {
        result = 3; // ERROR_SUCSESS
    }

    if (OpenProcessToken(hCurrProcess, TOKEN_ADJUST_PRIVILEGES, &hToken))
    {
        SetPrivilege(hToken, SE_DEBUG_NAME, FALSE);
        CloseHandle(hToken);
    }

    CloseHandle(hProc);
    CloseHandle(hDumpFile);
    CloseHandle(hCurrProcess);

    return result;
}

