provider "null" {}

resource "null_resource" "deploy_function" {
  provisioner "local-exec" {
    command = <<EOT
      faas-cli build -f stack.yaml --no-cache
      faas-cli push -f stack.yaml
      faas-cli deploy -f stack.yaml
    EOT
  }
}

resource "null_resource" "destroy_function" {
  provisioner "local-exec" {
    when = destroy

    command = "faas-cli remove hello-fn --gateway $OPENFAAS_URL"
  }
}