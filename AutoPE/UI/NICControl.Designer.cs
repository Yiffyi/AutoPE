namespace AutoPE.UI
{
    partial class NICControl
    {
        /// <summary> 
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary> 
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region 组件设计器生成的代码

        /// <summary> 
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            this.bLoad = new System.Windows.Forms.Button();
            this.label1 = new System.Windows.Forms.Label();
            this.cbAdapters = new System.Windows.Forms.ComboBox();
            this.gbNetL = new System.Windows.Forms.GroupBox();
            this.cbDHCP = new System.Windows.Forms.CheckBox();
            this.lCurGate = new System.Windows.Forms.Label();
            this.lCurMask = new System.Windows.Forms.Label();
            this.lCurIP = new System.Windows.Forms.Label();
            this.lCurMAC = new System.Windows.Forms.Label();
            this.groupBox1 = new System.Windows.Forms.GroupBox();
            this.tbDns1 = new System.Windows.Forms.TextBox();
            this.tbDns2 = new System.Windows.Forms.TextBox();
            this.bSetDns = new System.Windows.Forms.Button();
            this.gbNetL.SuspendLayout();
            this.groupBox1.SuspendLayout();
            this.SuspendLayout();
            // 
            // bLoad
            // 
            this.bLoad.Location = new System.Drawing.Point(404, 460);
            this.bLoad.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.bLoad.Name = "bLoad";
            this.bLoad.Size = new System.Drawing.Size(150, 46);
            this.bLoad.TabIndex = 9;
            this.bLoad.Text = "强制静态";
            this.bLoad.UseVisualStyleBackColor = true;
            this.bLoad.Click += new System.EventHandler(this.bLoad_Click);
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(32, 32);
            this.label1.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(58, 24);
            this.label1.TabIndex = 8;
            this.label1.Text = "网卡";
            // 
            // cbAdapters
            // 
            this.cbAdapters.FormattingEnabled = true;
            this.cbAdapters.Location = new System.Drawing.Point(126, 26);
            this.cbAdapters.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbAdapters.Name = "cbAdapters";
            this.cbAdapters.Size = new System.Drawing.Size(428, 32);
            this.cbAdapters.TabIndex = 7;
            // 
            // gbNetL
            // 
            this.gbNetL.Controls.Add(this.cbDHCP);
            this.gbNetL.Controls.Add(this.lCurGate);
            this.gbNetL.Controls.Add(this.lCurMask);
            this.gbNetL.Controls.Add(this.lCurIP);
            this.gbNetL.Controls.Add(this.lCurMAC);
            this.gbNetL.Location = new System.Drawing.Point(36, 78);
            this.gbNetL.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.gbNetL.Name = "gbNetL";
            this.gbNetL.Padding = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.gbNetL.Size = new System.Drawing.Size(522, 350);
            this.gbNetL.TabIndex = 6;
            this.gbNetL.TabStop = false;
            this.gbNetL.Text = "当前网络配置";
            // 
            // cbDHCP
            // 
            this.cbDHCP.AutoCheck = false;
            this.cbDHCP.AutoSize = true;
            this.cbDHCP.Location = new System.Drawing.Point(50, 60);
            this.cbDHCP.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbDHCP.Name = "cbDHCP";
            this.cbDHCP.Size = new System.Drawing.Size(198, 28);
            this.cbDHCP.TabIndex = 4;
            this.cbDHCP.Text = "DHCP 自动配置";
            this.cbDHCP.UseVisualStyleBackColor = true;
            // 
            // lCurGate
            // 
            this.lCurGate.AutoSize = true;
            this.lCurGate.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurGate.Location = new System.Drawing.Point(42, 260);
            this.lCurGate.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.lCurGate.Name = "lCurGate";
            this.lCurGate.Size = new System.Drawing.Size(161, 37);
            this.lCurGate.TabIndex = 3;
            this.lCurGate.Text = "lCurGate";
            // 
            // lCurMask
            // 
            this.lCurMask.AutoSize = true;
            this.lCurMask.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurMask.Location = new System.Drawing.Point(42, 206);
            this.lCurMask.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.lCurMask.Name = "lCurMask";
            this.lCurMask.Size = new System.Drawing.Size(161, 37);
            this.lCurMask.TabIndex = 2;
            this.lCurMask.Text = "lCurMask";
            // 
            // lCurIP
            // 
            this.lCurIP.AutoSize = true;
            this.lCurIP.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurIP.Location = new System.Drawing.Point(42, 152);
            this.lCurIP.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.lCurIP.Name = "lCurIP";
            this.lCurIP.Size = new System.Drawing.Size(125, 37);
            this.lCurIP.TabIndex = 1;
            this.lCurIP.Text = "lCurIP";
            // 
            // lCurMAC
            // 
            this.lCurMAC.AutoSize = true;
            this.lCurMAC.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurMAC.Location = new System.Drawing.Point(42, 98);
            this.lCurMAC.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.lCurMAC.Name = "lCurMAC";
            this.lCurMAC.Size = new System.Drawing.Size(143, 37);
            this.lCurMAC.TabIndex = 0;
            this.lCurMAC.Text = "lCurMAC";
            // 
            // groupBox1
            // 
            this.groupBox1.Controls.Add(this.tbDns2);
            this.groupBox1.Controls.Add(this.tbDns1);
            this.groupBox1.Location = new System.Drawing.Point(36, 437);
            this.groupBox1.Name = "groupBox1";
            this.groupBox1.Size = new System.Drawing.Size(255, 127);
            this.groupBox1.TabIndex = 10;
            this.groupBox1.TabStop = false;
            this.groupBox1.Text = "DNS 服务器";
            // 
            // tbDns1
            // 
            this.tbDns1.Location = new System.Drawing.Point(6, 34);
            this.tbDns1.Name = "tbDns1";
            this.tbDns1.Size = new System.Drawing.Size(242, 35);
            this.tbDns1.TabIndex = 0;
            // 
            // tbDns2
            // 
            this.tbDns2.Location = new System.Drawing.Point(6, 75);
            this.tbDns2.Name = "tbDns2";
            this.tbDns2.Size = new System.Drawing.Size(242, 35);
            this.tbDns2.TabIndex = 1;
            // 
            // bSetDns
            // 
            this.bSetDns.Location = new System.Drawing.Point(404, 521);
            this.bSetDns.Name = "bSetDns";
            this.bSetDns.Size = new System.Drawing.Size(150, 43);
            this.bSetDns.TabIndex = 11;
            this.bSetDns.Text = "设置 DNS";
            this.bSetDns.UseVisualStyleBackColor = true;
            this.bSetDns.Click += new System.EventHandler(this.bSetDns_Click);
            // 
            // NICControl
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(12F, 24F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.Controls.Add(this.bSetDns);
            this.Controls.Add(this.groupBox1);
            this.Controls.Add(this.bLoad);
            this.Controls.Add(this.label1);
            this.Controls.Add(this.cbAdapters);
            this.Controls.Add(this.gbNetL);
            this.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.Name = "NICControl";
            this.Size = new System.Drawing.Size(588, 593);
            this.gbNetL.ResumeLayout(false);
            this.gbNetL.PerformLayout();
            this.groupBox1.ResumeLayout(false);
            this.groupBox1.PerformLayout();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Button bLoad;
        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.ComboBox cbAdapters;
        private System.Windows.Forms.GroupBox gbNetL;
        private System.Windows.Forms.CheckBox cbDHCP;
        private System.Windows.Forms.Label lCurGate;
        private System.Windows.Forms.Label lCurMask;
        private System.Windows.Forms.Label lCurIP;
        private System.Windows.Forms.Label lCurMAC;
        private System.Windows.Forms.GroupBox groupBox1;
        private System.Windows.Forms.TextBox tbDns2;
        private System.Windows.Forms.TextBox tbDns1;
        private System.Windows.Forms.Button bSetDns;
    }
}
