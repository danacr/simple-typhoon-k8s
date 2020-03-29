FROM ubuntu:18.04


RUN apt-get update && apt-get install -y ca-certificates curl gettext-base git unzip wget

# install terraform
RUN wget https://releases.hashicorp.com/terraform/0.12.24/terraform_0.12.24_linux_amd64.zip && \
    unzip terraform_0.12.24_linux_amd64.zip -d /usr/local/bin/ && \
    chmod +x /usr/local/bin/terraform && \
    rm terraform_0.12.24_linux_amd64.zip

RUN useradd -rm -d /home/terraform -u 1000 terraform
USER terraform
WORKDIR /home/terraform

# Download the typhoon ct provider
RUN wget https://github.com/poseidon/terraform-provider-ct/releases/download/v0.4.0/terraform-provider-ct-v0.4.0-linux-amd64.tar.gz && \
    tar xzf terraform-provider-ct-v0.4.0-linux-amd64.tar.gz && \
    mkdir -p  ~/.terraform.d/plugins/ && \
    mv terraform-provider-ct-v0.4.0-linux-amd64/terraform-provider-ct ~/.terraform.d/plugins/terraform-provider-ct_v0.4.0 && \
    rm -r terraform-provider-ct*

COPY --chown=terraform . .