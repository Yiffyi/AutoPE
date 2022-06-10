namespace AutoPE.Model
{
    public class PEConfig
    {
        public string OsLoaderId { get; set; }
        public string RamdiskId { get; set; }
        public string Description { get; set; }
        public bool Display { get; set; }
        public string WimUrl { get; set; }
        public string SdiUrl { get; set; }
        public string WorkerUrl { get; set; }
    }
}