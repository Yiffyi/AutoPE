#AutoIt3Wrapper_OutFile_X64=bin\WatchDism.exe
#AutoIt3Wrapper_UseX64=Y
;#AutoIt3Wrapper_UseUpx=Y
#AutoIt3Wrapper_Res_CompanyName=Yiffyi
#AutoIt3Wrapper_Res_Description=[AutoPE] DISM 执行器
#AutoIt3Wrapper_Res_Fileversion_Use_Template=%YYYY.%MO.%DD
#AutoIt3Wrapper_Res_Fileversion=2023.10.4.0
#AutoIt3Wrapper_Res_HiDpi=Y
#AutoIt3Wrapper_Res_ProductName=AutoPE
#AutoIt3Wrapper_Res_ProductVersion=v2023b2
#AutoIt3Wrapper_Res_Language=2052
#AutoIt3Wrapper_Res_LegalCopyright=©️ 2022-2023 Yiffyi & 春晖中学
#AutoIt3Wrapper_Res_requestedExecutionLevel=requireAdministrator
#AutoIt3Wrapper_Run_Au3Stripper=Y
#RequireAdmin
#NoTrayIcon

#include "lib\Common.au3"

FileChangeDir(@ScriptDir)
InitLog("Debug")
FileDelete("tmp\dismRet")

Switch StringLower($CmdLine[1])
	Case "apply"
		Local $sCmd = 'DISM\DISM.exe /logpath:' & @ScriptDir & '\DISMApply.log' & ' /apply-image /imagefile:"' & $CmdLine[2] & '" /index:' & $CmdLine[3] & " /applydir:I:\"
		LogD("WatchDism: $sCmd="&$sCmd)
		Local $ret = RunWait($sCmd)
		If @error Then
			FileWrite("tmp\dismRet", @error)
			Err("RunWait DISM error: " & @error)
		Else
			FileWrite("tmp\dismRet", $ret)
		EndIf
		Exit $ret
	Case Else
		Err("Invalid action " &$CmdLine[1])
		Exit 1
EndSwitch
