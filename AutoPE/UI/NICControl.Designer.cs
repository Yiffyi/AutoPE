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
            this.gbNetL.SuspendLayout();
            this.SuspendLayout();
            // 
            // bLoad
            // 
            this.bLoad.Location = new System.Drawing.Point(204, 220);
            this.bLoad.Name = "bLoad";
            this.bLoad.Size = new System.Drawing.Size(75, 23);
            this.bLoad.TabIndex = 9;
            this.bLoad.Text = "强制静态";
            this.bLoad.UseVisualStyleBackColor = true;
            this.bLoad.Click += new System.EventHandler(this.bLoad_Click);
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(16, 16);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(29, 12);
            this.label1.TabIndex = 8;
            this.label1.Text = "网卡";
            // 
            // cbAdapters
            // 
            this.cbAdapters.FormattingEnabled = true;
            this.cbAdapters.Location = new System.Drawing.Point(63, 13);
            this.cbAdapters.Name = "cbAdapters";
            this.cbAdapters.Size = new System.Drawing.Size(216, 20);
            this.cbAdapters.TabIndex = 7;
            this.cbAdapters.SelectedIndexChanged += new System.EventHandler(this.cbAdapters_SelectedIndexChanged);
            // 
            // gbNetL
            // 
            this.gbNetL.Controls.Add(this.cbDHCP);
            this.gbNetL.Controls.Add(this.lCurGate);
            this.gbNetL.Controls.Add(this.lCurMask);
            this.gbNetL.Controls.Add(this.lCurIP);
            this.gbNetL.Controls.Add(this.lCurMAC);
            this.gbNetL.Location = new System.Drawing.Point(18, 39);
            this.gbNetL.Name = "gbNetL";
            this.gbNetL.Size = new System.Drawing.Size(261, 175);
            this.gbNetL.TabIndex = 6;
            this.gbNetL.TabStop = false;
            this.gbNetL.Text = "当前网络配置";
            // 
            // cbDHCP
            // 
            this.cbDHCP.AutoCheck = false;
            this.cbDHCP.AutoSize = true;
            this.cbDHCP.Location = new System.Drawing.Point(25, 30);
            this.cbDHCP.Name = "cbDHCP";
            this.cbDHCP.Size = new System.Drawing.Size(102, 16);
            this.cbDHCP.TabIndex = 4;
            this.cbDHCP.Text = "DHCP 自动配置";
            this.cbDHCP.UseVisualStyleBackColor = true;
            // 
            // lCurGate
            // 
            this.lCurGate.AutoSize = true;
            this.lCurGate.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurGate.Location = new System.Drawing.Point(21, 130);
            this.lCurGate.Name = "lCurGate";
            this.lCurGate.Size = new System.Drawing.Size(81, 19);
            this.lCurGate.TabIndex = 3;
            this.lCurGate.Text = "lCurGate";
            // 
            // lCurMask
            // 
            this.lCurMask.AutoSize = true;
            this.lCurMask.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurMask.Location = new System.Drawing.Point(21, 103);
            this.lCurMask.Name = "lCurMask";
            this.lCurMask.Size = new System.Drawing.Size(81, 19);
            this.lCurMask.TabIndex = 2;
            this.lCurMask.Text = "lCurMask";
            // 
            // lCurIP
            // 
            this.lCurIP.AutoSize = true;
            this.lCurIP.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurIP.Location = new System.Drawing.Point(21, 76);
            this.lCurIP.Name = "lCurIP";
            this.lCurIP.Size = new System.Drawing.Size(63, 19);
            this.lCurIP.TabIndex = 1;
            this.lCurIP.Text = "lCurIP";
            // 
            // lCurMAC
            // 
            this.lCurMAC.AutoSize = true;
            this.lCurMAC.Font = new System.Drawing.Font("Consolas", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.lCurMAC.Location = new System.Drawing.Point(21, 49);
            this.lCurMAC.Name = "lCurMAC";
            this.lCurMAC.Size = new System.Drawing.Size(72, 19);
            this.lCurMAC.TabIndex = 0;
            this.lCurMAC.Text = "lCurMAC";
            // 
            // NICControl
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.Controls.Add(this.bLoad);
            this.Controls.Add(this.label1);
            this.Controls.Add(this.cbAdapters);
            this.Controls.Add(this.gbNetL);
            this.Name = "NICControl";
            this.Size = new System.Drawing.Size(294, 254);
            this.gbNetL.ResumeLayout(false);
            this.gbNetL.PerformLayout();
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
    }
}
