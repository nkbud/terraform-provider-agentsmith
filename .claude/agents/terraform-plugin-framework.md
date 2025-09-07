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

Implement data source client functionality

Data sources use the optional Configure method to fetch configured clients from the provider. The provider configures the HashiCups client and the data source can save a reference to that client for its operations.

Implement data source schema

The data source uses the Schema method to define the acceptable configuration and state attribute names and types. The coffees data source will need to save a list of coffees with various attributes to the state.

Implement logging
10min
|
Terraform
Terraform

Reference this often? Create an account to bookmark tutorials.
In this tutorial, you will implement log messages in your provider and filter special values from the log output. Then you will manage log output to view those log statements when executing Terraform. To do this, you will:

Add log messages.
This creates provider-defined log messages in Terraform's logs.
Add structured log fields.
This enhances logging data with provider-defined key-value pairs for greater consistency across multiple logs and easier log viewing.
Add log filtering.
This redacts certain log messages or structured log field data from being included in the log output.
View all Terraform log output during commands.
This shows all Terraform logs in the terminal running a Terraform command.
Save Terraform log output to a file during commands.
This saves all Terraform logs to a file when running a Terraform command.
View specific Terraform log output.
This manages Terraform log output to show only certain logs.

Implement resource create and read
17min
|
Terraform
Terraform
Interactive
Interactive

Show Terminal

Reference this often? Create an account to bookmark tutorials.
In this tutorial, you will add create and read capabilities to a new order resource of a provider that interacts with the API of a fictional coffee-shop application called Hashicups. To do this, you will:

Define the initial resource type.
This prepares the resource to be added to the provider.
Add the resource to the provider.
This enables the resource for testing and Terraform configuration usage.
Implement the HashiCups client in the resource.
This retrieves the configured HashiCups client from the provider and makes it available for resource operations.
Define the resource's schema.
This prepares the resource to accept data from the Terraform configuration and store order information in the Terraform state.
Define the resource's data model.
This models the resource schema as a Go type so the data is accessible for other Go code.
Define the resource's create logic.
This handles calling the HashiCups API to create an order using the configuration saving Terraform state with the data.
Define the resource's read logic.
This handles calling the HashiCups API using the configured client and refreshing the Terraform state with the data.
Verify the resource's behavior.
This verifies that the resource behaves as expected when you refer to it in Terraform configuration.

Implement initial resource type

Providers use an implementation of the resource.Resource interface type as the starting point for a resource implementation.

This interface requires the following:

A Metadata method to define the resource type name, which is how the resource is used in Terraform configurations.
A Schema method to define the schema for any resource configuration, plan, and state data.
A Create method to define the logic which creates the resource and sets its initial Terraform state.
A Read method to define the logic which refreshes the Terraform state for the resource.
An Update method to define the logic which updates the resource and sets the updated Terraform state on success.
A Delete method to define the logic which deletes the resource and removes the Terraform state on success.

Implement log messages

Providers support logging through the tflog package of the github.com/hashicorp/terraform-plugin-log Go module. This package implements structured logging and filtering capabilities.

Open the internal/provider/provider.go file.

Update the top of the Configure method logic with the following.

internal/provider/provider.go
func (p *hashicupsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
tflog.Info(ctx, "Configuring HashiCups client")

    // Retrieve provider data from configuration
    var config hashicupsProviderModel
    /* ... */
Replace the import statement at the beginning of the file with the following.

internal/provider/provider.go
import (
"context"
"os"

    "github.com/hashicorp-demoapp/hashicups-client-go"
    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-log/tflog"
)

Implement structured log fields

The tflog package supports adding additional key-value pairs to logging for consistency and tracing flow. These pairs can be added for the rest of the provider request with the tflog.SetField() call or inline as a final parameter with any logging calls.

Open the internal/provider/provider.go file.

Inside your provider's Configure method, set three logging fields and a log message immediately before the hashicups.NewClient() call with the following.


Implement resource update
12min
|
Terraform
Terraform

Reference this often? Create an account to bookmark tutorials.
In this tutorial, you will add update capabilities to the order resource of a provider that interacts with the API of a fictional coffee-shop application called Hashicups. To do this, you will:

Verify your schema and model.
Verify the last_updated attribute is in the order resource schema and model. The provider will update this attribute to the current date time whenever the order resource is updated.
Implement resource update.
This update method uses the HashiCups client library to invoke a PUT request to the /orders/{orderId} endpoint with the updated order items in the request body. After the update is successful, it updates the resource's state.
Enhance plan output with plan modifier.
This clarifies the plan output of the id attribute to remove its difference by keeping the existing state value on updates.
Verify update functionality.
This ensures the resource is working as expected.

Implement resource delete
7min
|
Terraform
Terraform

Reference this often? Create an account to bookmark tutorials.
In this tutorial, you will add delete capabilities to the order resource of a provider that interacts with the API of a fictional coffee-shop application called Hashicups. To do this, you will:

Implement resource delete.
This delete method uses the HashiCups API client to invoke a DELETE request to the /orders/{orderId} endpoint. After the delete is successful, the framework automatically removes the resource from Terraform's state.
Verify delete functionality.
This ensures that the resource is working as expected.

Implement resource import
8min
|
Terraform
Terraform

Reference this often? Create an account to bookmark tutorials.
In this tutorial, you will add import capabilities to the order resource of a provider that interacts with the API of a fictional coffee-shop application called Hashicups. To do this, you will:

Implement resource import.
This import method takes the given order ID from the terraform import command and enables Terraform to begin managing the existing order.
Verify import functionality.
This ensures that resource import functionality is working as expected.

Implement documentation generation
8min
|
Terraform
Terraform

Reference this often? Create an account to bookmark tutorials.
In this tutorial, you will add documentation generation capabilities to a provider that interacts with the API of a fictional coffee-shop application called HashiCups. To do this, you will:

Add terraform-plugin-docs to the provider.
This enables the provider to automatically generate data source, resource, and function documentation.
Add schema descriptions.
This enhances the documentation to include a description for the provider itself, its data sources and resources, and each of their attributes.
Add configuration examples.
This enhances the documentation to include example Terraform configurations.
Add resource import documentation.
This enhances the resource documentation to include an example of how to import the resources that support it.
Run documentation generation.
This ensures that the documentation generation works as expected.
The terraform-plugin-docs Go module cmd/tfplugindocs command enables providers to implement documentation generation. The generation uses schema descriptions and conventionally placed files to produce provider documentation that is compatible with the Terraform Registry.

Add schema descriptions

The tfplugindocs tool will automatically include schema-based descriptions, if present in a data source, provider, or resource's schema. The schema.Schema type's Description field describes the data source, provider, or resource itself. Each attribute's or block's Description field describes that particular attribute or block. These descriptions should be tailored to practitioner usage and include any caveats or value expectations, such as special syntax.

Open the internal/provider/provider.go file.

Add documentation to your provider by replacing the entire Schema method with the following, which adds Description fields to the provider's schema and each of it's attributes.

Add configuration examples

The tfplugindocs tool will automatically include Terraform configuration examples from files with the following naming conventions:

Provider: examples/provider/provider.tf
Resources: examples/resources/TYPE/resource.tf
Data Sources: examples/data-sources/TYPE/data-source.tf
Functions: examples/functions/TYPE/function.tf
Replace TYPE with the name of the resource, data source, or function. For example: examples/resources/hashicups_order/resource.tf.

Open the examples/provider/provider.tf file and replace the existing code with the following.

Release and publish to the Terraform registry
10min
|
Terraform
Terraform

Reference this often? Create an account to bookmark tutorials.
The Terraform Registry is HashiCorp's public repository of Terraform providers and modules. Whenever you initialize a Terraform configuration, Terraform attempts to download the necessary providers from the Terraform Registry. The registry also hosts provider-specific documentation.

Navigate to your terraform-provider-hashicups directory.

Your code should match the 11-documentation-generation directory from the example repository.

Create an empty, public GitHub repository and name it terraform-provider-hashicups.

In your local terraform-provider-hashicups directory, complete the following steps to push the code to your new repository.

First, remove the current remote origin. This disassociates the local repository from hashicorp/terraform-provider-hashicups, so you can add a new origin.

$ git remote rm origin

Then, add a remote repository that points to your newly-created GitHub repository, replacing GH_ORG with your GitHub organization name or username.

$ git remote add origin https://github.com/GH_ORG/terraform-provider-hashicups

Next, move the existing code to the main branch.

$ git branch -M main

Finally, push the code to your GitHub repository.

$ git push -u origin main

Verify Terraform Registry Manifest file

The Registry Manifest provides additional information about the provider to the Terraform Registry.

Verify that the terraform-registry-manifest.json file exists in the root of your repository.

Notice that the contents of the file specify the version of the file itself, and the Terraform protocol version that the provider uses.

terraform-registry-manifest.json
{
"version": 1,
"metadata": {
"protocol_versions": ["6.0"]
}
}
Verify GoReleaser configuration

GoReleaser is a tool for building Go projects for multiple platforms, creating a checksums file, and signing the release.

Verify that the .goreleaser.yml file exists in the root of your repository.

This configuration generates and signs the necessary release artifacts to publish your provider to Terraform Registry. The GitHub Action workflow will run GoReleaser with this configuration. HashiCorp maintains this GoReleaser configuration file to use with any provider you release.

Verify GitHub Action workflow

Verify that the release.yml file exists in the .github/workflows/ directory

This GitHub Action will release new versions of your provider whenever you tag a commit on the main branch.

Generate GPG Signing Key

You need a GPG Signing Key to sign your provider binaries using GoReleaser and to verify your provider in the Terraform Registry. The Terraform Registry only supports RSA and DSA keys.

Generate your GPG key pair. When prompted to select which kind of key, respond with a 1 to select the RSA and RSA option.

$ gpg --full-generate-key
gpg (GnuPG) 2.3.6; Copyright (C) 2021 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Please select what kind of key you want:
(1) RSA and RSA
(2) DSA and Elgamal
(3) DSA (sign only)
(4) RSA (sign only)
(9) ECC (sign and encrypt) *default*
(10) ECC (sign only)
(14) Existing key from card
Your selection?

Enter 4096 when prompted for the key size.

RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (3072) 4096
Requested keysize is 4096 bits
Press Enter to accept the default option, indicating that the key does not expire.

Please specify how long the key should be valid.
0 = key does not expire
<n>  = key expires in n days
<n>w = key expires in n weeks
<n>m = key expires in n months
<n>y = key expires in n years
Key is valid for? (0)
Key does not expire at all

Confirm your selection by entering y.

Is this correct? (y/N) y
Create a user ID at the prompt. Use your own information, not the example information below. Leave the comment blank.

GnuPG needs to construct a user ID to identify your key.

Real name: Terraform Education
Email address: team-terraform-education@hashicorp.com
Comment:
When prompted, confirm your USER-ID by entering O. Save your USER-ID, you will use this later to generate your private and public keys.

You selected this USER-ID:
"Terraform Education <team-terraform-education@hashicorp.com>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? O
Finally, create and confirm your passphrase. This passphrase is used to decrypt your GPG private key, select a strong passphrase and keep it secure.

Generate GPG public key

Generate your GPG public key. Later, you will add this key to the Terraform Registry. When you publish a new version of your provider to the Terraform Registry, the registry will validate that the release is signed with this keypair and Terraform will verify the provider during terraform init.

Replace KEY_ID with the USER-ID you used when creating your GPG key (Your name and email address).

$ gpg --armor --export "KEY_ID"
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBGB91soBEACp/u6jJZAKeVahHGtR6jzDFcOvivhUFV2fuwBBW/jxqYrWEeEX
+rny8oChQjCABHG9bUxhSSqPyfCK/kUI3VK8Qrxpby6dQgqOFuF61P1mBI0BppLF
JGydv8J9SIIbYJOyajFzMLrvL+xvKD1AblRFBtQke8ts+gz9B+oW7SmAVMb3gM4n
##...
-----END PGP PUBLIC KEY BLOCK-----

Generate GPG private key

Generate your GPG private key. GoReleaser will use this to sign the provider releases.

Replace KEY_ID with the USER-ID you used when creating your GPG key. GPG will prompt you to enter the passphrase you created earlier.

$ gpg --armor --export-secret-keys "KEY_ID"
-----BEGIN PGP PRIVATE KEY BLOCK-----

lQdFBGB91soBEACp/u6jJZAKeVahHGtR6jzDFcOvivhUFV2fuwBBW/jxqYrWEeEX
+rny8oChQjCABHG9bUxhSSqPyfCK/kUI3VK8Qrxpby6dQgqOFuF61P1mBI0BppLF
JGydv8J9SIIbYJOyajFzMLrvL+xvKD1AblRFBtQke8ts+gz9B+oW7SmAVMb3gM4n
##...
-----END PGP PRIVATE KEY BLOCK-----

Note
Be sure to save your production GPG keys and passphrases and back them up in a secure location.
Add GitHub secrets for GitHub Action

The GitHub Action requires your GPG private key and passphrase to generate a release.

In your forked GitHub repository, go to Settings, then Secrets, then Actions, and click the New repository secret button to add the following repository secrets.

GPG_PRIVATE_KEY. This is the private GPG key you generated in the previous step. Include the -----BEGIN... and -----END lines.
PASSPHRASE. This is the passphrase for your GPG private key.
Add secrets to your forked GitHub repository for GoReleaser GH Actions

Create a provider release

The GitHub Action will trigger and create a release for your provider whenever a new valid version tag is pushed to the repository. Terraform provider versions must follow the Semantic Versioning standard (vMAJOR.MINOR.PATCH).

First, add your changes to a new commit.

$ git add .

Then, commit your changes.

$ git commit -m 'Add docs, goreleaser, and GH actions'

Next, create a new tag.

$ git tag v0.2.1

Finally, push the tag to GitHub.

$ git push origin v0.2.1

Verify provider release

In your forked repository, go to Actions. You should find the GitHub Action workflow you created earlier running.

Tip
You may need to acknowledge and approve workflow actions before the GitHub Action runs. If you need to do so, approve the workflow actions and manually trigger the run.
View GitHub Action creating release artifacts

Once the GitHub Action completes, you should find the release on the right-hand side of your main repository page.

New release for HashiCups provider

Add GPG public key to Terraform Registry

Go to the Terraform Registry and sign in with your GitHub account.

Next, go to User Settings, then Signing Keys. Select + New GPG Key and add the GPG Public signing key you generated in a previous step.

Add public GPG key to Terraform Registry

Click Save to add your public signing key.

Publish your provider to Terraform Registry

Since you cannot un-publish a provider from the Terraform Registry, you will not actually publish your HashiCups provider in this tutorial. However, this section walks you through the steps you would follow for a real provider. It is safe to follow most of the steps in this section, but do not click the Publish button at the end.

In the Terraform Registry, select Publish, then Providers from the top navigation bar.

Open providers page in Terraform Registry

Select your organization, then the GitHub repository containing your Terraform provider.

Select your provider category, and read and agree to the "Terms of Use" by checking the box.

Warning
Do not click Publish. Use this tutorial as a reference to publish your own custom Terraform provider. Please do not publish your copy of HashiCups to the Terraform Registry.
Terraform Registry publish provider confirmation page

Next steps

Congratulations! You have released your provider and walked through the steps to add it to the Terraform Registry.

Over the course of these tutorials, you re-created the HashiCups provider and learned how to create data sources, authenticate the provider to the HashiCups client, create resources with CRUD functionality, and import existing resources. In addition, you have released and published the HashiCups provider to the Terraform Registry.

A full list of official, partner, and community Terraform providers can be found on the Terraform Provider Registry. We encourage you to find a provider you are interested in and start contributing!

