package autope

type PlaybookPEStage struct {
	DiskpartScript string `toml:",multiline" comment:"分区脚本"`
	BootVolume     string `comment:"启动分区"`
	SystemVolume   string `comment:"系统分区"`
	DataVolume     string `comment:"数据分区"`

	FormatSystem bool `comment:"格式化系统分区"`
	FormatData   bool `comment:"格式化数据分区"`
	FormatBoot   bool `comment:"格式化启动分区"`

	ImagePath  string `comment:"镜像路径"`
	ImageIndex int    `comment:"镜像编号"`
}

type Playbook struct {
	PEStage PlaybookPEStage `comment:"PE阶段"`
}
