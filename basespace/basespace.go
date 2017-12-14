package basespace

type Run struct {
	Response RunResponse `json."Response"`
	ResponseStatus struct {
	} `json:"ResponseStatus"`
	Notifications []interface{} `json:"Notifications"`
}

type RunResponse struct {
	ReponseItems   []ReponseItem `json:"Items"`
	DisplayedCount int           `json:"DisplayedCount"`
	TotalCount     int           `json:"TotalCount"`
	Offset         int           `json:"Offset"`
	Limit          int           `json:"Limit"`
	SortDir        string        `json:"SortDir"`
	SortBy         string        `json:"SortBy"`
}

type UserOwnedBy struct {
	ID               string `json:"Id"`
	Href             string `json:"Href"`
	Name             string `json:"Name"`
	DateCreated      string `json:"DateCreated"`
	GravatarURL      string `json:"GravatarUrl"`
	HrefProperties   string `json:"HrefProperties"`
	ExternalDomainID string `json:"ExternalDomainId"`
}

type Instrument struct {
	InstrumentID int    `json:"InstrumentId"`
	Name         string `json:"Name"`
	SerialNumber string `json:"SerialNumber"`
}

type SequencingStats struct {
	Chemistry                     string  `json:"Chemistry"`
	ErrorRate                     float64 `json:"ErrorRate"`
	ErrorRateR1                   float64 `json:"ErrorRateR1"`
	ErrorRateR2                   int     `json:"ErrorRateR2"`
	Href                          string  `json:"Href"`
	IntensityCycle1               float64 `json:"IntensityCycle1"`
	IsIndexed                     bool    `json:"IsIndexed"`
	MaxCycleCalled                int     `json:"MaxCycleCalled"`
	MaxCycleExtracted             int     `json:"MaxCycleExtracted"`
	MaxCycleScored                int     `json:"MaxCycleScored"`
	NonIndexedErrorRate           float64 `json:"NonIndexedErrorRate"`
	NonIndexedIntensityCycle1     float64 `json:"NonIndexedIntensityCycle1"`
	NonIndexedPercentAligned      float64 `json:"NonIndexedPercentAligned"`
	NonIndexedPercentGtQ30        float64 `json:"NonIndexedPercentGtQ30"`
	NonIndexedProjectedTotalYield float64 `json:"NonIndexedProjectedTotalYield"`
	NonIndexedYieldTotal          float64 `json:"NonIndexedYieldTotal"`
	NumCyclesIndex1               int     `json:"NumCyclesIndex1"`
	NumCyclesIndex2               int     `json:"NumCyclesIndex2"`
	NumCyclesRead1                int     `json:"NumCyclesRead1"`
	NumCyclesRead2                int     `json:"NumCyclesRead2"`
	NumLanes                      int     `json:"NumLanes"`
	NumReads                      int     `json:"NumReads"`
	NumSurfaces                   int     `json:"NumSurfaces"`
	NumSwathsPerLane              int     `json:"NumSwathsPerLane"`
	NumTilesPerSwath              int     `json:"NumTilesPerSwath"`
	PercentAligned                float64 `json:"PercentAligned"`
	PercentGtQ30                  float64 `json:"PercentGtQ30"`
	PercentGtQ30R1                float64 `json:"PercentGtQ30R1"`
	PercentGtQ30R2                int     `json:"PercentGtQ30R2"`
	PercentPf                     float64 `json:"PercentPf"`
	PercentResynthesis            float64 `json:"PercentResynthesis"`
	PhasingR1                     float64 `json:"PhasingR1"`
	PhasingR2                     int     `json:"PhasingR2"`
	PrePhasingR1                  float64 `json:"PrePhasingR1"`
	PrePhasingR2                  int     `json:"PrePhasingR2"`
	ProjectedTotalYield           float64 `json:"ProjectedTotalYield"`
	ReadsPfTotal                  int     `json:"ReadsPfTotal"`
	ReadsTotal                    int     `json:"ReadsTotal"`
	YieldTotal                    float64 `json:"YieldTotal"`
}

type AnalysisSettings struct {
	ReverseComplementI5Index bool `json:"ReverseComplementI5Index"`
}

type ReponseItem struct {
	ID                      string           `json:"Id"`
	Href                    string           `json:"Href"`
	Name                    string           `json:"Name"`
	Number                  int              `json:"Number"`
	ExperimentName          string           `json:"ExperimentName"`
	Status                  string           `json:"Status"`
	ReagentBarcode          string           `json:"ReagentBarcode"`
	FlowcellBarcode         string           `json:"FlowcellBarcode"`
	DateCreated             string           `json:"DateCreated"`
	DateModified            string           `json:"DateModified"`
	UserOwnedBy             UserOwnedBy      `json:"UserOwnedBy"`
	TotalSize               int              `json:"TotalSize"`
	PlatformName            string           `json:"PlatformName"`
	Workflow                string           `json:"Workflow"`
	Instrument              Instrument       `json:"Instrument"`
	InstrumentName          string           `json:"InstrumentName"`
	InstrumentType          string           `json:"InstrumentType"`
	NumCyclesRead1          int              `json:"NumCyclesRead1"`
	NumCyclesRead2          int              `json:"NumCyclesRead2"`
	NumCyclesIndex1         int              `json:"NumCyclesIndex1"`
	NumCyclesIndex2         int              `json:"NumCyclesIndex2"`
	HrefBaseSpaceUI         string           `json:"HrefBaseSpaceUI"`
	HasCollaborators        bool             `json:"HasCollaborators"`
	SequencingStats         SequencingStats  `json:"SequencingStats"`
	AnalysisSettings        AnalysisSettings `json:"AnalysisSettings"`
	LaneAndQcStatus         string           `json:"LaneAndQcStatus"`
	LimsStatus              string           `json:"LimsStatus"`
	DateInstrumentStarted   string           `json:"DateInstrumentStarted"`
	DateInstrumentCompleted string           `json:"DateInstrumentCompleted"`
}

type Paging struct {
	DisplayedCount int    `json:"DisplayedCount"`
	TotalCount     int    `json:"TotalCount"`
	Offset         int    `json:"Offset"`
	Limit          int    `json:"Limit"`
	SortDir        string `json:"SortDir"`
	SortBy         string `json:"SortBy"`
}

type LaneIndexSum struct {
	IndexItems []IndexItem `json:"Items"`
	Paging     Paging      `json:"Paging"`
}

type IndexingSummary struct {
	TotalReads               int     `json:"TotalReads"`
	TotalPFReads             int     `json:"TotalPFReads"`
	TotalFractionMappedReads float64 `json:"TotalFractionMappedReads"`
	MappedReadsCV            float64 `json:"MappedReadsCV"`
	MinMappedReads           float64 `json:"MinMappedReads"`
	MaxMappedReads           float64 `json:"MaxMappedReads"`
}

type SampleSheet struct {
	SampleID    string `json:"SampleID"`
	ProjectName string `json:"ProjectName"`
}

type BioSample struct {
	ID             string         `json:"Id"`
	Href           string         `json:"Href"`
	UserOwnedBy    UserOwnedBy    `json:"UserOwnedBy"`
	BioSampleName  string         `json:"BioSampleName"`
	DefaultProject DefaultProject `json:"DefaultProject"`
	DateModified   string         `json:"DateModified"`
	DateCreated    string         `json:"DateCreated"`
	Status         string         `json:"Status"`
	LabStatus      string         `json:"LabStatus"`
}

type DefaultProject struct {
	ID               string      `json:"Id"`
	UserOwnedBy      UserOwnedBy `json:"UserOwnedBy"`
	Href             string      `json:"Href"`
	Name             string      `json:"Name"`
	Description      string      `json:"Description"`
	DateCreated      string      `json:"DateCreated"`
	DateModified     string      `json:"DateModified"`
	HasCollaborators bool        `json:"HasCollaborators"`
	TotalSize        int         `json:"TotalSize"`
}

type LibraryPrep struct {
	ID                      string `json:"Id"`
	Href                    string `json:"Href"`
	Name                    string `json:"Name"`
	ValidIndexingStrategies string `json:"ValidIndexingStrategies"`
	ValidReadTypes          string `json:"ValidReadTypes"`
	NumIndexCycles          int    `json:"NumIndexCycles"`
	DateModified            string `json:"DateModified"`
	State                   string `json:"State"`
	DefaultRead1Cycles      int    `json:"DefaultRead1Cycles"`
	DefaultRead2Cycles      int    `json:"DefaultRead2Cycles"`
}

type Library struct {
	ID           string      `json:"Id"`
	Href         string      `json:"Href"`
	Name         string      `json:"Name"`
	DateCreated  string      `json:"DateCreated"`
	DateModified string      `json:"DateModified"`
	Status       string      `json:"Status"`
	BioSample    BioSample   `json:"BioSample"`
	LibraryPrep  LibraryPrep `json:"LibraryPrep"`
	Project      Project     `json:"Project"`
	Biomolecule  string      `json:"Biomolecule"`
	UserOwnedBy  UserOwnedBy `json:"UserOwnedBy"`
}

type IndexingCount struct {
	ID             int         `json:"Id"`
	Index1         string      `json:"Index1"`
	Index2         string      `json:"Index2"`
	FractionMapped float64     `json:"FractionMapped"`
	Count          int         `json:"Count"`
	SampleSheet    SampleSheet `json:"SampleSheet"`
	BioSample      BioSample   `json:"BioSample"`
	Library        Library     `json:"Library"`
}

type LibraryPool struct {
	ID           string      `json:"Id"`
	Href         string      `json:"Href"`
	UserPoolID   string      `json:"UserPoolId"`
	UserOwnedBy  UserOwnedBy `json:"UserOwnedBy"`
	LibraryCount int         `json:"LibraryCount"`
	DateModified string      `json:"DateModified"`
	DateCreated  string      `json:"DateCreated"`
	Status       string      `json:"Status"`
}

type IndexItem struct {
	LaneNumber             int                    `json:"LaneNumber"`
	IndexingSummary        IndexingSummary        `json:"IndexingSummary"`
	IndexingCounts         []IndexingCount        `json:"IndexingCounts"`
	MappedReads            []float64              `json:"MappedReads"`
	LibraryPool            LibraryPool            `json:"LibraryPool"`
	UndeterminedIndexReads UndeterminedIndexReads `json:"UndeterminedIndexReads"`
}

type Application struct {
	ID                 string   `json:"Id"`
	Href               string   `json:"Href"`
	Name               string   `json:"Name"`
	CompanyName        string   `json:"CompanyName"`
	VersionNumber      string   `json:"VersionNumber"`
	ShortDescription   string   `json:"ShortDescription"`
	DateCreated        string   `json:"DateCreated"`
	PublishStatus      string   `json:"PublishStatus"`
	IsBillingActivated bool     `json:"IsBillingActivated"`
	Category           string   `json:"Category"`
	Classifications    []string `json:"Classifications"`
	AppFamilySlug      string   `json:"AppFamilySlug"`
	AppVersionSlug     string   `json:"AppVersionSlug"`
	Features           []string `json:"Features"`
	LockStatus         string   `json:"LockStatus"`
}

type UserCreatedBy struct {
	ID               string `json:"Id"`
	Href             string `json:"Href"`
	Name             string `json:"Name"`
	DateCreated      string `json:"DateCreated"`
	GravatarURL      string `json:"GravatarUrl"`
	HrefProperties   string `json:"HrefProperties"`
	ExternalDomainID string `json:"ExternalDomainId"`
}

type AppSession struct {
	ID               string        `json:"Id"`
	Name             string        `json:"Name"`
	Href             string        `json:"Href"`
	Application      Application   `json:"Application"`
	UserCreatedBy    UserCreatedBy `json:"UserCreatedBy"`
	ExecutionStatus  string        `json:"ExecutionStatus"`
	QcStatus         string        `json:"QcStatus"`
	StatusSummary    string        `json:"StatusSummary"`
	Purpose          string        `json:"Purpose"`
	DateCreated      string        `json:"DateCreated"`
	DateModified     string        `json:"DateModified"`
	DateCompleted    string        `json:"DateCompleted"`
	TotalSize        int           `json:"TotalSize"`
	DeliveryStatus   string        `json:"DeliveryStatus"`
	ContainsComments bool          `json:"ContainsComments"`
	HrefComments     string        `json:"HrefComments"`
}

type Project struct {
	ID               string      `json:"Id"`
	UserOwnedBy      UserOwnedBy `json:"UserOwnedBy"`
	Href             string      `json:"Href"`
	Name             string      `json:"Name"`
	Description      string      `json:"Description"`
	DateCreated      string      `json:"DateCreated"`
	DateModified     string      `json:"DateModified"`
	HasCollaborators bool        `json:"HasCollaborators"`
	TotalSize        int64       `json:"TotalSize"`
}

type CommonFastq struct {
	MaxLengthRead1      int  `json:"MaxLengthRead1"`
	MaxLengthRead2      int  `json:"MaxLengthRead2"`
	IsPairedEnd         bool `json:"IsPairedEnd"`
	TotalClustersPF     int  `json:"TotalClustersPF"`
	TotalClustersRaw    int  `json:"TotalClustersRaw"`
	TotalReadsPF        int  `json:"TotalReadsPF"`
	TotalReadsRaw       int  `json:"TotalReadsRaw"`
	MaxLengthIndexRead1 int  `json:"MaxLengthIndexRead1"`
	MaxLengthIndexRead2 int  `json:"MaxLengthIndexRead2"`
}

type Attributes struct {
	CommonFastq CommonFastq `json:"common_fastq"`
}

type UndeterminedIndexReads struct {
	ID                  string     `json:"Id"`
	Href                string     `json:"Href"`
	HrefFiles           string     `json:"HrefFiles"`
	Name                string     `json:"Name"`
	DateCreated         string     `json:"DateCreated"`
	DateModified        string     `json:"DateModified"`
	AppSession          AppSession `json:"AppSession"`
	Project             Project    `json:"DatasetType"`
	Attributes          Attributes `json:"Attributes"`
	QcStatus            string     `json:"QcStatus"`
	QcStatusSummary     string     `json:"QcStatusSummary"`
	UploadStatus        string     `json:"UploadStatus"`
	UploadStatusSummary string     `json:"UploadStatusSummary"`
	ValidationStatus    string     `json:"ValidationStatus"`
	HrefComments        string     `json:"HrefComments"`
	ContainsComments    bool       `json:"ContainsComments"`
}
