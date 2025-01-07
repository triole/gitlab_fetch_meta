package main

type tReposMeta []tRepoMeta

type tRepoMeta struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	ReadMeURL   string `json:"readme_url"`
}
