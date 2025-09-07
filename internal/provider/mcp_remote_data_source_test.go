package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMcpRemoteDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories, // Defined in provider_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccMcpRemoteDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.agentsmith_mcp_remote.test", "id"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_remote.test", "url", "http://localhost:8080/mcp"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_remote.test", "headers.%", "1"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_remote.test", "headers.X-My-Header", "my-value"),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_remote.test", "json", regexp.MustCompile(`"url":"http://localhost:8080/mcp"`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_remote.test", "json", regexp.MustCompile(`"transport":"sse"`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_remote.test", "json", regexp.MustCompile(`"description":"My remote server"`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_remote.test", "json", regexp.MustCompile(`"timeout":10000`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_remote.test", "json", regexp.MustCompile(`"headers":{"X-My-Header":"my-value"}`)),
				),
			},
		},
	})
}

const testAccMcpRemoteDataSourceConfig = `
data "agentsmith_mcp_remote" "test" {
  url         = "http://localhost:8080/mcp"
  transport   = "sse"
  description = "My remote server"
  timeout     = 10000
  headers = {
    "X-My-Header" = "my-value"
  }
}
`
