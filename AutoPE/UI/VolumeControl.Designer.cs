﻿namespace AutoPE.UI
{
    partial class VolumeControl
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
            this.label5 = new System.Windows.Forms.Label();
            this.cbDataVol = new System.Windows.Forms.ComboBox();
            this.label4 = new System.Windows.Forms.Label();
            this.label2 = new System.Windows.Forms.Label();
            this.cbBootVol = new System.Windows.Forms.ComboBox();
            this.cbSysVol = new System.Windows.Forms.ComboBox();
            this.cbFormatSys = new System.Windows.Forms.CheckBox();
            this.cbFormatData = new System.Windows.Forms.CheckBox();
            this.cbUseDiskpart = new System.Windows.Forms.CheckBox();
            this.bEditDiskpart = new System.Windows.Forms.Button();
            this.SuspendLayout();
            // 
            // label5
            // 
            this.label5.AutoSize = true;
            this.label5.Location = new System.Drawing.Point(26, 130);
            this.label5.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.label5.Name = "label5";
            this.label5.Size = new System.Drawing.Size(82, 24);
            this.label5.TabIndex = 20;
            this.label5.Text = "数据区";
            // 
            // cbDataVol
            // 
            this.cbDataVol.FormattingEnabled = true;
            this.cbDataVol.Location = new System.Drawing.Point(120, 124);
            this.cbDataVol.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbDataVol.Name = "cbDataVol";
            this.cbDataVol.Size = new System.Drawing.Size(320, 32);
            this.cbDataVol.TabIndex = 19;
            // 
            // label4
            // 
            this.label4.AutoSize = true;
            this.label4.Location = new System.Drawing.Point(26, 78);
            this.label4.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.label4.Name = "label4";
            this.label4.Size = new System.Drawing.Size(82, 24);
            this.label4.TabIndex = 18;
            this.label4.Text = "引导区";
            this.label4.Visible = false;
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(26, 26);
            this.label2.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(82, 24);
            this.label2.TabIndex = 17;
            this.label2.Text = "系统区";
            // 
            // cbBootVol
            // 
            this.cbBootVol.FormattingEnabled = true;
            this.cbBootVol.Location = new System.Drawing.Point(120, 72);
            this.cbBootVol.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbBootVol.Name = "cbBootVol";
            this.cbBootVol.Size = new System.Drawing.Size(320, 32);
            this.cbBootVol.TabIndex = 16;
            this.cbBootVol.Visible = false;
            // 
            // cbSysVol
            // 
            this.cbSysVol.FormattingEnabled = true;
            this.cbSysVol.Location = new System.Drawing.Point(120, 20);
            this.cbSysVol.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbSysVol.Name = "cbSysVol";
            this.cbSysVol.Size = new System.Drawing.Size(320, 32);
            this.cbSysVol.TabIndex = 15;
            // 
            // cbFormatSys
            // 
            this.cbFormatSys.AutoSize = true;
            this.cbFormatSys.Checked = true;
            this.cbFormatSys.CheckState = System.Windows.Forms.CheckState.Checked;
            this.cbFormatSys.Enabled = false;
            this.cbFormatSys.Location = new System.Drawing.Point(456, 26);
            this.cbFormatSys.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbFormatSys.Name = "cbFormatSys";
            this.cbFormatSys.Size = new System.Drawing.Size(114, 28);
            this.cbFormatSys.TabIndex = 21;
            this.cbFormatSys.Text = "格式化";
            this.cbFormatSys.UseVisualStyleBackColor = true;
            // 
            // cbFormatData
            // 
            this.cbFormatData.AutoSize = true;
            this.cbFormatData.Location = new System.Drawing.Point(456, 130);
            this.cbFormatData.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbFormatData.Name = "cbFormatData";
            this.cbFormatData.Size = new System.Drawing.Size(114, 28);
            this.cbFormatData.TabIndex = 22;
            this.cbFormatData.Text = "格式化";
            this.cbFormatData.UseVisualStyleBackColor = true;
            // 
            // cbUseDiskpart
            // 
            this.cbUseDiskpart.AutoSize = true;
            this.cbUseDiskpart.Location = new System.Drawing.Point(30, 189);
            this.cbUseDiskpart.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.cbUseDiskpart.Name = "cbUseDiskpart";
            this.cbUseDiskpart.Size = new System.Drawing.Size(306, 28);
            this.cbUseDiskpart.TabIndex = 23;
            this.cbUseDiskpart.Text = "使用 Diskpart 重新分区";
            this.cbUseDiskpart.UseVisualStyleBackColor = true;
            // 
            // bEditDiskpart
            // 
            this.bEditDiskpart.Location = new System.Drawing.Point(30, 233);
            this.bEditDiskpart.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.bEditDiskpart.Name = "bEditDiskpart";
            this.bEditDiskpart.Size = new System.Drawing.Size(292, 46);
            this.bEditDiskpart.TabIndex = 24;
            this.bEditDiskpart.Text = "编辑 Diskpart 脚本";
            this.bEditDiskpart.UseVisualStyleBackColor = true;
            this.bEditDiskpart.Click += new System.EventHandler(this.bEditDiskpart_Click);
            // 
            // VolumeControl
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(12F, 24F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.Controls.Add(this.bEditDiskpart);
            this.Controls.Add(this.cbUseDiskpart);
            this.Controls.Add(this.cbFormatData);
            this.Controls.Add(this.cbFormatSys);
            this.Controls.Add(this.label5);
            this.Controls.Add(this.cbDataVol);
            this.Controls.Add(this.label4);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.cbBootVol);
            this.Controls.Add(this.cbSysVol);
            this.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.Name = "VolumeControl";
            this.Size = new System.Drawing.Size(576, 297);
            this.Load += new System.EventHandler(this.VolumeControl_Load);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label label5;
        private System.Windows.Forms.ComboBox cbDataVol;
        private System.Windows.Forms.Label label4;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.ComboBox cbBootVol;
        private System.Windows.Forms.ComboBox cbSysVol;
        private System.Windows.Forms.CheckBox cbFormatSys;
        private System.Windows.Forms.CheckBox cbFormatData;
        private System.Windows.Forms.CheckBox cbUseDiskpart;
        private System.Windows.Forms.Button bEditDiskpart;
    }
}
