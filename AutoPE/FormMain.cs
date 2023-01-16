using AutoPE.Model;
using AutoPE.WMI;
using IniParser;
using IniParser.Model;
using Newtonsoft.Json;
using System;
using System.IO;
using System.Text;
using System.Windows.Forms;

namespace AutoPE
{
    public partial class FormMain : Form
    {
        private FileIniDataParser iniParser = new FileIniDataParser();
        private IniData iniCfg = new IniData();
        public FormMain()
        {
            InitializeComponent();
        }

        private void FormMain_Load(object sender, EventArgs e)
        {
            var rootConfig = JsonConvert.DeserializeObject<RootConfig>(File.ReadAllText("config.json"));
            fileControl1.LoadConfig(rootConfig.Inject);
            peControl1.LoadConfig(rootConfig.PE);
            nicControl1.Config = rootConfig.NIC;
        }
        public void SaveConfig()
        {
            var rootConfig = new RootConfig
            {
                PE = peControl1.Config,
                Inject = fileControl1.Config,
                NIC = nicControl1.Config
            };
            File.WriteAllText("config.json", JsonConvert.SerializeObject(rootConfig, Formatting.Indented));
        }
        public void WriteVolumeCmd(VolumeConfig vol, PEConfig pe)
        {
            if (!vol.UseDiskpart)
            {
                File.WriteAllText($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\volume.cmd", vol.CompileVolumeCmd());
            }
            else
            {
                File.WriteAllText($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\diskpart.txt", vol.DiskpartScript);
                iniCfg["Volume"]["UseDiskpart"] = "True";
            }

        }
        private void WriteNICCmd(VolumeConfig vol, PEConfig pe)
        {
            var nic = nicControl1.Nic;
            var nicCfg = nicControl1.Config;
            if (!(bool)nic["DHCPEnabled"] && nic.FirstIPAddress != null && nic.FirstSubnet != null)
            {
                var cmd = File.CreateText($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\network.cmd");

                cmd.WriteLine($"wmic path Win32_NetworkAdapterConfiguration where MACAddress='{nic["MACAddress"]}' call EnableStatic IPAddress={nic.FirstIPAddress} SubnetMask={nic.FirstSubnet}");

                if (nic.FirstDefaultIPGateway != null)
                    cmd.WriteLine($"wmic path Win32_NetworkAdapterConfiguration where MACAddress='{nic["MACAddress"]}' call SetGateways DefaultIPGateway={nic.FirstDefaultIPGateway}");

                string ns1 = "", ns2 = "";
                if (nic.DNSServerSearchOrder != null)
                {
                    if (nic.DNSServerSearchOrder.Length >= 1) ns1 = nic.DNSServerSearchOrder[0];
                    if (nic.DNSServerSearchOrder.Length >= 2) ns2 = nic.DNSServerSearchOrder[1];
                }
                if (nicCfg.DNSServer1 != null && nicCfg.DNSServer1.Length > 0) ns1 = nicCfg.DNSServer1;
                if (nicCfg.DNSServer2 != null && nicCfg.DNSServer2.Length > 0) ns2 = nicCfg.DNSServer2;
                cmd.WriteLine($"wmic path Win32_NetworkAdapterConfiguration where MACAddress='{nic["MACAddress"]}' call SetDNSServerSearchOrder DNSServerSearchOrder=({ns1}, {ns2})");
                cmd.Close();
            }
        }

        private void WriteFileConfig(VolumeConfig vol, PEConfig pe)
        {
            var cfg = fileControl1.Config;
            File.WriteAllText($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\systemAria2.txt", cfg.CompileSysImgAria2Input());
            File.WriteAllText($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\injectAria2.txt", cfg.CompileInjectAria2Input());
            iniCfg.Merge(cfg.CompileSysImgIni());
            iniCfg.Merge(cfg.PrepareTextFiles($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\", Encoding.ASCII));
        }

        public void WriteBCD(VolumeConfig vol, PEConfig pe)
        {
            var store = new BcdStore();
            store.CreateObject(pe.RamdiskId, BcdConsts.BCDO_TYPE_DEVICE);
            store.CreateObject(pe.OsLoaderId, BcdConsts.BCDO_TYPE_VISTA_OS_ENTRY);

            var ramdiskObj = new BcdObject(pe.RamdiskId);
            var osLoaderObj = new BcdObject(pe.OsLoaderId);
            var bootmgrObj = new BcdObject(BcdConsts.GUID_WINDOWS_BOOTMGR);

            ramdiskObj.SetPartitionDeviceElement(
                BcdConsts.BCDE_DEVICEOBJ_TYPE_SDIDEVICE,
                BcdConsts.BCDE_DEVICE_TYPE_PARTITION,
                "",

                NativeMethod.QueryDosDevice((string)vol.DataVol["DriveLetter"])
            );

            ramdiskObj.SetStringElement(BcdConsts.BCDE_DEVICEOBJ_TYPE_SDIPATH, @"\AutoPE\stage0\boot.sdi");

            osLoaderObj.SetFileDeviceElement(
                BcdConsts.BCDE_OSLOADER_OSDEVICE,

                BcdConsts.BCDE_DEVICE_TYPE_RAMDISK,
                pe.RamdiskId,
                @"\AutoPE\stage0\boot.wim",

                // Parent device
                BcdConsts.BCDE_DEVICE_TYPE_PARTITION,
                "",
                NativeMethod.QueryDosDevice((string)vol.DataVol["DriveLetter"])
            );

            osLoaderObj.SetFileDeviceElement(
                BcdConsts.BCDE_LIBRARY_TYPE_APPLICATIONDEVICE,

                BcdConsts.BCDE_DEVICE_TYPE_RAMDISK,
                pe.RamdiskId,
                @"\AutoPE\stage0\boot.wim",

                // Parent device
                BcdConsts.BCDE_DEVICE_TYPE_PARTITION,
                "",
                NativeMethod.QueryDosDevice((string)vol.DataVol["DriveLetter"])
            );

            osLoaderObj.SetStringElement(BcdConsts.BCDE_LIBRARY_TYPE_DESCRIPTION, pe.Description);
            osLoaderObj.SetStringElement(BcdConsts.BCDE_OSLOADER_SYSTEMROOT, @"\windows");
            osLoaderObj.SetBooleanElement(BcdConsts.BCDE_OSLOADER_WINPEMODE, true);

            bootmgrObj.SetObjectListElement(BcdConsts.BCDE_BOOTMGR_BOOTSEQUENCE, new string[] { pe.OsLoaderId });
        }

        public bool PrepareBootFiles(VolumeConfig vol, PEConfig pe)
        {
            try
            {
                Directory.Delete($"{vol.DataVol["DriveLetter"]}\\AutoPE\\", true);
            }
            catch (DirectoryNotFoundException) { }

            Directory.CreateDirectory($"{vol.DataVol["DriveLetter"]}\\AutoPE");
            Directory.CreateDirectory($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage0");
            Directory.CreateDirectory($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1");

            if (new FormDl(pe.SdiUrl, $"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage0\\boot.sdi").ShowDialog() != DialogResult.OK)
            {
                return false;
            }
            if (new FormDl(pe.WorkerUrl, $"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\worker.exe").ShowDialog() != DialogResult.OK)
            {
                return false;
            }
            if (new FormDl(pe.WimUrl, $"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage0\\boot.wim").ShowDialog() != DialogResult.OK)
            {
                return false;
            }
            return true;
        }
        public void WriteIniConfig(VolumeConfig vol, PEConfig pe)
        {
            iniParser.WriteFile($"{vol.DataVol["DriveLetter"]}\\AutoPE\\stage1\\config.ini", iniCfg, fileEncoding: Encoding.ASCII);
        }
        public void WriteVolumeLabel(VolumeConfig vol)
        {
            //vol.BootVol["Label"] = "[Boot]";
            vol.SystemVol["Label"] = "[DP] System";
            vol.DataVol["Label"] = "[DP] Data";
            //vol.BootVol.CommitChanges();
            vol.SystemVol.CommitChanges();
            vol.DataVol.CommitChanges();
        }
        private void bGo_Click(object sender, EventArgs e)
        {
            if (!PrepareBootFiles(volumeControl1.Config, peControl1.Config))
            {
                MessageBox.Show("操作取消了");
                return;
            }

            WriteFileConfig(volumeControl1.Config, peControl1.Config);
            WriteVolumeCmd(volumeControl1.Config, peControl1.Config);
            WriteNICCmd(volumeControl1.Config, peControl1.Config);
            WriteIniConfig(volumeControl1.Config, peControl1.Config);
            WriteVolumeLabel(volumeControl1.Config);
            WriteBCD(volumeControl1.Config, peControl1.Config);
            MessageBox.Show("重启时将自动进入 PE");
        }

        private void bSave_Click(object sender, EventArgs e)
        {
            SaveConfig();
        }
    }
}
