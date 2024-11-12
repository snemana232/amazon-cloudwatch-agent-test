// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT
variable "region" {
  type    = string
  default = "us-west-2"
}
variable "test_dir" {
  type    = string
  default = "./test/kueue"
}
variable "cwagent_image_repo" {
  type    = string
  default = "public.ecr.aws/cloudwatch-agent/cloudwatch-agent"
}
variable "cwagent_image_tag" {
  type    = string
  default = "latest"
}
variable "k8s_version" {
  type    = string
  default = "1.31"
}

variable "ami_type" {
  type    = string
  default = "AL2_x86_64"
}

variable "instance_type" {
  type    = string
  default = "m5.large"
}