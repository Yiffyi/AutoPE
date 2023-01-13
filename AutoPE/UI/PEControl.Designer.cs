namespace AutoPE.UI
{
    partial class PEControl
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
            this.tbWim = new System.Windows.Forms.TextBox();
            this.tbSdi = new System.Windows.Forms.TextBox();
            this.tbWorker = new System.Windows.Forms.TextBox();
            this.label1 = new System.Windows.Forms.Label();
            this.label2 = new System.Windows.Forms.Label();
            this.label3 = new System.Windows.Forms.Label();
            this.SuspendLayout();
            // 
            // tbWim
            // 
            this.tbWim.Location = new System.Drawing.Point(188, 18);
            this.tbWim.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.tbWim.Name = "tbWim";
            this.tbWim.Size = new System.Drawing.Size(584, 35);
            this.tbWim.TabIndex = 0;
            // 
            // tbSdi
            // 
            this.tbSdi.Location = new System.Drawing.Point(188, 72);
            this.tbSdi.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.tbSdi.Name = "tbSdi";
            this.tbSdi.Size = new System.Drawing.Size(584, 35);
            this.tbSdi.TabIndex = 1;
            // 
            // tbWorker
            // 
            this.tbWorker.Location = new System.Drawing.Point(188, 126);
            this.tbWorker.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.tbWorker.Name = "tbWorker";
            this.tbWorker.Size = new System.Drawing.Size(584, 35);
            this.tbWorker.TabIndex = 2;
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(34, 24);
            this.label1.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(142, 24);
            this.label1.TabIndex = 3;
            this.label1.Text = "PE 镜像 URL";
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(46, 78);
            this.label2.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(130, 24);
            this.label2.TabIndex = 4;
            this.label2.Text = "PE Sdi URL";
            // 
            // label3
            // 
            this.label3.AutoSize = true;
            this.label3.Location = new System.Drawing.Point(10, 132);
            this.label3.Margin = new System.Windows.Forms.Padding(6, 0, 6, 0);
            this.label3.Name = "label3";
            this.label3.Size = new System.Drawing.Size(166, 24);
            this.label3.TabIndex = 5;
            this.label3.Text = "PE Worker URL";
            // 
            // PEControl
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(12F, 24F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.Controls.Add(this.label3);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.label1);
            this.Controls.Add(this.tbWorker);
            this.Controls.Add(this.tbSdi);
            this.Controls.Add(this.tbWim);
            this.Margin = new System.Windows.Forms.Padding(6, 6, 6, 6);
            this.Name = "PEControl";
            this.Size = new System.Drawing.Size(782, 176);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.TextBox tbWim;
        private System.Windows.Forms.TextBox tbSdi;
        private System.Windows.Forms.TextBox tbWorker;
        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.Label label3;
    }
}
