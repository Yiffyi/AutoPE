#include-once
#include <InetConstants.au3>
#include <FileConstants.au3>
#include <Debug.au3>
#include <File.au3>
#include <WinAPIShPath.au3>

Global $bIsPE = True
RegRead("HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\WinPE", "Version")
If @error Then $bIsPE = False

Global $g_hLogFile, $g_sLogPath, $g_nLogLevel
Global Const $g_ProgramName = _WinAPI_PathRemoveExtension(@ScriptName), $g_ProgramTitle = "[AutoPE] " & $g_ProgramName

Func InitLog($sLevel, $sFileName = "")
	If $sFileName = "" Then
		$sFileName = @ScriptDir & "\logs\" & _WinAPI_PathRenameExtension(@ScriptName, ".log")
	EndIf
	Switch StringLower($sLevel)
		Case "debug"
			$g_nLogLevel = 0
		Case "info"
			$g_nLogLevel = 1
		Case "warning"
			$g_nLogLevel = 2
		Case "error"
			$g_nLogLevel = 3
		Case Else
			$g_nLogLevel = 0
	EndSwitch
	$g_sLogPath = $sFileName
	$g_hLogFile = FileOpen($sFileName, BitOR($FO_APPEND, $FO_CREATEPATH))
	OnAutoItExitRegister("_closeLog")
	_writeLog("[D]", "> Logger started <" & @CRLF & _DebugBugReportEnv())
	LogW("====== " & $g_ProgramName & "，启动！ ======")
EndFunc

Func _writeLog($sLevel, $sMsg)
	_FileWriteLog($g_hLogFile, $sLevel & ' ' & $sMsg)
EndFunc

Func _closeLog()
	FileClose($g_hLogFile)
EndFunc

Func LogD($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended) ;0
	If $g_nLogLevel <= 0 Then
		_writeLog("[D]", $s)
	EndIf
	SetError($_iCallerError, $_iCallerExtended)
EndFunc

Func LogI($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended) ;1
	If $g_nLogLevel <= 1 Then
		_writeLog("[I]", $s)
	EndIf
	SetError($_iCallerError, $_iCallerExtended)
EndFunc

Func LogW($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended) ;2
	If $g_nLogLevel <= 2 Then
		_writeLog("[W]", $s)
	EndIf
	SetError($_iCallerError, $_iCallerExtended)
EndFunc

Func LogE($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended) ;3
	If $g_nLogLevel <= 3 Then
		_writeLog("[E]", $s)
	EndIf
	SetError($_iCallerError, $_iCallerExtended)
EndFunc


Func Err($e)
	LogW($e)
	MsgBox($MB_ICONWARNING, $g_ProgramTitle, $e)
EndFunc


Func Failed($e)
	LogE($e)
	_closeLog()
	MsgBox($MB_ICONERROR, $g_ProgramTitle, $e)
	If MsgBox(BitOR($MB_ICONERROR, $MB_YESNO), $g_ProgramTitle, "是否将日志文件保存到U盘？") = $IDYES Then
		SplashTextOn($g_ProgramTitle, "等待U盘插入……")
		Local $aArray = DriveGetDrive($DT_REMOVABLE), $cnt = 1
		While $aArray[0] < 1 And $cnt <= 30
			Sleep(1000)
			$aArray = DriveGetDrive($DT_REMOVABLE)
			$cnt = $cnt + 1
		WEnd
		If $aArray[0] >= 1 Then
			For $i = 1 To $aArray[0]
				DirCopy(@ScriptDir & "\logs", $aArray[$i] & "\autope_failed_logs", $FC_OVERWRITE)
			Next
		Else
			SplashTextOn($g_ProgramTitle, "三十秒内未检测到U盘，停止等待")
		EndIf
	EndIf
	Exit 1
EndFunc


Func RunWaitCatch($cmd, $expected, $sThrow)
	LogD(StringFormat('RunWait: %s, %d, %s', $cmd, $expected, $sThrow))

	Local $ret = RunWait($cmd)
	If $ret <> $expected Or @error Then
		LogD("RunWait: error " & $ret)
		Err(StringFormat("命令运行错误: %s\n$cmd: %s\n$ret: %d, @error: %d", $sThrow, $cmd, $ret, @error))
		Return False
	EndIf

	LogD('RunWait: Finished, ' & $expected)
	Return True
EndFunc


Func Yes($s)
	If $s = "是" Then Return True
	Switch StringLower($s)
		Case "true", "t", "y", "yes"
			Return True
		Case Else
			Return False
	EndSwitch
EndFunc


Func No($s)
	If $s = "否" Then Return True
	Switch StringLower($s)
		Case "false", "f", "n", "no"
			Return True
		Case Else
			Return False
	EndSwitch
EndFunc


Func Lookfor($sPath)
	If StringLen($sSrc) = 0 Then Return ""

	If FileExists($sPath) Then
		Return $sPath
	Else
		LogW($sPath & " 没有在工作目录中找到 " & @WorkingDir)
		Local $a = DriveGetDrive($DT_ALL)
		For $i = 1 To $a[0]
			If FileExists($a[$i] & '\' & $sPath) Then
				LogI($sPath & " 找到了 " & $a[$i])
				Return $a[$i] & '\' & $sPath
			EndIf
		Next
		Err("没有找到 " & $sPath)
		Return ""
	EndIf
EndFunc


Func DownloadOrCopyFile($sSrc, $sDst)
	Local $sSchema = StringLeft($sSrc, 4)
	If $sSchema = "http" Or $sSchema = "ftp:" Then
		Local $hDownload = InetGet($sSrc, $sDst , $INET_FORCERELOAD + $INET_BINARYTRANSFER + $INET_FORCEBYPASS, $INET_DOWNLOADWAIT)
		Local $e = InetGetInfo($hDownload, $INET_DOWNLOADERROR)
		If $e Or @error Then
			Err("下载 " & $sSrc & " 时出错: " & $e & @error)
			FileDelete($sDst)
			Return False
		EndIf

		LogI($sSrc & " 已下载")
		Return True
	Else
		Local $s = Lookfor($sSrc)
		If StringLen($s) > 0 Then
			Return FileCopy($s, $sDst)
		Else
			Return False
		EndIf
	EndIf
EndFunc
