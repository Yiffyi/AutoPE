#include-once
#include <Debug.au3>
#include <WinAPI.au3>

#include "Common.au3"
#include "GoodIni.au3"

Func UNCMap($sUNCPath, $sDrive, $sUser, $sPwd)
	;Local $sUNCServerShare = _WinAPI_PathStripToRoot($sUNCPath), $sRelPath = _WinAPI_PathSkipRoot($sUNCPath)

	;_DebugReportVar("$sUNCServerShare", $sUNCServerShare)
	;_DebugReportVar("$sRelPath", $sRelPath)

	If DriveMapGet($sDrive) <> $sUNCPath Then
		DriveMapDel($sDrive)
		Local $ret = 0
		While True
			$ret = DriveMapAdd($sDrive, $sUNCPath, $DMA_DEFAULT, $sUser, $sPwd)
			LogD("DriveMapAdd: " & $ret)
			Switch String($ret)
				Case "0", ""
					LogW("映射共享错误，1s 后重试")
					Sleep(1000)
				Case "1"
					ExitLoop
				Case Else
					$sDrive = $ret
			EndSwitch
		WEnd
	EndIf
	LogI($sUNCPath & " 已挂载为 " & $sDrive)
	Return $sDrive
	;Return _WinAPI_PathAppend($sDrive, $sExeRelPath)
EndFunc


Func LoadUNCIni($sIniConfig)
	FixBadIniFile($sIniConfig)
	Local $aArray = IniReadSectionNames($sIniConfig)
	For $i = 1 To $aArray[0]
		$sSection = $aArray[$i]
		If StringLeft($sSection, 2) <> "共享" Then ContinueLoop
		Local $sUNCPath = IniRead($sIniConfig, $sSection, "地址", "")
		If StringLeft($sUNCPath, 2) <> "\\" Then
			Failed($sSection & "不是Windows共享")
		EndIf

		Local $sUser = IniRead($sIniConfig, $sSection, "用户名", "")
		Local $sPwd = IniRead($sIniConfig, $sSection, "密码", "")
		Local $sDrive = IniRead($sIniConfig, $sSection, "盘符", "*")

		UNCMap($sUNCPath, $sDrive, $sUser, $sPwd)
	Next
EndFunc
