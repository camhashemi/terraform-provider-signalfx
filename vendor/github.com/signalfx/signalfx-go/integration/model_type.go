/*
 * Integrations API
 *
 * APIs for creating, retrieving, updating, and deleting SignalFx integrations to the systems you use.<br> An integration provides SignalFx with information from the external system that you're connecting to. You'll need to retrieve this information from the external system before you use the API. Each external system is different, so to see a summary of its requirements and procedures, view its request body description. # Authentication To create, update, delete, or validate an integration, you need to authenticate your request using a session token associated with a SignalFx administrator. To **retrieve** an integration, your session token doesn't need to be associated with an administrator. You can also retrieve integrations using an org token.<br> In the web UI, session tokens are known as <strong>user access</strong> tokens, and org tokens are known as <strong>access tokens</strong>. <br> To learn more about authentication tokens, see the topic [Authentication Tokens](https://developers.signalfx.com/administration/access_tokens_overview.html) in the Developers Guide. # Supported service types SignalFx offers integrations for the following:<br>   * Data collection from other monitoring systems such as AWS CloudWatch   * Authentication using your existing Single Sign-On (**SSO**) system   * Sending alerts using your preferred messaging, chat, or incident management service <br> To use one of these integrations, you first register it with SignalFx. After that, you configure the integration to communicate between the system you're using and SignalFx. ## Data collection SignalFx integrations APIs support data collection for the following services:<br>   * Amazon Web Services (**AWS**)   * Google Cloud Platform (**GCP**)   * Microsoft Azure   * NewRelic  ## Authentication using SSO SignalFx integration APIs support SAML-based SSO integrations for the following services:<br>   * Microsoft Active Directory Federation Services (**ADFS**)   * Bitium   * Okta   * OneLogin   * PingOne  ## Alerts using message, chat, or incident management services SignalFx integration APIs support alert notifications using the following services: <br>   * BigPanda   * Office 365   * Opsgenie   * PagerDuty   * ServiceNow   * Slack   * VictorOps   * Webhook   * xMatters<br>  **NOTE:** You can't create Office 365 integrations using the API, and your ability to update them in a **PUT** request is limited, but you can retrieve their data or delete them. To create an Office 365 integration, use the the web UI. <br> # Viewing request body documentation The *request* body format for the following operations depends on the type of integration you use:<br>   * POST `/integration`   * PUT `/integration/{id}`<br>  The *response* body format for the following operations also depends on the type of integration you use:<br>   * GET `/integration`   * GET `/integration/{id}`  <br>  To see the request or response body format for an integration: <br>   1. Find the endpoint and method.   2. For a request body, find the section *REQUEST BODY SCHEMA*. For a     response body, find the section *RESPONSE SCHEMA*.   3. Scroll down to the `type` property.   4. At the end of the description for `type`, find the dropdown box that      contains the integration type. By default, it's set to *AWSCloudWatch*.   5. To see a complete list of integrations, click the down arrow. A list      with a vertical scroll bar appears.   6. Select the integration type from the list. The request body properties      for this integration type now appear.
 *
 * API version: 3.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package integration
// Type : Service that this integration object configures in SignalFx. This is an enumerated string that describes the service. SignalFx processes requests based on the type specified.<br> In addition, the \"type\" property controls OpenApi 3 validation. If you use a model for one service but specify the type for another service, OAS 3 validation programs may reject your request. <br> **For data collection services, the allowed values are:**<br>    * AWSCloudWatch   * Azure   * GCP   * NewRelic  **For SSO integrations, the allowed values are:**<br>   * ADFS   * AzureAD   * Bitium   * GoogleSaml   * Okta   * OneLogin   * PingOne  **For alerting services, the allowed values are:**<br>   * BigPanda   * Office365   * Opsgenie   * PagerDuty   * ServiceNow   * Slack   * VictorOps   * Webhook   * XMatters (note the capital \"X\")  ## Viewing request or response bodies To see the request or response body format, find it in the following dropdown box. To see a full list, click the down arrow and navigate with the vertical scrollbar.
type Type string

// List of Type
const (
	ADFS Type = "ADFS"
	AWS_CLOUD_WATCH Type = "AWSCloudWatch"
	AZURE Type = "Azure"
	AZURE_AD Type = "AzureAD"
	BITIUM Type = "Bitium"
	BIG_PANDA Type = "BigPanda"
	GCP Type = "GCP"
	GOOGLE_SAML Type = "GoogleSaml"
	OFFICE365 Type = "Office365"
	OKTA Type = "Okta"
	ONE_LOGIN Type = "OneLogin"
	OPSGENIE Type = "Opsgenie"
	PAGER_DUTY Type = "PagerDuty"
	PING_ONE Type = "PingOne"
	NEW_RELIC Type = "NewRelic"
	SERVICE_NOW Type = "ServiceNow"
	SLACK Type = "Slack"
	VICTOR_OPS Type = "VictorOps"
	WEBHOOK Type = "Webhook"
	X_MATTERS Type = "XMatters"
)