using AutoPE.WMI;
using System;
using System.Collections.Generic;
using System.Data;
using System.Linq;
using System.Windows.Forms;

namespace AutoPE.UI
{
    public partial class NICControl : UserControl
    {
        public NetworkAdapterConfiguration Nic;
        public NICControl()
        {
            InitializeComponent();

            cbAdapters.DisplayMember = "Key";
            cbAdapters.ValueMember = "Value";

            updateDataSource();
            updateAdapter();
        }
        private void updateDataSource()
        {
            var adapters = WMIObject
                .FromWQL(@"SELECT * From Win32_NetworkAdapterConfiguration WHERE IPEnabled = True")
                .Select(x => new NetworkAdapterConfiguration(x));
            cbAdapters.DataSource = adapters.Select(nic => new KeyValuePair<string, NetworkAdapterConfiguration>((string)nic["Caption"], nic)).ToArray();
        }

        private void updateAdapter()
        {
            NetworkAdapterConfiguration a = (NetworkAdapterConfiguration)cbAdapters.SelectedValue;
            Nic = a;
            cbDHCP.Checked = (bool)a["DHCPEnabled"];
            lCurMAC.Text = (string)a["MACAddress"];
            lCurIP.Text = a.FirstIPAddress ?? "无 IP";
            lCurMask.Text = a.FirstSubnet ?? "无 子网掩码";
            lCurGate.Text = a.FirstDefaultIPGateway ?? "无 默认网关";
        }

        private void bLoad_Click(object sender, EventArgs e)
        {
            var c = (NetworkAdapterConfiguration)cbAdapters.SelectedValue;

            c.EnableStatic(new string[] { lCurIP.Text }, new string[] { lCurMask.Text });
            c.SetGateways(new string[] { lCurGate.Text });
            updateDataSource();
        }

        private void cbAdapters_SelectedIndexChanged(object sender, EventArgs e)
        {
            updateAdapter();
        }
    }
}
