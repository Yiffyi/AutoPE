
namespace AutoPE
{
    partial class FormDl
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
            this.lFilename = new System.Windows.Forms.Label();
            this.pbDl = new System.Windows.Forms.ProgressBar();
            this.SuspendLayout();
            // 
            // lFilename
            // 
            this.lFilename.AutoSize = true;
            this.lFilename.Font = new System.Drawing.Font("微软雅黑", 15F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.lFilename.Location = new System.Drawing.Point(12, 54);
            this.lFilename.Name = "lFilename";
            this.lFilename.Size = new System.Drawing.Size(92, 27);
            this.lFilename.TabIndex = 0;
            this.lFilename.Text = "正在下载";
            // 
            // pbDl
            // 
            this.pbDl.Location = new System.Drawing.Point(12, 114);
            this.pbDl.Name = "pbDl";
            this.pbDl.Size = new System.Drawing.Size(395, 47);
            this.pbDl.TabIndex = 1;
            // 
            // FormDl
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(419, 173);
            this.Controls.Add(this.pbDl);
            this.Controls.Add(this.lFilename);
            this.Name = "FormDl";
            this.Text = "FormDl";
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label lFilename;
        private System.Windows.Forms.ProgressBar pbDl;
    }
}