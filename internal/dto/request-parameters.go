package dto

type RequestParameters struct {
	SplitID    int64
	DeviceID   int64
	NameSpace  string
	City       string
	Store      string
	Parameters []Parameter
}

type Parameter struct {
	ParamName string
	Value     string
}
