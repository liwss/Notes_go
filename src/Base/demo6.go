package Base

import "fmt"

type MusicEntry struct {
	Id 	string			//音乐的唯一Id
	Name 	string		//音乐名
	Artist	string		//艺术家名
	Source	string		//音乐位置
	Type 	string		//音乐类型（MP3和WAV）
}

type MusicManager struct {
	musics	[]MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 5)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func main() {
	a :=NewMusicManager()
	fmt.Println(a.Len())
}


