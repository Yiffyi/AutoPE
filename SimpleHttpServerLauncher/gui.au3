Global $iScale = RegRead("HKCU\Control Panel\Desktop\WindowMetrics", "AppliedDPI") / 96
#cs
	(GUI(?:Ctrl)?Create.*\d+)(?=[,)])   =  $1 * \$iScale
	(?<=\d)(?=[,)])  =  * \$iScale
#ce
#include <ButtonConstants.au3>
#include <EditConstants.au3>
#include <GUIConstantsEx.au3>
#include <StaticConstants.au3>
#include <WindowsConstants.au3>

#Region ### START Koda GUI section ### Form=C:\Users\HigherSY\source\repos\Yiffyi\AutoPE\SimpleHttpServerLauncher\Form1.kxf
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
#EndRegion ### END Koda GUI section ###


