namespace AutoPE.UI
{
    partial class FileControl
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
            this.tabControl1 = new System.Windows.Forms.TabControl();
            this.tbpSys = new System.Windows.Forms.TabPage();
            this.tbSysUrl = new System.Windows.Forms.TextBox();
            this.label12 = new System.Windows.Forms.Label();
            this.tbpNormal = new System.Windows.Forms.TabPage();
            this.bDelNormal = new System.Windows.Forms.Button();
            this.label10 = new System.Windows.Forms.Label();
            this.bAddNormal = new System.Windows.Forms.Button();
            this.label5 = new System.Windows.Forms.Label();
            this.label3 = new System.Windows.Forms.Label();
            this.tbNormalPath = new System.Windows.Forms.TextBox();
            this.tbNormalUrl = new System.Windows.Forms.TextBox();
            this.label2 = new System.Windows.Forms.Label();
            this.lbNormal = new System.Windows.Forms.ListBox();
            this.tbpText = new System.Windows.Forms.TabPage();
            this.cbTextAppend = new System.Windows.Forms.CheckBox();
            this.bDelText = new System.Windows.Forms.Button();
            this.label11 = new System.Windows.Forms.Label();
            this.tbTextPath = new System.Windows.Forms.TextBox();
            this.bAddText = new System.Windows.Forms.Button();
            this.label6 = new System.Windows.Forms.Label();
            this.label7 = new System.Windows.Forms.Label();
            this.label8 = new System.Windows.Forms.Label();
            this.tbTextContent = new System.Windows.Forms.TextBox();
            this.label9 = new System.Windows.Forms.Label();
            this.lbText = new System.Windows.Forms.ListBox();
            this.tbpCustom = new System.Windows.Forms.TabPage();
            this.tbCustomParam = new System.Windows.Forms.TextBox();
            this.label1 = new System.Windows.Forms.Label();
            this.tbCustomInput = new System.Windows.Forms.TextBox();
            this.label4 = new System.Windows.Forms.Label();
            this.tbSysIndex = new System.Windows.Forms.TextBox();
            this.tabControl1.SuspendLayout();
            this.tbpSys.SuspendLayout();
            this.tbpNormal.SuspendLayout();
            this.tbpText.SuspendLayout();
            this.tbpCustom.SuspendLayout();
            this.SuspendLayout();
            // 
            // tabControl1
            // 
            this.tabControl1.Controls.Add(this.tbpSys);
            this.tabControl1.Controls.Add(this.tbpNormal);
            this.tabControl1.Controls.Add(this.tbpText);
            this.tabControl1.Controls.Add(this.tbpCustom);
            this.tabControl1.Location = new System.Drawing.Point(3, 3);
            this.tabControl1.Name = "tabControl1";
            this.tabControl1.SelectedIndex = 0;
            this.tabControl1.Size = new System.Drawing.Size(441, 326);
            this.tabControl1.TabIndex = 0;
            // 
            // tbpSys
            // 
            this.tbpSys.Controls.Add(this.label4);
            this.tbpSys.Controls.Add(this.tbSysIndex);
            this.tbpSys.Controls.Add(this.tbSysUrl);
            this.tbpSys.Controls.Add(this.label12);
            this.tbpSys.Location = new System.Drawing.Point(4, 22);
            this.tbpSys.Name = "tbpSys";
            this.tbpSys.Padding = new System.Windows.Forms.Padding(3);
            this.tbpSys.Size = new System.Drawing.Size(433, 300);
            this.tbpSys.TabIndex = 3;
            this.tbpSys.Text = "系统镜像";
            this.tbpSys.UseVisualStyleBackColor = true;
            // 
            // tbSysUrl
            // 
            this.tbSysUrl.Location = new System.Drawing.Point(65, 56);
            this.tbSysUrl.Name = "tbSysUrl";
            this.tbSysUrl.Size = new System.Drawing.Size(362, 21);
            this.tbSysUrl.TabIndex = 1;
            // 
            // label12
            // 
            this.label12.AutoSize = true;
            this.label12.Location = new System.Drawing.Point(36, 59);
            this.label12.Name = "label12";
            this.label12.Size = new System.Drawing.Size(23, 12);
            this.label12.TabIndex = 0;
            this.label12.Text = "URL";
            // 
            // tbpNormal
            // 
            this.tbpNormal.Controls.Add(this.bDelNormal);
            this.tbpNormal.Controls.Add(this.label10);
            this.tbpNormal.Controls.Add(this.bAddNormal);
            this.tbpNormal.Controls.Add(this.label5);
            this.tbpNormal.Controls.Add(this.label3);
            this.tbpNormal.Controls.Add(this.tbNormalPath);
            this.tbpNormal.Controls.Add(this.tbNormalUrl);
            this.tbpNormal.Controls.Add(this.label2);
            this.tbpNormal.Controls.Add(this.lbNormal);
            this.tbpNormal.Location = new System.Drawing.Point(4, 22);
            this.tbpNormal.Name = "tbpNormal";
            this.tbpNormal.Padding = new System.Windows.Forms.Padding(3);
            this.tbpNormal.Size = new System.Drawing.Size(433, 300);
            this.tbpNormal.TabIndex = 0;
            this.tbpNormal.Text = "普通文件";
            this.tbpNormal.UseVisualStyleBackColor = true;
            // 
            // bDelNormal
            // 
            this.bDelNormal.Location = new System.Drawing.Point(340, 92);
            this.bDelNormal.Name = "bDelNormal";
            this.bDelNormal.Size = new System.Drawing.Size(75, 23);
            this.bDelNormal.TabIndex = 10;
            this.bDelNormal.Text = "删除";
            this.bDelNormal.UseVisualStyleBackColor = true;
            this.bDelNormal.Click += new System.EventHandler(this.bDelNormal_Click);
            // 
            // label10
            // 
            this.label10.AutoSize = true;
            this.label10.Location = new System.Drawing.Point(191, 68);
            this.label10.Name = "label10";
            this.label10.Size = new System.Drawing.Size(23, 12);
            this.label10.TabIndex = 9;
            this.label10.Text = "C:\\";
            // 
            // bAddNormal
            // 
            this.bAddNormal.Location = new System.Drawing.Point(89, 8);
            this.bAddNormal.Name = "bAddNormal";
            this.bAddNormal.Size = new System.Drawing.Size(37, 23);
            this.bAddNormal.TabIndex = 8;
            this.bAddNormal.Text = "+";
            this.bAddNormal.UseVisualStyleBackColor = true;
            this.bAddNormal.Click += new System.EventHandler(this.bAddNormal_Click);
            // 
            // label5
            // 
            this.label5.AutoSize = true;
            this.label5.Location = new System.Drawing.Point(132, 68);
            this.label5.Name = "label5";
            this.label5.Size = new System.Drawing.Size(53, 12);
            this.label5.TabIndex = 7;
            this.label5.Text = "保存路径";
            // 
            // label3
            // 
            this.label3.AutoSize = true;
            this.label3.Location = new System.Drawing.Point(132, 41);
            this.label3.Name = "label3";
            this.label3.Size = new System.Drawing.Size(23, 12);
            this.label3.TabIndex = 5;
            this.label3.Text = "URL";
            // 
            // tbNormalPath
            // 
            this.tbNormalPath.Location = new System.Drawing.Point(214, 65);
            this.tbNormalPath.Name = "tbNormalPath";
            this.tbNormalPath.Size = new System.Drawing.Size(201, 21);
            this.tbNormalPath.TabIndex = 3;
            // 
            // tbNormalUrl
            // 
            this.tbNormalUrl.Location = new System.Drawing.Point(191, 38);
            this.tbNormalUrl.Name = "tbNormalUrl";
            this.tbNormalUrl.Size = new System.Drawing.Size(224, 21);
            this.tbNormalUrl.TabIndex = 2;
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(6, 13);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(77, 12);
            this.label2.TabIndex = 1;
            this.label2.Text = "普通文件列表";
            // 
            // lbNormal
            // 
            this.lbNormal.FormattingEnabled = true;
            this.lbNormal.ItemHeight = 12;
            this.lbNormal.Location = new System.Drawing.Point(6, 38);
            this.lbNormal.Name = "lbNormal";
            this.lbNormal.Size = new System.Drawing.Size(120, 256);
            this.lbNormal.TabIndex = 0;
            // 
            // tbpText
            // 
            this.tbpText.Controls.Add(this.cbTextAppend);
            this.tbpText.Controls.Add(this.bDelText);
            this.tbpText.Controls.Add(this.label11);
            this.tbpText.Controls.Add(this.tbTextPath);
            this.tbpText.Controls.Add(this.bAddText);
            this.tbpText.Controls.Add(this.label6);
            this.tbpText.Controls.Add(this.label7);
            this.tbpText.Controls.Add(this.label8);
            this.tbpText.Controls.Add(this.tbTextContent);
            this.tbpText.Controls.Add(this.label9);
            this.tbpText.Controls.Add(this.lbText);
            this.tbpText.Location = new System.Drawing.Point(4, 22);
            this.tbpText.Name = "tbpText";
            this.tbpText.Padding = new System.Windows.Forms.Padding(3);
            this.tbpText.Size = new System.Drawing.Size(433, 300);
            this.tbpText.TabIndex = 1;
            this.tbpText.Text = "文本文件";
            this.tbpText.UseVisualStyleBackColor = true;
            // 
            // cbTextAppend
            // 
            this.cbTextAppend.AutoSize = true;
            this.cbTextAppend.Location = new System.Drawing.Point(191, 249);
            this.cbTextAppend.Name = "cbTextAppend";
            this.cbTextAppend.Size = new System.Drawing.Size(84, 16);
            this.cbTextAppend.TabIndex = 22;
            this.cbTextAppend.Text = "添加到末尾";
            this.cbTextAppend.UseVisualStyleBackColor = true;
            // 
            // bDelText
            // 
            this.bDelText.Location = new System.Drawing.Point(346, 270);
            this.bDelText.Name = "bDelText";
            this.bDelText.Size = new System.Drawing.Size(75, 23);
            this.bDelText.TabIndex = 21;
            this.bDelText.Text = "删除";
            this.bDelText.UseVisualStyleBackColor = true;
            this.bDelText.Click += new System.EventHandler(this.bDelText_Click);
            // 
            // label11
            // 
            this.label11.AutoSize = true;
            this.label11.Location = new System.Drawing.Point(191, 221);
            this.label11.Name = "label11";
            this.label11.Size = new System.Drawing.Size(23, 12);
            this.label11.TabIndex = 20;
            this.label11.Text = "C:\\";
            // 
            // tbTextPath
            // 
            this.tbTextPath.Location = new System.Drawing.Point(215, 218);
            this.tbTextPath.Name = "tbTextPath";
            this.tbTextPath.Size = new System.Drawing.Size(206, 21);
            this.tbTextPath.TabIndex = 19;
            // 
            // bAddText
            // 
            this.bAddText.Location = new System.Drawing.Point(89, 8);
            this.bAddText.Name = "bAddText";
            this.bAddText.Size = new System.Drawing.Size(37, 23);
            this.bAddText.TabIndex = 18;
            this.bAddText.Text = "+";
            this.bAddText.UseVisualStyleBackColor = true;
            this.bAddText.Click += new System.EventHandler(this.bAddText_Click);
            // 
            // label6
            // 
            this.label6.AutoSize = true;
            this.label6.Location = new System.Drawing.Point(132, 250);
            this.label6.Name = "label6";
            this.label6.Size = new System.Drawing.Size(53, 12);
            this.label6.TabIndex = 17;
            this.label6.Text = "冲突处理";
            // 
            // label7
            // 
            this.label7.AutoSize = true;
            this.label7.Location = new System.Drawing.Point(132, 221);
            this.label7.Name = "label7";
            this.label7.Size = new System.Drawing.Size(53, 12);
            this.label7.TabIndex = 14;
            this.label7.Text = "本地路径";
            // 
            // label8
            // 
            this.label8.AutoSize = true;
            this.label8.Location = new System.Drawing.Point(132, 38);
            this.label8.Name = "label8";
            this.label8.Size = new System.Drawing.Size(29, 12);
            this.label8.TabIndex = 13;
            this.label8.Text = "内容";
            // 
            // tbTextContent
            // 
            this.tbTextContent.Location = new System.Drawing.Point(191, 35);
            this.tbTextContent.Multiline = true;
            this.tbTextContent.Name = "tbTextContent";
            this.tbTextContent.Size = new System.Drawing.Size(230, 177);
            this.tbTextContent.TabIndex = 10;
            // 
            // label9
            // 
            this.label9.AutoSize = true;
            this.label9.Location = new System.Drawing.Point(6, 13);
            this.label9.Name = "label9";
            this.label9.Size = new System.Drawing.Size(77, 12);
            this.label9.TabIndex = 9;
            this.label9.Text = "文本文件列表";
            // 
            // lbText
            // 
            this.lbText.FormattingEnabled = true;
            this.lbText.ItemHeight = 12;
            this.lbText.Location = new System.Drawing.Point(6, 38);
            this.lbText.Name = "lbText";
            this.lbText.Size = new System.Drawing.Size(120, 256);
            this.lbText.TabIndex = 8;
            // 
            // tbpCustom
            // 
            this.tbpCustom.Controls.Add(this.tbCustomParam);
            this.tbpCustom.Controls.Add(this.label1);
            this.tbpCustom.Controls.Add(this.tbCustomInput);
            this.tbpCustom.Location = new System.Drawing.Point(4, 22);
            this.tbpCustom.Name = "tbpCustom";
            this.tbpCustom.Size = new System.Drawing.Size(433, 300);
            this.tbpCustom.TabIndex = 2;
            this.tbpCustom.Text = "自定义文件";
            this.tbpCustom.UseVisualStyleBackColor = true;
            // 
            // tbCustomParam
            // 
            this.tbCustomParam.Location = new System.Drawing.Point(98, 10);
            this.tbCustomParam.Name = "tbCustomParam";
            this.tbCustomParam.Size = new System.Drawing.Size(332, 21);
            this.tbCustomParam.TabIndex = 2;
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(3, 13);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(89, 12);
            this.label1.TabIndex = 1;
            this.label1.Text = "Aria2 额外参数";
            // 
            // tbCustomInput
            // 
            this.tbCustomInput.Location = new System.Drawing.Point(3, 37);
            this.tbCustomInput.Multiline = true;
            this.tbCustomInput.Name = "tbCustomInput";
            this.tbCustomInput.Size = new System.Drawing.Size(427, 260);
            this.tbCustomInput.TabIndex = 0;
            // 
            // label4
            // 
            this.label4.AutoSize = true;
            this.label4.Location = new System.Drawing.Point(6, 86);
            this.label4.Name = "label4";
            this.label4.Size = new System.Drawing.Size(53, 12);
            this.label4.TabIndex = 1;
            this.label4.Text = "镜像编号";
            // 
            // tbSysIndex
            // 
            this.tbSysIndex.Location = new System.Drawing.Point(65, 83);
            this.tbSysIndex.Name = "tbSysIndex";
            this.tbSysIndex.Size = new System.Drawing.Size(100, 21);
            this.tbSysIndex.TabIndex = 2;
            // 
            // FileControl
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.Controls.Add(this.tabControl1);
            this.Name = "FileControl";
            this.Size = new System.Drawing.Size(447, 332);
            this.tabControl1.ResumeLayout(false);
            this.tbpSys.ResumeLayout(false);
            this.tbpSys.PerformLayout();
            this.tbpNormal.ResumeLayout(false);
            this.tbpNormal.PerformLayout();
            this.tbpText.ResumeLayout(false);
            this.tbpText.PerformLayout();
            this.tbpCustom.ResumeLayout(false);
            this.tbpCustom.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.TabControl tabControl1;
        private System.Windows.Forms.TabPage tbpNormal;
        private System.Windows.Forms.TabPage tbpText;
        private System.Windows.Forms.TabPage tbpCustom;
        private System.Windows.Forms.TextBox tbCustomInput;
        private System.Windows.Forms.Label label5;
        private System.Windows.Forms.Label label3;
        private System.Windows.Forms.TextBox tbNormalPath;
        private System.Windows.Forms.TextBox tbNormalUrl;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.ListBox lbNormal;
        private System.Windows.Forms.TextBox tbCustomParam;
        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.Label label6;
        private System.Windows.Forms.Label label7;
        private System.Windows.Forms.Label label8;
        private System.Windows.Forms.TextBox tbTextContent;
        private System.Windows.Forms.Label label9;
        private System.Windows.Forms.ListBox lbText;
        private System.Windows.Forms.Button bAddNormal;
        private System.Windows.Forms.Button bAddText;
        private System.Windows.Forms.Label label10;
        private System.Windows.Forms.Label label11;
        private System.Windows.Forms.TextBox tbTextPath;
        private System.Windows.Forms.Button bDelNormal;
        private System.Windows.Forms.Button bDelText;
        private System.Windows.Forms.CheckBox cbTextAppend;
        private System.Windows.Forms.TabPage tbpSys;
        private System.Windows.Forms.TextBox tbSysUrl;
        private System.Windows.Forms.Label label12;
        private System.Windows.Forms.Label label4;
        private System.Windows.Forms.TextBox tbSysIndex;
    }
}
