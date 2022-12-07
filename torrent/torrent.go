package torrent

import (
	"errors"
	"io"
	"strings"

	dgoBencode "github.com/Dizzrt/dgo-bencode"
)

var (
	ErrData = errors.New("data error")
)

type TorrentInfo struct {
	Name        string
	Hash        string
	Pieces      string
	Length      int
	PieceLength int
	Files       []map[string]any
}

type Torrent struct {
	Announce     string
	AnnounceList [][]string
	CreationDate int
	CreatedBy    string
	Comment      string
	Info         TorrentInfo
	IsMutiFiles  bool
}

func (this *Torrent) Init(data map[string]any) *Torrent {
	_announce, _ := data["announce"].(string)
	this.Announce = _announce

	_announceList, _ := data["announce-list"].([][]string)
	this.AnnounceList = _announceList

	_creationDate, _ := data["creation date"].(int)
	this.CreationDate = _creationDate

	_creationBy, _ := data["created by"].(string)
	this.CreatedBy = _creationBy

	_comment, _ := data["comment"].(string)
	this.Comment = _comment

	info, ok := data["info"].(map[string]any)
	if ok {
		_name, _ := info["name"].(string)
		this.Info.Name = _name

		_pieces, _ := info["pieces"].(string)
		this.Info.Pieces = _pieces

		_length, _ := info["length"].(int)
		this.Info.Length = _length

		_pieceLength, _ := info["piece length"].(int)
		this.Info.PieceLength = _pieceLength

		_hash, _ := info["filehash"].(string)
		this.Info.Hash = _hash

		_files, ok := info["files"].([]map[string]any)
		if ok {
			this.IsMutiFiles = true
			this.Info.Files = _files
		}
	}

	return this
}

func New(data map[string]any) *Torrent {
	return new(Torrent).Init(data)
}

func NewWithReader(r io.Reader) (bt *Torrent, err error) {
	rawData, err := dgoBencode.BencodeDecodeFromReader(r)
	if err != nil {
		return nil, err
	}

	data, ok := rawData[0].(map[string]any)
	if ok {
		bt = New(data)
		return
	} else {
		return nil, ErrData
	}
}

func NewWithString(str string) (bt *Torrent, err error) {
	return NewWithReader(strings.NewReader(str))
}
