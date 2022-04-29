package policy

import (
	"fmt"
	"strings"

	"github.com/vranyes/goranger/utils"
)

type PolicyResource struct {
	Values      []string `json:"values"`
	IsRecursive bool     `json:"isRecursive"`
	IsExcludes  bool     `json:"isExcludes"`
}

type PolicyItemAccess struct {
	Type      string `json:"type"`
	IsAllowed bool   `json:"isAllowed"`
}

type PolicyItemCondition struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type PolicyItem struct {
	Groups        []string              `json:"groups"`
	Accesses      []PolicyItemAccess    `json:"accesses"`
	DelegateAdmin bool                  `json:"delegateAdmin"`
	Users         []string              `json:"users"`
	Conditions    []PolicyItemCondition `json:"conditions"`
	Roles         []string              `json:"roles"`
}

type Interval struct {
	Minutes int `json:"minutes"`
	Hours   int `json:"hours"`
	Days    int `json:"days"`
}

type Schedule struct {
	Month      string `json:"month"`
	Year       string `json:"year"`
	DayOfWeek  string `json:"dayOfWeek"`
	Hour       string `json:"hour"`
	Minute     string `json:"minute"`
	DayOfMonth string `json:"dayOfMonth"`
}

type ValidityRecurrence struct {
	Interval Interval `json:"interval"`
	Schedule Schedule `json:"schedule"`
}

type ValiditySchedule struct {
	Recurrences []ValidityRecurrence `json:"recurrences"`
	StartTime   string               `json:"startTime"`
	TimeZone    string               `json:"timeZone"`
	EndTime     string               `json:"endTime"`
}

type ItemDataMaskInfo struct {
	DataMaskType  string `json:"dataMaskType"`
	ValueExpr     string `json:"valueExpr"`
	ConditionExpr string `json:"conditionExpr"`
}

type PolicyItemRowFilterInfo struct {
	FilterExpr string `json:"filterExpr"`
}

type PolicyItemDataMaskInfo struct {
	DataMaskType  string `json:"dataMaskType"`
	ValueExpr     string `json:"valueExpr"`
	ConditionExpr string `json:"conditionExpr"`
}

type DataMaskPolicyItem struct {
	DataMaskInfo  PolicyItemDataMaskInfo `json:"dataMaskInfo"`
	Groups        []string               `json:"groups"`
	Accesses      []PolicyItemAccess     `json:"accesses"`
	DelegateAdmin bool                   `json:"delegateAdmin"`
	Users         []string               `json:"users"`
	Conditions    []PolicyItemCondition  `json:"conditions"`
	Roles         []string               `json:"roles"`
}

type RowFilterPolicyItem struct {
	RowFilterInfo []PolicyItemRowFilterInfo `json:"rowFilterInfo"`
	Groups        []string                  `json:"groups"`
	Accesses      []PolicyItemAccess        `json:"accesses"`
	DelegateAdmin bool                      `json:"delegateAdmin"`
	Users         []string                  `json:"users"`
	Conditions    []PolicyItemCondition     `json:"conditions"`
	Roles         []string                  `json:"roles"`
}

type Policy struct {
	AllowExceptions      []PolicyItem              `json:"allowExceptions"`
	Conditions           []PolicyItemCondition     `json:"conditions"`
	DataMaskPolicyItems  []DataMaskPolicyItem      `json:"dataMaskPolicyItems"`
	DenyExceptions       []PolicyItem              `json:"denyExceptions"`
	DenyPolicyItems      []PolicyItem              `json:"denyPolicyItems"`
	Description          string                    `json:"description"`
	IsAuditEnabled       bool                      `json:"isAuditEnabled"`
	IsDenyAllElse        bool                      `json:"isDenyAllElse"`
	Name                 string                    `json:"name"`
	PolicyItems          []PolicyItem              `json:"policyItems"`
	PolicyLabels         []string                  `json:"policyLabels"`
	PolicyPriority       int                       `json:"policyPriority"`
	PolicyType           int                       `json:"policyType"`
	Resources            map[string]PolicyResource `json:"resources"`
	ResourceSignature    string                    `json:"resourceSignature"`
	RowFilterPolicyItems []RowFilterPolicyItem     `json:"rowFilterPolicyItems"`
	Service              string                    `json:"service"`
	ServiceType          string                    `json:"serviceType"`
	ValiditySchedules    []ValiditySchedule        `json:"validitySchedules"`
	ZoneName             string                    `json:"zoneName"`

	// Base level options, set only when Unmarshalling from ranger
	CreateTime int64
	Id         int
	UpdatedBy  string
	Version    uint
	CreatedBy  string
	Guid       string
	UpdateTime int64
	IsEnabled  bool
}

func (p Policy) DisplayPolicy() {
	policyName := utils.NullIfEmpty(p.Name)
	createdBy := utils.NullIfEmpty(p.CreatedBy)

	fmt.Println(fmt.Sprintf("Name: %s, Created By: %s", policyName, createdBy))
	for resource, values := range p.Resources {
		fmt.Println(fmt.Sprintf("%s %s", resource, strings.Join(values.Values, "")))
	}

	for _, policy := range p.PolicyItems {
		var accesses strings.Builder
		var length = len(policy.Accesses)
		for i, access := range policy.Accesses {
			accesses.WriteString(access.Type)
			if length != 1 {
				if length >= i+2 {
					accesses.WriteString(", ")
				}
				if length == i+2 {
					accesses.WriteString("and ")
				}
			}
		}
		fmt.Println(fmt.Sprintf(
			"User %s can %s",
			strings.Join(policy.Users, ""), accesses.String()))
	}
	fmt.Println()
}
