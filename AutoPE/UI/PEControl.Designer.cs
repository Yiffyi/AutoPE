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
            this.tbWim.Location = new System.Drawing.Point(94, 54);
            this.tbWim.Name = "tbWim";
            this.tbWim.Size = new System.Drawing.Size(294, 21);
            this.tbWim.TabIndex = 0;
            // 
            // tbSdi
            // 
            this.tbSdi.Location = new System.Drawing.Point(94, 81);
            this.tbSdi.Name = "tbSdi";
            this.tbSdi.Size = new System.Drawing.Size(294, 21);
            this.tbSdi.TabIndex = 1;
            // 
            // tbWorker
            // 
            this.tbWorker.Location = new System.Drawing.Point(94, 108);
            this.tbWorker.Name = "tbWorker";
            this.tbWorker.Size = new System.Drawing.Size(294, 21);
            this.tbWorker.TabIndex = 2;
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(17, 57);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(71, 12);
            this.label1.TabIndex = 3;
            this.label1.Text = "PE 镜像 URL";
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(23, 84);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(65, 12);
            this.label2.TabIndex = 4;
            this.label2.Text = "PE Sdi URL";
            // 
            // label3
            // 
            this.label3.AutoSize = true;
            this.label3.Location = new System.Drawing.Point(5, 111);
            this.label3.Name = "label3";
            this.label3.Size = new System.Drawing.Size(83, 12);
            this.label3.TabIndex = 5;
            this.label3.Text = "PE Worker URL";
            // 
            // PEControl
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.Controls.Add(this.label3);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.label1);
            this.Controls.Add(this.tbWorker);
            this.Controls.Add(this.tbSdi);
            this.Controls.Add(this.tbWim);
            this.Name = "PEControl";
            this.Size = new System.Drawing.Size(391, 313);
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
