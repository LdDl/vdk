package format

import (
	"github.com/LdDl/vdk/av/avutil"
	"github.com/LdDl/vdk/format/aac"
	"github.com/LdDl/vdk/format/flv"
	"github.com/LdDl/vdk/format/mp4"
	"github.com/LdDl/vdk/format/rtmp"
	"github.com/LdDl/vdk/format/rtsp"
	"github.com/LdDl/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
