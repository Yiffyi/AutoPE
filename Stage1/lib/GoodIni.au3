#include-once

#include <FileConstants.au3>

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

Func CreateGoodIniFile($sIniFilePath)
	Local $hIni = FileOpen($sIniFilePath, BitOR($FO_OVERWRITE, $FO_CREATEPATH, $FO_UNICODE))
	FileWrite($hIni, @CRLF)
	FileClose($hIni)
EndFunc
