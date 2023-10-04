#AutoIt3Wrapper_OutFile_X64=bin\WindeployShim.exe
#AutoIt3Wrapper_Res_Description=[AutoPE] WindeployShim

#AutoIt3Wrapper_UseX64=Y
;#AutoIt3Wrapper_UseUpx=Y
#AutoIt3Wrapper_Res_CompanyName=Yiffyi
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
#include "lib\GoodIni.au3"
#include "lib\WinSplashText.au3"

;Run("CmdHolder.exe")
;MsgBox(0, "", "WindeployShim!")
FileChangeDir(@ScriptDir)
InitLog("Debug")

Global $sIniConfig = @ScriptDir & "\Main.ini"
FixBadIniFile($sIniConfig)

WinSetOnTop("[CLASS:FirstUXWndClass]", "", $WINDOWS_NOONTOP)
If Yes(IniRead($sIniConfig, "Debug?", "Debug!", "false")) Then
	Run("cmd.exe")
	SetWinSplashText("[Shim] 前面的系统，以后再来安装吧！")
	Err("已暂停")
EndIf

Local $sDots = ".", $nAnimLoop = 0
While $nAnimLoop < 5
	SetWinSplashText("欢迎使用 AutoPE " & $sDots)
	If StringLen($sDots) < 3 Then
		$sDots = $sDots & "."
	Else
		$sDots = "."
	EndIf
	Sleep(500)
	$nAnimLoop = $nAnimLoop + 1
WEnd

RunWait(@ScriptDir & "\Stage3.exe PreWindeploy", @ScriptDir)

Local $sNext = RegRead("HKEY_LOCAL_MACHINE\SYSTEM\Setup", "_CmdLine")

If StringLen($sNext) > 0 Then
	RunWait($sNext)
EndIf

RunWait(@ScriptDir & "\Stage3.exe PostWindeploy", @ScriptDir)

$sNext = RegRead("HKEY_LOCAL_MACHINE\SYSTEM\Setup", "CmdLine")
If StringLen($sNext) > 0 And StringLeft($sNext, 9) <> "\$~AutoPE" Then
	RegWrite("HKEY_LOCAL_MACHINE\SYSTEM\Setup", "_CmdLine", "REG_SZ", $sNext)
	RegWrite("HKEY_LOCAL_MACHINE\SYSTEM\Setup", "CmdLine", "REG_SZ", "\$~AutoPE\WindeployShim.exe")
Else
	RegDelete("HKEY_LOCAL_MACHINE\SYSTEM\Setup", "_CmdLine")
EndIf


;LoadNetCfgFromIni(@ScriptDir & "\NetworkCfg.ini")

;_WinAPI_RegLoadKey($HKEY_USERS, "OfflineSOFTWARE", "I:\Windows\System32\config\SOFTWARE")
;RegWrite("HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon", "Userinit", "REG_SZ", "\$~AutoPE\UserinitShim.exe")
;_WinAPI_RegUnloadKey($HKEY_USERS, "OfflineSOFTWARE")

;MsgBox(0, "", "Windeploy exited!")
