using AutoPE.WMI;
using System.Text;

namespace AutoPE.Model
{
    public class VolumeConfig
    {
        public Volume SystemVol { get; set; }
        public bool FormatSystem { get; set; }
        public Volume DataVol { get; set; }
        public bool FormatData { get; set; }
        //public Volume BootVol { get; set; }
        public bool UseDiskpart { get; set; }
        public string DiskpartScript { get; set; }

        public string CompileVolumeCmd()
        {
            var b = new StringBuilder();

            // 由格式化时的时间决定. 格式化后即改变
            //var bootSerial = (uint)BootVol["SerialNumber"];
            var systemSerial = (uint)SystemVol["SerialNumber"];
            var dataSerial = (uint)DataVol["SerialNumber"];

            if (FormatSystem) b.AppendLine($"wmic path Win32_Volume where Label='[DP] System' call Format FileSystem=NTFS QuickFormat=true Label='[DP] System'");

            b.AppendLine($"wmic path Win32_Volume where Label='[DP] System' set DriveLetter=I:");

            if (FormatData) b.AppendLine($"wmic path Win32_Volume where Label='[DP] Data' call Format FileSystem=NTFS QuickFormat=true Label='[DP] Data'");

            b.AppendLine($"wmic path Win32_Volume where Label='[DP] Data' set DriveLetter=J:");

            //if (bootSerial != systemSerial && bootSerial != dataSerial)
            //{
            //    // b.AppendLine($"wmic path Win32_Volume where SerialNumber='{bootSerial}' call Format FileSystem=FAT32 QuickFormat=true Label='Boot'");
            //    // b.AppendLine($"wmic path Win32_Volume where Label='Boot' set DriveLetter=K:");
            //    b.AppendLine($"wmic path Win32_Volume where Label='[DP] Boot' set DriveLetter=K:");
            //}

            return b.ToString();
        }
    }
}
