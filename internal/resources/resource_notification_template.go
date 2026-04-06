package resources

import (
	"context"
	"errors"
	"regexp"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Values match notification-srv / cpentity.CommMethod JSON (lowercase).
var notificationTemplateCommMethods = []string{"email", "sms", "ivr", "push"}

// Values match notify.MessageFormat JSON (lowercase).
var notificationTemplateMessageFormats = []string{"html", "text", "media"}

// Values match basetype.ObjectOwner for templates (lowercase).
var notificationTemplateOwners = []string{"client", "admin", "core", "system"}

type NotificationTemplateResource struct {
	BaseResource
}

func NewNotificationTemplateResource() resource.Resource {
	return &NotificationTemplateResource{
		BaseResource: NewBaseResource(
			BaseResourceConfig{
				Name:   RESOURCE_NOTIFICATION_TEMPLATE,
				Schema: &notificationTemplateSchema,
			},
		),
	}
}

type notificationTemplateModel struct {
	ID                  types.String `tfsdk:"id"`
	GroupID             types.String `tfsdk:"group_id"`
	TemplateKey         types.String `tfsdk:"template_key"`
	CommunicationMethod types.String `tfsdk:"communication_method"`
	Locale              types.String `tfsdk:"locale"`
	MessageFormat       types.String `tfsdk:"message_format"`
	Description         types.String `tfsdk:"description"`
	Subject             types.String `tfsdk:"subject"`
	Content             types.String `tfsdk:"content"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	Owner               types.String `tfsdk:"owner"`
	ProcessingType      types.String `tfsdk:"processing_type"`
	UsageType           types.String `tfsdk:"usage_type"`
	VerificationType    types.String `tfsdk:"verification_type"`
	Number              types.Int64  `tfsdk:"number"`
	UserGroupIDs        types.Set    `tfsdk:"user_group_ids"`
}

var notificationTemplateSchema = schema.Schema{
	MarkdownDescription: "Manages a **notification template** via **notification-srv** (`/templates`). " +
		"This is separate from `cidaas_template` (legacy `templates-srv`).\n\n" +
		"**Import:** `terraform import cidaas_notification_template.NAME <template_document_id>`",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Document id of the template (`_id` from the API).",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"group_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Template group id (e.g. `default`, `developer`).",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"template_key": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Template type key (template type id).",
			Validators: []validator.String{
				stringvalidator.RegexMatches(
					regexp.MustCompile(`^[A-Z0-9_-]+$`),
					"must be uppercase letters, digits, underscores, or hyphens",
				),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"communication_method": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Channel: `email`, `sms`, `ivr`, or `push` (notification-srv JSON values).",
			Validators: []validator.String{
				stringvalidator.OneOf(notificationTemplateCommMethods...),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"locale": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "BCP47 locale (e.g. `en`, `de`).",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"message_format": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Message format: `html`, `text`, or `media` (notification-srv `messageFormat`).",
			Validators: []validator.String{
				stringvalidator.OneOf(notificationTemplateMessageFormats...),
			},
		},
		"description": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Description of the template (10–600 characters per API validation).",
			Validators: []validator.String{
				stringvalidator.LengthBetween(10, 600),
			},
		},
		"subject": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Email subject. Required for `email` channel in practice.",
		},
		"content": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Template body (HTML/text per message_format).",
		},
		"enabled": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
			MarkdownDescription: "Whether the template is enabled.",
		},
		"owner": schema.StringAttribute{
			Optional: true,
			Computed: true,
			MarkdownDescription: "Object owner: `client`, `admin`, `core`, or `system`. " +
				"Omit to adopt the value returned by the API (e.g. seeded `default` templates are often `admin`). " +
				"If set, apply returns that value even when the API stores another owner for system templates.",
			Validators: []validator.String{
				stringvalidator.OneOf(notificationTemplateOwners...),
			},
		},
		"processing_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Processing type when required by the template type (e.g. `CODE`, `LINK`).",
		},
		"usage_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Usage type when required by the template type.",
		},
		"verification_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Verification type when required by the template type.",
		},
		"number": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Optional template number suffix in id generation.",
		},
		"user_group_ids": schema.SetAttribute{
			ElementType:         types.StringType,
			Optional:            true,
			MarkdownDescription: "Restrict template access to these user groups.",
		},
	},
}

func (r *NotificationTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan notificationTemplateModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	apiReq, diags := notificationTemplateToAPI(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	docID := cidaas.SyntheticTemplateDocumentID(apiReq)
	var res *cidaas.NotificationsSrvTemplateModel
	var err error
	if docID != "" {
		// Seeded templates (e.g. default group) already exist → update. New developer-only rows → GET 404 then POST.
		existing, getErr := r.cidaasClient.NotificationsSrvTemplate.GetAllowNotFound(ctx, docID)
		if getErr != nil {
			resp.Diagnostics.AddError("failed to create notification template", util.FormatErrorMessage(getErr))
			return
		}
		if existing != nil {
			res, err = r.cidaasClient.NotificationsSrvTemplate.Update(ctx, docID, apiReq)
		} else {
			res, err = r.cidaasClient.NotificationsSrvTemplate.Create(ctx, apiReq)
			if err != nil && errors.Is(err, cidaas.ErrNotificationTemplateAlreadyExists) {
				res, err = r.cidaasClient.NotificationsSrvTemplate.Update(ctx, docID, apiReq)
			}
		}
	} else {
		res, err = r.cidaasClient.NotificationsSrvTemplate.Create(ctx, apiReq)
		if err != nil && errors.Is(err, cidaas.ErrNotificationTemplateAlreadyExists) {
			resp.Diagnostics.AddError("failed to create notification template", "template already exists but document id could not be computed; use terraform import with the existing _id")
			return
		}
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to create notification template", util.FormatErrorMessage(err))
		return
	}
	state := notificationTemplateFromAPI(res)
	state.Owner = notificationTemplateOwnerAfterApply(plan.Owner, res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NotificationTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state notificationTemplateModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.NotificationsSrvTemplate.Get(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read notification template", util.FormatErrorMessage(err))
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("notification template not found", state.ID.ValueString())
		return
	}
	out := notificationTemplateFromAPI(res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &out)...)
}

func (r *NotificationTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state notificationTemplateModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	apiReq, diags := notificationTemplateToAPI(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.NotificationsSrvTemplate.Update(ctx, plan.ID.ValueString(), apiReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to update notification template", util.FormatErrorMessage(err))
		return
	}
	newState := notificationTemplateFromAPI(res)
	newState.Owner = notificationTemplateOwnerAfterApply(plan.Owner, res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *NotificationTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state notificationTemplateModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	privilegedNoOp, err := r.cidaasClient.NotificationsSrvTemplate.Delete(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete notification template", util.FormatErrorMessage(err))
		return
	}
	if privilegedNoOp {
		resp.Diagnostics.AddWarning(
			"notification template not removed in Cidaas",
			"The API does not allow deleting this template (e.g. seeded or system-owned). Terraform state was cleared anyway; apply will recreate or adopt the same template by id on the next create.",
		)
	}
}

func (r *NotificationTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func notificationTemplateToAPI(ctx context.Context, m notificationTemplateModel) (cidaas.NotificationsSrvTemplateModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	out := cidaas.NotificationsSrvTemplateModel{
		ID:                  m.ID.ValueString(),
		GroupID:             m.GroupID.ValueString(),
		TemplateKey:         m.TemplateKey.ValueString(),
		CommunicationMethod: m.CommunicationMethod.ValueString(),
		Locale:              m.Locale.ValueString(),
		MessageFormat:       m.MessageFormat.ValueString(),
		Description:         m.Description.ValueString(),
		Content:             m.Content.ValueString(),
		Enabled:             m.Enabled.ValueBool(),
	}
	if !m.Owner.IsNull() {
		out.Owner = m.Owner.ValueString()
	}
	if !m.Subject.IsNull() {
		out.Subject = m.Subject.ValueString()
	}
	if !m.ProcessingType.IsNull() {
		out.ProcessingType = m.ProcessingType.ValueString()
	}
	if !m.UsageType.IsNull() {
		out.UsageType = m.UsageType.ValueString()
	}
	if !m.VerificationType.IsNull() {
		out.VerificationType = m.VerificationType.ValueString()
	}
	if !m.Number.IsNull() {
		n := int(m.Number.ValueInt64())
		out.Number = &n
	}
	if !m.UserGroupIDs.IsNull() && !m.UserGroupIDs.IsUnknown() {
		diags.Append(m.UserGroupIDs.ElementsAs(ctx, &out.UserGroupIDs, false)...)
	}
	return out, diags
}

// notificationTemplateFromAPI maps API JSON to Terraform state. Owner is filled from the API; callers of
// Create/Update then overwrite m.Owner with notificationTemplateOwnerAfterApply so apply matches the plan.
func notificationTemplateFromAPI(data *cidaas.NotificationsSrvTemplateModel) notificationTemplateModel {
	m := notificationTemplateModel{
		ID:                  util.StringValueOrNull(&data.ID),
		GroupID:             util.StringValueOrNull(&data.GroupID),
		TemplateKey:         util.StringValueOrNull(&data.TemplateKey),
		CommunicationMethod: types.StringValue(data.CommunicationMethod),
		Locale:              types.StringValue(data.Locale),
		MessageFormat:       types.StringValue(data.MessageFormat),
		Description:         types.StringValue(data.Description),
		Content:             types.StringValue(data.Content),
		Enabled:             types.BoolValue(data.Enabled),
	}
	if data.Subject != "" {
		m.Subject = types.StringValue(data.Subject)
	} else {
		m.Subject = types.StringNull()
	}
	if data.Owner != "" {
		m.Owner = types.StringValue(data.Owner)
	} else {
		m.Owner = types.StringNull()
	}
	if data.ProcessingType != "" {
		m.ProcessingType = types.StringValue(data.ProcessingType)
	} else {
		m.ProcessingType = types.StringNull()
	}
	if data.UsageType != "" {
		m.UsageType = types.StringValue(data.UsageType)
	} else {
		m.UsageType = types.StringNull()
	}
	if data.VerificationType != "" {
		m.VerificationType = types.StringValue(data.VerificationType)
	} else {
		m.VerificationType = types.StringNull()
	}
	if data.Number != nil {
		m.Number = types.Int64Value(int64(*data.Number))
	} else {
		m.Number = types.Int64Null()
	}
	if len(data.UserGroupIDs) > 0 {
		m.UserGroupIDs = util.SetValueOrNull(data.UserGroupIDs)
	} else {
		m.UserGroupIDs = types.SetNull(types.StringType)
	}
	return m
}

// notificationTemplateOwnerAfterApply returns the owner value Terraform must store after Create/Update so the
// result matches the plan: known plan values win (avoid inconsistent apply when the API still stores e.g. admin
// for seeded templates); unknown (omitted) uses the API; explicit null stays null.
func notificationTemplateOwnerAfterApply(planOwner types.String, data *cidaas.NotificationsSrvTemplateModel) types.String {
	switch {
	case !planOwner.IsNull() && !planOwner.IsUnknown():
		return planOwner
	case planOwner.IsNull():
		return types.StringNull()
	default: // unknown in plan — adopt API after apply
		if data != nil && data.Owner != "" {
			return types.StringValue(data.Owner)
		}
		return types.StringNull()
	}
}
