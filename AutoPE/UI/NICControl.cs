using AutoPE.Model;
using AutoPE.WMI;
using System;
using System.Data;
using System.Linq;
using System.Windows.Forms;

namespace AutoPE.UI
{
    public partial class NICControl : UserControl
    {
        public NetworkAdapterConfiguration Nic => (NetworkAdapterConfiguration)_nicSource.Current;
        private BindingSource _nicSource = new BindingSource();

        private NICConfig _config;
        public NICConfig Config
        {
            get => _config;
            set
            {
                _config = value;
                tbDns1.DataBindings.Add("Text", Config, "DNSServer1", false, DataSourceUpdateMode.OnPropertyChanged, "");
                tbDns2.DataBindings.Add("Text", Config, "DNSServer2", false, DataSourceUpdateMode.OnPropertyChanged, "");
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
            c.SetDNSServerSearchOrder(new string[] { c.DNSServer1, c.DNSServer2 });
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
