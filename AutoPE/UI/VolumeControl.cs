using AutoPE.Model;
using AutoPE.WMI;
using System;
using System.Collections.Generic;
using System.Data;
using System.Linq;
using System.Windows.Forms;

namespace AutoPE.UI
{
    public partial class VolumeControl : UserControl
    {
        public VolumeConfig Config;
        private List<Volume> volumes;
        private BindingSource sysDataSource = new BindingSource();
        private BindingSource dataDataSource = new BindingSource();
        private BindingSource bootDataSource = new BindingSource();
        public VolumeControl()
        {
            Config = new VolumeConfig
            {
                FormatSystem = true
            };
            InitializeComponent();
        }

        private void VolumeControl_Load(object sender, EventArgs e)
        {
            volumes = WMIObject.FromWQL("SELECT * FROM Win32_Volume WHERE DriveType = 3").Select(raw => new Volume(raw)).ToList();
            sysDataSource.CurrentChanged += (o, _) => { Config.SystemVol = (o as BindingSource).Current as Volume; };
            sysDataSource.DataSource = volumes;

            dataDataSource.CurrentChanged += (o, _) => { Config.DataVol = (o as BindingSource).Current as Volume; };
            dataDataSource.DataSource = volumes;

            //bootDataSource.CurrentChanged += (o, _) => { Config.BootVol = (o as BindingSource).Current as Volume; };
            //bootDataSource.DataSource = volumes;

            cbSysVol.DisplayMember = "HumanReadableCaption";
            cbSysVol.DataSource = sysDataSource;

            cbBootVol.DisplayMember = "HumanReadableCaption";
            cbBootVol.DataSource = bootDataSource;

            cbDataVol.DisplayMember = "HumanReadableCaption";
            cbDataVol.DataSource = dataDataSource;

            // 预选分区
            //cbBootVol.SelectedIndex = volumes.FindIndex(v => v.HasBootmgr);
            cbSysVol.SelectedIndex = volumes.FindIndex(v => v.HasWindows);

            var maxFree = volumes.SkipWhile(v => v.HasWindows).OrderByDescending(v => (ulong)v["FreeSpace"]).First();
            cbDataVol.SelectedIndex = volumes.FindIndex(v => v == maxFree);

            cbFormatSys.DataBindings.Add("Checked", Config, "FormatSystem");
            cbFormatData.DataBindings.Add("Checked", Config, "FormatData");
        }

        private void bEditDiskpart_Click(object sender, EventArgs e)
        {
            var frm = new FormEditDiskpart(Config.DiskpartScript);
            if (frm.ShowDialog() == DialogResult.OK)
            {
                Config.DiskpartScript = frm.DiskpartScript;
            }
        }
    }
}
