cd /d %~dp0
md %SYSTEMDRIVE%\WinREMount
xcopy /H /Y %SYSTEMDRIVE%\Recovery\WindowsRE\Winre.wim winre.wim*
xcopy /H /Y %SYSTEMDRIVE%\Recovery\WindowsRE\boot.sdi boot.sdi*
attrib -S -H winre.wim
attrib -S -H boot.sdi

Dism /Mount-Image /ImageFile:"%~dp0\Winre.wim" /index:1 /MountDir:"%SYSTEMDRIVE%\WinREMount"

takeown /f %SYSTEMDRIVE%\WinREMount\sources /R
icacls %SYSTEMDRIVE%\WinREMount\sources /T /C /RESET

del /s /f /q %SYSTEMDRIVE%\WinREMount\sources
del /s /f /q %SYSTEMDRIVE%\WinREMount\Windows\System32\winpeshl.ini

xcopy /Y startnet.cmd %SYSTEMDRIVE%\WinREMount\Windows\System32\
xcopy /Y lookup.cmd %SYSTEMDRIVE%\WinREMount\Windows\System32\
pause
