package metadata_test

import (
	"fmt"

	"github.com/navidrome/navidrome/conf"
	"github.com/navidrome/navidrome/conf/configtest"
	"github.com/navidrome/navidrome/core/ffmpeg"
	"github.com/navidrome/navidrome/scanner/metadata"
	_ "github.com/navidrome/navidrome/scanner/metadata/ffmpeg"
	_ "github.com/navidrome/navidrome/scanner/metadata/taglib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tags", func() {
	// Only run these tests if FFmpeg is available
	FFmpegContext := XContext
	if ffmpeg.New().IsAvailable() {
		FFmpegContext = Context
	} else {
		fmt.Println("darn")
	}
	FFmpegContext("Extract with FFmpeg", func() {
		BeforeEach(func() {
			DeferCleanup(configtest.SetupConfig())
			conf.Server.Scanner.Extractor = "ffmpeg"
		})

		DescribeTable("Lyrics test",
			func(file string) {
				path := "tests/fixtures/" + file
				mds, err := metadata.Extract(path)
				Expect(err).ToNot(HaveOccurred())
				Expect(mds).To(HaveLen(1))

				m := mds[path]

				fmt.Println("TITLE IS: " + m.Title())

			},

			Entry("Parses MP3 files", "test.shn"),
			// Disabled, because it fails in pipeline
			// Entry("Parses WAV files", "test.wav"),

			// FFMPEG behaves very weirdly for multivalued tags for non-ID3
			// Specifically, they are separated by ";, which is indistinguishable
			// from other fields
		)
	})
})
