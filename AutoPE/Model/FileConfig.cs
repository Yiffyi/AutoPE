using IniParser.Model;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.IO;
using System.Text;

namespace AutoPE.Model
{
    public class FileConfig
    {
        public string SystemImgUrl { get; set; }
        public int SystemImgIndex { get; set; } = 1;

        public class NormalFile
        {
            public string Url { get; set; }
            public string Path { get; set; }
        }

        public List<NormalFile> NormalFiles;

        public class TextFile
        {
            public string Content { get; set; }
            public string Path { get; set; }
            public bool Append { get; set; }
        }

        public List<TextFile> TextFiles;

        public string CustomInput { get; set; }

        public string CompileSysImgAria2Input()
        {
            StringBuilder b = new StringBuilder();
            b.AppendLine(SystemImgUrl);
            b.AppendLine("  out=System.esd");

            return b.ToString();
        }

        public IniData CompileSysImgIni()
        {
            IniData ini = new IniData();
            ini["SystemImg"]["ImgIndex"] = SystemImgIndex.ToString();
            return ini;
        }

        public string CompileInjectAria2Input()
        {
            StringBuilder b = new StringBuilder();
            foreach (NormalFile f in NormalFiles)
            {
                b.AppendLine(f.Url);
                b.AppendLine($"  out={f.Path}");
            }

            return b.ToString();
        }

        public IniData PrepareTextFiles(string folder, Encoding e)
        {
            IniData ini = new IniData();
            for (int i = 0; i < TextFiles.Count; i++)
            {
                var f = TextFiles[i];
                File.WriteAllText(Path.Combine(folder, $"{i}.txt"), f.Content, e);
                ini[$"Text{i}"]["SrcPath"] = $"{i}.txt";
                ini[$"Text{i}"]["DstPath"] = f.Path;
                ini[$"Text{i}"]["Append"] = f.Append.ToString();
            }
            return ini;
        }
    }
}
