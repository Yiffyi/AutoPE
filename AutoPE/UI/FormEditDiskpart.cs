using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace AutoPE.UI
{
    public partial class FormEditDiskpart : Form
    {
        public string DiskpartScript { get; set; }
        public FormEditDiskpart(string previous)
        {
            InitializeComponent();
            DiskpartScript = previous;
        }

        private void bOK_Click(object sender, EventArgs e)
        {
            DialogResult = DialogResult.OK;
        }

        private void FormEditDiskpart_Load(object sender, EventArgs e)
        {
            tbScript.DataBindings.Add("Text", this, "DiskpartScript", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
        }
    }
}
