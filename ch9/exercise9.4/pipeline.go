package pipeline

func pipeline(stages int) (chan<- interface{}, <-chan interface{}) {
	if stages < 1 {
		return nil, nil
	}

	in := make(chan interface{})
	out := in

	for i := 0; i < stages; i++ {
		// 不断新建 channel
		// 再把值从 prev 取出再发送到新建 channel
		prev := out // 前进
		next := make(chan interface{})
		// stages 个 goroutine
		go func() {
			// prev -> next
			// 从 prev 中取值传送给 next
			for v := range prev {
				next <- v
			}
			close(next)
		}()
		// 前进，开始下一轮迭代
		out = next
	}
	// 最后，in 指向流水线中的第一个，out 指向最后一个
	return in, out
}
