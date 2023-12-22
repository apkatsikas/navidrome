package playback

import (
	"context"
	"fmt"
	"testing"

	"github.com/navidrome/navidrome/core/playback/mpv"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/stretchr/testify/mock"
)

// import (
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// )

// var _ = BeforeSuite(func() {

// })

// var _ = Describe("Auth", func() {

// 	BeforeEach(func() {

// 	})

// 	Describe("do", func() {
// 		It("does", func() {

// 			Expect(err).To(HaveOccurred())
// 		})
// 	})
// })

type mockPlaybackServer struct {
	mock.Mock
	PlaybackServer
}

func (m *mockPlaybackServer) GetMediaFile(id string) (*model.MediaFile, error) {
	return &model.MediaFile{}, nil
}

func TestThingz(t *testing.T) {
	// SHould I be using log here?
	ctx := log.NewContext(context.Background())
	ps := &mockPlaybackServer{}

	mediaIds := []string{"1234", "5678", "90"}
	device := NewPlaybackDevice(
		&playbackServer{},
		"name",
		"device name",
		func(playbackDoneChannel chan bool, deviceName string, mf model.MediaFile) (*mpv.MpvTrack, error) {
			return &mpv.MpvTrack{
				//
			}, nil
		},
	)
	device.ParentPlaybackServer = ps

	setStatus, err := device.Set(ctx, mediaIds)
	if err != nil {
		t.Error(err)
	}
	startStatus, err := device.Start(ctx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(setStatus)
	fmt.Println(startStatus)
	fmt.Println(ps.Calls)
}
