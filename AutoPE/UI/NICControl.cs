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
        public NetworkAdapterConfiguration Nic => (NetworkAdapterConfiguration)_nicSource.Current;
        private BindingSource _nicSource = new BindingSource();
        public NICControl()
        {
            InitializeComponent();

            updateDataSource();

            cbAdapters.DisplayMember = "Caption";
            cbAdapters.DataSource = _nicSource;

            cbDHCP.DataBindings.Add("Checked", _nicSource, "DHCPEnabled");
            lCurMAC.DataBindings.Add("Text", _nicSource, "MACAddress");
            lCurIP.DataBindings.Add("Text", _nicSource, "FirstIPAddress", false, DataSourceUpdateMode.OnPropertyChanged, "无 IP");
            lCurMask.DataBindings.Add("Text", _nicSource, "FirstSubnet", false, DataSourceUpdateMode.OnPropertyChanged, "无 子网掩码");
            lCurGate.DataBindings.Add("Text", _nicSource, "FirstDefaultIPGateway", false, DataSourceUpdateMode.OnPropertyChanged, "无 默认网关");
        }
        private void updateDataSource()
        {
            var adapters = WMIObject
                .FromWQL(@"SELECT * From Win32_NetworkAdapterConfiguration WHERE IPEnabled = True")
                .Select(x => new NetworkAdapterConfiguration(x));
            _nicSource.DataSource = adapters;
        }

        private void bLoad_Click(object sender, EventArgs e)
        {
            var c = (NetworkAdapterConfiguration)cbAdapters.SelectedValue;

            c.EnableStatic(new string[] { lCurIP.Text }, new string[] { lCurMask.Text });
            c.SetGateways(new string[] { lCurGate.Text });
            updateDataSource();
        }
    }
}
