#pragma compile(x64, true)
#pragma compile(Console, true)
#NoTrayIcon
#RequireAdmin
Global Const $SD_REBOOT = 2
Global Const $MB_ICONERROR = 16
Global Const $MB_ICONINFORMATION = 64
Global Const $FC_OVERWRITE = 1
Global Const $FC_CREATEPATH = 8
Global Const $FO_APPEND = 1
Global Const $FO_BINARY = 16
Global $sIniConfig = "stage1\config.ini"
FileInstall("aria2c.exe", "aria2c.exe")
StartNetwork()
SetupVolumes()
ApplyImage()
ProcessInjectFiles()
FixBoot()
WriteBackConfig()
MsgBox($MB_ICONINFORMATION, "[AutoPE] Worker", "=== 工作结束 ===" & @CRLF & "10s 后重启...", 10)
Shutdown($SD_REBOOT)
Func StartNetwork()
If FileExists("stage1\network.cmd") Then
TryRunWaitCatch("stage1\network.cmd", 0, "ERR_NETWORK")
Else
RunWait("wpeutil WaitForNetwork")
EndIf
RunWait("wpeutil DisableFirewall")
EndFunc
Func SetupVolumes()
Local $sUseDiskpart = IniRead($sIniConfig, "Volume", "UseDiskpart", "False")
If $sUseDiskpart == "True" Then
TryRunWaitCatch("diskpart /s stage1\diskpart.txt", 0, "ERR_DISKPART")
Else
TryRunWaitCatch("stage1\volume.cmd", 0, "ERR_SETUP_VOLUME")
EndIf
EndFunc
Func ApplyImage()
TryRunWaitCatch("aria2c -i stage1\systemAria2.txt -d J:\AutoPE\stage2 -x 4 --file-allocation=falloc", 0, "ERR_ARIA2_SYSIMG")
Local $sImgIndex = IniRead($sIniConfig, "SystemImg", "ImgIndex", "1")
TryRunWaitCatch("dism /apply-image /imagefile:J:\AutoPE\stage2\System.esd" & " /index:" & $sImgIndex & " /applydir:I:\", 0, "ERR_DISM_APPLY")
EndFunc
Func FixBoot()
TryRunWaitCatch("bcdboot I:\windows /l zh-cn", 0, "ERR_BCDBOOT")
EndFunc
Func ProcessInjectFiles()
TryRunWaitCatch("aria2c -i stage1\injectAria2.txt -d I:\ --file-allocation=falloc", 0, "ERR_ARIA2_INJ")
Local $sections = IniReadSectionNames($sIniConfig)
For $i = 1 To $sections[0]
If StringLeft($sections[$i], 4) = "Text" Then
Local $cfg = IniReadSection($sIniConfig, $sections[$i])
Local $sDstPath = "I:\" & IniRead($sIniConfig, $sections[$i], "DstPath", "")
Local $sSrcPath = "stage1\" & IniRead($sIniConfig, $sections[$i], "SrcPath", "")
Local $sAppend = IniRead($sIniConfig, $sections[$i], "Append", "")
D('$sDstPath = ' & $sDstPath, 'Inject')
D('$sAppend = ' & $sAppend, 'Inject')
If $sAppend = "True" Then
Local $hSrcFile = FileOpen($sSrcPath, $FO_BINARY)
If $hSrcFile = -1 Then Failed("ERR_SRCFILE_OPEN")
Local $hDstFile = FileOpen($sDstPath, $FO_BINARY + $FO_APPEND)
If $hDstFile = -1 Then Failed("ERR_DSTFILE_OPEN")
FileWrite($hDstFile, FileRead($hSrcFile))
FileClose($hDstFile)
FileClose($hSrcFile)
Else
FileCopy($sSrcPath, $sDstPath, $FC_OVERWRITE + $FC_CREATEPATH)
EndIf
EndIf
Next
EndFunc
Func WriteBackConfig()
FileCopy("stage1\network.cmd", "J:\AutoPE\stage2\network.cmd")
EndFunc
Func D($msg, $pos, $arg = "")
ConsoleWrite('@@ ' & $pos & '(' & $arg & '): ' & $msg & @CRLF & '    > @error: ' & @error & @CRLF)
EndFunc
Func TryRunWaitCatch($cmd, $expected_ret, $sErr)
D('Starting', 'RunWait', $cmd)
$ret = RunWait($cmd)
D('Finished, $ret = ' & $ret, 'RunWait', $cmd)
If $ret <> $expected_ret Then Failed($sErr & "_" & $ret)
EndFunc
Func Failed($where)
MsgBox($MB_ICONERROR, "[DP] Worker", "部署系统遇到了错误" & @CRLF & "CODE: " & $where)
Exit 1
EndFunc
