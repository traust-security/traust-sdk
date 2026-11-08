// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// PqcClassification — Values of `pqc_classification` in pqc-blockers.schema.json.
type PqcClassification string

const (
	PqcClassificationShorKeyEstablishment PqcClassification = "shor-key-establishment"
	PqcClassificationShorSignature        PqcClassification = "shor-signature"
	PqcClassificationClock2030Parameter   PqcClassification = "clock-2030-parameter"
	PqcClassificationClassicallyBroken    PqcClassification = "classically-broken"
	PqcClassificationHndlExposure         PqcClassification = "hndl-exposure"
	PqcClassificationPqcBlockerConfig     PqcClassification = "pqc-blocker-config"
	PqcClassificationPqcAdoption          PqcClassification = "pqc-adoption"
)

// PqcClassificationValues returns all valid PqcClassification values.
func PqcClassificationValues() []PqcClassification {
	return []PqcClassification{
		PqcClassificationShorKeyEstablishment,
		PqcClassificationShorSignature,
		PqcClassificationClock2030Parameter,
		PqcClassificationClassicallyBroken,
		PqcClassificationHndlExposure,
		PqcClassificationPqcBlockerConfig,
		PqcClassificationPqcAdoption,
	}
}
