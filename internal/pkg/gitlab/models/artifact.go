package models

type Artifact struct {
	FileType   string `json:"file_type"`
	Size       int64  `json:"size"`
	Filename   string `json:"filename"`
	FileFormat string `json:"file_format"`
}
