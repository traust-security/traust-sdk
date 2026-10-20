// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// PeachBoundaryType — Values of `boundary_type` in report.schema.json.
type PeachBoundaryType string

const (
	PeachBoundaryTypeHardwareSeparation     PeachBoundaryType = "hardware_separation"
	PeachBoundaryTypeHardwareVirtualization PeachBoundaryType = "hardware_virtualization"
	PeachBoundaryTypeContainerization       PeachBoundaryType = "containerization"
	PeachBoundaryTypeDataSegmentation       PeachBoundaryType = "data_segmentation"
	PeachBoundaryTypeNetworkSegmentation    PeachBoundaryType = "network_segmentation"
	PeachBoundaryTypeIdentitySegmentation   PeachBoundaryType = "identity_segmentation"
)

// PeachBoundaryTypeValues returns all valid PeachBoundaryType values.
func PeachBoundaryTypeValues() []PeachBoundaryType {
	return []PeachBoundaryType{
		PeachBoundaryTypeHardwareSeparation,
		PeachBoundaryTypeHardwareVirtualization,
		PeachBoundaryTypeContainerization,
		PeachBoundaryTypeDataSegmentation,
		PeachBoundaryTypeNetworkSegmentation,
		PeachBoundaryTypeIdentitySegmentation,
	}
}
