cd /d %~dp0
Dism /Cleanup-Image /Image:"%SYSTEMDRIVE%\WinREMount" /StartComponentCleanup /ResetBase
Dism /Unmount-Image /MountDir:"%SYSTEMDRIVE%\WinREMount" /commit
Dism /Export-Image /SourceImageFile:"%~dp0\winre.wim" /SourceIndex:1 /DestinationImageFile:"%~dp0\boot.wim"
pause