package goexiftool

import "time"

type E struct {
	EF []EXIF
}

type EXIF struct {
	Aperture                 float64
	ApertureValue            float64
	ApplicationRecordVersion int64
	BitsPerSample            int64
	BrightnessValue          float64
	CodedCharacterSet        string
	ColorComponents          int64
	ColorSpace               string
	ComponentsConfiguration  string
	Compression              string
	CreateDate               string
	CurrentIPTCDigest        string
	DateTimeOriginal         string
	Directory                string
	EncodingProcess          string
	EnvelopeRecordVersion    int64
	ExifByteOrder            string
	ExifImageHeight          int64
	ExifImageWidth           int64
	ExifToolVersion          float64
	ExifVersion              string
	ExposureCompensation     float64
	ExposureMode             string
	ExposureProgram          string
	ExposureTime             string
	FileAccessDate           string
	FileCreateDate           string
	FileModifyDate           string
	FileName                 string
	FilePermissions          string
	FileSize                 string
	FileType                 string
	FileTypeExtension        string
	Flash                    string
	FlashpixVersion          string
	FNumber                  float64
	FocalLength              string
	FocalLength35efl         string
	//GPSInfo                         // embedded GPSInfo
	GPSAltitude         string //float32
	GPSAltitudeRef      string //float32
	GPSDateStamp        string
	GPSLatitude         string //float32
	GPSLatitudeRef      string
	GPSLongitude        string //float32
	GPSLongitudeRef     string
	GPSMapDatum         string
	GPSProcessingMethod string
	GPSDateTime         string //time.Time
	GPSVersionID        string
	ImageDescription    string `bson:"image_description"` //` varchar(254) DEFAULT NULL,
	ImageHeight         int64
	ImageSize           string
	ImageUniqueID       string
	ImageWidth          int64
	InteropIndex        string
	InteropVersion      string
	IPTCDigest          string
	ISO                 int64
	LightSource         string
	LightValue          float64
	Make                string
	MaxApertureValue    float64
	Megapixels          float64
	MeteringMode        string
	MIMEType            string
	Model               string
	ModifyDate          string
	Orientation         string
	RegionInfo          RegionInfo
	ResolutionUnit      string
	SceneCaptureType    string
	SceneType           string
	SensingMethod       string
	ShutterSpeed        string
	ShutterSpeedValue   string
	Software            string
	SourceFile          string
	ThumbnailImage      string
	ThumbnailLength     int64
	ThumbnailOffset     int64
	UserComment         string
	WhiteBalance        string
	XMPToolkit          string
	XResolution         float64
	YCbCrPositioning    string
	YCbCrSubSampling    string
	YResolution         float64
}
type RegionInfo struct {
	AppliedToDimensions AppliedToDimensions
	RegionList          []RegionList
}

type AppliedToDimensions struct {
	H    int64
	Unit string
	W    int64
}

type RegionList struct {
	Area Area
	Name string
	Type string
}
type Area struct {
	H    float64
	Unit string
	W    float64
	X    float64
	Y    float64
}
type GPSInfo struct {
	GPSAltitude         float32
	GPSAltitudeRef      float32
	GPSDateStamp        string
	GPSLatitude         float32
	GPSLatitudeRef      string
	GPSLongitude        float32
	GPSLongitudeRef     string
	GPSMapDatum         string
	GPSProcessingMethod string
	GPSDateTime         time.Time
	GPSVersionID        string
}
