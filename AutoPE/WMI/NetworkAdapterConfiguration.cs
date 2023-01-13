using System;
using System.Management;

namespace AutoPE.WMI
{
    public class NetworkAdapterConfiguration : WMIObject
    {
        public NetworkAdapterConfiguration(ManagementObject rawObject) : base(rawObject)
        {
        }

        /*public static List<NetworkAdapter> GetInstances(bool onlyEnabled = false, string macAddress = null, string caption = null)
{
   ManagementClass mc = new ManagementClass("Win32_NetworkAdapterConfiguration");
   ManagementObjectCollection moc = mc.GetInstances();

   var result = new List<NetworkAdapter>();
   foreach (ManagementObject mo in moc)
   {
       if (onlyEnabled && !(bool)mo["IPEnabled"])
           continue;

       var adapter = new NetworkAdapter(mo);

       if (macAddress != null && adapter.MACAddress != macAddress) continue;
       if (caption != null && adapter.Caption != caption) continue;

       result.Add(adapter);
   }
   return result;
}*/
        public string Caption => (string)RawObject["Caption"];
        public string FirstIPAddress => (RawObject["IPAddress"] as string[])?[0];
        public string FirstSubnet => (RawObject["IPSubnet"] as string[])?[0];
        public string FirstDefaultIPGateway => (RawObject["DefaultIPGateway"] as string[])?[0];
        public bool DHCPEnabled => (bool)RawObject["DHCPEnabled"];
        public string MACAddress => (string)RawObject["MACAddress"];

        public string[] DNSServerSearchOrder => (string[])RawObject["DNSServerSearchOrder"];
        public string DNSServer1 => DNSServerSearchOrder?[0];
        public string DNSServer2 => DNSServerSearchOrder?[1];

        public uint EnableDHCP() =>
            CallMethod<UInt32>("EnableDHCP");

        public uint EnableStatic(string[] ipAddr, string[] subnetMask) =>
            CallMethod<UInt32>("EnableStatic", new
            {
                IPAddress = ipAddr,
                SubnetMask = subnetMask
            });

        public uint SetGateways(string[] gateway) =>
            CallMethod<UInt32>("SetGateways", new
            {
                DefaultIPGateway = gateway
            });

        public uint SetDNSServerSearchOrder(string[] dnsServer) =>
            CallMethod<UInt32>("SetDNSServerSearchOrder", new
            {
                DNSServerSearchOrder = dnsServer
            });
    }
}
