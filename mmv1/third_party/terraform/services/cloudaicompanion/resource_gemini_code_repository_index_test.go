package gemini_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccGeminiCodeRepositoryIndex_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id": os.Getenv("GOOGLE_PROJECT"),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGeminiCodeRepositoryIndex_basic(context),
			},
			{
				ResourceName:            "google_gemini_code_repository_index.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"code_repository_index_id", "labels", "location", "terraform_labels"},
			},
			{
				Config: testAccGeminiCodeRepositoryIndex_update(context),
			},
			{
				ResourceName:            "google_gemini_code_repository_index.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"code_repository_index_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccGeminiCodeRepositoryIndex_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gemini_code_repository_index" "example" {
  provider = google-beta
  labels = {"label1": "value1"}
  location = "us-central1"
  code_repository_index_id = "test-cri-index-example"
  kms_key = "projects/projectExample/locations/locationExample/keyRings/keyRingExample/cryptoKeys/cryptoKeyExample"
}
`, context)
}
func testAccGeminiCodeRepositoryIndex_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gemini_code_repository_index" "example" {
  provider = google-beta
  labels = {"label1": "value1", "label2": "value2"}
  location = "us-central1"
  code_repository_index_id = "test-cri-index-example"
  kms_key = "projects/projectExample/locations/locationExample/keyRings/keyRingExample/cryptoKeys/cryptoKeyExample"
}
`, context)
}
