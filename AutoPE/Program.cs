﻿using System;
using System.Windows.Forms;

namespace AutoPE
{
    static class Program
    {
        public static int BUILD = 4;
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>
        [STAThread]
        static void Main()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            Application.Run(new FormMain());
        }
    }
}
