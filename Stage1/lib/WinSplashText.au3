#include-once
#include <WindowsConstants.au3>
;#include <WinAPIGdi.au3>
#include <WinAPIGdiDC.au3>
#include <WinAPISysWin.au3>

Func SetWinSplashText($sText)
	Local $hWnd = WinGetHandle("[CLASS:FirstUXWndClass]")
	ControlSetText($hWnd, "", "Static1", $sText)
	;Won't work
	;_WinAPI_RedrawWindow($hWnd, 0, 0, $RDW_UPDATENOW + $RDW_INVALIDATE + $RDW_ALLCHILDREN)
	Local $hDC = _WinAPI_GetDC($hWnd)
	_WinAPI_SetBkColor($hDC, 0xFFFFFF)
	_WinAPI_SetBkMode($hDC, $TRANSPARENT)
	_WinAPI_InvalidateRect($hWnd, 0, True)
	_WinAPI_UpdateWindow($hWnd)
	_WinAPI_ReleaseDC($hWnd, $hDC)
EndFunc
