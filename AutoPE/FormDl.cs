using System;
using System.ComponentModel;
using System.IO;
using System.Net;
using System.Windows.Forms;

namespace AutoPE
{
    public partial class FormDl : Form
    {
        private readonly WebClient _client;
        private readonly Uri _address;
        private readonly string _filePath;

        public FormDl(string address, string filePath)
        {
            InitializeComponent();
            _client = new WebClient();
            _address = new Uri(address);
            _filePath = filePath;
            _client.DownloadFileCompleted += downloadFileCompleted;
            _client.DownloadProgressChanged += downloadProgressChanged;

            var d = Path.GetDirectoryName(filePath);
            if (!Directory.Exists(d))
            {
                Directory.CreateDirectory(d);
            }

            lFilename.Text = $"正在下载: {filePath}";
            _client.DownloadFileAsync(_address, _filePath);
        }

        private void downloadFileCompleted(object sender, AsyncCompletedEventArgs e)
        {
            pbDl.Value = 100;
            _client.Dispose();
            if (e.Error != null)
            {
                MessageBox.Show($"下载时发生错误: {e.Error.Message}");
                DialogResult = DialogResult.No;
            }
            else
            {
                DialogResult = DialogResult.OK;
            }
        }

        private void downloadProgressChanged(object sender, DownloadProgressChangedEventArgs e)
        {
            if (e.ProgressPercentage % 5 == 0)
            {
                pbDl.Value = e.ProgressPercentage;
            }
        }
    }
}
