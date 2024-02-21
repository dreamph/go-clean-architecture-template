package constants

const (
	PdfExtension    = ".pdf"
	ExcelExtension  = ".xlsx"
	JpgExtension    = ".jpg"
	JpegExtension   = ".jpeg"
	PngExtension    = ".png"
	P12Extension    = ".p12"
	CertExtension   = ".cert"
	JrxmlExtension  = ".jrxml"
	JasperExtension = ".jasper"
	DocxExtension   = ".docx"
	ZipExtension    = ".zip"
)

const (
	Comma      = ","
	Pipe       = "|"
	Slash      = "/"
	Dash       = "-"
	Colon      = ":"
	UnderScore = "_"
)

const (
	SeparateFormatDataComma = ", "
)

const (
	ValueEmpty = ""
)

const (
	EN = "EN"
	TH = "TH"
)

const (
	MobilePlatformAndroid = "ANDROID"
	MobilePlatformIOS     = "IOS"
)

// Status
const (
	StatusActive   = int32(20)
	StatusInActive = int32(10)
)

type SortBy int

const (
	SortByAsc  = SortBy(1)
	SortByDesc = SortBy(2)
)
