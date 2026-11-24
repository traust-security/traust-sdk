// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
