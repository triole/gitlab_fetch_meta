package main

type tReposMeta []tRepoMeta

type tRepoMeta struct {
	ID            int    `json:"id"`
	Description   string `json:"description"`
	ReadMeURL     string `json:"readme_url"`
	MetaFilenames []string
	MetaData      tMetadata
}

type tMetadata map[string]interface{}

func (rm *tRepoMeta) initMeta() {
	for _, el := range rm.MetaFilenames {
		println(el)
	}
}
