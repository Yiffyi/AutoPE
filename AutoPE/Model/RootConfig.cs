﻿namespace AutoPE.Model
{
    internal class RootConfig
    {
        public FileConfig Inject { get; set; }
        public PEConfig PE { get; set; }
        public NICConfig NIC { get; set; }
        // public VolumeConfig Volume { get; set; }
    }
}
