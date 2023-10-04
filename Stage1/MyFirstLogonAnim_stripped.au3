#NoTrayIcon
Global Const $DLG_MOVEABLE = 16
Local $sMsg = "海内存知己，天涯若比邻。" & @CRLF & @CRLF & "Pretend to be FirstLogonAnim.exe: "
SplashTextOn("[AutoPE] MyFirstLogonAnim", $sMsg & "0s", 500, 200, 0, 0, $DLG_MOVEABLE)
Local $nCur = 0
While True
ControlSetText("[AutoPE] MyFirstLogonAnim", "", "Static1", $sMsg & $nCur & "s")
$nCur = $nCur + 1
Sleep(1000)
WEnd
