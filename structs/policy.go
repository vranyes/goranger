package structs

type policyResource struct {
	Values      []string
	IsRecursive bool
	IsExclusive bool
}

type policyItemAccess struct {
	Type      string
	IsAllowed bool
}

type policyItemCondition struct {
	Type   string
	Values []string
}

type policyItem struct {
	Groups        []string
	Accesses      []policyItemAccess
	DelegateAdmin bool
	Users         []string
	Conditions    []policyItemCondition
	Roles         []string
}

type interval struct {
	Minutes int
	Hours   int
	Days    int
}

type schedule struct {
	Month      string
	Year       string
	DayOfWeek  string
	Hour       string
	Minute     string
	DayOfMonth string
}

type validityRecurrence struct {
	Interval interval
	Schedule schedule
}

type validityScuedule struct {
	Recurrences []validityRecurrence
	StartTime   string
	TimeZone    string
	EndTime     string
}

type itemDataMaskInfo struct {
	DataMaskType  string
	ValueExpr     string
	ConditionExpr string
}

type policyItemRowFilterInfo struct {
	FilterExpr string
}

type policyItemDataMaskInfo struct {
	DataMaskType  string
	ValueExpr     string
	ConditionExpr string
}

type dataMaskPolicyItem struct {
	DataMaskInfo  policyItemDataMaskInfo
	Groups        []string
	Accessess     []policyItemAccess
	DelegateAdmin bool
	Users         []string
	Conditions    []policyItemCondition
	Roles         []string
}

type rowFilterPolicyItem struct {
	RowFilterInfo []policyItemRowFilterInfo
	Groups        []string
	Accessess     []policyItemAccess
	DelegateAdmin bool
	Users         []string
	Conditions    []policyItemCondition
	Roles         []string
}

type policyItemConditions struct {
	Values []string
	Type   string
}

type Policy struct {
	DenyExceptions       []policyItem
	AllowExceptions      []policyItem
	ValidityScuedules    []validityScuedule
	PolicyPriority       int
	PolicyType           int
	IsAuditEnabled       bool
	ZoneName             string
	DenyPolicyItems      []policyItem
	Service              string
	Resources            map[string]policyResource
	DataMaskPolicyItems  []dataMaskPolicyItem
	RowFilterPolicyItems []rowFilterPolicyItem
	PolicyItems          []policyItem
	PolicyLabels         []string
	ResourceSignature    string
	Name                 string
	ServiceType          string
	Description          string
	Conditions           []policyItemConditions

	// Base level options, set only when Unmarshalling from ranger
	CreateTime int
	Id         int
	UpdatedBy  string
	Version    uint
	CreatedBy  string
	Guid       string
	UpdateTime int
	IsEnabled  bool
}
