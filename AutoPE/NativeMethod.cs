using System.Runtime.InteropServices;
using System.Text;

namespace AutoPE
{
    internal class NativeMethod
    {
        [DllImport("kernel32.dll", SetLastError = true)]
        static extern uint QueryDosDevice(string lpDeviceName, StringBuilder lpTargetPath, int ucchMax);

        public static string QueryDosDevice(string driveLetter)
        {
            StringBuilder realPath = new StringBuilder(256);
            QueryDosDevice(driveLetter, realPath, 256);
            return realPath.ToString();
        }
    }
}
