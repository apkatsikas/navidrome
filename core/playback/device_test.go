package playback

import (
	"context"
	"fmt"
	"testing"

	"github.com/dexterlb/mpvipc"
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

type mockMpvipc struct {
	mock.Mock
}

// Open simulates opening a connection
func (m *mockMpvipc) Open() error {
	return nil
}

// ListenForEvents simulates listening for events
func (m *mockMpvipc) ListenForEvents(events chan<- *mpvipc.Event, stop <-chan struct{}) {

}

// NewEventListener simulates creating a new event listener
func (m *mockMpvipc) NewEventListener() (chan *mpvipc.Event, chan struct{}) {
	return nil, nil
}

// Call simulates calling an arbitrary command
func (m *mockMpvipc) Call(arguments ...interface{}) (interface{}, error) {
	// Implement your mock behavior for command execution
	return nil, nil
}

// Set simulates setting a property
func (m *mockMpvipc) Set(property string, value interface{}) error {
	// Implement your mock behavior for setting a property
	return nil
}

// Get simulates getting a property
func (m *mockMpvipc) Get(property string) (interface{}, error) {
	m.Called()
	// Implement your mock behavior for getting a property
	return nil, nil
}

// Close simulates closing the connection
func (m *mockMpvipc) Close() error {
	return nil
}

// IsClosed simulates checking if the connection is closed
func (m *mockMpvipc) IsClosed() bool {
	return false
}

// WaitUntilClosed simulates waiting until the connection is closed
func (m *mockMpvipc) WaitUntilClosed() {

}

func TestThingz(t *testing.T) {
	// SHould I be using log here?
	ctx := log.NewContext(context.Background())
	ps := &mockPlaybackServer{}
	mockMpv := &mockMpvipc{}
	mockMpv.On("Get").Return(true, nil)

	mediaIds := []string{"1234", "5678", "90"}
	device := NewPlaybackDevice(
		&playbackServer{},
		"name",
		"device name",
		func(playbackDoneChannel chan bool, deviceName string, mf model.MediaFile) (*mpv.MpvTrack, error) {
			return &mpv.MpvTrack{
				Conn: mockMpv,
			}, nil
		},
	)
	device.ParentPlaybackServer = ps

	setStatus, err := device.Set(ctx, mediaIds)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(setStatus.Playing)

	startStatus, err := device.Start(ctx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(startStatus.Playing)

	fmt.Println(mockMpv.AssertExpectations(t))
}
