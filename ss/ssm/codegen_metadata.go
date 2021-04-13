//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ praxisgen -metadata=ss/ssm/restful_doc -output=ss/ssm -pkg=ssm -target=1.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssm

import (
	"regexp"

	"gopkg.in/rightscale/rsc.v9/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"Execution": &metadata.Resource{
		Name: "Execution",
		Description: `An Execution is a launched instance of a CloudApp. Executions can be created from the catalog
by launching an Application, from Designer by launching a Template, or directly in Manager
by using the API and sending the CAT source or CAT Compiled source.
Executions are represented in RightScale Cloud Management by a deployment -- the resources
defined in the CAT are all created in the Deployment. Any action on a running CloudApp should
be made on its Execution resource.
Making changes to any resource directly in the CM deployment
may result in undesired behavior since the Execution only refreshes certain information as a
result of running an Operation on an Execution. For example, if a Server is replaced in CM
instead of through Self-Service, the new Server's information won' be available in
Self-Service.`,
		Identifier: "application/vnd.rightscale.self_service.execution",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "api_resources",
				FieldName: "ApiResources",
				FieldType: "[]*Resource",
			},

			&metadata.Attribute{
				Name:      "available_actions",
				FieldName: "AvailableActions",
				FieldType: "[]string",
			},

			&metadata.Attribute{
				Name:      "available_operations",
				FieldName: "AvailableOperations",
				FieldType: "[]*OperationDefinition",
			},

			&metadata.Attribute{
				Name:      "available_operations_info",
				FieldName: "AvailableOperationsInfo",
				FieldType: "[]*OperationInfo",
			},

			&metadata.Attribute{
				Name:      "compilation_href",
				FieldName: "CompilationHref",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "configuration_options",
				FieldName: "ConfigurationOptions",
				FieldType: "[]*ConfigurationOption",
			},

			&metadata.Attribute{
				Name:      "cost",
				FieldName: "Cost",
				FieldType: "*CostStruct",
			},

			&metadata.Attribute{
				Name:      "created_by",
				FieldName: "CreatedBy",
				FieldType: "*User",
			},

			&metadata.Attribute{
				Name:      "current_schedule",
				FieldName: "CurrentSchedule",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "dependencies",
				FieldName: "Dependencies",
				FieldType: "[]*CatDependency",
			},

			&metadata.Attribute{
				Name:      "deployment",
				FieldName: "Deployment",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "deployment_url",
				FieldName: "DeploymentUrl",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "description",
				FieldName: "Description",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "ends_at",
				FieldName: "EndsAt",
				FieldType: "*time.Time",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "latest_notification",
				FieldName: "LatestNotification",
				FieldType: "*Notification",
			},

			&metadata.Attribute{
				Name:      "latest_notifications",
				FieldName: "LatestNotifications",
				FieldType: "[]*Notification",
			},

			&metadata.Attribute{
				Name:      "launched_from",
				FieldName: "LaunchedFrom",
				FieldType: "*LaunchedFrom",
			},

			&metadata.Attribute{
				Name:      "launched_from_summary",
				FieldName: "LaunchedFromSummary",
				FieldType: "map[string]interface{}",
			},

			&metadata.Attribute{
				Name:      "links",
				FieldName: "Links",
				FieldType: "*ExecutionLinks",
			},

			&metadata.Attribute{
				Name:      "name",
				FieldName: "Name",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "next_action",
				FieldName: "NextAction",
				FieldType: "*ScheduledAction",
			},

			&metadata.Attribute{
				Name:      "outputs",
				FieldName: "Outputs",
				FieldType: "[]*Output",
			},

			&metadata.Attribute{
				Name:      "running_operations",
				FieldName: "RunningOperations",
				FieldType: "[]*Operation",
			},

			&metadata.Attribute{
				Name:      "schedule_required",
				FieldName: "ScheduleRequired",
				FieldType: "bool",
			},

			&metadata.Attribute{
				Name:      "scheduled",
				FieldName: "Scheduled",
				FieldType: "bool",
			},

			&metadata.Attribute{
				Name:      "schedules",
				FieldName: "Schedules",
				FieldType: "[]*Schedule",
			},

			&metadata.Attribute{
				Name:      "source",
				FieldName: "Source",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "status",
				FieldName: "Status",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List information about the Executions, or use a filter to only return certain Executions. A view can be used for various levels of detail.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/executions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by status, syntax is ["status==running"]`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of execution IDs to retrieve`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded", "index", "tiny"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by status, syntax is ["status==running"]`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of execution IDs to retrieve`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded", "index", "tiny"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show details for a given Execution. A view can be used for various levels of detail.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/executions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded", "source"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new execution from a CAT, a compiled CAT, an Application in the Catalog, or a Template in Designer`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "application_href",
						Description: `The href of the Application in Catalog from which to create the Execution. This attribute is mutually exclusive with: source, compiled_cat, compilation_href and template_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compilation_href",
						Description: `The href of the Compilation from which to create the Execution. This attribute is mutually exclusive with: source, compiled_cat, template_href and application_href. NOTE: This requires :designer role at least.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[cat_parser_gem_version]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[compiler_ver]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[conditions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[definitions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[imports]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[long_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[mappings]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[operations]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[outputs]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[package]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[parameters]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[permissions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[resources]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[rs_ca_ver]",
						Description: ``,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[short_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[source]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "current_schedule",
						Description: `The currently selected schedule name, or nil for CloudApps using the '24/7' schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "defer_launch",
						Description: `Whether or not to defer launching the execution. Setting this value to true will keep the execution in not_started state until it is explicitly launched or the first scheduled start operation occurs.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `The description for the execution. The short_description of the Template will be used if none is provided.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ends_at",
						Description: `The day on which the CloudApp should be automatically terminated`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name for the Execution. The Template name will be used if none is provided. This will be used as the name of the deployment (appended with a unique ID).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][name]",
						Description: `Name of configuration option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][type]",
						Description: `Type of configuration option.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"string", "number", "list"},
					},
					&metadata.ActionParam{
						Name:        "options[][value]",
						Description: `Configuration option value, a string, integer or array of strings depending on type`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedule_required",
						Description: `Whether the CloudApp requires a schedule. If set to false, allows user to pick from '24/7' schedule in the UI`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][action]",
						Description: `The name of the action to be run. When the value is "run", the "operation" struct should contain the name of the operation to run as well as any options needed by the operation.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"launch", "start", "stop", "terminate", "run"},
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][email]",
						Description: `User email`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][id]",
						Description: `User id`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][compilation_href]",
						Description: `The HREF of the compilation used to create this execution`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][cost][unit]",
						Description: `Currency used for the cost value`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][cost][updated_at]",
						Description: `Timestamp of last cost refresh`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][cost][value]",
						Description: `Amount of instance usage in CloudApp deployment, only available roughly 24 hours after launch, empty if not available`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][email]",
						Description: `User email`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][id]",
						Description: `User id`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][current_schedule]",
						Description: `The currently selected schedule name, or nil for CloudApps using the '24/7' schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][deployment]",
						Description: `CloudApp deployment href`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][deployment_url]",
						Description: `URL of the CloudApp deployment in the Cloud Management Dashboard`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][description]",
						Description: `Description of execution`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][ends_at]",
						Description: `The time of the next 'terminate' ScheduledAction (if any).`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][href]",
						Description: `Execution href`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][id]",
						Description: `User id`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][kind]",
						Description: `The kind of this resource, always self_service#execution`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][latest_notification][category]",
						Description: `Notification category, info or error`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"info", "error", "status_update"},
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][href]",
						Description: `Execution href`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][id]",
						Description: `User id`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][kind]",
						Description: `The kind of this resource, always self_service#execution`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][latest_notification][message]",
						Description: `Notification content`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][latest_notification][read]",
						Description: `Whether notification was marked as read (not currently used)`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][latest_notification][timestamps][created_at]",
						Description: `Creation timestamp`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][launched_from][type]",
						Description: `The type of the value (one of: application, template, compiled_cat, source, compilation`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][cost][value]",
						Description: `Amount of instance usage in CloudApp deployment, only available roughly 24 hours after launch, empty if not available`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][launched_from_summary]",
						Description: `How the CloudApp was launched, either from Application, Template, source, or compiled_cat`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][action]",
						Description: `The name of the action to be run. When the value is "run", the "operation" struct should contain the name of the operation to run as well as any options needed by the operation.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"launch", "start", "stop", "terminate", "run"},
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][email]",
						Description: `User email`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][id]",
						Description: `User id`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][execution_schedule]",
						Description: `Indicates ScheduledActions that were created by the system as part of an execution schedule.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][first_occurrence]",
						Description: `The time and day of the first occurrence when the action will be ran, similar to the "DTSTART" property specified by iCal. Used (in conjunction with timezone) to determine the time of day for the "next_occurrence". Can be set to the future or past. DateTimes should be passed as ISO-8601 formatted time strings.  All DateTimes are converted to UTC when returned.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][href]",
						Description: `Execution href`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][id]",
						Description: `User id`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][kind]",
						Description: `The kind of this resource, always self_service#execution`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][next_occurrence]",
						Description: `The Date/Time for the next occurrence. Since "DateTime implies a timezone offset (but no DST preference), the "timezone" parameter will be used to determine the DST preference. All DateTimes are converted to UTC when returned.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][recurrence]",
						Description: `iCal recurrence rule (RRULE) as described by RFC 5545. Expresses the days on which the action will be run. Optionally a "last occurrence" date can be set by passing the iCal "UNTIL" parameter in the rule (date-time must be passed in ISO-8601 format). If omitted, the action will only be run once, on the "first_occurrence".`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][recurrence_description]",
						Description: `Read-only attribute that gets automatically generated from the recurrence definition`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][latest_notification][timestamps][created_at]",
						Description: `Creation timestamp`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][timezone]",
						Description: `The timezone in which the "first_occurrence" and "next_occurrence" times will be interpreted. Used to determine when Daylight Savings Time changes occur. Supports standardized "tzinfo" names [found here](http://www.iana.org/time-zones).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][schedule_required]",
						Description: `Whether the CloudApp requires a schedule. If set to false, allows user to pick from '24/7' schedule in the UI`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][scheduled]",
						Description: `Indicates whether or not an execution has a scheduled start action`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][source]",
						Description: `Original CAT source`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][status]",
						Description: `Execution status.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"not_started", "launching", "starting", "enabling", "running", "disabling", "disabled", "terminating", "stopping", "waiting_for_operations", "canceling_operations", "stopped", "terminated", "failed", "provisioning", "decommissioning", "decommissioned"},
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][latest_notification][timestamps][created_at]",
						Description: `Creation timestamp`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][execution_schedule]",
						Description: `Indicates ScheduledActions that were created by the system as part of an execution schedule.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][first_occurrence]",
						Description: `The time and day of the first occurrence when the action will be ran, similar to the "DTSTART" property specified by iCal. Used (in conjunction with timezone) to determine the time of day for the "next_occurrence". Can be set to the future or past. DateTimes should be passed as ISO-8601 formatted time strings.  All DateTimes are converted to UTC when returned.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][href]",
						Description: `Execution href`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][id]",
						Description: `User id`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][kind]",
						Description: `The kind of this resource, always self_service#execution`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][next_occurrence]",
						Description: `The Date/Time for the next occurrence. Since "DateTime implies a timezone offset (but no DST preference), the "timezone" parameter will be used to determine the DST preference. All DateTimes are converted to UTC when returned.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][created_by][name]",
						Description: `User name, usually of the form "First Last"`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][recurrence]",
						Description: `iCal recurrence rule (RRULE) as described by RFC 5545. Expresses the days on which the action will be run. Optionally a "last occurrence" date can be set by passing the iCal "UNTIL" parameter in the rule (date-time must be passed in ISO-8601 format). If omitted, the action will only be run once, on the "first_occurrence".`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][recurrence_description]",
						Description: `Read-only attribute that gets automatically generated from the recurrence definition`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][latest_notification][timestamps][created_at]",
						Description: `Creation timestamp`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions[][execution][next_action][timezone]",
						Description: `The timezone in which the "first_occurrence" and "next_occurrence" times will be interpreted. Used to determine when Daylight Savings Time changes occur. Supports standardized "tzinfo" names [found here](http://www.iana.org/time-zones).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][created_from]",
						Description: `optional HREF of the Schedule resource used to create this schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][description]",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][name]",
						Description: `The name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `The raw CAT source from which to create the Execution. The CAT will be compiled first and then launched if successful. This attribute is mutually exclusive with: compiled_cat, template_href, compilation_href and application_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `The href of the Template in Designer from which to create the Execution. This attribute is mutually exclusive with: source, compiled_cat, compilation_href and application_href. NOTE: This requires :designer role at least.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "application_href",
						Description: `The href of the Application in Catalog from which to create the Execution. This attribute is mutually exclusive with: source, compiled_cat, compilation_href and template_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compilation_href",
						Description: `The href of the Compilation from which to create the Execution. This attribute is mutually exclusive with: source, compiled_cat, template_href and application_href. NOTE: This requires :designer role at least.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat",
						Description: `The compiled CAT source from which to create the Execution. This attribute is mutually exclusive with: source, template_href and application_href.`,
						Type:        "*CompiledCAT",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "current_schedule",
						Description: `The currently selected schedule name, or nil for CloudApps using the '24/7' schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "defer_launch",
						Description: `Whether or not to defer launching the execution. Setting this value to true will keep the execution in not_started state until it is explicitly launched or the first scheduled start operation occurs.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `The description for the execution. The short_description of the Template will be used if none is provided.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ends_at",
						Description: `The day on which the CloudApp should be automatically terminated`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name for the Execution. The Template name will be used if none is provided. This will be used as the name of the deployment (appended with a unique ID).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options",
						Description: `The configuration options of the Execution. These are the values provided for the CloudApp parameters.`,
						Type:        "[]*ConfigurationOption",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedule_required",
						Description: `Whether the CloudApp requires a schedule. If set to false, allows user to pick from '24/7' schedule in the UI`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scheduled_actions",
						Description: `The inital ScheduledActions to apply to the Execution.`,
						Type:        "[]*ScheduledActionParam",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules",
						Description: `The schedules available to the CloudApp`,
						Type:        "[]*Schedule",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `The raw CAT source from which to create the Execution. The CAT will be compiled first and then launched if successful. This attribute is mutually exclusive with: compiled_cat, template_href, compilation_href and application_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `The href of the Template in Designer from which to create the Execution. This attribute is mutually exclusive with: source, compiled_cat, compilation_href and application_href. NOTE: This requires :designer role at least.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "patch",
				Description: `Updates an execution end date or selected schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/api/manager/projects/%s/executions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "current_schedule",
						Description: `The name of the schedule to select, or nil to use the '24/7' schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ends_at",
						Description: `The day on which the CloudApp should be automatically terminated`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "current_schedule",
						Description: `The name of the schedule to select, or nil to use the '24/7' schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ends_at",
						Description: `The day on which the CloudApp should be automatically terminated`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `No description provided for delete.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/manager/projects/%s/executions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "force",
						Description: `Force delete execution, bypassing state checks (only available to designers and admins).
Note: using this option only deletes the CloudApp from Self-Service and does not modify or terminate resources in any way. Any cloud resources running must be manually destroyed.`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "force",
						Description: `Force delete execution, bypassing state checks (only available to designers and admins).
Note: using this option only deletes the CloudApp from Self-Service and does not modify or terminate resources in any way. Any cloud resources running must be manually destroyed.`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete several executions from the database. Note: if an execution has not successfully been terminated, there may still be associated cloud resources running.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/manager/projects/%s/executions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "force",
						Description: `Force delete execution, bypassing state checks (only available to designers and admins).
Note: using this option only deletes the CloudApp from Self-Service and does not modify or terminate resources in any way. Any cloud resources running must be manually destroyed.`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "force",
						Description: `Force delete execution, bypassing state checks (only available to designers and admins).
Note: using this option only deletes the CloudApp from Self-Service and does not modify or terminate resources in any way. Any cloud resources running must be manually destroyed.`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "download",
				Description: `Download the CAT source for the execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/executions/%s/download",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)/download`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_version",
						Description: `The API version (only valid value is currently "1.0")`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_version",
						Description: `The API version (only valid value is currently "1.0")`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "launch",
				Description: `Launch an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/%s/actions/launch",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)/actions/launch`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "start",
				Description: `Start an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/%s/actions/start",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)/actions/start`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "stop",
				Description: `Stop an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/%s/actions/stop",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)/actions/stop`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "terminate",
				Description: `Terminate an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/%s/actions/terminate",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)/actions/terminate`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "multi_launch",
				Description: `Launch several Executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/actions/launch",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/actions/launch`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to launch`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to launch`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_start",
				Description: `Start several Executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/actions/start",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/actions/start`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to start`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to start`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_stop",
				Description: `Stop several Executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/actions/stop",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/actions/stop`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to stop`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to stop`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_terminate",
				Description: `Terminate several Executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/actions/terminate",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/actions/terminate`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to terminate`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to terminate`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "run",
				Description: `Runs an Operation on an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/%s/actions/run",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/([^/]+)/actions/run`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "configuration_options[][name]",
						Description: `Name of configuration option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "configuration_options[][type]",
						Description: `Type of configuration option.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"string", "number", "list"},
					},
					&metadata.ActionParam{
						Name:        "configuration_options[][value]",
						Description: `Configuration option value, a string, integer or array of strings depending on type`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the operation to run`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "configuration_options",
						Description: `The configuration options of the operation. These are the values provided for the CloudApp parameters that this operation depends on.`,
						Type:        "[]*ConfigurationOption",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the operation to run`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_run",
				Description: `Runs an Operation on several Executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/executions/actions/run",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/executions/actions/run`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to run`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "configuration_options[][name]",
						Description: `Name of configuration option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "configuration_options[][type]",
						Description: `Type of configuration option.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"string", "number", "list"},
					},
					&metadata.ActionParam{
						Name:        "configuration_options[][value]",
						Description: `Configuration option value, a string, integer or array of strings depending on type`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the operation to run`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to run`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "configuration_options",
						Description: `The configuration options of the operation. These are the values provided for the CloudApp parameters that this operation depends on.`,
						Type:        "[]*ConfigurationOption",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the operation to run`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
		Links: map[string]string{
			"latest_notifications": "",
			"running_operations":   "",
		},
	},
	"Notification": &metadata.Resource{
		Name: "Notification",
		Description: `The Notification resource represents a system notification that an action has occurred. Generally
these Notifications are the start and completion of Operations. Currently notifications are only
available via the API/UI and are not distributed externally to users.`,
		Identifier: "application/vnd.rightscale.self_service.notification",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "category",
				FieldName: "Category",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "execution",
				FieldName: "Execution",
				FieldType: "*Execution",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "links",
				FieldName: "Links",
				FieldType: "*NotificationLinks",
			},

			&metadata.Attribute{
				Name:      "message",
				FieldName: "Message",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "read",
				FieldName: "Read",
				FieldType: "bool",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the most recent 50 Notifications. Use the filter parameter to specify specify Executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/notifications",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/notifications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The Notification IDs to return`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The Notification IDs to return`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Get details for a specific Notification`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/notifications/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/notifications/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
		Links: map[string]string{
			"execution": "",
		},
	},
	"Operation": &metadata.Resource{
		Name: "Operation",
		Description: `Operations represent actions that can be taken on an Execution.
When a CloudApp is launched, a sequence of Operations is run as [explained here](http://docs.rightscale.com/ss/reference/ss_CAT_file_language.html#operations) in the Operations section
While a CloudApp is running, users may launch any custom Operations as defined in the CAT.
Once a CAT is Terminated, a sequence of Operations is run as [explained here](http://docs.rightscale.com/ss/reference/ss_CAT_file_language.html#operations) in the Operations section`,
		Identifier: "application/vnd.rightscale.self_service.operation",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "configuration_options",
				FieldName: "ConfigurationOptions",
				FieldType: "[]*ConfigurationOption",
			},

			&metadata.Attribute{
				Name:      "created_by",
				FieldName: "CreatedBy",
				FieldType: "*User",
			},

			&metadata.Attribute{
				Name:      "execution",
				FieldName: "Execution",
				FieldType: "*Execution",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "label",
				FieldName: "Label",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "links",
				FieldName: "Links",
				FieldType: "*OperationLinks",
			},

			&metadata.Attribute{
				Name:      "name",
				FieldName: "Name",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "status",
				FieldName: "Status",
				FieldType: "*StatusStruct",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Get the list of 50 most recent Operations (usually filtered by Execution).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/operations",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/operations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution ID or status`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `IDs of operations to filter on`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `The maximum number of operations to retrieve. The maximum (and default) limit is 50.If a limit of more than 50 is specified, only 50 operations will be returned`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution ID or status`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `IDs of operations to filter on`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `The maximum number of operations to retrieve. The maximum (and default) limit is 50.If a limit of more than 50 is specified, only 50 operations will be returned`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Get the details for a specific Operation`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/operations/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/operations/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Trigger an Operation to run by specifying the Execution ID and the name of the Operation.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/operations",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/operations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "execution_id",
						Description: `The Execution ID on which to run the operation`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the operation to run`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][name]",
						Description: `Name of configuration option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][type]",
						Description: `Type of configuration option.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"string", "number", "list"},
					},
					&metadata.ActionParam{
						Name:        "options[][value]",
						Description: `Configuration option value, a string, integer or array of strings depending on type`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "execution_id",
						Description: `The Execution ID on which to run the operation`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the operation to run`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options",
						Description: `The configuration options of the operation. These are the values provided for the CloudAPP parameters that this operation depends on.`,
						Type:        "[]*ConfigurationOption",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
		Links: map[string]string{
			"execution": "",
		},
	},
	"ScheduledAction": &metadata.Resource{
		Name: "ScheduledAction",
		Description: `ScheduledActions describe a set of timed occurrences for an action to be run (at most once per day).
Recurrence Rules are based off of the [RFC 5545](https://tools.ietf.org/html/rfc5545) iCal spec, and timezones are from the standard [tzinfo database](http://www.iana.org/time-zones).
All DateTimes must be passed in [ISO-8601 format](https://en.wikipedia.org/wiki/ISO_8601)`,
		Identifier: "application/vnd.rightscale.self_service.scheduled_action",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "action",
				FieldName: "Action",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "created_by",
				FieldName: "CreatedBy",
				FieldType: "*User",
			},

			&metadata.Attribute{
				Name:      "execution",
				FieldName: "Execution",
				FieldType: "*Execution",
			},

			&metadata.Attribute{
				Name:      "execution_schedule",
				FieldName: "ExecutionSchedule",
				FieldType: "bool",
			},

			&metadata.Attribute{
				Name:      "first_occurrence",
				FieldName: "FirstOccurrence",
				FieldType: "*time.Time",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "links",
				FieldName: "Links",
				FieldType: "*ScheduledActionLinks",
			},

			&metadata.Attribute{
				Name:      "name",
				FieldName: "Name",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "next_occurrence",
				FieldName: "NextOccurrence",
				FieldType: "*time.Time",
			},

			&metadata.Attribute{
				Name:      "operation",
				FieldName: "Operation",
				FieldType: "*OperationStruct",
			},

			&metadata.Attribute{
				Name:      "recurrence",
				FieldName: "Recurrence",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "recurrence_description",
				FieldName: "RecurrenceDescription",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},

			&metadata.Attribute{
				Name:      "timezone",
				FieldName: "Timezone",
				FieldType: "string",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List ScheduledAction resources in the project. The list can be filtered to a given execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/scheduled_actions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/scheduled_actions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by execution id or execution creator (user) id.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by execution id or execution creator (user) id.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Retrieve given ScheduledAction resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/manager/projects/%s/scheduled_actions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/scheduled_actions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new ScheduledAction resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/scheduled_actions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/scheduled_actions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "action",
						Description: `The name of the action to be run. When the value is "run", the "operation" struct should contain the name of the operation to run as well as any options needed by the operation.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						ValidValues: []string{"launch", "start", "stop", "terminate", "run"},
					},
					&metadata.ActionParam{
						Name:        "execution_id",
						Description: `Id of the Execuion.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "first_occurrence",
						Description: `The time and day of the first occurrence when the action will be ran, similar to the "DTSTART" property specified by iCal. Used (in conjunction with timezone) to determine the time of day for the "next_occurrence". Can be set to the future or past. DateTimes should be passed as ISO-8601 formatted time strings.  All DateTimes are converted to UTC when returned.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The human-readable name for the ScheduledAction.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "operation[configuration_options][][name]",
						Description: `Name of configuration option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "operation[configuration_options][][type]",
						Description: `Type of configuration option.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"string", "number", "list"},
					},
					&metadata.ActionParam{
						Name:        "operation[configuration_options][][value]",
						Description: `Configuration option value, a string, integer or array of strings depending on type`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "operation[configuration_options][][name]",
						Description: `Name of configuration option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "recurrence",
						Description: `iCal recurrence rule (RRULE) as described by RFC 5545. Expresses the days on which the action will be run. Optionally a "last occurrence" date can be set by passing the iCal "UNTIL" parameter in the rule (date-time must be passed in ISO-8601 format). If omitted, the action will only be run once, on the "first_occurrence".`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "timezone",
						Description: `The timezone in which the "first_occurrence" and "next_occurrence" times will be interpreted. Used to determine when Daylight Savings Time changes occur. Supports standardized "tzinfo" names [found here](http://www.iana.org/time-zones).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "action",
						Description: `The name of the action to be run. When the value is "run", the "operation" struct should contain the name of the operation to run as well as any options needed by the operation.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						ValidValues: []string{"launch", "start", "stop", "terminate", "run"},
					},
					&metadata.ActionParam{
						Name:        "execution_id",
						Description: `Id of the Execuion.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "first_occurrence",
						Description: `The time and day of the first occurrence when the action will be ran, similar to the "DTSTART" property specified by iCal. Used (in conjunction with timezone) to determine the time of day for the "next_occurrence". Can be set to the future or past. DateTimes should be passed as ISO-8601 formatted time strings.  All DateTimes are converted to UTC when returned.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The human-readable name for the ScheduledAction.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "operation",
						Description: `When scheduling a "run" action, contains details on the operation to run`,
						Type:        "*OperationStruct",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "recurrence",
						Description: `iCal recurrence rule (RRULE) as described by RFC 5545. Expresses the days on which the action will be run. Optionally a "last occurrence" date can be set by passing the iCal "UNTIL" parameter in the rule (date-time must be passed in ISO-8601 format). If omitted, the action will only be run once, on the "first_occurrence".`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "timezone",
						Description: `The timezone in which the "first_occurrence" and "next_occurrence" times will be interpreted. Used to determine when Daylight Savings Time changes occur. Supports standardized "tzinfo" names [found here](http://www.iana.org/time-zones).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "patch",
				Description: `Updates the 'next_occurrence' property of a ScheduledAction.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/api/manager/projects/%s/scheduled_actions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/scheduled_actions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "next_occurrence",
						Description: `The Date/Time for the next occurrence, useful for delaying a single occurrence. DateTimes should be passed as ISO-8601 formatted time strings.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "next_occurrence",
						Description: `The Date/Time for the next occurrence, useful for delaying a single occurrence. DateTimes should be passed as ISO-8601 formatted time strings.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete a ScheduledAction.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/manager/projects/%s/scheduled_actions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/scheduled_actions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "skip",
				Description: `Skips the requested number of ScheduledAction occurrences. If no count is provided, one occurrence is skipped. On success, the next_occurrence view of the updated ScheduledAction is returned.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/manager/projects/%s/scheduled_actions/%s/actions/skip",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/api/manager/projects/([^/]+)/scheduled_actions/([^/]+)/actions/skip`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "count",
						Description: `The number of scheduled occurrences to skip. If not provided, the default count is 1.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "count",
						Description: `The number of scheduled occurrences to skip. If not provided, the default count is 1.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
		Links: map[string]string{
			"execution": "",
		},
	},
}
