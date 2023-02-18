package config

type Upload struct {
	PublicHost string `mapstructure:"public_host" json:"public_host" yaml:"public_host"`
	PublicPort string `mapstructure:"public_port" json:"public_port" yaml:"public_port"`
	UploadRoot string `mapstructure:"upload_root" json:"upload_root" yaml:"upload_root"`
	//VideoPlayURL string `mapstructure:"video_play_url" json:"video_play_url" yaml:"video_play_url"`
	//CoverShowURL string `mapstructure:"cover_show_url" json:"cover_show_url" yaml:"cover_show_url"`
}
