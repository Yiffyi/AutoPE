#include-once
#include <WinAPI.au3>

#include "Common.au3"

Func _WinAPI_RegUnloadKey($hKey, $sSubkey)
    _HiveLoad_Privilege()
    If @error Then Return SetError(@error)

    Local $aRet = DllCall('advapi32.dll', 'long', 'RegUnLoadKeyW', 'handle', $hKey, 'wstr', $sSubkey)
    If @error Then Return SetError(@error, @extended, 0)
    Return SetError($aRet[0], 0, $aRet[0] = 0)
EndFunc

Func _WinAPI_RegLoadKey($hKey, $sSubkey, $sFile)
    _HiveLoad_Privilege()
    If @error Then Return SetError(@error)

    Local $aRet = DllCall('advapi32.dll', 'long', 'RegLoadKeyW', 'handle', $hKey, 'wstr', $sSubkey, 'wstr', $sFile)
    If @error Then Return SetError(@error, @extended, 0)
    Return SetError($aRet[0], 0, $aRet[0] = 0)
EndFunc

Func _HiveLoad_Privilege()
	Local $aPrivileges[2] = [$SE_BACKUP_NAME, $SE_RESTORE_NAME]

	; Enable "SeBackupPrivilege" and "SeRestorePrivilege" privileges to save and restore registry hive
	Local $hToken = _WinAPI_OpenProcessToken(BitOR($TOKEN_ADJUST_PRIVILEGES, $TOKEN_QUERY))
	Local $aAdjust
	_WinAPI_AdjustTokenPrivileges($hToken, $aPrivileges, $SE_PRIVILEGE_ENABLED, $aAdjust)
	If @error Or @extended Then
		Failed('_HiveLoad_Privilege' & @TAB & '权限提升错误')
		SetError(@error, @extended)
	EndIf
	_WinAPI_CloseHandle($hToken)
EndFunc
