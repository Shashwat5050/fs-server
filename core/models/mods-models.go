package models

import "time"

type CfGame struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Logo struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

type Screenshot struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

type Hash struct {
	Value string `json:"value"`
	Algo  int    `json:"algo"`
}

type GameVersion struct {
	GameVersionName        string    `json:"gameVersionName"`
	GameVersionPadded      string    `json:"gameVersionPadded"`
	GameVersion            string    `json:"gameVersion"`
	GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
	GameVersionTypeID      int       `json:"gameVersionTypeId"`
}

type Module struct {
	Name string `json:"name"`
}

type File struct {
	ID                   int           `json:"id"`
	GameID               int           `json:"gameId"`
	DisplayName          string        `json:"displayName"`
	FileName             string        `json:"fileName"`
	Hashes               []Hash        `json:"hashes"`
	FileDate             time.Time     `json:"fileDate"`
	FileLength           int           `json:"fileLength"`
	DownloadURL          string        `json:"downloadUrl"`
	GameVersions         []string      `json:"gameVersions"`
	SortableGameVersions []GameVersion `json:"sortableGameVersions"`
	Dependencies         []any         `json:"dependencies"`
	AlternateFileID      int           `json:"alternateFileId"`
	IsServerPack         bool          `json:"isServerPack"`
	ServerPackFileID     int           `json:"serverPackFileId"`
	Modules              []Module      `json:"modules"`
}

type FileIndex struct {
	GameVersion       string `json:"gameVersion"`
	FileID            int    `json:"fileId"`
	Filename          string `json:"filename"`
	ReleaseType       int    `json:"releaseType"`
	GameVersionTypeID int    `json:"gameVersionTypeId"`
	ModLoader         int    `json:"modLoader,omitempty"`
}

type Mod struct {
	ID                            int          `json:"id"`
	GameID                        int          `json:"gameId"`
	Name                          string       `json:"name"`
	Summary                       string       `json:"summary"`
	Logo                          Logo         `json:"logo"`
	Screenshots                   []Screenshot `json:"screenshots"`
	MainFileID                    int          `json:"mainFileId"`
	LatestFiles                   []File       `json:"latestFiles"`
	LatestFilesIndexes            []FileIndex  `json:"latestFilesIndexes"`
	LatestEarlyAccessFilesIndexes []any        `json:"latestEarlyAccessFilesIndexes"`
	DateCreated                   time.Time    `json:"dateCreated"`
	DateModified                  time.Time    `json:"dateModified"`
	DateReleased                  time.Time    `json:"dateReleased"`
}
