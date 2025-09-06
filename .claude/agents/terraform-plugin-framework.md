---
name: terraform-plugin
description: when the task requires usage of the terraform-plugin-framework
model: inherit
color: purple
---

The plugin framework is HashiCorp’s recommended way develop to Terraform Plugins on protocol version 6 or protocol version 5.

We recommend using the framework to develop new providers because it offers significant advantages as compared to Terraform Plugin SDKv2. We also recommend migrating existing providers to the framework when possible. Refer to Plugin Framework Benefits for higher level details about how the framework makes provider development easier and Plugin Framework Features for a detailed functionality comparison between the SDKv2 and the framework.

# Get Started

Try the Terraform Plugin Framework tutorials.
Clone the terraform-provider-scaffolding-framework template repository on GitHub.

Key Concepts:
Provider Servers encapsulate all Terraform plugin details and handle all calls for provider, resource, and data source operations by implementing the Terraform Plugin Protocol. 
They are implemented as binaries that the Terraform CLI downloads, starts, and stops.
Providers are the top level abstraction that define the available resources and data sources for practitioners to use and may accept its own configuration, such as authentication information.
Schemas define available fields for provider, resource, or provisioner configuration block, and give Terraform metadata about those fields.
Resources are an abstraction that allow Terraform to manage infrastructure objects, such as a compute instance, an access policy, or disk. Providers act as a translation layer between Terraform and an API, offering one or more resources for practitioners to define in a configuration.
Data Sources are an abstraction that allow Terraform to reference external data. Providers have data sources that tell Terraform how to request external data and how to convert the response into a format that practitioners can interpolate.
Functions are an abstraction that allow Terraform to reference computational logic. Providers can implement their own custom logic functions to augment the Terraform configuration language built-in functions.

Test and Publish:
Learn to write acceptance tests for your provider.
Learn to publish your provider to the Terraform Registry.

Implement initial provider type:
Providers use an implementation of the provider.Provider interface type as the starting point for all implementation details.

This interface requires the following:
A Metadata method to define the provider type name for inclusion in each data source and resource type name. For example, a resource type named "hashicups_order" would have a provider type name of "hashicups".
A Schema method to define the schema for provider-level configuration. Later in these tutorials, you will update this method to accept a HashiCups API token and endpoint.
A Configure method to configure shared clients for data source and resource implementations.
A DataSources method to define the provider's data sources.
A Resources method to define the provider's resources.

Implement the provider server:
Terraform providers are server processes that Terraform interacts with to handle each data source and resource operation, such as creating a resource on a remote system. Later in these tutorials, you will connect those Terraform operations to a locally running HashiCups API.

Serving a provider follows these steps:
Starts a provider server process. By implementing the main function, which is the code execution starting point for Go language programs, a long-running server will listen for Terraform requests.
Framework provider servers also support optional functionality such as enabling support for debugging tools. You will not implement this functionality in these tutorials.

Prepare Terraform for local provider install:
Terraform installs providers and verifies their versions and checksums when you run terraform init. Terraform will download your providers from either the provider registry or a local registry. However, while building your provider you will want to test Terraform configuration against a local development build of the provider. The development build will not have an associated version number or an official set of checksums listed in a provider registry.
Terraform allows you to use local provider builds by setting a dev_overrides block in a configuration file called .terraformrc. This block overrides all other configured installation methods.
Terraform searches for the .terraformrc file in your home directory and applies any configuration settings you set.

A configured client will then be available for any data source or resource to use. To do this, you will:

# Client

Define the provider schema.
This prepares the provider to accept Terraform configuration for client authentication and host information.
Define the provider data model.
This models the provider schema as a Go type so the data is accessible for other Go code.
Define the provider configure method.
This reads the Terraform configuration using the data model or checks environment variables if data is missing from the configuration. It raises errors if any necessary client configuration is missing. The configured client is then created and made available for data sources and resources.
Verify configuration behaviors.
This ensures the expected provider configuration behaviors.

Implement provider schema
The Plugin Framework uses a provider's Schema method to define the acceptable configuration attribute names and types. The HashiCups client needs a host, username, and password to be properly configured. The Terraform Plugin Framework types package contains schema and data model types that can work with Terraform's null, unknown, or known values.

Implement provider data model
The Terraform Plugin Framework uses Go struct types with tfsdk struct field tags to map schema definitions into Go types with the actual data. The types within the struct must align with the types in the schema.

Implement temporary data source

Provider configuration only occurs if there is a valid data source or resource supported by the provider and used in a Terraform configuration. For now, create a temporary data source implementation so you can verify the provider configuration behaviors. Later tutorials will guide you through the concepts and implementation details of real data sources and resources.

In this tutorial, you will implement a data source to read the list of coffees from the HashiCups API and save it in Terraform’s state. To do this, you will:

Define the initial data source type.
This prepares the data source to be added to the provider.
Add data source to provider.
This enables the data source for testing and Terraform configuration usage.
Implement the HashiCups client in the data source.
This retrieves the configured HashiCups client from the provider and makes it available for data source operations.
Define the data source schema.
This prepares the data source to set Terraform state with the list of coffees.
Define the data source data model.
This models the data source schema as a Go type so the data is accessible for other Go code.
Define the data source read logic.
This handles calling the HashiCups API using the configured client and setting the Terraform state with the data.
Verify data source behavior.
This ensures the expected data source behavior.

This interface requires the following:

Metadata method. This defines the data source type name, which is how the data source is referred to in Terraform configurations.
Schema method. This defines the schema for any data source configuration and state data.
Read method. This defines the logic which sets the Terraform state for the data source.