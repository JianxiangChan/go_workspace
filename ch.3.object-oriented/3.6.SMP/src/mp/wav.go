//一个包拆分开来写
package mp

import (
	"fmt"
	"time"
)

//定义一个MP3Player类
type WAVPlayer struct {
	stat    int
	process int
}

//实现WAVPlayer方法
func (p *WAVPlayer) Play(source string) {
	fmt.Println("Play WAV music", source)

	p.process = 0

	for p.process < 100 {
		time.Sleep(100 * time.Millisecond) //假装正在播放
		fmt.Println(".")
		p.process += 10
	}

	fmt.Println("\nFinished playing", source)
}
