Global Const $FC_OVERWRITE = 1
Global Const $MB_SYSTEMMODAL = 4096
Global $__g_aTimers_aTimerIDs[1][3]
Func _Timer_SetTimer($hWnd, $iElapse = 250, $sTimerFunc = "", $iTimerID = -1)
#Au3Stripper_Ignore_Funcs=$sTimerFunc
Local $aCall[1] = [0], $pTimerFunc = 0, $hCallBack = 0, $iIndex = $__g_aTimers_aTimerIDs[0][0] + 1
If $iTimerID = -1 Then
ReDim $__g_aTimers_aTimerIDs[$iIndex + 1][3]
$__g_aTimers_aTimerIDs[0][0] = $iIndex
$iTimerID = $iIndex + 1000
For $x = 1 To $iIndex
If $__g_aTimers_aTimerIDs[$x][0] = $iTimerID Then
$iTimerID = $iTimerID + 1
$x = 0
EndIf
Next
If $sTimerFunc <> "" Then
$hCallBack = DllCallbackRegister($sTimerFunc, "none", "hwnd;uint;uint_ptr;dword")
If $hCallBack = 0 Then Return SetError(-1, -1, 0)
$pTimerFunc = DllCallbackGetPtr($hCallBack)
If $pTimerFunc = 0 Then Return SetError(-1, -2, 0)
EndIf
$aCall = DllCall("user32.dll", "uint_ptr", "SetTimer", "hwnd", $hWnd, "uint_ptr", $iTimerID, "uint", $iElapse, "ptr", $pTimerFunc)
If @error Or Not $aCall[0] Then Return SetError(@error + 10, @extended, 0)
$__g_aTimers_aTimerIDs[$iIndex][0] = $aCall[0]
$__g_aTimers_aTimerIDs[$iIndex][1] = $iTimerID
$__g_aTimers_aTimerIDs[$iIndex][2] = $hCallBack
Else
For $x = 1 To $iIndex - 1
If $__g_aTimers_aTimerIDs[$x][0] = $iTimerID Then
If IsHWnd($hWnd) Then $iTimerID = $__g_aTimers_aTimerIDs[$x][1]
$hCallBack = $__g_aTimers_aTimerIDs[$x][2]
If $hCallBack <> 0 Then
$pTimerFunc = DllCallbackGetPtr($hCallBack)
If $pTimerFunc = 0 Then Return SetError(-1, -12, 0)
EndIf
$aCall = DllCall("user32.dll", "uint_ptr", "SetTimer", "hwnd", $hWnd, "uint_ptr", $iTimerID, "uint", $iElapse, "ptr", $pTimerFunc)
If @error Or Not $aCall[0] Then Return SetError(@error + 20, @extended, 0)
ExitLoop
EndIf
Next
EndIf
Return $aCall[0]
EndFunc
Global $iScale = RegRead("HKCU\Control Panel\Desktop\WindowMetrics", "AppliedDPI") / 96
Global Const $GUI_EVENT_CLOSE = -3
$Form1 = GUICreate("春晖系统部署 - 服务端", 286 * $iScale, 326 * $iScale, 304 * $iScale, 171 * $iScale)
GUISetFont(8, 400, 0, "Segoe UI")
$Label1 = GUICtrlCreateLabel("服务路径：", 14 * $iScale, 98 * $iScale, 69 * $iScale, 17 * $iScale)
$iDir = GUICtrlCreateInput("@ScriptDir", 84 * $iScale, 98 * $iScale, 121 * $iScale, 21 * $iScale)
$Label2 = GUICtrlCreateLabel("春晖系统部署 - 服务端 - v2023 b1", 28 * $iScale, 42 * $iScale, 239 * $iScale, 22 * $iScale)
GUICtrlSetFont(-1, 12, 400, 0, "方正姚体")
$bBrowse = GUICtrlCreateButton("浏览", 217 * $iScale, 91 * $iScale, 61 * $iScale, 35 * $iScale)
$Label3 = GUICtrlCreateLabel("端口号：", 26 * $iScale, 147 * $iScale, 56 * $iScale, 17 * $iScale)
$iPort = GUICtrlCreateInput("80", 84 * $iScale, 147 * $iScale, 121 * $iScale, 21 * $iScale)
$Label4 = GUICtrlCreateLabel("状态：", 38 * $iScale, 189 * $iScale, 43 * $iScale, 17 * $iScale)
$lStatus = GUICtrlCreateLabel("已停止", 84 * $iScale, 189 * $iScale, 43 * $iScale, 17 * $iScale)
GUICtrlSetColor(-1, 0xFF0000)
$bStart = GUICtrlCreateButton("启动", 14 * $iScale, 287 * $iScale, 75 * $iScale, 25 * $iScale)
$bStop = GUICtrlCreateButton("停止", 196 * $iScale, 287 * $iScale, 75 * $iScale, 25 * $iScale)
$lRead = GUICtrlCreateLabel("读：0 字节", 84 * $iScale, 210 * $iScale, 65 * $iScale, 17 * $iScale)
$lWrite = GUICtrlCreateLabel("写：0 字节", 84 * $iScale, 231 * $iScale, 65 * $iScale, 17 * $iScale)
$lMemory = GUICtrlCreateLabel("内存占用：0 字节", 84 * $iScale, 252 * $iScale, 104 * $iScale, 17 * $iScale)
GUISetState(@SW_SHOW)
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
MsgBox($MB_SYSTEMMODAL, "春晖系统部署", "没有选择目录")
Else
FileChangeDir($sFolder)
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
