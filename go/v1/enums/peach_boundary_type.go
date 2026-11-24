// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
