package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMcpStdioDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories, // Defined in provider_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccMcpStdioDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.agentsmith_mcp_stdio.test", "id"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_stdio.test", "command", "my-server"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_stdio.test", "args.#", "2"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_stdio.test", "args.0", "--port"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_stdio.test", "args.1", "8080"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_stdio.test", "env.%", "1"),
					resource.TestCheckResourceAttr("data.agentsmith_mcp_stdio.test", "env.FOO", "bar"),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_stdio.test", "json", regexp.MustCompile(`"command":"my-server"`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_stdio.test", "json", regexp.MustCompile(`"transport":"stdio"`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_stdio.test", "json", regexp.MustCompile(`"args":\["--port","8080"\]`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_stdio.test", "json", regexp.MustCompile(`"description":"My test server"`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_stdio.test", "json", regexp.MustCompile(`"timeout":5000`)),
					resource.TestMatchResourceAttr("data.agentsmith_mcp_stdio.test", "json", regexp.MustCompile(`"env":{"FOO":"bar"}`)),
				),
			},
		},
	})
}

const testAccMcpStdioDataSourceConfig = `
data "agentsmith_mcp_stdio" "test" {
  command     = "my-server"
  args        = ["--port", "8080"]
  description = "My test server"
  timeout     = 5000
  env = {
    "FOO" = "bar"
  }
}
`
