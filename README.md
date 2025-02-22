# ms: A Go library to deal with Mass Spectrometry data.

Includes the unthermo library that can read Thermo Finnigan RAW files without external API's,
a successor to [unfinnigan](http://code.google.com/p/unfinnigan/wiki/FileLayoutOverview).

## Usable tools:

* XIC, prints the mass chromatogram for a given m/z.
    * [linux amd64 executable](https://bitbucket.org/proteinspector/ms/downloads/xic)
    * [windows amd64 executable](https://bitbucket.org/proteinspector/ms/downloads/xic.exe)
    * [macos amd64 executable](https://bitbucket.org/proteinspector/ms/downloads/xicmac)
    * Documentation over [here](https://bitbucket.org/proteinspector/ms/src/master/unthermo/tools/xic.go)

## reader.go

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type [AuditTag](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-1078) 

```
type AuditTag struct {
	Time     uint64   //8 bytes Windows 64-bit timestamp
	Tag1     audittag //50 bytes
	Tag2     audittag
	Unknown1 uint32 //4 bytes
}
```


#### type [AutoSamplerInfo](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-832) 

```
type AutoSamplerInfo struct {
	Preamble AutoSamplerPreamble
	Text     PascalString
}
```


AutoSamplerInfo comes from the sampling device

#### func (\*AutoSamplerInfo) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-824) 

```
func (data *AutoSamplerInfo) Read(r io.Reader, v Version)
```


#### type [AutoSamplerPreamble](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-837) 

```
type AutoSamplerPreamble struct {
	Unknown1      uint32
	Unknown2      uint32
	NumberOfWells uint32
	Unknown3      uint32
	Unknown4      uint32
	Unknown15     uint32
}
```


#### type [CDataPacket](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-609) 

```
type CDataPacket struct {
	Value float64
	Time  float64
}
```


CDataPackets are the data from Chromatography machines

#### func (\*CDataPacket) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-622) 

```
func (data *CDataPacket) Read(r io.Reader, v Version)
```


#### type [CDataPackets](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-614) 

```
type CDataPackets []CDataPacket
```


#### func (CDataPackets) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-616) 

```
func (data CDataPackets) Read(r io.Reader, v Version)
```


#### type [CIndexEntries](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-566) 

```
type CIndexEntries []CIndexEntry
```


#### func (CIndexEntries) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-568) 

```
func (data CIndexEntries) Read(r io.Reader, v Version)
```


#### type [CIndexEntry](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-549) 

```
type CIndexEntry struct {
	Offset32 uint32
	Index    uint32
	Event    uint16
	Unknown1 uint16
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
	Unknown5 float64
	Time     float64
	Unknown6 float64
	Unknown7 float64
	Value    float64

	Offset uint64
}
```


#### func (\*CIndexEntry) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-583) 

```
func (data *CIndexEntry) Read(r io.Reader, v Version)
```


#### func (CIndexEntry) [Size](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-574) 

```
func (data CIndexEntry) Size(v Version) uint64
```


#### type [CentroidedPeak](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-251) 

```
type CentroidedPeak struct {
	Mz        float32
	Abundance float32
}
```


A peak itself

#### type [File](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-17) 

```
type File struct {
	// contains filtered or unexported fields
}
```


File is an in-memory representation of the Thermo RAW file

#### func [Open](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-28) 

```
func Open(fn string) (file File, err error)
```


Open opens the supplied filename and reads the indices from the RAW file in memory. Multiple files may be read concurrently.

#### func (\*File) [AllScans](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-78) 

```
func (rf *File) AllScans(fun func(scan ms.Scan))
```


AllScans is a convenience function that runs over all spectra in the raw file

On every encountered MS Scan, the function fun is called

#### func (\*File) [Chromatography](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-1127) 

```
func (rf *File) Chromatography(instr int) (cdata CDataPackets)
```


Experimental: read out chromatography data from a connected instrument

#### func (\*File) [Close](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-69) 

```
func (rf *File) Close() error
```


Close closes the RAW file

#### func (\*File) [ComputeMeanSpectrum](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-111) 

```
func (rf *File) ComputeMeanSpectrum() (s ms.Spectrum)
```


Computes mean spectrum from profile data

#### func (\*File) [NScans](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-106) 

```
func (rf *File) NScans() int
```


NScans returns the number of scans in the index

#### func (\*File) [Scan](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-87) 

```
func (rf *File) Scan(sn int) (scan ms.Scan)
```


Scan returns the scan at the scan number in argument

#### type [FileHeader](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-1057) 

```
type FileHeader struct {
	Magic      uint16    //2 bytes
	Signature  signature //18 bytes
	Unknown1   uint32    //4 bytes
	Unknown2   uint32    //4 bytes
	Unknown3   uint32    //4 bytes
	Unknown4   uint32    //4 bytes
	Version    Version   //4 bytes
	AuditStart AuditTag  //112 bytes
	AuditEnd   AuditTag  //112 bytes
	Unknown5   uint32    //4 bytes
	Unknown6   [60]byte  //60 bytes
	Tag        headertag //1028 bytes
}
```


The Thermo fileheaders most valuable piece of info is the file version. It determines the reading strategy for some data structures that changed over time

#### func (\*FileHeader) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-1106) 

```
func (data *FileHeader) Read(r io.Reader, v Version)
```


#### type [FractionCollector](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-374) 

```
type FractionCollector struct {
	Lowmz  float64
	Highmz float64
}
```


#### type [InfoPreamble](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-956) 

```
type InfoPreamble struct {
	Methodfilepresent uint32
	Year              uint16
	Month             uint16
	Weekday           uint16
	Day               uint16
	Hour              uint16
	Minute            uint16
	Second            uint16
	Millisecond       uint16

	Unknown1        uint32
	DataAddr32      uint32
	NControllers    uint32
	NControllers2   uint32
	Unknown2        uint32
	Unknown3        uint32
	RunHeaderAddr32 []uint32
	Unknown4        []uint32
	Unknown5        []uint32
	Padding1        [764]byte //760 bytes, 756 bytes in v57

	DataAddr      uint64
	Unknown6      uint64
	RunHeaderAddr []uint64
	Unknown7      []uint64
	Padding2      [1032]byte //1024 bytes, 1008 bytes in v64
}
```


#### type [InjectionData](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-887) 

```
type InjectionData struct {
	Unknown1                    uint32
	Rownumber                   uint32
	Unknown2                    uint32
	Vial                        [6]uint16 //utf-16
	Injectionvolume             float64
	SampleWeight                float64
	SampleVolume                float64
	InternationalStandardAmount float64
	Dilutionfactor              float64
}
```


#### type [PacketHeader](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-257) 

```
type PacketHeader struct {
	Unknown1           uint32
	ProfileSize        uint32
	PeaklistSize       uint32
	Layout             uint32
	DescriptorListSize uint32
	UnknownStreamSize  uint32
	TripletStreamSize  uint32
	Unknown2           uint32
	Lowmz              float32
	Highmz             float32
}
```


A Header containing info about how many peaks/profile points were registered

#### type [PascalString](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-1097) 

```
type PascalString struct {
	Length int32
	Text   []uint16
}
```


#### func (PascalString) [String](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-1102) 

```
func (t PascalString) String() string
```


#### type [PeakDescriptor](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-238) 

```
type PeakDescriptor struct {
	Index  uint16
	Flags  uint8
	Charge uint8
}
```


A struct containing more info about the peaks

#### type [PeakList](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-245) 

```
type PeakList struct {
	Count uint32
	Peaks []CentroidedPeak
}
```


The data structure holding the peaks

#### type [Profile](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-271) 

```
type Profile struct {
	FirstValue float64
	Step       float64
	PeakCount  uint32
	Nbins      uint32
	Chunks     []ProfileChunk
}
```


The structure containing the profile-mode points

#### type [ProfileChunk](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-280) 

```
type ProfileChunk struct {
	Firstbin uint32
	Nbins    uint32
	Fudge    float32
	Signal   []float32
}
```


Profile points are collected in chunks with adjacent signal points

#### type [RawFileInfo](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-946) 

```
type RawFileInfo struct {
	Preamble InfoPreamble
	Heading1 PascalString
	Heading2 PascalString
	Heading3 PascalString
	Heading4 PascalString
	Heading5 PascalString
	Unknown1 PascalString
}
```


RawFileInfo contains the addresses of the different RunHeaders, (header of the data that each connected instrument produced) also the acquisition date

#### func (\*RawFileInfo) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-985) 

```
func (data *RawFileInfo) Read(r io.Reader, v Version)
```


#### type [Reaction](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-366) 

```
type Reaction struct {
	Precursormz float64
	Unknown1    float64
	Energy      float64
	Unknown2    uint32
	Unknown3    uint32
}
```


#### type [RunHeader](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-631) 

```
type RunHeader struct {
	SampleInfo        SampleInfo
	Filename1         filename
	Filename2         filename
	Filename3         filename
	Filename4         filename
	Filename5         filename
	Filename6         filename
	Unknown1          float64
	Unknown2          float64
	Filename7         filename
	Filename8         filename
	Filename9         filename
	Filename10        filename
	Filename11        filename
	Filename12        filename
	Filename13        filename
	ScantrailerAddr32 uint32
	ScanparamsAddr32  uint32
	Unknown3          uint32
	Unknown4          uint32
	Nsegs             uint32
	Unknown5          uint32
	Unknown6          uint32
	OwnAddr32         uint32
	Unknown7          uint32
	Unknown8          uint32

	ScanindexAddr   uint64
	DataAddr        uint64
	InstlogAddr     uint64
	ErrorlogAddr    uint64
	Unknown9        uint64
	ScantrailerAddr uint64
	ScanparamsAddr  uint64
	Unknown10       uint32
	Unknown11       uint32
	OwnAddr         uint64

	Unknown12 uint32
	Unknown13 uint32
	Unknown14 uint32
	Unknown15 uint32
	Unknown16 uint32
	Unknown17 uint32
	Unknown18 uint32
	Unknown19 uint32
	Unknown20 uint32
	Unknown21 uint32
	Unknown22 uint32
	Unknown23 uint32
	Unknown24 uint32
	Unknown25 uint32
	Unknown26 uint32
	Unknown27 uint32
	Unknown28 uint32
	Unknown29 uint32
	Unknown30 uint32
	Unknown31 uint32
	Unknown32 uint32
	Unknown33 uint32
	Unknown34 uint32
	Unknown35 uint32

	Unknown36 [8]byte
	Unknown37 uint32
	Device    PascalString
	Model     PascalString
	SN        PascalString
	SWVer     PascalString
	Tag1      PascalString
	Tag2      PascalString
	Tag3      PascalString
	Tag4      PascalString
}
```


RunHeaders contain all data addresses for data that a certain machine connected to the Mass Spectrometer (including the MS itself) has acquired. Also SN data is available

#### func (\*RunHeader) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-707) 

```
func (data *RunHeader) Read(r io.Reader, v Version)
```


#### type [SampleInfo](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-800) 

```
type SampleInfo struct {
	Unknown1        uint32
	Unknown2        uint32
	FirstScanNumber uint32
	LastScanNumber  uint32
	InstlogLength   uint32
	Unknown3        uint32
	Unknown4        uint32
	ScanindexAddr   uint32 //unused in 64-bit versions
	DataAddr        uint32
	InstlogAddr     uint32
	ErrorlogAddr    uint32
	Unknown5        uint32
	MaxSignal       float64
	Lowmz           float64
	Highmz          float64
	Starttime       float64
	Endtime         float64
	Unknown6        [56]byte
	Tag1            [44]uint16
	Tag2            [20]uint16
	Tag3            [160]uint16
}
```


SampleInfo contains some other info

#### type [ScanDataPacket](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-228) 

```
type ScanDataPacket struct {
	Header         PacketHeader
	Profile        Profile
	PeakList       PeakList
	DescriptorList []PeakDescriptor
	Unknown        []float32
	Triplets       []float32
}
```


An MS scan packet, containing Centroid Peak or Profile intensities

#### func (\*ScanDataPacket) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-295) 

```
func (data *ScanDataPacket) Read(r io.Reader, v Version)
```


#### type [ScanDataPackets](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-287) 

```
type ScanDataPackets []ScanDataPacket
```


#### func (ScanDataPackets) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-289) 

```
func (data ScanDataPackets) Read(r io.Reader, v Version)
```


#### type [ScanEvent](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-348) 

```
type ScanEvent struct {
	Preamble [132]uint8 //128 bytes from v63 on, 120 in v62, 80 in v57, 41 below that
	//Preamble[6] == ms-level
	//Preamble[40] == analyzer
	Nprecursors uint32

	Reaction []Reaction

	Unknown1 [13]uint32
	MZrange  [3]FractionCollector
	Nparam   uint32

	Unknown2 [4]float64
	A        float64
	B        float64
	C        float64
}
```


ScanEvents are encoded headers of the MS scans, their Preamble contain the MS level, type of ionization etc. Events themselves contain range, and conversion parameters from Hz to m/z

#### func (ScanEvent) [Convert](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-454) 

```
func (data ScanEvent) Convert(v float64) float64
```


Convert Hz values to m/z

#### func (\*ScanEvent) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-387) 

```
func (data *ScanEvent) Read(r io.Reader, v Version)
```


#### type [ScanEvents](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-379) 

```
type ScanEvents []ScanEvent
```


#### func (ScanEvents) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-381) 

```
func (data ScanEvents) Read(r io.Reader, v Version)
```


#### type [ScanIndex](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-489) 

```
type ScanIndex []ScanIndexEntry
```


#### func (ScanIndex) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-491) 

```
func (data ScanIndex) Read(r io.Reader, v Version)
```


#### type [ScanIndexEntry](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-470) 

```
type ScanIndexEntry struct {
	Offset32       uint32
	Index          uint32
	Scanevent      uint16
	Scansegment    uint16
	Next           uint32
	Unknown1       uint32
	DataPacketSize uint32
	Time           float64
	Totalcurrent   float64
	Baseintensity  float64
	Basemz         float64
	Lowmz          float64
	Highmz         float64
	Offset         uint64
	Unknown2       uint32
	Unknown3       uint32
}
```


#### func (\*ScanIndexEntry) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-508) 

```
func (data *ScanIndexEntry) Read(r io.Reader, v Version)
```


#### func (ScanIndexEntry) [Size](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-497) 

```
func (data ScanIndexEntry) Size(v Version) uint64
```


#### type [SequencerRow](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-849) 

```
type SequencerRow struct {
	Injection  InjectionData
	Unknown1   PascalString
	Unknown2   PascalString
	ID         PascalString
	Comment    PascalString
	Userlabel1 PascalString
	Userlabel2 PascalString
	Userlabel3 PascalString
	Userlabel4 PascalString
	Userlabel5 PascalString
	Instmethod PascalString
	Procmethod PascalString
	Filename   PascalString
	Path       PascalString

	Vial     PascalString
	Unknown3 PascalString
	Unknown4 PascalString
	Unknown5 uint32

	Unknown6  PascalString
	Unknown7  PascalString
	Unknown8  PascalString
	Unknown9  PascalString
	Unknown10 PascalString
	Unknown11 PascalString
	Unknown12 PascalString
	Unknown13 PascalString
	Unknown14 PascalString
	Unknown15 PascalString
	Unknown16 PascalString
	Unknown17 PascalString
	Unknown18 PascalString
	Unknown19 PascalString
	Unknown20 PascalString
}
```


SequencerRow contains more information about what the autosampler did

#### func (\*SequencerRow) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-899) 

```
func (data *SequencerRow) Read(r io.Reader, v Version)
```


#### type [TrailerLength](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-337) 

```
type TrailerLength uint32
```


I currently have no idea what TrailerLength is

#### func (\*TrailerLength) [Read](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-339) 

```
func (data *TrailerLength) Read(r io.Reader, v Version)
```


#### type [Version](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/unthermo/reader.go#lines-1110) 

```
type Version uint32
```

## data.go

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type [Analyzer](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/data.go#lines-29) 

```
type Analyzer int
```


Analyzer is the mass analyzer

```
const (
	ITMS Analyzer = iota
	TQMS
	SQMS
	TOFMS
	FTMS
	Sector
	Undefined
)
```


The analyzer types are documented in literature

#### type [Peak](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/data.go#lines-5) 

```
type Peak struct {
	Mz float64
	I  float32
}
```


Peak represents an ion peak

#### type [Scan](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/data.go#lines-14) 

```
type Scan struct {
	Analyzer Analyzer
	MSLevel  uint8
	//Spectrum is a function forcing the read of a spectrum,
	//which is "delayed" for efficiency reasons. If it was not delayed
	//and Spectrum were a data structure, it would always have to
	//be read, which is very expensive. Now if only another property of
	//Scan (cheaper to obtain) is requested, resources are saved.
	Spectrum func() Spectrum
	//PrecursorMzs is only filled with mz values at MSx scans.
	PrecursorMzs []float64
	Time         float64
}
```


Scan represents the peak acquisition event of the mass spectrometer

#### type [Spectrum](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/data.go#lines-11) 

```
type Spectrum []Peak
```


A Spectrum is a collection of peaks

#### func (Spectrum) [Len](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/data.go#lines-44) 

```
func (a Spectrum) Len() int
```


#### func (Spectrum) [Less](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/data.go#lines-46) 

```
func (a Spectrum) Less(i, j int) bool
```


#### func (Spectrum) [Swap](https://bitbucket.org/lomereiter/ms/src/0764cbb94863/data.go#lines-45) 

```
func (a Spectrum) Swap(i, j int)
```

## Tools

A set of tools using the unthermo library

The labelq tool inserts iTRAQ reporter ions from HCD scans in CID spectra.

```
Program output is MGF formatted MS2 spectra

```


The peakstats tool outputs a few data about the peaks of supplied ions:

*   Mass, Time at Maximum, Maximal intensity of ions found in the LC/MS map
*   Full width at half maximum of this maximal peak

The printspectrum tool prints out the spectrum (mz and intensity values) of a Thermo RAW File

```
Every line of the output is a peak registered by the mass spectrometer
characterized by an m/z value in Da and an intensity in the mass spectrometer's unit of abundance

```


The XIC tool prints mass chromatograms for a specific m/z.

```
For the m/z given, it prints the peak with highest intensity in interval
[mz-tol ppm,mz+tol ppm] for every MS-1 scan.

Every line contains the retention time and intensity of a found peak

Example:
    xic -mz 361.1466 -tol 2.5 -raw rawfile.raw

Output:
    0.003496666666666667 10500.583
    0.015028333333333333 11793.04
    0.03391333333333333 10178.598
    0.05393333333333334 10671.821
    0.07350833333333334 11572.251

```
