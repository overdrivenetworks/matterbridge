// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InsightIdentity undocumented
type InsightIdentity struct {
	// Object is the base model of InsightIdentity
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Address undocumented
	Address *string `json:"address,omitempty"`
}

// InsightValueDouble undocumented
type InsightValueDouble struct {
	// UserExperienceAnalyticsInsightValue is the base model of InsightValueDouble
	UserExperienceAnalyticsInsightValue
	// Value undocumented
	Value *float64 `json:"value,omitempty"`
}

// InsightValueInt undocumented
type InsightValueInt struct {
	// UserExperienceAnalyticsInsightValue is the base model of InsightValueInt
	UserExperienceAnalyticsInsightValue
	// Value undocumented
	Value *int `json:"value,omitempty"`
}
