#RequireAdmin
#NoTrayIcon
Global Const $FO_APPEND = 1
Global Const $FO_OVERWRITE = 2
Global Const $FO_CREATEPATH = 8
Global Const $FO_UNICODE = 32
Global Const $OPT_MATCHSTART = 1
Global Const $WINDOWS_NOONTOP = 0
Global Const $MB_ICONWARNING = 48
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
LogW("====== " & _WinAPI_PathRemoveExtension(@ScriptName) & "，启动！ ======")
EndFunc
Func _writeLog($sLevel, $sMsg)
_FileWriteLog($g_hLogFile, $sLevel & ' ' & $sMsg)
EndFunc
Func _closeLog()
FileClose($g_hLogFile)
EndFunc
Func LogW($s, Const $_iCallerError = @error, Const $_iCallerExtended = @extended)
If $g_nLogLevel <= 2 Then
_writeLog("[W]", $s)
EndIf
SetError($_iCallerError, $_iCallerExtended)
EndFunc
Func Err($e)
LogW($e)
MsgBox($MB_ICONWARNING, "[AutoPE]", $e)
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
Global Const $TRANSPARENT = 1
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
