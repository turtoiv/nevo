# nevo

This project makes use of Ookla speedtest functionality and Netflix fast.com functionality. It provides both upload and download speed values for Ookla but only download speed for Netflix.

Two libraries were used, one for speedtest and one for fast.com:<br>
1)Library used for speedtest: go get https://github.com/kylegrantlucas/speedtest<br>
2)Library used for fast.com: go get  -u github.com/ddo/fast

The library exports one function, StartSpeedTest which receives as parameter the type of the provider for the speed calculation. For now only two providers are supported: Ookla and Netflix. Any other value will not be accepted and the function will fail.

Usage:

func main() { <br>
  StartSpeedTest(Ookla) <br>
} <br>

There are also some tests added, to run them go test -v <br>

To run the application go run speed_ookla.go <br>
