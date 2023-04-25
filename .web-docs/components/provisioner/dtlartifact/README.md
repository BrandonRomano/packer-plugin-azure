Type: `azure-dtlartifact`

The Azure DevTest Labs provisioner can be used to apply an artifact to a VM - See [Add an artifact to a VM](https://docs.microsoft.com/en-us/azure/devtest-labs/add-artifact-vm)


## Configuration Reference

There are many configuration options available for the builder. We'll start
with authentication parameters, then go over the Azure ARM builder specific
options. In addition to the options listed here, a [communicator](/packer/docs/templates/legacy_json_templates/communicator) can be configured for this builder.

### Authentication options

<!-- Code generated from the comments of the Config struct in builder/azure/common/client/config.go; DO NOT EDIT MANUALLY -->

Config allows for various ways to authenticate Azure clients.  When
`client_id` and `subscription_id` are specified in addition to one and only
one of the following: `client_secret`, `client_jwt`, `client_cert_path` --
Packer will use the specified Azure Active Directory (AAD) Service Principal
(SP).  If only `use_interactive_auth` is specified, Packer will try to
interactively log on the current user (tokens will be cached).  If none of
these options are specified, Packer will attempt to use the Managed Identity
and subscription of the VM that Packer is running on.  This will only work if
Packer is running on an Azure VM with either a System Assigned Managed
Identity or User Assigned Managed Identity.

<!-- End of code generated from the comments of the Config struct in builder/azure/common/client/config.go; -->


#### Managed Identity

If you're running Packer on an Azure VM with a [managed identity](/packer/integrations/BrandonRomano/azure#azure-managed-identity) you don't need to specify any additional configuration options. As Packer will attempt to use the Managed Identity and subscription of the VM that Packer is running on.

#### Interactive User Authentication

To use interactive user authentication, you should specify `subscription_id` only.
Packer will use cached credentials or redirect you to a website to log in.

#### Service Principal

To use a [service principal](/packer/integrations/BrandonRomano/azure#azure-active-directory-service-principal)
you should specify `subscription_id`, `client_id` and one of `client_secret`,
`client_cert_path` or `client_jwt`.

- `subscription_id` (string) - Subscription under which the build will be
  performed. **The service principal specified in `client_id` must have full
  access to this subscription, unless build_resource_group_name option is
  specified in which case it needs to have owner access to the existing
  resource group specified in build_resource_group_name parameter.**

- `client_id` (string) - The Active Directory service principal associated with
  your builder.

- `client_secret` (string) - The password or secret for your service principal.

- `client_cert_path` (string) - The location of a PEM file containing a
  certificate and private key for service principal.

- `client_cert_token_timeout` (duration string | ex: "1h30m12s") - How long to set the expire time on the token created when using
  `client_cert_path`.

- `client_jwt` (string) - The bearer JWT assertion signed using a certificate
  associated with your service principal principal. See [Azure Active
  Directory docs](https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-certificate-credentials)
  for more information.

## Azure DevTest Labs provisioner specific options


### Required:

<!-- Code generated from the comments of the Config struct in provisioner/azure-dtlartifact/provisioner.go; DO NOT EDIT MANUALLY -->

- `dtl_artifacts` ([]DtlArtifact) - Dtl Artifacts

- `lab_name` (string) - Name of the existing lab where the virtual machine exist.

- `lab_resource_group_name` (string) - Name of the resource group where the lab exist.

- `vm_name` (string) - Name of the virtual machine within the DevTest lab.

<!-- End of code generated from the comments of the Config struct in provisioner/azure-dtlartifact/provisioner.go; -->


### Optional:

<!-- Code generated from the comments of the Config struct in provisioner/azure-dtlartifact/provisioner.go; DO NOT EDIT MANUALLY -->

- `polling_duration_timeout` (duration string | ex: "1h5m2s") - The default PollingDuration for azure is 15mins, this property will override
  that value. See [Azure DefaultPollingDuration](https://godoc.org/github.com/Azure/go-autorest/autorest#pkg-constants)
  If your Packer build is failing on the
  ARM deployment step with the error `Original Error:
  context deadline exceeded`, then you probably need to increase this timeout from
  its default of "15m" (valid time units include `s` for seconds, `m` for
  minutes, and `h` for hours.)

- `azure_tags` (map[string]\*string) - Azure Tags

<!-- End of code generated from the comments of the Config struct in provisioner/azure-dtlartifact/provisioner.go; -->


#### DtlArtifact
<!-- Code generated from the comments of the DtlArtifact struct in provisioner/azure-dtlartifact/provisioner.go; DO NOT EDIT MANUALLY -->

- `artifact_name` (string) - Artifact Name

- `artifact_id` (string) - Artifact Id

- `parameters` ([]ArtifactParameter) - Parameters

<!-- End of code generated from the comments of the DtlArtifact struct in provisioner/azure-dtlartifact/provisioner.go; -->


#### ArtifactParmater
<!-- Code generated from the comments of the ArtifactParameter struct in provisioner/azure-dtlartifact/provisioner.go; DO NOT EDIT MANUALLY -->

- `name` (string) - Name

- `value` (string) - Value

- `type` (string) - Type

<!-- End of code generated from the comments of the ArtifactParameter struct in provisioner/azure-dtlartifact/provisioner.go; -->


## Basic Example

```hcl
source "null" "example" {
  communicator = "none"
}

build {
  sources = ["source.null.example"]

  provisioner "azure-dtlartifact" {
    lab_name                          = "packer-test"
    lab_resource_group_name           = "packer-test"
    vm_name                          = "packer-test-vm"
    dtl_artifacts {
        artifact_name = "linux-apt-package"
        parameters {
          name  = "packages"
          value = "vim"
        }
        parameters {
          name  = "update"
          value = "true"
        }
        parameters {
          name  = "options"
          value = "--fix-broken"
        }
    }
  }
}
```
