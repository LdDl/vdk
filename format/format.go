package format

import (
	"github.com/morozka/vdk/av/avutil"
	"github.com/morozka/vdk/format/aac"
	"github.com/morozka/vdk/format/flv"
	"github.com/morozka/vdk/format/mp4"
	"github.com/morozka/vdk/format/rtmp"
	"github.com/morozka/vdk/format/rtsp"
	"github.com/morozka/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
