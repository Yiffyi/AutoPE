#AutoIt3Wrapper_OutFile_X64=bin\MyFirstLogonAnim.exe
#AutoIt3Wrapper_Res_Description=[AutoPE] 覆盖 FirstLogonAnim

#AutoIt3Wrapper_UseX64=Y
;#AutoIt3Wrapper_UseUpx=Y
#AutoIt3Wrapper_Res_CompanyName=Yiffyi
#AutoIt3Wrapper_Res_Fileversion_Use_Template=%YYYY.%MO.%DD
#AutoIt3Wrapper_Res_Fileversion=2023.8.8.0
#AutoIt3Wrapper_Res_HiDpi=Y
#AutoIt3Wrapper_Res_ProductName=AutoPE
#AutoIt3Wrapper_Res_ProductVersion=v2023b2
#AutoIt3Wrapper_Res_Language=2052
#AutoIt3Wrapper_Res_LegalCopyright=©️ 2022-2023 Yiffyi & 春晖中学
;#AutoIt3Wrapper_Res_requestedExecutionLevel=requireAdministrator
#AutoIt3Wrapper_Run_Au3Stripper=Y

;#RequireAdmin
#NoTrayIcon
#include <AutoItConstants.au3>

Local $sMsg = "海内存知己，天涯若比邻。" & @CRLF & @CRLF & "Pretend to be FirstLogonAnim.exe: "
SplashTextOn("[AutoPE] MyFirstLogonAnim", $sMsg & "0s", 500, 200, 0, 0, $DLG_MOVEABLE)

Local $nCur = 0
While True
	ControlSetText("[AutoPE] MyFirstLogonAnim", "", "Static1", $sMsg & $nCur & "s")
	$nCur = $nCur + 1
	Sleep(1000)
WEnd
