package main

import (
	"encoding/xml"
	"strings"
)

const header = "<?xml version=\"1.0\" standalone=\"yes\"?>\n"
const warning = "<!-- warning this file is automatically generated -->\n"

type Chart struct {
	XMLName xml.Name `xml:"chart"`

	Pages []Page
	Plots []Plot
}

type Page struct {
	XMLName xml.Name `xml:"page"`

	Id     string `xml:"id,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
	Png    string `xml:"png,attr,omitempty"`
	Width  string `xml:"width,attr,omitempty"`
	Plots  []Plot `xml:"plot"`

	Border    *Border    `xml:"border,omitempty"`
	Copyright *Copyright `xml:"copyright,omitempty"`
	Date      *Date      `xml:"date,omitempty"`
	Label     *Label     `xml:"label,omitempty"`
	Name      *Name      `xml:"name,omitempty"`
}

type Plot struct {
	XMLName xml.Name `xml:"plot"`

	Id      string `xml:"id,attr,omitempty"`
	Clip    string `xml:"clip,attr,omitempty"`
	Height  string `xml:"height,attr,omitempty"`
	Caption string `xml:"label,attr,omitempty"`
	Length  string `xml:"length,attr,omitempty"`
	Max     string `xml:"max,attr,omitempty"`
	Min     string `xml:"min,attr,omitempty"`
	Missing string `xml:"missing,attr,omitempty"`
	Overlap string `xml:"overlap,attr,omitempty"`
	Pdf     string `xml:"pdf,attr,omitempty"`
	Png     string `xml:"png,attr,omitempty"`
	Width   string `xml:"width,attr,omitempty"`
	X       string `xml:"x,attr,omitempty"`
	Y       string `xml:"y,attr,omitempty"`

	Border    *Border    `xml:",omitempty"`
	Copyright *Copyright `xml:",omitempty"`
	Date      *Date      `xml:",omitempty"`
	Dates     []Date     `xml:",omitempty"`
	Label     *Label     `xml:",omitempty"`
	Name      *Name      `xml:",omitempty"`
	Streams   []Stream   `xml:",omitempty"`
	Title     *Title     `xml:",omitempty"`
	XGrid     *XGrid     `xml:",omitempty"`
	XLabel    *XLabel    `xml:",omitempty"`
	XTick     *XTick     `xml:",omitempty"`
	YTick     *YTick     `xml:",omitempty"`
	YGrid     *YGrid     `xml:",omitempty"`
	YLabel    *YLabel    `xml:",omitempty"`
}

type Border struct {
	XMLName xml.Name `xml:"border"`

	Bg       string `xml:"bg,attr,omitempty"`
	Colour   string `xml:"colour,attr,omitempty"`
	Gradient string `xml:"gradient,attr,omitempty"`
	Pen      string `xml:"pen,attr,omitempty"`
}

type Copyright struct {
	XMLName xml.Name `xml:"copyright"`

	Aligned string `xml:"aligned,attr,omitempty"`
	Colour  string `xml:"colour,attr,omitempty"`
	Font    string `xml:"font,attr,omitempty"`
}

type Date struct {
	XMLName xml.Name `xml:"date"`

	Aligned string `xml:"aligned,attr,omitempty"`
	Colour  string `xml:"colour,attr,omitempty"`
	Font    string `xml:"font,attr,omitempty"`
	String  string `xml:"string,attr,omitempty"`
	Time    string `xml:"time,attr,omitempty"`
}

type Label struct {
	XMLName xml.Name `xml:"label"`

	Colour string `xml:"colour,attr,omitempty"`
	Font   string `xml:"font,attr,omitempty"`
	String string `xml:"string,attr,omitempty"`
}

type Stream struct {
	XMLName xml.Name `xml:"stream"`

	Alt      string `xml:"alt,attr,omitempty"`
	Auto     string `xml:"auto,attr,omitempty"`
	Colour   string `xml:"colour,attr,omitempty"`
	Equalise string `xml:"equalise,attr,omitempty"`
	Font     string `xml:"font,attr,omitempty"`
	Format   string `xml:"format,attr,omitempty"`
	Gain     string `xml:"gain,attr,omitempty"`
	High     string `xml:"high,attr,omitempty"`
	Id       string `xml:"id,attr,omitempty"`
	Srcname  string `xml:"label,attr,omitempty"`
	Low      string `xml:"low,attr,omitempty"`
	Map      string `xml:"map,attr,omitempty"`
	Missing  string `xml:"missing,attr,omitempty"`
	Reverse  string `xml:"reverse,attr,omitempty"`
	Row      string `xml:"row,attr,omitempty"`
	Rrd      string `xml:"rrd,attr,omitempty"`
	Style    string `xml:"style,attr,omitempty"`

	Copyright *Copyright `xml:",omitempty"`
	ColourMap *ColourMap `xml:",omitempty"`
	Date      *Date      `xml:",omitempty"`
	Label     *Label     `xml:",omitempty"`
	Name      *Name      `xml:",omitempty"`
	Scale     *Scale     `xml:",omitempty"`
	Scalebar  *Scalebar  `xml:",omitempty"`
	Tag       *Tag       `xml:",omitempty"`
	Tags      []Tag      `xml:",omitempty"`
	Title     *Title     `xml:",omitempty"`
	XGrid     *XGrid     `xml:",omitempty"`
	XGrids    []XGrid    `xml:",omitempty"`
	XLabel    *XLabel    `xml:",omitempty"`
	XTick     *XTick     `xml:",omitempty"`
	YGrid     *YGrid     `xml:",,omitempty"`
	YLabel    *YLabel    `xml:",omitempty"`
	YTick     *YTick     `xml:",omitempty"`

	lag float64
}

type ColourMap struct {
	XMLName xml.Name `xml:"colourmap"`

	Border *Border `xml:",omitempty"`
	Title  *Title  `xml:",omitempty"`
	YGrid  *YGrid  `xml:"ygrid,omitempty"`
	YTick  *YTick  `xml:",omitempty"`
}

type Title struct {
	XMLName xml.Name `xml:"title"`

	Colour string `xml:"colour,attr,omitempty"`
	Font   string `xml:"font,attr,omitempty"`
	String string `xml:"string,attr,omitempty"`
}

type XGrid struct {
	XMLName xml.Name `xml:"xgrid"`

	Colour string `xml:"colour,attr,omitempty"`
	Hint   string `xml:"hint,attr,omitempty"`
	IsDate string `xml:"isdate,attr,omitempty"`
	Pen    string `xml:"pen,attr,omitempty"`
	Step   string `xml:"step,attr,omitempty"`
}

type XLabel struct {
	XMLName xml.Name `xml:"xlabel"`

	Colour string `xml:"colour,attr,omitempty"`
	Font   string `xml:"font,attr,omitempty"`
	String string `xml:"string,attr,omitempty"`
}

type XTick struct {
	XMLName xml.Name `xml:"xtick"`

	Colour  string `xml:"colour,attr,omitempty"`
	Factor  string `xml:"factor,attr,omitempty"`
	Font    string `xml:"font,attr,omitempty"`
	IsDate  string `xml:"isdate,attr,omitempty"`
	Max     string `xml:"max,attr,omitempty"`
	Min     string `xml:"min,attr,omitempty"`
	Step    string `xml:"step,attr,omitempty"`
	String  string `xml:"string,attr,omitempty"`
	YOffset string `xml:"yoffset,attr,omitempty"`
}

type YGrid struct {
	XMLName xml.Name `xml:"ygrid"`

	Colour string `xml:"colour,attr,omitempty"`
	Hint   string `xml:"hint,attr,omitempty"`
	IsDate string `xml:"isdate,attr,omitempty"`
	Pen    string `xml:"pen,attr,omitempty"`
	Step   string `xml:"step,attr,omitempty"`
}

type YTick struct {
	XMLName xml.Name `xml:"ytick"`

	Aligned string `xml:"aligned,attr,omitempty"`
	Colour  string `xml:"colour,attr,omitempty"`
	Factor  string `xml:"factor,attr,omitempty"`
	Font    string `xml:"font,attr,omitempty"`
	Hint    string `xml:"hint,attr,omitempty"`
	Max     string `xml:"max,attr,omitempty"`
	Min     string `xml:"min,attr,omitempty"`
	Rotated string `xml:"rotated,attr,omitempty"`
	Step    string `xml:"step,attr,omitempty"`
	String  string `xml:"string,attr,omitempty"`
	XOffset string `xml:"xoffset,attr,omitempty"`
}

type YLabel struct {
	XMLName xml.Name `xml:"ylabel"`

	Colour string `xml:"colour,attr,omitempty"`
	Font   string `xml:"font,attr,omitempty"`
	Hint   string `xml:"hint,attr,omitempty"`
	Step   string `xml:"step,attr,omitempty"`
	String string `xml:"string,attr,omitempty"`
}

type Name struct {
	XMLName xml.Name `xml:"name"`

	Box     string `xml:"box,attr,omitempty"`
	Colour  string `xml:"colour,attr,omitempty"`
	Font    string `xml:"font,attr,omitempty"`
	Pad     string `xml:"pad,attr,omitempty"`
	String  string `xml:"string,attr,omitempty"`
	XOffset string `xml:"xoffset,attr,omitempty"`
	YOffset string `xml:"yoffset,attr,omitempty"`
}

type Tag struct {
	XMLName xml.Name `xml:"tag"`

	Aligned string `xml:"aligned,attr,omitempty"`
	Box     string `xml:"box,attr,omitempty"`
	Colour  string `xml:"colour,attr,omitempty"`
	Font    string `xml:"font,attr,omitempty"`
	Rotated string `xml:"rotated,attr,omitempty"`
	String  string `xml:"string,attr,omitempty"`
	XOffset string `xml:"xoffset,attr,omitempty"`
	YOffset string `xml:"yoffset,attr,omitempty"`
}

type Scalebar struct {
	XMLName xml.Name `xml:"scalebar"`

	Length  string `xml:"length,attr,omitempty"`
	Stroke  string `xml:"stroke,attr,omitempty"`
	Width   string `xml:"width,attr,omitempty"`
	YOffset string `xml:"yoffset,attr,omitempty"`

	Scale Scale
}

type Scale struct {
	XMLName xml.Name `xml:"scale"`

	Aligned string `xml:"aligned,attr,omitempty"`
	Colour  string `xml:"colour,attr,omitempty"`
	Font    string `xml:"font,attr,omitempty"`
	String  string `xml:"string,attr,omitempty"`
}

type Streams []Stream

func (s Streams) Len() int           { return len(s) }
func (s Streams) Swap(a, b int)      { s[a], s[b] = s[b], s[a] }
func (s Streams) Less(a, b int) bool { return s[a].lag > s[b].lag }

func (c *Chart) Marshal() ([]byte, error) {

	res, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return nil, err
	}
	str := string(res)
	for _, t := range []string{"border", "copyright", "date", "label", "name", "tag", "title", "xlabel", "ylabel", "xgrid", "xtick", "ygrid", "ytick", "scale", "stream"} {
		str = strings.Replace(str, "></"+t+">", " />", -1)
		str = strings.Replace(str, "&#39;", "'", -1)
	}
	res = append([]byte(header+warning), append([]byte(str), '\n')...)

	return res, nil
}
