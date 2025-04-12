package artifactregistry

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceArtifactRegistryRepositoryFile() *schema.Resource {

	return &schema.Resource{
		Read: DataSourceArtifactRegistryRepositoryFileRead,

		Schema: map[string]*schema.Schema{
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Project ID of the project.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The region of the artifact registry repository. For example, "us-west1".`,
			},
			"repository_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The last part of the repository name to fetch from.`,
			},
			"file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The file name to fetch.`,
			},
		},
	}
}

func DataSourceArtifactRegistryRepositoryFileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	// download repository file
	// https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.files/download
	urlRequest, err := tpgresource.ReplaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/files/{{file}}:download")
	if err != nil {
		return fmt.Errorf("Error setting api endpoint")
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		RawURL:    urlRequest,
		UserAgent: userAgent,
	})
	if err != nil {
		return fmt.Errorf("Error retrieving repository file: %s", err)
	}

	// set the schema data using the response
	if err := d.Set("file", res["file"]); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}

	id, err := tpgresource.ReplaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/files/{{file}}")
	if err != nil {
		return fmt.Errorf("Error constructing the data source id: %s", err)
	}

	d.SetId(id)

	return nil
}
