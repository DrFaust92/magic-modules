package artifactregistry_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccDataSourceArtifactRegistryRepositoryFile(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	resourceName := "data.google_artifact_registry_repository_file.test"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceArtifactRegistryRepositoryFileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "file"),
				),
			},
		},
	})
}

const testAccDataSourceArtifactRegistryRepositoryFileConfig = `
data "google_artifact_registry_repository_file" "test" {
	project       = "cloudrun"
	location      = "us"
	repository_id = "container"
	file          = "hello:latest"
}
`
