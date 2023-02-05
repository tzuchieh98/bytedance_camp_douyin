package config

type Video struct {
	VideoUploadPath string `mapstructure:"video_upload_path" json:"video_upload_path" yaml:"video_upload_path"`
	CoverUploadPath string `mapstructure:"cover_upload_path" json:"cover_upload_path" yaml:"cover_upload_path"`
	VideoPlayURL    string `mapstructure:"video_play_url" json:"video_play_url" yaml:"video_play_url"`
	CoverShowURL    string `mapstructure:"cover_show_url" json:"cover_show_url" yaml:"cover_show_url"`
}
