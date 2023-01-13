using AutoPE.Model;
using AutoPE.WMI;
using System;
using System.Collections.Generic;
using System.Data;
using System.Linq;
using System.Net;
using System.Windows.Forms;

namespace AutoPE.UI
{
    public partial class NICControl : UserControl
    {
        public NetworkAdapterConfiguration Nic => (NetworkAdapterConfiguration)_nicSource.Current;
        private BindingSource _nicSource = new BindingSource();

        public NICConfig Config
        {
            get
            {
                return new NICConfig
                {
                    DNSServer1 = tbDns1.Text == Nic.DNSServer1 ? null : tbDns1.Text,
                    DNSServer2 = tbDns2.Text == Nic.DNSServer2 ? null : tbDns2.Text
                };
            }
        }
        public NICControl()
        {
            InitializeComponent();

            updateDataSource();

            cbAdapters.DisplayMember = "Caption";
            cbAdapters.DataSource = _nicSource;

            cbDHCP.DataBindings.Add("Checked", _nicSource, "DHCPEnabled");
            lCurMAC.DataBindings.Add("Text", _nicSource, "MACAddress");
            lCurIP.DataBindings.Add("Text", _nicSource, "FirstIPAddress", false, DataSourceUpdateMode.Never, "无 IP");
            lCurMask.DataBindings.Add("Text", _nicSource, "FirstSubnet", false, DataSourceUpdateMode.Never, "无 子网掩码");
            lCurGate.DataBindings.Add("Text", _nicSource, "FirstDefaultIPGateway", false, DataSourceUpdateMode.Never, "无 默认网关");
            tbDns1.DataBindings.Add("Text", _nicSource, "DNSServer1", false, DataSourceUpdateMode.Never, "");
            tbDns2.DataBindings.Add("Text", _nicSource, "DNSServer2", false, DataSourceUpdateMode.Never, "");
        }

        public void LoadConfig(NICConfig cfg)
        {
            if (cfg.DNSServer1 != null) { tbDns1.Text = cfg.DNSServer1; }
            if (cfg.DNSServer2 != null) { tbDns2.Text = cfg.DNSServer2; }
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

        private void bSetDns_Click(object sender, EventArgs e)
        {
            var c = (NetworkAdapterConfiguration)cbAdapters.SelectedValue;
            c.SetDNSServerSearchOrder(new string[] { tbDns1.Text, tbDns2.Text });
            updateDataSource();
        }
    }
}
