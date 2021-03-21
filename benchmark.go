package main

import (
	"fmt"
)

/* Benchmark test for ookla and netflix
** 10 runs for each provider and the average values are computed and displayed
*/
func main() {
	result := computeBenchmark(Ookla, 10)
	fmt.Printf("Ookla benchmark : Download: %3.2f Mbps | Upload: %3.2f Mbps\n", result[0], result[1])

	result = computeBenchmark(Netflix, 10)
	fmt.Printf("Netflix benchmark : Download: %3.2f Mbps | Upload: %3.2f Mbps\n", result[0], result[1])
}

func computeBenchmark(providerType ProviderType, count int) [2]float64 {
	var result [2]float64
	var provider Provider

	if providerType == Ookla {
		provider = &OoklaProvider{}
	}
	if providerType == Netflix {
		provider = &NetflixProvider{}
	}

	uploadAverage := 0.0
	downloadAverage := 0.0
	for i := 0; i < count; i++  {
		err := provider.getLinkSpeed() 
		if err == nil {
			downloadAverage += provider.getDownloadSpeed()
			uploadAverage += provider.getUploadSpeed()
		}
	}

	downloadAverage = float64(downloadAverage)/float64(count)
	uploadAverage = float64(uploadAverage)/float64(count)
	result[0] = downloadAverage
	result[1] = uploadAverage

	return result
}
