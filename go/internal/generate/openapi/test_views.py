import unittest

from views import openapi, pascal, refs, view_component


class ViewComponentTest(unittest.TestCase):
    def test_names_are_pascal_case(self):
        self.assertEqual(pascal("findings-rollup"), "FindingsRollup")
        self.assertEqual(pascal("top_finding"), "TopFinding")

    def test_a_nullable_type_list_becomes_a_nullable_type(self):
        self.assertEqual(
            openapi({"type": ["string", "null"], "format": "date-time"}, "P"),
            {"type": "string", "format": "date-time", "nullable": True},
        )

    def test_a_null_enum_member_becomes_nullability(self):
        out = openapi({"type": ["string", "null"], "enum": ["high", None]}, "P")
        self.assertEqual(out, {"type": "string", "enum": ["high"], "nullable": True})

    def test_a_nullable_reference_rides_an_allof(self):
        out = openapi({"oneOf": [{"$ref": "#/$defs/assessment"}, {"type": "null"}]}, "Cov")
        self.assertEqual(out, {"allOf": [{"$ref": "#/components/schemas/CovAssessment"}], "nullable": True})

    def test_validation_only_keywords_are_dropped(self):
        node = {
            "type": "array",
            "items": {"type": "string"},
            "contains": {"const": "x"},
            "maxContains": 1,
            "not": {"pattern": "y"},
            "allOf": [{"if": {}, "then": {}}],
        }
        self.assertEqual(openapi(node, "P"), {"type": "array", "items": {"type": "string"}})

    def test_a_definition_composed_by_reference_is_that_definition(self):
        node = {"description": "d", "allOf": [{"$ref": "#/$defs/breakdown"}], "propertyNames": {}}
        self.assertEqual(
            openapi(node, "B"), {"description": "d", "$ref": "#/components/schemas/BBreakdown"}
        )

    def test_structure_in_an_allof_stops_the_generator(self):
        with self.assertRaises(SystemExit):
            openapi({"allOf": [{"properties": {"a": {}}}, {"$ref": "#/$defs/x"}]}, "P")

    def test_a_reference_deep_into_defs_stops_the_generator(self):
        with self.assertRaises(SystemExit):
            openapi({"$ref": "#/$defs/metric_is/coverage"}, "P")

    def test_numbers_are_doubles_and_a_bare_string_enum_is_a_string(self):
        self.assertEqual(openapi({"type": "number"}, "P"), {"type": "number", "format": "double"})
        self.assertEqual(openapi({"enum": ["a", "b"]}, "P"), {"enum": ["a", "b"], "type": "string"})

    def test_a_reference_to_another_view_is_that_views_component(self):
        base = "https://example.com/traust-contracts/views_schema/v1/findings-rollup.schema.json"
        self.assertEqual(openapi({"$ref": base}, "E"), {"$ref": "#/components/schemas/FindingsRollupInsight"})
        self.assertEqual(
            openapi({"$ref": base + "#/$defs/severity_counts"}, "E"),
            {"$ref": "#/components/schemas/FindingsRollupInsightSeverityCounts"},
        )

    def test_a_nullable_reference_to_another_view_rides_an_allof(self):
        ref = "https://example.com/traust-contracts/views_schema/v1/pqc-portfolio.schema.json"
        out = openapi({"oneOf": [{"$ref": ref}, {"type": "null"}]}, "E")
        self.assertEqual(out, {"allOf": [{"$ref": "#/components/schemas/PqcPortfolioInsight"}], "nullable": True})

    def test_only_a_view_or_its_defs_is_a_component(self):
        base = "https://example.com/traust-contracts/views_schema/v1/findings-rollup.schema.json"
        self.assertIsNone(view_component(base + "#/$defs/top_finding/properties/cvss_score"))
        self.assertIsNone(view_component("https://example.com/traust-contracts/artifacts_schema/v1/x.schema.json"))
        self.assertIsNone(view_component("#/$defs/local"))

    def test_refs_lists_every_reference(self):
        node = {"properties": {"a": {"$ref": "#/x/A"}, "b": {"items": {"$ref": "#/x/B"}}}}
        self.assertEqual(sorted(refs(node)), ["#/x/A", "#/x/B"])


if __name__ == "__main__":
    unittest.main()
