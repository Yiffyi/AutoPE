#AutoIt3Wrapper_UseX64=Y
#AutoIt3Wrapper_UseUpx=Y
#AutoIt3Wrapper_Res_CompanyName=Yiffyi
#AutoIt3Wrapper_Res_Description=AutoPE HTTP 服务启动器
#AutoIt3Wrapper_Res_Fileversion_Use_Template=%YYYY.%MO.%DD
#AutoIt3Wrapper_Res_Fileversion=2023.1.28.0
#AutoIt3Wrapper_Res_HiDpi=Y
#AutoIt3Wrapper_Res_ProductName=AutoPE
#AutoIt3Wrapper_Res_ProductVersion=v2023b1
#AutoIt3Wrapper_Res_Language=2052
#AutoIt3Wrapper_Res_LegalCopyright=©️ 2022-2023 Yiffyi & 春晖中学
#AutoIt3Wrapper_Res_requestedExecutionLevel=requireAdministrator
#AutoIt3Wrapper_Run_Au3Stripper=Y

#include <AutoItConstants.au3>
#include <FileConstants.au3>
#include <MsgBoxConstants.au3>
#include <Timers.au3>
#include "gui.au3"

Local $sFolder = @WorkingDir
_Timer_SetTimer($Form1, 1000, "UpdateProcessState")
GUICtrlSetData($iDir, @WorkingDir)
While 1
	$nMsg = GUIGetMsg()
	Switch $nMsg
		Case $GUI_EVENT_CLOSE
			Exit

		Case $bBrowse
			$sFolder = FileSelectFolder("选择一个文件夹", "")
			If @error Then
				; Display the error message.
				MsgBox($MB_SYSTEMMODAL, "春晖系统部署", "没有选择目录")
			Else
				FileChangeDir($sFolder)
				; Display the selected folder.
				GUICtrlSetData($iDir, @WorkingDir)
				MsgBox($MB_SYSTEMMODAL, "春晖系统部署", "你选择了:" & @CRLF & $sFolder)
			EndIf

		Case $bStart
			FileInstall("caddy_windows_amd64.exe", "caddy.exe", $FC_OVERWRITE)
			Run(@ComSpec & ' /c "caddy.exe file-server --browse --listen :' & GUICtrlRead($iPort) & ' || pause"')
		Case $bStop
			ProcessClose("caddy.exe")
	EndSwitch
WEnd

Func UpdateProcessState($hWnd, $iMsg, $iIDTimer, $iTime)
	Local $aMemory = ProcessGetStats("caddy.exe", $PROCESS_STATS_MEMORY)
	Local $aIO = ProcessGetStats("caddy.exe", $PROCESS_STATS_IO)
	If @error Then
		GUICtrlSetData($lStatus, "已停止")
		GUICtrlSetColor($lStatus, 0xFF0000) ;红色
		Return
	EndIf

	GUICtrlSetData($lStatus, "运行中")
	GUICtrlSetColor($lStatus, 0x008000) ;绿色

	GUICtrlSetData($lRead, "读：" & ByteSuffix($aIO[3]))
	GUICtrlSetData($lWrite, "写：" & ByteSuffix($aIO[4]))
	GUICtrlSetData($lMemory, "内存占用：" & ByteSuffix($aMemory[0]))
EndFunc

Func ByteSuffix($iBytes)
	Local $iIndex = 0, $aArray = [' 字节', ' KB', ' MB', ' GB', ' TB', ' PB', ' EB', ' ZB', ' YB']
	While $iBytes > 1023
		$iIndex += 1
		$iBytes /= 1024
	WEnd
	Return Round($iBytes) & $aArray[$iIndex]
EndFunc
