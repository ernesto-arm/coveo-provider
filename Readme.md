## Terraform Provider for COVEO
Created this provider to be able to manage coveo source (push mainly) for each environment in terraform.
Feel free to add more functionality.

*Disclaimer:* This functionality is only intended for fulfilling my needs of creating/destroying PUSH sources in terraform.
I haven't test any other source type.

#### Supported functionality:
- Source (CRUD)

#### Build 
``go build -o terraform-provider-coveo.exe``

#### Test
``terraform init``
``terraform apply``

#### Update dependencies
``go get -u github/ernesto-arm/go-coveo/...``

# TODO
- At the moment the coveo upstream service is pointing to a fork of https://github.com/coveooss/go-coveo under my account.
  There is a PR to merge it into the original one, when accepted ideally I should migrate to the original so will be easy to extends.
  