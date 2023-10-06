#RequireAdmin
#NoTrayIcon
Global Const $FC_OVERWRITE = 1
Global Const $FO_APPEND = 1
Global Const $FO_OVERWRITE = 2
Global Const $FO_CREATEPATH = 8
Global Const $FO_UNICODE = 32
Global Const $OPT_MATCHSTART = 1
Global Const $UBOUND_DIMENSIONS = 0
Global Const $UBOUND_ROWS = 1
Global Const $UBOUND_COLUMNS = 2
Global Const $DT_REMOVABLE = "REMOVABLE"
Global Const $DMA_DEFAULT = 0
Global Const $MB_YESNO = 4
Global Const $MB_ICONERROR = 16
Global Const $MB_ICONWARNING = 48
Global Const $IDYES = 6
Global Const $STR_STRIPLEADING = 1
Global Const $STR_STRIPTRAILING = 2
Global Const $STR_ENTIRESPLIT = 1
Global Const $STR_NOCOUNT = 2
Func _SendMessage($hWnd, $iMsg, $wParam = 0, $lParam = 0, $iReturn = 0, $wParamType = "wparam", $lParamType = "lparam", $sReturnType = "lresult")
Local $aCall = DllCall("user32.dll", $sReturnType, "SendMessageW", "hwnd", $hWnd, "uint", $iMsg, $wParamType, $wParam, $lParamType, $lParam)
If @error Then Return SetError(@error, @extended, "")
If $iReturn >= 0 And $iReturn <= 4 Then Return $aCall[$iReturn]
Return $aCall
EndFunc
Global Const $FORMAT_MESSAGE_ALLOCATE_BUFFER = 0x00000100
Global Const $FORMAT_MESSAGE_IGNORE_INSERTS = 0x00000200
Global Const $FORMAT_MESSAGE_FROM_SYSTEM = 0x00001000
Func _WinAPI_FormatMessage($iFlags, $pSource, $iMessageID, $iLanguageID, ByRef $pBuffer, $iSize, $vArguments)
Local $sBufferType = "struct*"
If IsString($pBuffer) Then $sBufferType = "wstr"
Local $aCall = DllCall("kernel32.dll", "dword", "FormatMessageW", "dword", $iFlags, "struct*", $pSource, "dword", $iMessageID, "dword", $iLanguageID, $sBufferType, $pBuffer, "dword", $iSize, "ptr", $vArguments)
If @error Then Return SetError(@error, @extended, 0)
If Not $aCall[0] Then Return SetError(10, _WinAPI_GetLastError(), 0)
If $sBufferType = "wstr" Then $pBuffer = $aCall[5]
Return $aCall[0]
EndFunc
Func _WinAPI_GetLastError(Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
Local $aCall = DllCall("kernel32.dll", "dword", "GetLastError")
Return SetError($_iCallerError, $_iCallerExtended, $aCall[0])
EndFunc
Func _WinAPI_GetLastErrorMessage(Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
Local $iLastError = _WinAPI_GetLastError()
Local $tBufferPtr = DllStructCreate("ptr")
Local $nCount = _WinAPI_FormatMessage(BitOR($FORMAT_MESSAGE_ALLOCATE_BUFFER, $FORMAT_MESSAGE_FROM_SYSTEM, $FORMAT_MESSAGE_IGNORE_INSERTS), 0, $iLastError, 0, $tBufferPtr, 0, 0)
If @error Then Return SetError(-@error, @extended, "")
Local $sText = ""
Local $pBuffer = DllStructGetData($tBufferPtr, 1)
If $pBuffer Then
If $nCount > 0 Then
Local $tBuffer = DllStructCreate("wchar[" &($nCount + 1) & "]", $pBuffer)
$sText = DllStructGetData($tBuffer, 1)
If StringRight($sText, 2) = @CRLF Then $sText = StringTrimRight($sText, 2)
EndIf
DllCall("kernel32.dll", "handle", "LocalFree", "handle", $pBuffer)
EndIf
Return SetError($_iCallerError, $_iCallerExtended, $sText)
EndFunc
Func _WinAPI_SetLastError($iErrorCode, Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
DllCall("kernel32.dll", "none", "SetLastError", "dword", $iErrorCode)
Return SetError($_iCallerError, $_iCallerExtended, Null)
EndFunc
Global Const $__g_sReportWindowText_Debug = "Debug Window hidden text"
Global $__g_sReportTitle_Debug = "AutoIt Debug Report"
Global $__g_iReportType_Debug = 0
Global $__g_bReportWindowWaitClose_Debug = True, $__g_bReportWindowClosed_Debug = True
Global $__g_hReportEdit_Debug = 0
Global $__g_sReportCallBack_Debug
Global $__g_bReportTimeStamp_Debug = False
Func _DebugBugReportEnv(Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
Local $sAutoItX64, $sAdminMode, $sCompiled, $sOsServicePack, $sMUIlang, $sKBLayout, $sCPUArch
If @AutoItX64 Then $sAutoItX64 = "/X64"
If IsAdmin() Then $sAdminMode = ", AdminMode"
If @Compiled Then $sCompiled = ", Compiled"
If @OSServicePack Then $sOsServicePack = "/" & StringReplace(@OSServicePack, "Service Pack ", "SP")
If @OSLang <> @MUILang Then $sMUIlang = ", MUILang: " & @MUILang
If @OSLang <> StringRight(@KBLayout, 4) Then $sKBLayout = ", Keyboard: " & @KBLayout
If @OSArch <> @CPUArch Then $sCPUArch = ", CPUArch: " & @CPUArch
Return SetError($_iCallerError, $_iCallerExtended, "AutoIt: " & @AutoItVersion & $sAutoItX64 & $sAdminMode & $sCompiled & ", OS: " & @OSVersion & $sOsServicePack & "/" & @OSArch & ", OSLang: " & @OSLang & $sMUIlang & $sKBLayout & $sCPUArch & @CRLF & "  Script: " & @ScriptFullPath)
EndFunc
Func _DebugReport($sData, $bLastError = False, $bExit = False, Const $_iCallerError = @error, $_iCallerExtended = @extended)
If $__g_iReportType_Debug <= 0 Or $__g_iReportType_Debug > 6 Then Return SetError($_iCallerError, $_iCallerExtended, 0)
Local $iLastError = _WinAPI_GetLastError()
__Debug_ReportWrite($sData, $bLastError, $iLastError)
If $bExit Then Exit
_WinAPI_SetLastError($iLastError)
If $bLastError Then $_iCallerExtended = $iLastError
Return SetError($_iCallerError, $_iCallerExtended, 1)
EndFunc
Func __Debug_ReportWindowCreate()
Local $nOld = Opt("WinDetectHiddenText", $OPT_MATCHSTART)
Local $bExists = WinExists($__g_sReportTitle_Debug, $__g_sReportWindowText_Debug)
If $bExists Then
If $__g_hReportEdit_Debug = 0 Then
$__g_hReportEdit_Debug = ControlGetHandle($__g_sReportTitle_Debug, $__g_sReportWindowText_Debug, "Edit1")
$__g_bReportWindowWaitClose_Debug = False
EndIf
EndIf
Opt("WinDetectHiddenText", $nOld)
$__g_bReportWindowClosed_Debug = False
If Not $__g_bReportWindowWaitClose_Debug Then Return 0
Local Const $WS_OVERLAPPEDWINDOW = 0x00CF0000
Local Const $WS_HSCROLL = 0x00100000
Local Const $WS_VSCROLL = 0x00200000
Local Const $ES_READONLY = 2048
Local Const $EM_LIMITTEXT = 0xC5
Local Const $GUI_HIDE = 32
Local $w = 580, $h = 380
GUICreate($__g_sReportTitle_Debug, $w, $h, -1, -1, $WS_OVERLAPPEDWINDOW)
Local $idLabelHidden = GUICtrlCreateLabel($__g_sReportWindowText_Debug, 0, 0, 1, 1)
GUICtrlSetState($idLabelHidden, $GUI_HIDE)
Local $idEdit = GUICtrlCreateEdit("", 4, 4, $w - 8, $h - 8, BitOR($WS_HSCROLL, $WS_VSCROLL, $ES_READONLY))
$__g_hReportEdit_Debug = GUICtrlGetHandle($idEdit)
GUICtrlSetBkColor($idEdit, 0xFFFFFF)
GUICtrlSendMsg($idEdit, $EM_LIMITTEXT, 0, 0)
GUISetState()
$__g_bReportWindowWaitClose_Debug = True
Return 1
EndFunc
#Au3Stripper_Ignore_Funcs=__Debug_ReportWindowWrite
Func __Debug_ReportWindowWrite($sData)
If $__g_bReportWindowClosed_Debug Then __Debug_ReportWindowCreate()
Local Const $WM_GETTEXTLENGTH = 0x000E
Local Const $EM_SETSEL = 0xB1
Local Const $EM_REPLACESEL = 0xC2
Local $nLen = _SendMessage($__g_hReportEdit_Debug, $WM_GETTEXTLENGTH, 0, 0, 0, "int", "int")
_SendMessage($__g_hReportEdit_Debug, $EM_SETSEL, $nLen, $nLen, 0, "int", "int")
_SendMessage($__g_hReportEdit_Debug, $EM_REPLACESEL, True, $sData, 0, "int", "wstr")
EndFunc
Func __Debug_ReportNotepadCreate()
Local $bExists = WinExists($__g_sReportTitle_Debug)
If $bExists Then
If $__g_hReportEdit_Debug = 0 Then
$__g_hReportEdit_Debug = WinGetHandle($__g_sReportTitle_Debug)
Return 0
EndIf
EndIf
Local $pNotepad = Run("Notepad.exe")
$__g_hReportEdit_Debug = WinWait("[CLASS:Notepad]")
If $pNotepad <> WinGetProcess($__g_hReportEdit_Debug) Then
Return SetError(3, 0, 0)
EndIf
WinActivate($__g_hReportEdit_Debug)
WinSetTitle($__g_hReportEdit_Debug, "", String($__g_sReportTitle_Debug))
Return 1
EndFunc
#Au3Stripper_Ignore_Funcs=__Debug_ReportNotepadWrite
Func __Debug_ReportNotepadWrite($sData)
If $__g_hReportEdit_Debug = 0 Then __Debug_ReportNotepadCreate()
ControlCommand($__g_hReportEdit_Debug, "", "Edit1", "EditPaste", String($sData))
EndFunc
Func __Debug_ReportWrite($sData, $bLastError = False, $iLastError = 0)
Local $sError = ""
If $__g_bReportTimeStamp_Debug And($sData <> "") Then $sData = @YEAR & "/" & @MON & "/" & @MDAY & " " & @HOUR & ":" & @MIN & ":" & @SEC & " " & $sData
If $bLastError Then
$sError = " LastError = " & $iLastError & " : (" & _WinAPI_GetLastErrorMessage() & ")" & @CRLF
EndIf
$sData &= $sError
$sData = StringReplace($sData, "'", "''")
Local Static $sERROR_CODE = ">Error code:"
If StringInStr($sData, $sERROR_CODE) Then
$sData = StringReplace($sData, $sERROR_CODE, @TAB & $sERROR_CODE)
If(StringInStr($sData, $sERROR_CODE & " 0") = 0) Then
$sData = StringReplace($sData, $sERROR_CODE, $sERROR_CODE & @TAB & @TAB & @TAB & @TAB)
EndIf
EndIf
Execute($__g_sReportCallBack_Debug & "'" & $sData & "')")
Return
EndFunc
Func _ArrayFromString($sArrayStr, $sDelim_Col = "|", $sDelim_Row = @CRLF, $bForce2D = False, $iStripWS = $STR_STRIPLEADING + $STR_STRIPTRAILING)
If $sDelim_Col = Default Then $sDelim_Col = "|"
If $sDelim_Row = Default Then $sDelim_Row = @CRLF
If $bForce2D = Default Then $bForce2D = False
If $iStripWS = Default Then $iStripWS = $STR_STRIPLEADING + $STR_STRIPTRAILING
Local $aRow, $aCol = StringSplit($sArrayStr, $sDelim_Row, $STR_ENTIRESPLIT + $STR_NOCOUNT)
$aRow = StringSplit($aCol[0], $sDelim_Col, $STR_ENTIRESPLIT + $STR_NOCOUNT)
If UBound($aCol) = 1 And Not $bForce2D Then
For $m = 0 To UBound($aRow) - 1
$aRow[$m] =($iStripWS ? StringStripWS($aRow[$m], $iStripWS) : $aRow[$m])
Next
Return $aRow
EndIf
Local $aRet[UBound($aCol)][UBound($aRow)]
For $n = 0 To UBound($aCol) - 1
$aRow = StringSplit($aCol[$n], $sDelim_Col, $STR_ENTIRESPLIT + $STR_NOCOUNT)
If UBound($aRow) > UBound($aRet, 2) Then Return SetError(1)
For $m = 0 To UBound($aRow) - 1
$aRet[$n][$m] =($iStripWS ? StringStripWS($aRow[$m], $iStripWS) : $aRow[$m])
Next
Next
Return $aRet
EndFunc
Func _ArrayToString(Const ByRef $aArray, $sDelim_Col = "|", $iStart_Row = Default, $iEnd_Row = Default, $sDelim_Row = @CRLF, $iStart_Col = Default, $iEnd_Col = Default)
If $sDelim_Col = Default Then $sDelim_Col = "|"
If $sDelim_Row = Default Then $sDelim_Row = @CRLF
If $iStart_Row = Default Then $iStart_Row = -1
If $iEnd_Row = Default Then $iEnd_Row = -1
If $iStart_Col = Default Then $iStart_Col = -1
If $iEnd_Col = Default Then $iEnd_Col = -1
If Not IsArray($aArray) Then Return SetError(1, 0, -1)
Local $iDim_1 = UBound($aArray, $UBOUND_ROWS) - 1
If $iDim_1 = -1 Then Return ""
If $iStart_Row = -1 Then $iStart_Row = 0
If $iEnd_Row = -1 Then $iEnd_Row = $iDim_1
If $iStart_Row < -1 Or $iEnd_Row < -1 Then Return SetError(3, 0, -1)
If $iStart_Row > $iDim_1 Or $iEnd_Row > $iDim_1 Then Return SetError(3, 0, "")
If $iStart_Row > $iEnd_Row Then Return SetError(4, 0, -1)
Local $sRet = ""
Switch UBound($aArray, $UBOUND_DIMENSIONS)
Case 1
For $i = $iStart_Row To $iEnd_Row
$sRet &= $aArray[$i] & $sDelim_Col
Next
Return StringTrimRight($sRet, StringLen($sDelim_Col))
Case 2
Local $iDim_2 = UBound($aArray, $UBOUND_COLUMNS) - 1
If $iDim_2 = -1 Then Return ""
If $iStart_Col = -1 Then $iStart_Col = 0
If $iEnd_Col = -1 Then $iEnd_Col = $iDim_2
If $iStart_Col < -1 Or $iEnd_Col < -1 Then Return SetError(5, 0, -1)
If $iStart_Col > $iDim_2 Or $iEnd_Col > $iDim_2 Then Return SetError(5, 0, -1)
If $iStart_Col > $iEnd_Col Then Return SetError(6, 0, -1)
Local $iDelimColLen = StringLen($sDelim_Col)
For $i = $iStart_Row To $iEnd_Row
For $j = $iStart_Col To $iEnd_Col
$sRet &= $aArray[$i][$j] & $sDelim_Col
Next
$sRet = StringTrimRight($sRet, $iDelimColLen) & $sDelim_Row
Next
Return StringTrimRight($sRet, StringLen($sDelim_Row))
Case Else
Return SetError(2, 0, -1)
EndSwitch
Return 1
EndFunc
Func _FileWriteLog($sLogPath, $sLogMsg, $iFlag = -1)
Local $iOpenMode = $FO_APPEND
Local $sMsg = @YEAR & "-" & @MON & "-" & @MDAY & " " & @HOUR & ":" & @MIN & ":" & @SEC & " : " & $sLogMsg
If $iFlag = Default Then $iFlag = -1
If $iFlag <> -1 Then
$iOpenMode = $FO_OVERWRITE
$sMsg &= @CRLF & FileRead($sLogPath)
EndIf
Local $hFileOpen = $sLogPath
If IsString($sLogPath) Then $hFileOpen = FileOpen($sLogPath, $iOpenMode)
If $hFileOpen = -1 Then Return SetError(1, 0, 0)
Local $iReturn = FileWriteLine($hFileOpen, $sMsg)
If IsString($sLogPath) Then $iReturn = FileClose($hFileOpen)
If $iFlag <> -1 And Not IsString($sLogPath) Then SetExtended(1)
If $iReturn = 0 Then Return SetError(2, 0, 0)
Return $iReturn
EndFunc
Func _WinAPI_PathFindExtension($sFilePath)
Local $aCall = DllCall('shlwapi.dll', 'wstr', 'PathFindExtensionW', 'wstr', $sFilePath)
If @error Then Return SetError(@error, @extended, '')
Return $aCall[0]
EndFunc
Func _WinAPI_PathIsExe($sFilePath)
Local $aCall = DllCall('shell32.dll', 'bool', 'PathIsExe', 'wstr', $sFilePath)
If @error Then Return SetError(@error, @extended, False)
Return $aCall[0]
EndFunc
Func _WinAPI_PathRemoveExtension($sFilePath)
Local $aCall = DllCall('shlwapi.dll', 'none', 'PathRemoveExtensionW', 'wstr', $sFilePath)
If @error Then Return SetError(@error, @extended, '')
Return $aCall[1]
EndFunc
Func _WinAPI_PathRenameExtension($sFilePath, $sExt)
Local $tPath = DllStructCreate('wchar[260]')
DllStructSetData($tPath, 1, $sFilePath)
Local $aCall = DllCall('shlwapi.dll', 'bool', 'PathRenameExtensionW', 'struct*', $tPath, 'wstr', $sExt)
If @error Or Not $aCall[0] Then Return SetError(@error, @extended, '')
Return DllStructGetData($tPath, 1)
EndFunc
Global $bIsPE = True
RegRead("HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\WinPE", "Version")
If @error Then $bIsPE = False
Global $g_hLogFile, $g_sLogPath, $g_nLogLevel
Global Const $g_ProgramName = _WinAPI_PathRemoveExtension(@ScriptName)
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
Func LogD($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
If $g_nLogLevel <= 0 Then
_writeLog("[D]", $s)
EndIf
SetError($_iCallerError, $_iCallerExtended)
EndFunc
Func LogI($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
If $g_nLogLevel <= 1 Then
_writeLog("[I]", $s)
EndIf
SetError($_iCallerError, $_iCallerExtended)
EndFunc
Func LogW($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
If $g_nLogLevel <= 2 Then
_writeLog("[W]", $s)
EndIf
SetError($_iCallerError, $_iCallerExtended)
EndFunc
Func LogE($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
If $g_nLogLevel <= 3 Then
_writeLog("[E]", $s)
EndIf
SetError($_iCallerError, $_iCallerExtended)
EndFunc
Func Err($e)
LogW($e)
MsgBox($MB_ICONWARNING, "[AutoPE] " & $g_ProgramName, $e)
EndFunc
Func Failed($e)
LogE($e)
_closeLog()
MsgBox($MB_ICONERROR, "[AutoPE] " & $g_ProgramName, $e)
If MsgBox(BitOR($MB_ICONERROR, $MB_YESNO), "[AutoPE] " & $g_ProgramName, "是否将日志文件保存到U盘？") = $IDYES Then
SplashTextOn("[AutoPE] " & $g_ProgramName, "等待U盘插入……")
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
SplashTextOn("[AutoPE] " & $g_ProgramName, "三十秒内未检测到U盘，停止等待")
EndIf
EndIf
Exit 1
EndFunc
Func FixBadIniFile($sIniFilePath)
If FileGetEncoding($sIniFilePath) <> $FO_UNICODE Then
Local $sContent = FileRead($sIniFilePath)
Local $hIni = FileOpen($sIniFilePath, BitOR($FO_OVERWRITE, $FO_UNICODE))
If $hIni = -1 Then Return False
FileWrite($hIni, $sContent)
FileClose($hIni)
EndIf
Return True
EndFunc
Global Const $HGDI_ERROR = Ptr(-1)
Global Const $INVALID_HANDLE_VALUE = Ptr(-1)
Global Const $KF_EXTENDED = 0x0100
Global Const $KF_ALTDOWN = 0x2000
Global Const $KF_UP = 0x8000
Global Const $LLKHF_EXTENDED = BitShift($KF_EXTENDED, 8)
Global Const $LLKHF_ALTDOWN = BitShift($KF_ALTDOWN, 8)
Global Const $LLKHF_UP = BitShift($KF_UP, 8)
Func _WinAPI_GetFullPathName($sFilePath)
Local $aCall = DllCall('kernel32.dll', 'dword', 'GetFullPathNameW', 'wstr', $sFilePath, 'dword', 4096, 'wstr', '', 'ptr', 0)
If @error Or Not $aCall[0] Then Return SetError(@error, @extended, '')
Return $aCall[3]
EndFunc
Func _WinAPI_GetDC($hWnd)
Local $aCall = DllCall("user32.dll", "handle", "GetDC", "hwnd", $hWnd)
If @error Then Return SetError(@error, @extended, 0)
Return $aCall[0]
EndFunc
Func _WinAPI_ReleaseDC($hWnd, $hDC)
Local $aCall = DllCall("user32.dll", "int", "ReleaseDC", "hwnd", $hWnd, "handle", $hDC)
If @error Then Return SetError(@error, @extended, False)
Return $aCall[0]
EndFunc
Func _WinAPI_SetBkColor($hDC, $iColor)
Local $aCall = DllCall("gdi32.dll", "INT", "SetBkColor", "handle", $hDC, "INT", $iColor)
If @error Then Return SetError(@error, @extended, -1)
Return $aCall[0]
EndFunc
Func _WinAPI_SetBkMode($hDC, $iBkMode)
Local $aCall = DllCall("gdi32.dll", "int", "SetBkMode", "handle", $hDC, "int", $iBkMode)
If @error Then Return SetError(@error, @extended, 0)
Return $aCall[0]
EndFunc
Func _WinAPI_InvalidateRect($hWnd, $tRECT = 0, $bErase = True)
If Not IsHWnd($hWnd) Then $hWnd = GUICtrlGetHandle($hWnd)
Local $aCall = DllCall("user32.dll", "bool", "InvalidateRect", "hwnd", $hWnd, "struct*", $tRECT, "bool", $bErase)
If @error Then Return SetError(@error, @extended, False)
Return $aCall[0]
EndFunc
Func _WinAPI_UpdateWindow($hWnd)
Local $aCall = DllCall("user32.dll", "bool", "UpdateWindow", "hwnd", $hWnd)
If @error Then Return SetError(@error, @extended, False)
Return $aCall[0]
EndFunc
Func RegSetNetCfg($sHive, $sTargetDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
Local $i = 1
While 1
$sSubKey = RegEnumKey($sHive & "\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}", $i)
$i = $i + 1
If @error Then ExitLoop
If Not StringIsDigit($sSubKey) Then ContinueLoop
LogD("$sSubKey " & $sSubKey)
$sPath = $sHive & "\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}\" & $sSubKey
$sNetCfgId = RegRead($sPath, "NetCfgInstanceId")
If @error Then ContinueLoop
$sDeviceId = RegRead($sPath, "DeviceInstanceID")
If @error or $sDeviceId <> $sTargetDeviceId Then ContinueLoop
LogD("$sNetCfgId " & $sNetCfgId)
LogD("$sDeviceId " & $sDeviceId)
$sPath = $sHive & "\Services\Tcpip\Parameters\Interfaces\" & $sNetCfgId
RegWrite($sPath, "EnableDHCP", "REG_DWORD", 0)
RegWrite($sPath, "IPAddress", "REG_MULTI_SZ", _ArrayToString($aIPAddr, @LF))
RegWrite($sPath, "SubnetMask", "REG_MULTI_SZ", _ArrayToString($aSubnet, @LF))
RegWrite($sPath, "DefaultGateway", "REG_MULTI_SZ", _ArrayToString($aDefGateway, @LF))
RegWrite($sPath, "NameServer", "REG_SZ", _ArrayToString($aDNS, ','))
Return True
WEnd
Return False
EndFunc
Func LoadNetCfgFromIni($sPath)
LogD("$sPath="&$sPath)
If Not FileExists($sPath) Then Return
Local $aArray = IniReadSectionNames($sPath)
For $i = 1 To $aArray[0]
$sSection = $aArray[$i]
If StringLeft($sSection, 6) <> "NetCfg" Then ContinueLoop
Local $sDeviceId = IniRead($sPath, $sSection, "DeviceId", "")
Local $aIPAddr = _ArrayFromString(IniRead($sPath, $sSection, "IPAddr", ""))
Local $aSubnet = _ArrayFromString(IniRead($sPath, $sSection, "Subnet", ""))
Local $aDefGateway = _ArrayFromString(IniRead($sPath, $sSection, "DefGateway", ""))
Local $aDNS = _ArrayFromString(IniRead($sPath, $sSection, "DNS", ""))
Local $bSet = RegSetNetCfg("HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet", $sDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
If $bSet Then RestartAdapter($sDeviceId)
Next
EndFunc
Func RestartAdapter($sDeviceId)
Local $hDLL = DllOpen("cfgmgr32.dll")
If $hDLL = -1 Then
LogE("unable to load Cfgmgr32.dll: " & @error)
Return False
EndIf
Local $rLocate = DllCall($hDLL, "DWORD", "CM_Locate_DevNodeW", "DWORD*", 0, "WSTR", $sDeviceId, "DWORD", 0)
If $rLocate[0] <> 0 Then
LogE("unable to locate DevNode: " & $rLocate[0])
Return False
EndIf
Local $ret = DllCall($hDLL, "DWORD", "CM_Disable_DevNode", "DWORD", $rLocate[1], "ULONG", 0)
If $ret[0] <> 0 Then
LogE("unable to disable DevNode: " & $ret[0])
Return False
EndIf
$ret = DllCall($hDLL, "DWORD", "CM_Enable_DevNode", "DWORD", $rLocate[1], "ULONG", 0)
If $ret[0] <> 0 Then
LogE("unable to enable DevNode: " & $ret[0])
Return False
EndIf
Return True
EndFunc
Func UNCMap($sUNCPath, $sDrive, $sUser, $sPwd)
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
Func RunPackEntry($sDir)
LogD("Searching " & $sDir)
Local $hSearch = FileFindFirstFile($sDir & "\AutoPEPack.*")
Local $sFileName = ""
While True
$sFileName = FileFindNextFile($hSearch)
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
$sTmp = _WinAPI_GetFullPathName($sTmp)
$sPackPath = _WinAPI_GetFullPathName($sPackPath)
LogI("Applying pack: " & $sPackPath)
DirRemove($sTmp, 1)
DirCreate($sTmp)
If RunWait($s7z & ' x "' & $sPackPath & '"', $sTmp) <> 0 Then
DirRemove($sTmp, 1)
Return False
EndIf
If RunPackEntry($sTmp) <> 0 Then
LogE($sPackPath & " failed")
EndIf
Return DirRemove($sTmp, 1)
EndFunc
Func ApplyISOPack($sPackPath, $sWinCDEmu)
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
$ret = RunWait(StringFormat('%s /unmount "%s"', $sWinCDEmu, $sPackPath))
If $ret <> 0 Then
Err("Could not detach " & $sPackPath & ": " & $ret)
Return False
EndIf
Return True
EndFunc
Global Const $TRANSPARENT = 1
Func SetWinSplashText($sText)
Local $hWnd = WinGetHandle("[CLASS:FirstUXWndClass]")
ControlSetText($hWnd, "", "Static1", $sText)
Local $hDC = _WinAPI_GetDC($hWnd)
_WinAPI_SetBkColor($hDC, 0xFFFFFF)
_WinAPI_SetBkMode($hDC, $TRANSPARENT)
_WinAPI_InvalidateRect($hWnd, 0, True)
_WinAPI_UpdateWindow($hWnd)
_WinAPI_ReleaseDC($hWnd, $hDC)
EndFunc
Global $sIniConfig = @ScriptDir & "\Main.ini"
FileChangeDir(@ScriptDir)
InitLog("Debug")
FixBadIniFile($sIniConfig)
FileInstall("7za.exe", "tmp\7za.exe")
Func _apply($sList)
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
Switch StringLower($CmdLine[1])
Case "prewindeploy"
SetWinSplashText("[AutoPE] PreWindeploy")
_apply("tmp\preWindeployPacks")
Case "postwindeploy"
SetWinSplashText("[AutoPE] PostWindeploy")
_apply("tmp\postWinDeployPacks")
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
Case "audit"
_apply("tmp\auditPacks")
Case "firstlogon"
_apply("tmp\firstLogonPacks")
Case Else
MsgBox(0, "[AutoPE] Stage3", $CmdLine[1] & " not implemented")
EndSwitch
