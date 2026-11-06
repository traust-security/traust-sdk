// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
