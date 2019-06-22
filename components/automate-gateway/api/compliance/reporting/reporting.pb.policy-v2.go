// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/compliance/reporting/reporting.proto

package reporting

import policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"

func init() {
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListReports", "compliance:reporting:reports", "compliance:reports:list", "POST", "/compliance/reporting/reports", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListReportIds", "compliance:reporting:report-ids", "compliance:reportids:list", "POST", "/compliance/reporting/report-ids", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ReadReport", "compliance:reporting:reports:{id}", "compliance:reports:get", "POST", "/compliance/reporting/reports/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListSuggestions", "compliance:reporting:suggestions", "compliance:reportSuggestions:list", "POST", "/compliance/reporting/suggestions", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*SuggestionRequest); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "type":
					return m.Type
				case "text":
					return m.Text
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListProfiles", "compliance:reporting:profiles", "compliance:reportProfiles:list", "POST", "/compliance/reporting/profiles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ReadNode", "compliance:reporting:nodes:{id}", "compliance:reportNodes:get", "GET", "/compliance/reporting/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListNodes", "compliance:reporting:nodes", "compliance:reportNodes:list", "POST", "/compliance/reporting/nodes/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/compliance/reporting/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/LicenseUsageNodes", "compliance:reporting:licenseusage", "compliance:reportingLicenseUsage:list", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
