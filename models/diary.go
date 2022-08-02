package models

type Diary struct {
	UrlArchive string `db:"url_archive"`
	Cod        int64  `db:"cod"`
	Processed  bool   `db:"processed"`
}
