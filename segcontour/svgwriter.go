package segcontour

import (
	"github.com/aphecetche/pigiron/geo"
	"github.com/aphecetche/pigiron/mapping"
)

// ShowFlags is a option struct to select what to show / hide in SVG outputs
type ShowFlags struct {
	des         bool
	dualsampas  bool
	pads        bool
	padchannels bool
}

func svgDualSampaPads(w *geo.SVGWriter, dualSampaPads *[][]geo.Polygon) {
	w.GroupStart("pads")
	for _, dsp := range *dualSampaPads {
		for _, p := range dsp {
			w.Polygon(&p)
		}
	}
	w.GroupEnd()
}

func svgDetectionElements(w *geo.SVGWriter, de *geo.Contour) {
	w.GroupStart("detectionelements")
	w.Contour(de)
	w.GroupEnd()
}

// SVGSegmentation creates a SVG representation of segmentation
func SVGSegmentation(seg *mapping.Segmentation, w *geo.SVGWriter, show ShowFlags) {

	dualSampaPads := getPadPolygons(seg)
	deContour := GetSegmentationEnvelop(seg)

	if show.des {
		svgDetectionElements(w, &deContour)
	}
	if show.pads {
		svgDualSampaPads(w, &dualSampaPads)
	}
}