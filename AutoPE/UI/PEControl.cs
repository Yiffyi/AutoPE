using AutoPE.Model;
using System.Windows.Forms;

namespace AutoPE.UI
{
    public partial class PEControl : UserControl
    {
        public PEConfig Config;
        public PEControl()
        {
            InitializeComponent();
        }

        public void LoadConfig(PEConfig cfg)
        {
            Config = cfg;

            tbWim.DataBindings.Add("Text", Config, "WimUrl", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
            tbSdi.DataBindings.Add("Text", Config, "SdiUrl", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
            tbWorker.DataBindings.Add("Text", Config, "WorkerUrl", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
        }
    }
}
