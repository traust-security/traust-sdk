import unittest

from normalize import normalize_problems


class ProblemNormalizationTest(unittest.TestCase):
    def test_error_media_share_one_typed_problem_schema(self):
        success = {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/RunOrder"}}}}
        spec = {
            "paths": {
                "/v1/orders/start": {
                    "post": {
                        "responses": {
                            "200": success,
                            "401": {
                                "description": "Unauthorized",
                                "content": {
                                    "application/json": {"schema": {"$ref": "#/components/schemas/Problem"}},
                                    "application/problem+json": {},
                                },
                            },
                            "422": {
                                "content": {
                                    "application/json": {"schema": {"$ref": "#/components/schemas/HTTPValidationError"}}
                                }
                            },
                        }
                    }
                }
            }
        }
        normalize_problems(spec)
        responses = spec["paths"]["/v1/orders/start"]["post"]["responses"]
        self.assertIs(responses["200"], success)
        self.assertEqual(responses["401"]["description"], "Unauthorized")
        for status in ("401", "422"):
            self.assertEqual(
                responses[status]["content"],
                {"application/json": {"schema": {"$ref": "#/components/schemas/Problem"}}},
            )


if __name__ == "__main__":
    unittest.main()
