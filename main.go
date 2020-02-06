package main

import "time"
import "unsafe"

import "github.com/moutend/go-wca"

func main() {
	for {
		var level float32
		var mute bool
		var err error

		// volume level of other applications
		ssem, err := ssEnumer()
		if err != nil {
			panic(nil)
		}
		var sscount int
		if err = ssem.GetCount(&sscount); err != nil {
			panic(err)
		}
		for i, _ := range make([]interface{}, sscount) {
			var session *wca.IAudioSessionControl
			if err = ssem.GetSession(i, &session); err != nil {
				panic(err)
			}
			dispatch, err := session.QueryInterface(wca.IID_ISimpleAudioVolume)
			session.Release()
			if err != nil {
				panic(err)
			}
			spvolume := (*wca.ISimpleAudioVolume)(unsafe.Pointer(dispatch))

			// level
			if err = spvolume.GetMasterVolume(&level); err != nil {
				panic(err)
			}
			if level != 1.00 {
				if err = spvolume.SetMasterVolume(1.00, nil); err != nil {
					panic(err)
				}
			}

			// mute
			if err = spvolume.GetMute(&mute); err != nil {
				panic(err)
			}
			if mute {
				if err = spvolume.SetMute(false, nil); err != nil {
					panic(err)
				}
			}
			spvolume.Release()
		}
		ssem.Release()

		// volume level of master channel
		epvolume, err := epVolume()
		if err != nil {
			panic(err)
		}

		// level
		if err = epvolume.GetMasterVolumeLevelScalar(&level); err != nil {
			panic(err)
		}
		if level != 1.00 {
			if err = epvolume.SetMasterVolumeLevelScalar(1.00, nil); err != nil {
				panic(err)
			}
		}

		//mute
		if err = epvolume.GetMute(&mute); err != nil {
			panic(err)
		}
		if mute {
			if err = epvolume.SetMute(false, nil); err != nil {
				panic(err)
			}
		}
		epvolume.Release()

		time.Sleep(time.Second)
	}
}
