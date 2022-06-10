using AutoPE.Model;
using System;
using System.Windows.Forms;

namespace AutoPE.UI
{
    public partial class FileControl : UserControl
    {
        public FileConfig Config;
        private BindingSource normalSource = new BindingSource();
        private BindingSource textSource = new BindingSource();
        public FileControl()
        {
            InitializeComponent();
        }

        public void LoadConfig(FileConfig config)
        {
            Config = config;

            tbSysUrl.DataBindings.Add("Text", Config, "SystemImgUrl", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
            tbSysIndex.DataBindings.Add("Text", Config, "SystemImgIndex");

            normalSource.DataSource = Config.NormalFiles;
            textSource.DataSource = Config.TextFiles;

            lbNormal.DataSource = normalSource;
            lbNormal.DisplayMember = "Url";

            tbNormalUrl.DataBindings.Add("Text", normalSource, "Url", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
            tbNormalPath.DataBindings.Add("Text", normalSource, "Path", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);

            lbText.DataSource = textSource;
            lbText.DisplayMember = "Path";

            tbTextContent.DataBindings.Add("Text", textSource, "Content", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
            tbTextPath.DataBindings.Add("Text", textSource, "Path", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
            cbTextAppend.DataBindings.Add("Checked", textSource, "Append", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);

            tbCustomInput.DataBindings.Add("Text", Config, "CustomInput", false, DataSourceUpdateMode.OnPropertyChanged, string.Empty);
        }

        private void bAddNormal_Click(object sender, EventArgs e)
        {
            normalSource.Add(new FileConfig.NormalFile
            {
                Url = "???"
            });
        }

        private void bAddText_Click(object sender, EventArgs e)
        {
            textSource.Add(new FileConfig.TextFile
            {
                Path = "???"
            });
        }

        private void bDelNormal_Click(object sender, EventArgs e)
        {
            normalSource.RemoveAt(lbNormal.SelectedIndex);
        }

        private void bDelText_Click(object sender, EventArgs e)
        {
            textSource.RemoveAt(lbText.SelectedIndex);
        }
    }
}
