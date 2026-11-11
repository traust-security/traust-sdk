// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
