package types1

type FileInfo struct {
	Filepath       string `json:"filepath"`
	Filetype       string `json:"filetype"`
	Lastupdatetime string `json:"lastupdatetime"`
	Filename       string `json:"filename"`
	Filesize       int    `json:"filesize"`
}
