#AutoIt3Wrapper_OutFile_X64=bin\Stage3.exe
#AutoIt3Wrapper_Res_Description=[AutoPE] 分阶段执行程序

#AutoIt3Wrapper_UseX64=Y
;#AutoIt3Wrapper_UseUpx=Y
#AutoIt3Wrapper_Res_CompanyName=Yiffyi
#AutoIt3Wrapper_Res_Fileversion_Use_Template=%YYYY.%MO.%DD
#AutoIt3Wrapper_Res_Fileversion=2023.10.6.0
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
#include "lib\NetworkCfg.au3"
#include "lib\UNC.au3"
#include "lib\Pack.au3"
#include "lib\WinSplashText.au3"

Global $sIniConfig = @ScriptDir & "\Main.ini"
FileChangeDir(@ScriptDir)
InitLog("Debug")

FixBadIniFile($sIniConfig)

;MsgBox(0, "[AutoPE] Stage3 " & @UserName, "Stage3 working!")
;MsgBox(0, "[AutoPE] Stage3", $CmdLine[1])

FileInstall("7za.exe", "tmp\7za.exe")
;FileInstall("WinCDEmu.exe", "tmp\WinCDEmu.exe")

Func _apply($sList)
	;RunWait("tmp\WinCDEmu.exe /install")
	;If Not FileExists("tmp\imdiskinst.exe") Then
	;	FileInstall("imdiskinst.exe", "tmp\imdiskinst.exe")
	;	EnvSet("IMDISK_SILENT_SETUP", "1")
	;	RunWait("tmp\imdiskinst.exe")
	;EndIf
	If Not FileExists("tmp\WinCDEmu.exe") Then
		FileInstall("WinCDEmu.exe", "tmp\WinCDEmu.exe")
		RunWait("tmp\WinCDEmu.exe /install")
	EndIf

	If FileExists($sList) Then
		Local $aPacks = FileReadToArray($sList)
		For $i = 0 To @extended-1
			Local $sPackPath = $aPacks[$i]
			SetWinSplashText("[AutoPE] 正在加载 " & $sPackPath)
			Switch StringLower(_WinAPI_PathFindExtension($sPackPath))
				Case ".7z"
					Apply7zPack("tmp\packs\" & $sPackPath, "tmp\7za.exe", "tmp\workdir")
				Case ".iso"
					ApplyISOPack("tmp\packs\" & $sPackPath, "tmp\WinCDEmu.exe")
				Case Else
					LogE($sPackPath & " 包文件拓展名不正确")
			EndSwitch
		Next

		If FileDelete($sList) = 0 Then
			LogW("_apply: 无法删除 " & $sList)
		EndIf
	Else
		LogW("_apply: " & $sList & " 未找到")
	EndIf

EndFunc

;WinSetState("[CLASS:FirstUXWndClass]", "", @SW_HIDE)
Switch StringLower($CmdLine[1])
	Case "prewindeploy"
		SetWinSplashText("[AutoPE] PreWindeploy")
		;_apply("tmp\preWindeployPacks")
	Case "postwindeploy"
		SetWinSplashText("[AutoPE] PostWindeploy")
		;_apply("tmp\postWinDeployPacks")
	#cs
	Case "preoobe"
		SetWinSplashText("[AutoPE] PreOOBE")
		LoadNetCfgFromIni(@ScriptDir & "\NetworkCfg.ini")
		LoadUNCIni($sIniConfig)
		_apply("tmp\preOobePacks")
		SetWinSplashText("[AutoPE] PreOOBE done.")
	Case "postoobe"
		SetWinSplashText("[AutoPE] PostOOBE")
		_apply("tmp\postOobePacks")
		SetWinSplashText("[AutoPE] PostOOBE done.")
	#ce
	Case "specialize"
		SetWinSplashText("[AutoPE] Specialize")
		LoadNetCfgFromIni(@ScriptDir & "\NetworkCfg.ini")
		LoadUNCIni($sIniConfig)
		_apply("tmp\specializePacks")
		SetWinSplashText("[AutoPE] Specialize done.")
	Case "audit"
		_apply("tmp\auditPacks")
	Case "firstlogon"
		_apply("tmp\firstLogonPacks")
	Case Else
		MsgBox(0, "[AutoPE] Stage3", $CmdLine[1] & " not implemented")
EndSwitch

;WinSetState("[CLASS:FirstUXWndClass]", "", @SW_SHOW)
