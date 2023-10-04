
namespace AutoPE
{
    partial class FormMain
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            AutoPE.Model.NICConfig nicConfig1 = new AutoPE.Model.NICConfig();
            this.bGo = new System.Windows.Forms.Button();
            this.bSave = new System.Windows.Forms.Button();
            this.lVersion = new System.Windows.Forms.Label();
            this.peControl1 = new AutoPE.UI.PEControl();
            this.fileControl1 = new AutoPE.UI.FileControl();
            this.volumeControl1 = new AutoPE.UI.VolumeControl();
            this.nicControl1 = new AutoPE.UI.NICControl();
            this.SuspendLayout();
            // 
            // bGo
            // 
            this.bGo.Location = new System.Drawing.Point(140, 11);
            this.bGo.Margin = new System.Windows.Forms.Padding(4, 4, 4, 4);
            this.bGo.Name = "bGo";
            this.bGo.Size = new System.Drawing.Size(112, 34);
            this.bGo.TabIndex = 15;
            this.bGo.Text = "Go";
            this.bGo.UseVisualStyleBackColor = true;
            this.bGo.Click += new System.EventHandler(this.bGo_Click);
            // 
            // bSave
            // 
            this.bSave.Location = new System.Drawing.Point(18, 11);
            this.bSave.Margin = new System.Windows.Forms.Padding(4, 4, 4, 4);
            this.bSave.Name = "bSave";
            this.bSave.Size = new System.Drawing.Size(112, 34);
            this.bSave.TabIndex = 17;
            this.bSave.Text = "保存";
            this.bSave.UseVisualStyleBackColor = true;
            this.bSave.Click += new System.EventHandler(this.bSave_Click);
            // 
            // lVersion
            // 
            this.lVersion.AutoSize = true;
            this.lVersion.Location = new System.Drawing.Point(918, 20);
            this.lVersion.Margin = new System.Windows.Forms.Padding(2, 0, 2, 0);
            this.lVersion.Name = "lVersion";
            this.lVersion.Size = new System.Drawing.Size(134, 18);
            this.lVersion.TabIndex = 18;
            this.lVersion.Text = "客户端 v2023 b";
            // 
            // peControl1
            // 
            this.peControl1.Location = new System.Drawing.Point(474, 60);
            this.peControl1.Margin = new System.Windows.Forms.Padding(9);
            this.peControl1.Name = "peControl1";
            this.peControl1.Size = new System.Drawing.Size(586, 135);
            this.peControl1.TabIndex = 16;
            // 
            // fileControl1
            // 
            this.fileControl1.Location = new System.Drawing.Point(474, 240);
            this.fileControl1.Margin = new System.Windows.Forms.Padding(9);
            this.fileControl1.Name = "fileControl1";
            this.fileControl1.Size = new System.Drawing.Size(832, 498);
            this.fileControl1.TabIndex = 13;
            // 
            // volumeControl1
            // 
            this.volumeControl1.Location = new System.Drawing.Point(32, 518);
            this.volumeControl1.Margin = new System.Windows.Forms.Padding(9);
            this.volumeControl1.Name = "volumeControl1";
            this.volumeControl1.Size = new System.Drawing.Size(427, 220);
            this.volumeControl1.TabIndex = 12;
            // 
            // nicControl1
            // 
            nicConfig1.DNSServer1 = null;
            nicConfig1.DNSServer2 = null;
            nicConfig1.PreferredNICName = null;
            this.nicControl1.Config = nicConfig1;
            this.nicControl1.Location = new System.Drawing.Point(18, 60);
            this.nicControl1.Margin = new System.Windows.Forms.Padding(9);
            this.nicControl1.Name = "nicControl1";
            this.nicControl1.Size = new System.Drawing.Size(441, 440);
            this.nicControl1.TabIndex = 11;
            // 
            // FormMain
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(9F, 18F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.AutoSize = true;
            this.ClientSize = new System.Drawing.Size(1306, 777);
            this.Controls.Add(this.lVersion);
            this.Controls.Add(this.bSave);
            this.Controls.Add(this.peControl1);
            this.Controls.Add(this.bGo);
            this.Controls.Add(this.fileControl1);
            this.Controls.Add(this.volumeControl1);
            this.Controls.Add(this.nicControl1);
            this.Margin = new System.Windows.Forms.Padding(4, 4, 4, 4);
            this.Name = "FormMain";
            this.Text = "春晖中学系统部署客户端";
            this.Load += new System.EventHandler(this.FormMain_Load);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion
        private UI.NICControl nicControl1;
        private UI.VolumeControl volumeControl1;
        private UI.FileControl fileControl1;
        private System.Windows.Forms.Button bGo;
        private UI.PEControl peControl1;
        private System.Windows.Forms.Button bSave;
        private System.Windows.Forms.Label lVersion;
    }
}