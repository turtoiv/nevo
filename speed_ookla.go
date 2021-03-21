package main;

import (
	"fmt"
	"errors"
	"github.com/kylegrantlucas/speedtest"
	"github.com/ddo/go-fast"
)

type Provider interface {
	getLinkSpeed() error
	getDownloadSpeed() float64
	getUploadSpeed() float64
}

type ProviderType int
const (
	Ookla = iota
	Netflix
)

type OoklaProvider struct {
	downloadSpeed float64
	uploadSpeed float64
}

type NetflixProvider struct {
	downloadSpeed float64
	uploadSpeed float64
}

/* This function is exported from the library
** it provides the functionality to check the speed test
* Param: providerType - which provider will be used to test the speed 
*/
func StartSpeedTest(providerType ProviderType) error {
	var provider Provider
	if providerType != Ookla && providerType != Netflix {
		fmt.Printf("unkown provider")
		return errors.New("Unknown provider")
	}

	if providerType == Ookla {
		provider = &OoklaProvider{}
	}
	if providerType == Netflix {
		provider = &NetflixProvider{}
	}

	err := provider.getLinkSpeed()
	if err == nil {
		fmt.Printf("Download: %3.2f Mbps | Upload: %3.2f Mbps\n", provider.getDownloadSpeed(), provider.getUploadSpeed())
	}

	return err
}

/* Implementation to get download and upload speed 
** using Ookla speedtest
*/
func (oklProvider *OoklaProvider) getLinkSpeed() error{
	client, err := speedtest.NewDefaultClient()

	if err != nil {
		fmt.Printf("error creating client: %v", err)
		return err
	}

	server, err := client.GetServer("")
	if err != nil {
		fmt.Printf("error getting server: %v", err)
		return err
	}

	umbps, err := client.Upload(server)
	if err != nil {
		fmt.Printf("error getting upload: %v", err)
		return err
	}

	dmbps, err := client.Download(server)
	if err != nil {
		fmt.Printf("error getting download: %v", err)
		return err
	}

	oklProvider.downloadSpeed = dmbps
	oklProvider.uploadSpeed = umbps
	return nil
}

/* implementation to get download speed using fast.com
** upload speed is not provided by Measure function so this function
** will compute only the download speed
** the measurement is an average with values obtained from a number of servers
*/
func (nflProvider *NetflixProvider) getLinkSpeed() error {
	fastCom := fast.New()
	err := fastCom.Init()

	if err != nil {
		fmt.Printf("error initializing library")
		return err
	}

	urls, err := fastCom.GetUrls()

	if err != nil {
		fmt.Printf("error getting test servers")
		return err
	}

	KbpsChan := make(chan float64)
	count := 0
	average := 0.0

	go func() {
		for Kbps := range KbpsChan {
			//fmt.Printf("%.2f Kbps %.2f Mbps\n", Kbps, Kbps/1000)
			average += float64(Kbps)/float64(1000);
			count++;
		}
	}()
	err = fastCom.Measure(urls, KbpsChan)
	nflProvider.downloadSpeed = float64(average)/float64(count);

	return nil
}

func (oklProvider *OoklaProvider) getDownloadSpeed() float64 { 
	return oklProvider.downloadSpeed
}

func (oklProvider *OoklaProvider) getUploadSpeed() float64 {
	return oklProvider.uploadSpeed
}

func (nflProvider *NetflixProvider) getDownloadSpeed() float64 {
	return nflProvider.downloadSpeed
}
func (nflProvider *NetflixProvider) getUploadSpeed() float64 {
	return nflProvider.uploadSpeed
}