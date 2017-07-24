package mp

import "fmt"

//定义了一个Player类
//这里没有直接将MusicEntry做完参数传入，因为其包含了一些多余的信息，播放
//音乐只需要知道source和type这两个就行了，即音乐的位置和音乐的类型
type Player interface {
	Play(source string)
}

//Player类的方法
func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		//指定MP3的方法
		p = &MP3Player{}
	case "WAV":
		//指定WAV的方法
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
	}
	p.Play(source)
}
