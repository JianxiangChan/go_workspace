//一个包拆分开来写
package mp

import (
	"fmt"
	"time"
)

//定义一个MP3Player类
type MP3Player struct {
	stat    int
	process int
}

//实现MP3Player方法
func (p *MP3Player) Play(source string) {
	fmt.Println("Play MP3 music", source)

	p.process = 0

	for p.process < 100 {
		time.Sleep(100 * time.Millisecond) //假装正在播放
		fmt.Println(".")
		p.process += 10
	}

	fmt.Println("\nFinished playing", source)
}
