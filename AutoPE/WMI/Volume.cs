using System;
using System.Management;
using System.Text;

namespace AutoPE.WMI
{
    public class Volume : WMIObject
    {
        public Volume(ManagementObject rawObject) : base(rawObject)
        {
            hrCaption = new Lazy<string>(() =>
            {
                StringBuilder b = new StringBuilder();
                if (HasWindows)
                {
                    b.Append("系统 - ");
                }
                else if (HasBootmgr)
                {
                    b.Append("引导区 - ");
                }
                else
                {
                    b.Append("普通分区 - ");
                }

                if (this["Label"] == null)
                {
                    b.Append("本地磁盘");
                }
                else
                {
                    b.Append(this["Label"]);
                }

                if (this["DriveLetter"] != null)
                {
                    b.AppendFormat(" ({0})", this["DriveLetter"]);
                }
                else
                {
                    b.AppendFormat(" ({0})", this["DeviceID"]);
                }

                return b.ToString();
            });
        }

        // https://stackoverflow.com/questions/56784915/python-wmi-can-we-really-say-that-c-is-always-the-boot-drive
        // 系统分区 (Windows 安装位置)
        // 反人类设定
        public bool HasWindows => (bool)this["BootVolume"];

        // 引导分区 (BCD 所在位置)
        // 反人类设定
        public bool HasBootmgr => (bool)this["SystemVolume"];

        private Lazy<string> hrCaption;
        public string HumanReadableCaption => hrCaption.Value;

        public uint Format(string FileSystem, bool QuickFormat, uint ClusterSize = 4096, string Label = "", bool EnableCompression = false) =>
            CallMethod<UInt32>("Format", new
            {
                FileSystem = FileSystem,
                QuickFormat = QuickFormat,
                Label = Label,
                EnableCompression = EnableCompression
            });
    }
}
