Feature: My awesome module is working

Scenario: Accessing my module instance via SSH
  When the terraform module "my-awesome-module" is deployed 
  Then I can access the instance via SSH

