// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ExtendedKeyUsage undocumented
type ExtendedKeyUsage struct {
	// Object is the base model of ExtendedKeyUsage
	Object
	// Name Extended Key Usage Name
	Name *string `json:"name,omitempty"`
	// ObjectIdentifier Extended Key Usage Object Identifier
	ObjectIdentifier *string `json:"objectIdentifier,omitempty"`
}

// ExtendedPlacePropertiesModel undocumented
type ExtendedPlacePropertiesModel struct {
	// Object is the base model of ExtendedPlacePropertiesModel
	Object
	// PriceRange undocumented
	PriceRange *string `json:"priceRange,omitempty"`
	// MenuURL undocumented
	MenuURL *string `json:"menuUrl,omitempty"`
	// BusinessURL undocumented
	BusinessURL *string `json:"businessUrl,omitempty"`
	// FormattedAddress undocumented
	FormattedAddress *string `json:"formattedAddress,omitempty"`
	// OpeningHoursSpecifications undocumented
	OpeningHoursSpecifications []OpeningHoursSpecification `json:"openingHoursSpecifications,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
}
