#AutoIt3Wrapper_Run_Au3Stripper=Y
#AutoIt3Wrapper_UseX64=Y
#AutoIt3Wrapper_Change2CUI=Y
#AutoIt3Wrapper_Run_Debug_Mode=Y
#NoTrayIcon
#RequireAdmin
#include <AutoItConstants.au3>
#include <MsgBoxConstants.au3>
#include <FileConstants.au3>
#include "Base64.au3"

#cs
请将系统盘盘符设为 I
将数据盘盘符设为 J
#ce

Global $sIniConfig = "stage1\config.ini"

FileInstall("aria2c.exe", "aria2c.exe")

StartNetwork()
SetupVolumes()
ApplyImage()
ProcessInjectFiles()
FixBoot()
WriteBackConfig()

MsgBox($MB_ICONINFORMATION, "[DP] Worker", "=== 工作结束 ===" & @CRLF & "10s 后重启...", 10)
Shutdown($SD_REBOOT)

Func StartNetwork()
	;RunWait("wpeutil InitializeNetwork /NoWait")
	If FileExists("stage1\network.cmd") Then
		TryRunWaitCatch("stage1\network.cmd", 0, "ERR_NETWORK")
	Else
		RunWait("wpeutil WaitForNetwork")
	EndIf
	RunWait("wpeutil DisableFirewall")
EndFunc   ;==>StartNetwork


Func SetupVolumes()
	Local $sUseDiskpart = IniRead($sIniConfig, "Volume", "UseDiskpart", "False")
	If $sUseDiskpart == "True" Then
		TryRunWaitCatch("diskpart /s stage1\diskpart.txt", 0, "ERR_DISKPART")
	Else
		TryRunWaitCatch("stage1\volume.cmd", 0, "ERR_SETUP_VOLUME")
	EndIf
EndFunc   ;==>SetupVolumes


Func ApplyImage()
	; 注意参数空格
	; 下载
	TryRunWaitCatch("aria2c -i stage1\systemAria2.txt -d J:\AutoPE\stage2 -x 4 --file-allocation=falloc", 0, "ERR_ARIA2_SYSIMG")
	Local $sImgIndex = IniRead($sIniConfig, "SystemImg", "ImgIndex", "1")
	TryRunWaitCatch("dism /apply-image /imagefile:J:\AutoPE\stage2\System.esd" & " /index:" & $sImgIndex & " /applydir:I:\", 0, "ERR_DISM_APPLY")
EndFunc   ;==>ApplyImage

Func FixBoot()
;~ If FileExists("K:\") Then
;~    $ret = RunWait("bcdboot I:\windows /s K: /l zh-cn")
;~    If $ret <> 0 Then Failed("ERR_BCDBOOT_RET")
;~ Else
	TryRunWaitCatch("bcdboot I:\windows /l zh-cn", 0, "ERR_BCDBOOT")
	; $ret = RunWait("bootsect /nt60 I: /mbr")
	; If $ret <> 0 Then Failed("ERR_BCDSECT_RET")
;~ EndIf
EndFunc   ;==>FixBoot


Func ProcessInjectFiles()
	TryRunWaitCatch("aria2c -i stage1\injectAria2.txt -d I:\ --file-allocation=falloc", 0, "ERR_ARIA2_INJ")
	Local $sections = IniReadSectionNames($sIniConfig)

	For $i = 1 To $sections[0]
		If StringLeft($sections[$i], 4) = "Text" Then
			Local $cfg = IniReadSection($sIniConfig, $sections[$i])
			Local $sDstPath = "I:\" & IniRead($sIniConfig, $sections[$i], "DstPath", "")
			D('$sDstPath = ' & $sDstPath, 'Inject')
			Local $dContent = base64(IniRead($sIniConfig, $sections[$i], "Base64Content", ""), False)
			Local $sAppend = IniRead($sIniConfig, $sections[$i], "Append", "")
			D('$sAppend = ' & $sAppend, 'Inject')

			Local $fileMode = $FO_OVERWRITE
			If $sAppend = "True" Then
				$fileMode = $FO_APPEND
			EndIf

			Local $hDstFile = FileOpen($sDstPath, $fileMode + $FO_CREATEPATH)
			If $hDstFile = -1 Then Failed("ERR_DSTFILE_OPEN")

			FileWrite($hDstFile, $dContent)
			FileClose($hDstFile)
		EndIf
	Next

	#cs
	; 1 = success
	Local $retCopy = FileCopy("packages/*.7z", "C:\WindSys_Setup\WTDR\WTDR.Pack\Pack")
	If $retCopy <> 1 Then Failed("ERR_PKG_COPY")
	#ce

	#cs
	Local $hSrcINI = FileOpen("append.ini", $FO_READ)
	If $hSrcINI = -1 Then Failed("ERR_SRCINI_OPEN")

	Local $hDstINI = FileOpen("I:\WindSys_Setup\WTDR\WTDR.Pack\Pack_Config.ini", $FO_APPEND)
	If $hDstINI = -1 Then Failed("ERR_DSTINI_OPEN")

	Local $sContent = FileRead($hSrcINI)
	FileWrite($hDstINI, $sContent)

	FileClose($hSrcINI)
	FileClose($hDstINI)
	#ce
EndFunc   ;==>ProcessInjectFiles

Func WriteBackConfig()
	;DirCreate("J:\DPFiles\worker") ;创建了却不加 $FC_OVERWRITE 会出错
	FileCopy("stage1\network.cmd", "J:\AutoPE\stage2\network.cmd")
	;DirCopy("X:\DPFiles\config", "J:\DPFiles\config", $FC_OVERWRITE)
EndFunc   ;==>WriteBackConfig

Func D($msg, $pos, $arg = "")
	ConsoleWrite('@@ ' & $pos & '(' & $arg & '): ' & $msg & @CRLF & '    > @error: ' & @error & @CRLF)
EndFunc

Func TryRunWaitCatch($cmd, $expected_ret, $sErr)
	D('Starting', 'RunWait', $cmd) ;### Debug Console
	$ret = RunWait($cmd)
	D('Finished, $ret = ' & $ret, 'RunWait', $cmd)
	If $ret <> $expected_ret Then Failed($sErr & "_" & $ret)
EndFunc   ;==>TryRunWaitCatch

Func Failed($where)
	MsgBox($MB_ICONERROR, "[DP] Worker", "部署系统遇到了错误" & @CRLF & "CODE: " & $where)
	Exit 1
EndFunc   ;==>Failed
