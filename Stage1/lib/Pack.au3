#include-once

#include <WinAPIFiles.au3>
#include <AutoItConstants.au3>

#include "Common.au3"

Func RunPackEntry($sDir)
	LogD("Searching " & $sDir)
	Local $hSearch = FileFindFirstFile($sDir & "\AutoPEPack.*")
	; Assign a Local variable the empty string which will contain the files names found.
	Local $sFileName = ""
	While True
		$sFileName = FileFindNextFile($hSearch)
		; If there is no more file matching the search.
		If @error Then ExitLoop
		LogD("$sFileName=" & $sFileName)
		LogD("@extended=" & @extended)
		LogD("_WinAPI_PathIsExe=" & _WinAPI_PathIsExe($sFileName))
		If Not @extended And _WinAPI_PathIsExe($sFileName) Then
			Return RunWait($sDir & '\' & $sFileName, $sDir)
		EndIf
	WEnd
	FileClose($hSearch)
	$hSearch = FileFindFirstFile($sDir & "\*")
	$sFileName = ""
	While True
		$sFileName = FileFindNextFile($hSearch)
		; If there is no more file matching the search.
		If @error Then ExitLoop
		LogD("$sFileName=" & $sFileName)
		If @extended Then
			Local $ret = RunPackEntry($sDir & '\' & $sFileName)
			If $ret <> -1 Then
				Return $ret
			EndIf
		EndIf
	WEnd
	FileClose($hSearch)
	Return -1
EndFunc


Func Apply7zPack($sPackPath, $s7z, $sTmp)
	$sTmp = _WinAPI_GetFullPathName($sTmp) ; RunWait Workdir only accepts full path
	$sPackPath = _WinAPI_GetFullPathName($sPackPath)

	LogI("Applying pack: " & $sPackPath)

	DirRemove($sTmp, 1)
	DirCreate($sTmp)

	If RunWait($s7z & ' x "' & $sPackPath & '"', $sTmp) <> 0 Then
		;Err("Extracting pack failed")
		DirRemove($sTmp, 1)
		Return False
	EndIf
	If RunPackEntry($sTmp) <> 0 Then
		LogE($sPackPath & " failed")
	EndIf
	Return DirRemove($sTmp, 1)
EndFunc

Func ApplyISOPack($sPackPath, $sWinCDEmu)
	;$sTmp = _WinAPI_GetFullPathName($sTmp)
	;DirRemove($sTmp, 1)
	;DirCreate($sTmp)
	;Local $ret = RunWait(StringFormat('imdisk -a -f "%s" -m "%s"', $sPackPath, $sTmp))
	Local $ret = RunWait(StringFormat('%s "%s" /wait', $sWinCDEmu, $sPackPath))
	If $ret <> 0 Then
		Err("Could not mount " & $sPackPath & ": " & $ret)
		Return False
	EndIf

	Local $aArray = DriveGetDrive("CDROM")
	Local $ret = -1
	For $i = 1 To $aArray[0]
		$ret = RunPackEntry($aArray[$i] & '\')
		If $ret <> -1 Then
			ExitLoop
		EndIf
	Next

	If $ret <> -1 Then
		If $ret <> 0 Then
			Err($sPackPath & " failed")
		EndIf
	Else
		Err($sPackPath & " pack entry not found")
	EndIf

	;$ret = RunWait(StringFormat('imdisk -D -m "%s"', $sTmp))
	$ret = RunWait(StringFormat('%s /unmount "%s"', $sWinCDEmu, $sPackPath))
	If $ret <> 0 Then
		Err("Could not detach " & $sPackPath & ": " & $ret)
		Return False
	EndIf

	Return True
EndFunc

