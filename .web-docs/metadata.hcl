# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "TODO"
  description = "TODO"
  identifier = "packer/BrandonRomano/azure"
  component {
    type = "builder"
    name = "Azure chroot"
    slug = "chroot"
  }
  component {
    type = "builder"
    name = "Azure arm"
    slug = "arm"
  }
  component {
    type = "builder"
    name = "Azure DevTest Lab"
    slug = "dtl"
  }
  component {
    type = "provisioner"
    name = "Azure DevTest Lab"
    slug = "dtlartifact"
  }
}
