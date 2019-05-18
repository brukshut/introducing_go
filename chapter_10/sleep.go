package main

/*
func After(d Duration) <-chan Time
    After waits for the duration to elapse and then sends the current time on
    the returned channel. It is equivalent to NewTimer(d).C. The underlying
    Timer is not recovered by the garbage collector until the timer fires. If
    efficiency is a concern, use NewTimer instead and call Timer.Stop if the
    timer is no longer needed.
*/

import ("fmt"; "time")

func sleep(sec time.Duration) time.Time {
	return <-time.After(time.Second * sec)
}

func main() {
  msg := sleep(3)
	fmt.Println(msg)

  msg = sleep(5)
	fmt.Println(msg)

  msg = sleep(1)
	fmt.Println(msg)

  // unbuffered channels are synchronous
  // buffered channel with size 1 is synchronous?
  chnl := make(chan time.Time, 1)
  chnl <- sleep(3)
}

