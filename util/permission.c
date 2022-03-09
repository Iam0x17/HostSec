#include <permission.h>

BOOL IsElevated()
{
    BOOL bIsElevated = FALSE;
    HANDLE hToken = NULL;

    if (OpenProcessToken(GetCurrentProcess(), TOKEN_QUERY, &hToken)) {

        struct {
            DWORD TokenIsElevated;
        } te;
        DWORD dwReturnLength = 0;

        if (GetTokenInformation(hToken,(TOKEN_INFORMATION_CLASS)20, &te, sizeof(te), &dwReturnLength)) {
            if (dwReturnLength == sizeof(te))
                bIsElevated = te.TokenIsElevated;
        }
        CloseHandle(hToken);
    }
    return bIsElevated;
}