package main

import "github.com/go-ole/go-ole"
import "github.com/moutend/go-wca"

// default audio device
func mmDevice() (*wca.IMMDevice, error) {
	var enumer *wca.IMMDeviceEnumerator
	if err := wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &enumer); err != nil {
		panic(err)
	}
	defer enumer.Release()

	var device *wca.IMMDevice
	return device, enumer.GetDefaultAudioEndpoint(wca.ERender, wca.EMultimedia, &device)
}

func ssEnumer() (*wca.IAudioSessionEnumerator, error) {
	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		panic(err)
	}
	defer ole.CoUninitialize()

	device, err := mmDevice()
	if err != nil {
		panic(err)
	}
	defer device.Release()

	var manager *wca.IAudioSessionManager2
	if err = device.Activate(wca.IID_IAudioSessionManager2, wca.CLSCTX_ALL, nil, &manager); err != nil {
		panic(err)
	}
	defer manager.Release()

	var enumer *wca.IAudioSessionEnumerator
	return enumer, manager.GetSessionEnumerator(&enumer)
}

func epVolume() (*wca.IAudioEndpointVolume, error) {
	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		panic(err)
	}
	defer ole.CoUninitialize()

	device, err := mmDevice()
	if err != nil {
		panic(err)
	}
	defer device.Release()

	var volume *wca.IAudioEndpointVolume
	return volume, device.Activate(wca.IID_IAudioEndpointVolume, wca.CLSCTX_ALL, nil, &volume)
}
