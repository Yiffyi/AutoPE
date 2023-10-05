#include-once
#include <Array.au3>
#include <WinAPIReg.au3>

#include "Common.au3"
#include "HiveLoad.au3"

Func SaveNetCfg($sKey, $sDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
	IniWrite(@ScriptDir & "\NetworkCfg.ini", "NetCfg" & $sKey, "DeviceId", $sDeviceId)

	Local $sIPAddr = _ArrayToString($aIPAddr)
	IniWrite(@ScriptDir & "\NetworkCfg.ini", "NetCfg" & $sKey, "IPAddr", $sIPAddr)

	Local $sSubnet = _ArrayToString($aSubnet)
	IniWrite(@ScriptDir & "\NetworkCfg.ini", "NetCfg" & $sKey, "Subnet", $sSubnet)

	Local $sDefGateway = _ArrayToString($aDefGateway)
	IniWrite(@ScriptDir & "\NetworkCfg.ini", "NetCfg" & $sKey, "DefGateway", $sDefGateway)

	Local $sDNS = _ArrayToString($aDNS)
	IniWrite(@ScriptDir & "\NetworkCfg.ini", "NetCfg" & $sKey, "DNS", $sDNS)
EndFunc


Func WMISetNetCfg($sDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
	Local $objCol = ObjGet("WinMgmts:").InstancesOf("Win32_NetworkAdapterSetting")
	For $i in $objCol
		$e = ObjGet("WinMgmts:" & $i.Element)
		$s = ObjGet("WinMgmts:" & $i.Setting)
		LogD("WMI: $e.PNPDeviceID=" & $e.PNPDeviceID)
		If $e.PNPDeviceID = $sDeviceId Then
			$ret = $s.EnableStatic($aIPAddr, $aSubnet)
			LogD("$s.EnableStatic="&$ret)
			$ret = $s.SetGateways($aDefGateway)
			LogD("$s.SetGateways="&$ret)
			$ret = $s.SetDNSServerSearchOrder($aDNS)
			LogD("$s.SetDNSServerSearchOrder"&$ret)
		EndIf
	Next
	LogI("应用网络配置，等待 3s")
	Sleep(3000)
EndFunc


Func RegSetNetCfg($sHive, $sTargetDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
	Local $i = 1
	While 1
		$sSubKey = RegEnumKey($sHive & "\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}", $i)
		$i = $i + 1
		If @error Then ExitLoop
		If Not StringIsDigit($sSubKey) Then ContinueLoop
		LogD("$sSubKey " & $sSubKey)

		$sPath = $sHive & "\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}\" & $sSubKey
		$sNetCfgId = RegRead($sPath, "NetCfgInstanceId")
		If @error Then ContinueLoop
		;https://learn.microsoft.com/en-us/windows-hardware/drivers/install/device-instance-ids
		$sDeviceId = RegRead($sPath, "DeviceInstanceID")
		If @error or $sDeviceId <> $sTargetDeviceId Then ContinueLoop

		LogD("$sNetCfgId " & $sNetCfgId)
		LogD("$sDeviceId " & $sDeviceId)

		$sPath = $sHive & "\Services\Tcpip\Parameters\Interfaces\" & $sNetCfgId
		RegWrite($sPath, "EnableDHCP", "REG_DWORD", 0)
		RegWrite($sPath, "IPAddress", "REG_MULTI_SZ", _ArrayToString($aIPAddr, @LF))
		RegWrite($sPath, "SubnetMask", "REG_MULTI_SZ", _ArrayToString($aSubnet, @LF))
		RegWrite($sPath, "DefaultGateway", "REG_MULTI_SZ", _ArrayToString($aDefGateway, @LF))
		RegWrite($sPath, "NameServer", "REG_SZ", _ArrayToString($aDNS, ','))

		Return True
	WEnd
	Return False
EndFunc

Func LoadNetCfgFromIni($sPath)
	LogD("$sPath="&$sPath)
	If Not FileExists($sPath) Then Return
	Local $aArray = IniReadSectionNames($sPath)
	For $i = 1 To $aArray[0]
		$sSection = $aArray[$i]
		If StringLeft($sSection, 6) <> "NetCfg" Then ContinueLoop
		Local $sDeviceId = IniRead($sPath, $sSection, "DeviceId", "")
		Local $aIPAddr = _ArrayFromString(IniRead($sPath, $sSection, "IPAddr", ""))
		;If $bDebug Then _DebugArrayDisplay($aIPAddr, "$aIPAddr")

		Local $aSubnet = _ArrayFromString(IniRead($sPath, $sSection, "Subnet", ""))
		;If $bDebug Then _DebugArrayDisplay($aSubnet, "$aSubnet")

		Local $aDefGateway = _ArrayFromString(IniRead($sPath, $sSection, "DefGateway", ""))
		;If $bDebug Then _DebugArrayDisplay($aDefGateway, "$aDefGateway")

		Local $aDNS = _ArrayFromString(IniRead($sPath, $sSection, "DNS", ""))
		;If $bDebug Then _DebugArrayDisplay($aDNS, "$aDNS")

		;WMISetNetCfg($sDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
		Local $bSet = RegSetNetCfg("HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet", $sDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
		If $bSet Then RestartAdapter($sDeviceId)
	Next
EndFunc


Func ExtractNetCfgFromRegistry($sHive)
	Local $i = 1
	While 1
		$sSubKey = RegEnumKey($sHive & "\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}", $i)
		$i = $i + 1
		If @error Then ExitLoop
		If Not StringIsDigit($sSubKey) Then ContinueLoop
		LogD("$sSubKey " & $sSubKey)

		$sPath = $sHive & "\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}\" & $sSubKey
		$sNetCfgId = RegRead($sPath, "NetCfgInstanceId")
		If @error Then ContinueLoop
		;https://learn.microsoft.com/en-us/windows-hardware/drivers/install/device-instance-ids
		$sDeviceId = RegRead($sPath, "DeviceInstanceID")
		If @error Then ContinueLoop

		If StringLeft($sDeviceId, 3) = "SWD" Or StringLeft($sDeviceId, 10) = "ROOT\KDNIC" Then ContinueLoop ; 过滤各类 WAN Miniport 及 Kernel Debug Adapter

		LogD("$sNetCfgId " & $sNetCfgId)
		LogD("$sDeviceId " & $sDeviceId)

		$sPath = $sHive & "\Services\Tcpip\Parameters\Interfaces\" & $sNetCfgId
		$dwDhcpEnabled = RegRead($sPath, "EnableDHCP")
		If @error Or $dwDhcpEnabled Then ContinueLoop
		LogD("$dwDhcpEnabled " & $dwDhcpEnabled)

		$aIPAddr = StringSplit(RegRead($sPath, "IPAddress"), @LF, $STR_NOCOUNT)
		;If $bDebug Then _DebugArrayDisplay($aIPAddr, "$aIPAddr")

		$aSubnet = StringSplit(RegRead($sPath, "SubnetMask"), @LF, $STR_NOCOUNT)
		;If $bDebug Then _DebugArrayDisplay($aSubnet, "$aSubnet")

		$aDefGateway = StringSplit(RegRead($sPath, "DefaultGateway"), @LF, $STR_NOCOUNT)
		;If $bDebug Then _DebugArrayDisplay($aDefGateway, "$aDefGateway")

		$aDNS = StringSplit(RegRead($sPath, "NameServer"), ',', $STR_NOCOUNT)
		;If $bDebug Then _DebugArrayDisplay($aDNS, "$aDNS")

		;If Not $bOnline Then SetNetCfg($sDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
		;If $bOnline Then
		SaveNetCfg($sNetCfgId, $sDeviceId, $aIPAddr, $aSubnet, $aDefGateway, $aDNS)
	WEnd
EndFunc

Func ExtractNetCfgFromOffline()
	Local $aArray = DriveGetDrive($DT_FIXED)
	For $i = 1 To $aArray[0]
		If $aArray[$i] = EnvGet("SystemDrive") Then ContinueLoop
		Local $sHivePath = $aArray[$i] & "\Windows\System32\config\SYSTEM"
		If FileExists($sHivePath) Then
			LogD("Network: $sDrive " & $aArray[$i])
			$ret = _WinAPI_RegLoadKey($HKEY_USERS, "OfflineWindows", $sHivePath)

			ExtractNetCfgFromRegistry("HKEY_USERS\OfflineWindows\ControlSet001")

			_WinAPI_RegUnloadKey($HKEY_USERS, "OfflineWindows")
		EndIf
	Next
EndFunc

Func RestartAdapter($sDeviceId)
	Local $hDLL = DllOpen("cfgmgr32.dll") ; MUST, otherwise error occured on second DllCall
	If $hDLL = -1 Then
		LogE("unable to load Cfgmgr32.dll: " & @error)
		Return False
	EndIf

	Local $rLocate = DllCall($hDLL, "DWORD", "CM_Locate_DevNodeW", "DWORD*", 0, "WSTR", $sDeviceId, "DWORD", 0)
	If $rLocate[0] <> 0 Then
		LogE("unable to locate DevNode: " & $rLocate[0])
		Return False
	EndIf

	Local $ret = DllCall($hDLL, "DWORD", "CM_Disable_DevNode", "DWORD", $rLocate[1], "ULONG", 0)
	If $ret[0] <> 0 Then
		LogE("unable to disable DevNode: " & $ret[0])
		Return False
	EndIf

	;Sleep(2000)

	$ret = DllCall($hDLL, "DWORD", "CM_Enable_DevNode", "DWORD", $rLocate[1], "ULONG", 0)
	If $ret[0] <> 0 Then
		LogE("unable to enable DevNode: " & $ret[0])
		Return False
	EndIf

	Return True
	;DllClose($hDLL)
EndFunc

