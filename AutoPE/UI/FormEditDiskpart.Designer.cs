namespace AutoPE.UI
{
    partial class FormEditDiskpart
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
            this.label1 = new System.Windows.Forms.Label();
            this.tbScript = new System.Windows.Forms.TextBox();
            this.bOK = new System.Windows.Forms.Button();
            this.SuspendLayout();
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(12, 9);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(413, 12);
            this.label1.TabIndex = 0;
            this.label1.Text = "请将系统区盘符分配为 I，数据区盘符分配为 J，引导区（如果有）分配为 K";
            // 
            // tbScript
            // 
            this.tbScript.Location = new System.Drawing.Point(12, 43);
            this.tbScript.Multiline = true;
            this.tbScript.Name = "tbScript";
            this.tbScript.Size = new System.Drawing.Size(413, 310);
            this.tbScript.TabIndex = 1;
            // 
            // bOK
            // 
            this.bOK.Location = new System.Drawing.Point(350, 359);
            this.bOK.Name = "bOK";
            this.bOK.Size = new System.Drawing.Size(75, 23);
            this.bOK.TabIndex = 2;
            this.bOK.Text = "确定";
            this.bOK.UseVisualStyleBackColor = true;
            this.bOK.Click += new System.EventHandler(this.bOK_Click);
            // 
            // FormEditDiskpart
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(437, 391);
            this.Controls.Add(this.bOK);
            this.Controls.Add(this.tbScript);
            this.Controls.Add(this.label1);
            this.Name = "FormEditDiskpart";
            this.Text = "FormEditDiskpart";
            this.Load += new System.EventHandler(this.FormEditDiskpart_Load);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.TextBox tbScript;
        private System.Windows.Forms.Button bOK;
    }
}