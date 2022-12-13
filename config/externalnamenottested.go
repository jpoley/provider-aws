/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{

	// amp
	//

	// amplify
	//
	// Amplify domain association can be imported using app_id and domain_name: d2ypk4k47z8u6/example.com
	"aws_amplify_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .parameters.app_id }}/{{ .external_name }}"),

	// apprunner
	//
	// App Runner Custom Domain Associations can be imported by using the domain_name and service_arn separated by a comma (,)
	"aws_apprunner_custom_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .external_name }},{{ .parameters.service_arn }}"),

	// aws_appsync_domain_name can be imported using the AppSync domain name
	"aws_appsync_domain_name": config.ParameterAsIdentifier("domain_name"),
	// aws_appsync_domain_name_api_association can be imported using the AppSync domain name
	"aws_appsync_domain_name_api_association": config.ParameterAsIdentifier("domain_name"),

	// batch
	//
	// AWS Batch compute can be imported using the compute_environment_name
	"aws_batch_compute_environment": config.ParameterAsIdentifier("compute_environment_name"),
	// Batch Job Definition can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-definition/sample
	"aws_batch_job_definition": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-definition/{{ .external_name }}"),
	// Batch Job Queue can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-queue/sample
	"aws_batch_job_queue": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-queue/{{ .external_name }}"),

	// ce
	//
	// aws_ce_cost_category can be imported using the id
	"aws_ce_cost_category": config.IdentifierFromProvider,

	// cloudformation
	//
	// Cloudformation Stacks Instances imported using the StackSet name, target AWS account ID, and target AWS: example,123456789012,us-east-1
	"aws_cloudformation_stack_set_instance": config.IdentifierFromProvider,
	// aws_cloudformation_type can be imported with their type version Amazon Resource Name (ARN)
	"aws_cloudformation_type": config.IdentifierFromProvider,

	// cloudhsmv2
	//
	// CloudHSM v2 Clusters can be imported using the cluster id
	"aws_cloudhsm_v2_cluster": config.IdentifierFromProvider,
	// HSM modules can be imported using their HSM ID
	"aws_cloudhsm_v2_hsm": config.IdentifierFromProvider,

	// cloudwatchlogs
	//
	// CloudWatch Logs destinations can be imported using the name
	"aws_cloudwatch_log_destination": config.NameAsIdentifier,
	// CloudWatch Logs destination policies can be imported using the destination_name
	"aws_cloudwatch_log_destination_policy": config.ParameterAsIdentifier("destination_name"),
	// CloudWatch Logs subscription filter can be imported using the log group name and subscription filter name separated by |
	"aws_cloudwatch_log_subscription_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_group_name }}|{{ .external_name }}"),

	// codeartifact
	//
	// CodeArtifact Domain can be imported using the CodeArtifact Domain arn
	"aws_codeartifact_domain": config.IdentifierFromProvider,
	// CodeArtifact Domain Permissions Policies can be imported using the CodeArtifact Domain ARN
	"aws_codeartifact_domain_permissions_policy": config.IdentifierFromProvider,
	// CodeArtifact Repository can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository": config.IdentifierFromProvider,
	// CodeArtifact Repository Permissions Policies can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository_permissions_policy": config.IdentifierFromProvider,

	// codebuild
	//
	// CodeBuild Project can be imported using the name
	"aws_codebuild_project": config.NameAsIdentifier,
	// CodeBuild Report Group can be imported using the CodeBuild Report Group arn: arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
	"aws_codebuild_report_group": config.TemplatedStringAsIdentifier("name", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:report-group/{{ .external_name }}"),
	// CodeBuild Resource Policy can be imported using the CodeBuild Resource Policy arn
	"aws_codebuild_resource_policy": config.IdentifierFromProvider,
	// CodeBuild Source Credential can be imported using the CodeBuild Source Credential arn: arn:aws:codebuild:us-west-2:123456789:token:github
	"aws_codebuild_source_credential": config.TemplatedStringAsIdentifier("", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:token:{{ .parameters.token }}"),
	// CodeBuild Webhooks can be imported using the CodeBuild Project name
	"aws_codebuild_webhook": config.ParameterAsIdentifier("project_name"),

	// cognitoidp
	//
	// Cognito User Groups can be imported using the user_pool_id/name attributes concatenated
	"aws_cognito_user_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.user_pool_id }}/{{ .external_name }}"),
	// No import
	"aws_cognito_user_in_group": config.IdentifierFromProvider,

	// configservice
	//
	// Config aggregate authorizations can be imported using account_id:region
	"aws_config_aggregate_authorization": config.TemplatedStringAsIdentifier("", "{{ .parameters.account_id }}:{{ .parameters.region }}"),
	// Config Organization Conformance Packs can be imported using the name
	"aws_config_organization_conformance_pack": config.NameAsIdentifier,
	// Config Organization Custom Rules can be imported using the name
	"aws_config_organization_custom_rule": config.NameAsIdentifier,
	// Config Organization Managed Rules can be imported using the name
	"aws_config_organization_managed_rule": config.NameAsIdentifier,

	// connect
	//
	// Amazon Connect User Hierarchy Groups can be imported using the instance_id and hierarchy_group_id separated by a colon (:)
	"aws_connect_user_hierarchy_group": config.IdentifierFromProvider,

	// datapipeline
	//
	// aws_datapipeline_pipeline_definition can be imported using the id
	"aws_datapipeline_pipeline_definition": config.IdentifierFromProvider,

	// datasync
	//
	// aws_datasync_agent can be imported by using the DataSync Agent Amazon Resource Name (ARN)
	"aws_datasync_agent": config.IdentifierFromProvider,
	// aws_datasync_location_efs can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_efs": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_lustre_file_system can be imported by using the DataSync-ARN#FSx-Lustre-ARN
	"aws_datasync_location_fsx_lustre_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_openzfs_file_system can be imported by using the DataSync-ARN#FSx-openzfs-ARN
	"aws_datasync_location_fsx_openzfs_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_windows_file_system can be imported by using the DataSync-ARN#FSx-Windows-ARN
	"aws_datasync_location_fsx_windows_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_hdfs can be imported by using the Amazon Resource Name (ARN)
	"aws_datasync_location_hdfs": config.IdentifierFromProvider,
	// aws_datasync_location_nfs can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_nfs": config.IdentifierFromProvider,
	// aws_datasync_location_s3 can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_s3": config.IdentifierFromProvider,
	// aws_datasync_location_smb can be imported by using the Amazon Resource Name (ARN)
	"aws_datasync_location_smb": config.IdentifierFromProvider,
	// aws_datasync_task can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_task": config.IdentifierFromProvider,

	// directconnect
	//
	// No import
	"aws_dx_bgp_peer": config.IdentifierFromProvider,
	// No import
	"aws_dx_connection_confirmation": config.IdentifierFromProvider,
	// Direct Connect gateway associations can be imported using dx_gateway_id together with associated_gateway_id
	// TODO: associated_gateway_id parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_dx_gateway_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.dx_gateway_id }}/{{ .parameters.associated_gateway_id }}"),
	//
	"aws_dx_gateway_association_proposal": config.IdentifierFromProvider,
	// No import
	"aws_dx_hosted_connection": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface_accepter": config.ParameterAsIdentifier("virtual_interface_id"),
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface_accepter": config.ParameterAsIdentifier("virtual_interface_id"),
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface_accepter": config.ParameterAsIdentifier("virtual_interface_id"),
	// Direct Connect private virtual interfaces can be imported using the vif id
	"aws_dx_private_virtual_interface": config.IdentifierFromProvider,

	// dlm
	//
	// DLM lifecycle policies can be imported by their policy ID
	"aws_dlm_lifecycle_policy": config.IdentifierFromProvider,

	// dms
	//
	// Certificates can be imported using the certificate_id
	"aws_dms_certificate": config.ParameterAsIdentifier("certificate_id"),
	// Endpoints can be imported using the endpoint_id
	"aws_dms_endpoint": config.ParameterAsIdentifier("endpoint_id"),
	// Event subscriptions can be imported using the name
	"aws_dms_event_subscription": config.NameAsIdentifier,
	// Replication instances can be imported using the replication_instance_id
	"aws_dms_replication_instance": config.ParameterAsIdentifier("replication_instance_id"),
	// Replication subnet groups can be imported using the replication_subnet_group_id
	"aws_dms_replication_subnet_group": config.ParameterAsIdentifier("replication_subnet_group_id"),
	// Replication tasks can be imported using the replication_task_id
	"aws_dms_replication_task": config.ParameterAsIdentifier("replication_task_id"),

	// ds
	//
	// Conditional forwarders can be imported using the directory id and remote_domain_name: d-1234567890:example.com
	"aws_directory_service_conditional_forwarder": config.TemplatedStringAsIdentifier("", "{{ .parameters.directory_id }}:{{ .parameters.remote_domain_name }}"),
	// DirectoryService directories can be imported using the directory id
	"aws_directory_service_directory": config.IdentifierFromProvider,
	// Directory Service Log Subscriptions can be imported using the directory id
	"aws_directory_service_log_subscription": config.ParameterAsIdentifier("directory_id"),

	// ec2
	//
	// No import
	"aws_ami_from_instance": config.IdentifierFromProvider,
	//
	"aws_ec2_client_vpn_authorization_rule": config.IdentifierFromProvider,
	// AWS Client VPN endpoints can be imported using the id value found via aws ec2 describe-client-vpn-endpoints
	"aws_ec2_client_vpn_endpoint": config.IdentifierFromProvider,
	// AWS Client VPN network associations can be imported using the endpoint ID and the association ID. Values are separated by a ,
	"aws_ec2_client_vpn_network_association": config.IdentifierFromProvider,
	// AWS Client VPN routes can be imported using the endpoint ID, target subnet ID, and destination CIDR block. All values are separated by a ,
	"aws_ec2_client_vpn_route": config.TemplatedStringAsIdentifier("", "{{ .parameters.client_vpn_endpoint_id }},{{ .parameters.target_vpc_subnet_id }},{{ .parameters.destination_cidr_block }}"),
	// aws_ec2_fleet can be imported by using the Fleet identifier
	"aws_ec2_fleet": config.IdentifierFromProvider,
	// aws_ec2_local_gateway_route can be imported by using the EC2 Local Gateway Route Table identifier and destination CIDR block separated by underscores (_)
	"aws_ec2_local_gateway_route": config.TemplatedStringAsIdentifier("", "{{ .parameters.local_gateway_route_table_id }}_{{ .parameters.destination_cidr_block }}"),
	// aws_ec2_local_gateway_route_table_vpc_association can be imported by using the Local Gateway Route Table VPC Association identifier
	"aws_ec2_local_gateway_route_table_vpc_association": config.IdentifierFromProvider,
	// aws_ec2_tag can be imported by using the EC2 resource identifier and key, separated by a comma (,)
	"aws_ec2_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_id }}_{{ .parameters.key }}"),
	// Traffic mirror sessions can be imported using the id
	"aws_ec2_traffic_mirror_session": config.IdentifierFromProvider,
	// Traffic mirror targets can be imported using the id
	"aws_ec2_traffic_mirror_target": config.IdentifierFromProvider,
	// Internet Gateway Attachments can be imported using the id
	"aws_internet_gateway_attachment": config.IdentifierFromProvider,
	// No import
	"aws_network_acl_association": config.IdentifierFromProvider,
	// VPC Endpoint Services can be imported using ID of the connection, which is the VPC Endpoint Service ID and VPC Endpoint ID separated by underscore (_)
	"aws_vpc_endpoint_connection_accepter": config.TemplatedStringAsIdentifier("", "{{ .parameters.vpc_endpoint_service_id }}_{{ .parameters.vpc_endpoint_id }}"),
	// VPC Endpoint Policies can be imported using the id
	"aws_vpc_endpoint_policy": config.IdentifierFromProvider,
	// No import
	"aws_vpc_endpoint_security_group_association": config.IdentifierFromProvider,
	// IPAMs can be imported using the delegate account id
	"aws_vpc_ipam_organization_admin_account": config.ParameterAsIdentifier("delegated_admin_account_id"),
	// IPAMs can be imported using the <cidr>_<ipam-pool-id>
	"aws_vpc_ipam_pool_aws_default_network_acl": config.IdentifierFromProvider,
	// No import
	"aws_vpc_ipam_preview_next_cidr": config.IdentifierFromProvider,
	// aws_vpc_ipv6_cidr_block_association can be imported by using the VPC CIDR Association ID
	"aws_vpc_ipv6_cidr_block_association": config.IdentifierFromProvider,
	// VPC Peering Connection Accepters can be imported by using the Peering Connection ID
	"aws_vpc_peering_connection_accepter": config.ParameterAsIdentifier("vpc_peering_connection_id"),
	// VPC Peering Connection Options can be imported using the vpc peering id
	"aws_vpc_peering_connection_options": config.ParameterAsIdentifier("vpc_peering_connection_id"),

	// securityhub
	//
	// imported using the account ID
	"aws_securityhub_invite_accepter": FormattedIdentifierFromProvider("", "master_id"),
	// imported using the AWS account ID
	"aws_securityhub_organization_admin_account": FormattedIdentifierFromProvider("", "admin_account_id"),
	// imported using the AWS account ID
	// no Terraform argument specifies the AWS account ID and
	// Terraform resource ID is the AWS account ID for the resource
	"aws_securityhub_organization_configuration": config.IdentifierFromProvider,
	// no import documentation
	"aws_securityhub_standards_control": config.IdentifierFromProvider,

	// servicecatalog
	//
	// no import documentation
	"aws_servicecatalog_organizations_access": config.IdentifierFromProvider,
	// imported using the provisioned product ID,
	// which has provider-generated random parts:
	// pp-dnigbtea24ste
	"aws_servicecatalog_provisioned_product": config.IdentifierFromProvider,

	// servicediscovery
	//
	// imported using the namespace ID,
	// which has provider-generated random parts:
	// ns-1234567890
	"aws_service_discovery_http_namespace": config.IdentifierFromProvider,
	// imported using the service ID and instance ID:
	// 0123456789/i-0123
	"aws_service_discovery_instance": FormattedIdentifierFromProvider("/", "service_id", "instance_id"),

	// elasticache
	//
	// ElastiCache Security Groups can be imported by name
	"aws_elasticache_security_group": config.NameAsIdentifier,
	// ElastiCache Global Replication Groups can be imported using the global_replication_group_id,
	// which is an attribute reported in the state.
	// TODO: we need to check the value of a global_replication_group_id to
	// see if further normalization is possible
	"aws_elasticache_global_replication_group": config.IdentifierFromProvider,
	// ElastiCache user group associations can be imported using the user_group_id and user_id:
	// userGoupId1,userId
	"aws_elasticache_user_group_association": FormattedIdentifierFromProvider(",", "user_group_id", "user_id"),

	// ram
	//
	// RAM Principal Associations can be imported using their Resource Share ARN and the principal separated by a comma:
	// arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,123456789012
	"aws_ram_principal_association": FormattedIdentifierFromProvider(",", "resource_share_arn", "principal"),
	// RAM Resource Associations can be imported using their Resource Share ARN and Resource ARN separated by a comma:
	// arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,arn:aws:ec2:eu-west-1:123456789012:subnet/subnet-12345678
	"aws_ram_resource_association": FormattedIdentifierFromProvider(",", "resource_share_arn", "resource_arn"),
	// Resource shares can be imported using the arn of the resource share:
	// aws_ram_resource_share.example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12
	// TODO: validation may kick in, in which case we can use config.IdentifierFromProvider
	"aws_ram_resource_share": TemplatedStringAsIdentifierWithNoName("arn:aws:ram:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:resource-share/{{ .external_name }}"),
	// Resource share accepters can be imported using the resource share ARN:
	// arn:aws:ram:us-east-1:123456789012:resource-share/c4b56393-e8d9-89d9-6dc9-883752de4767
	"aws_ram_resource_share_accepter": FormattedIdentifierFromProvider("", "share_arn"),

	// sns
	//
	// SNS platform applications can be imported using the ARN:
	// arn:aws:sns:us-west-2:0123456789012:app/GCM/gcm_application
	"aws_sns_platform_application": config.TemplatedStringAsIdentifier("name", "arn:aws:sns:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:app/GCM/{{ .external_name }}"),
	// no import documentation is provided
	// TODO: we will need to check if normalization is possible
	"aws_sns_sms_preferences": config.IdentifierFromProvider,
	// SNS Topic Policy can be imported using the topic ARN:
	// arn:aws:sns:us-west-2:0123456789012:my-topic
	"aws_sns_topic_policy": FormattedIdentifierFromProvider("", "arn"),

	// ecs
	//
	// ECS Task Sets can be imported via the task_set_id, service, and cluster separated by commas (,):
	// ecs-svc/7177320696926227436,arn:aws:ecs:us-west-2:123456789101:service/example/example-1234567890,arn:aws:ecs:us-west-2:123456789101:cluster/example
	// TODO: validation may kick in, in which case we can use config.IdentifierFromProvider
	"aws_ecs_task_set": TemplatedStringAsIdentifierWithNoName("{{ .external_name }},{{ .parameters.service }},{{ .parameters.cluster }}"),

	// grafana
	//
	// Grafana workspace license association can be imported using the workspace's id
	"aws_grafana_license_association": FormattedIdentifierFromProvider("", "workspace_id"),

	// gamelift
	//
	// GameLift Game Server Group can be imported using the name
	"aws_gamelift_game_server_group": config.ParameterAsIdentifier("game_server_group_name"),
}
