// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
