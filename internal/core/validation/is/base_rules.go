package is

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var (
	Email            = is.Email
	EmailFormat      = is.EmailFormat
	URL              = is.URL
	RequestURL       = is.RequestURL
	RequestURI       = is.RequestURI
	Alpha            = is.Alpha
	Digit            = is.Digit
	Alphanumeric     = is.Alphanumeric
	UTFLetter        = is.UTFLetter
	UTFDigit         = is.UTFDigit
	UTFLetterNumeric = is.UTFLetterNumeric
	UTFNumeric       = is.UTFNumeric
	LowerCase        = is.LowerCase
	UpperCase        = is.UpperCase
	Hexadecimal      = is.Hexadecimal
	HexColor         = is.HexColor
	RGBColor         = is.RGBColor
	Int              = is.Int
	Float            = is.Float
	UUIDv3           = is.UUIDv3
	UUIDv4           = is.UUIDv4
	UUIDv5           = is.UUIDv5
	UUID             = is.UUID
	CreditCard       = is.CreditCard
	ISBN10           = is.ISBN10
	ISBN13           = is.ISBN13
	ISBN             = is.ISBN
	JSON             = is.JSON
	ASCII            = is.ASCII
	PrintableASCII   = is.PrintableASCII
	Multibyte        = is.Multibyte
	FullWidth        = is.FullWidth
	HalfWidth        = is.HalfWidth
	VariableWidth    = is.VariableWidth
	Base64           = is.Base64
	DataURI          = is.DataURI
	E164             = is.E164
	CountryCode2     = is.CountryCode2
	CountryCode3     = is.CountryCode3
	CurrencyCode     = is.CurrencyCode
	DialString       = is.DialString
	MAC              = is.MAC
	IP               = is.IP
	IPv4             = is.IPv4
	IPv6             = is.IPv6
	Subdomain        = is.Subdomain
	Domain           = is.Domain
	DNSName          = is.DNSName
	Host             = is.Host
	Port             = is.Port
	MongoID          = is.MongoID
	Latitude         = is.Latitude
	Longitude        = is.Longitude
	SSN              = is.SSN
	Semver           = is.Semver
)
