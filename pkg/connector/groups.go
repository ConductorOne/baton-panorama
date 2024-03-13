package connector

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-panorama/pkg/panorama"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	grant "github.com/conductorone/baton-sdk/pkg/types/grant"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
)

type groupBuilder struct {
	resourceType *v2.ResourceType
	client       *panorama.Client
}

const memberEntitlement = "member"

func (o *groupBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return groupResourceType
}

func groupResource(group *panorama.Group) (*v2.Resource, error) {
	profile := map[string]interface{}{
		"name": group.Name,
	}

	groupTraits := []rs.GroupTraitOption{
		rs.WithGroupProfile(profile),
	}

	resource, err := rs.NewGroupResource(group.Name, groupResourceType, group.Name, groupTraits)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (o *groupBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	groups, _, err := o.client.ListGroups(ctx)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to list groups")
	}

	var resources []*v2.Resource
	for _, group := range groups {
		resource, err := groupResource(&group) // #nosec G601
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create group resource")
		}

		resources = append(resources, resource)
	}

	return resources, "", nil, nil
}

func (o *groupBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	assigmentOptions := []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDescription("Is member of group"),
		ent.WithDisplayName(fmt.Sprintf("Member of %s", resource.DisplayName)),
	}

	entitlement := ent.NewAssignmentEntitlement(resource, memberEntitlement, assigmentOptions...)
	rv = append(rv, entitlement)

	return rv, "", nil, nil
}

func (o *groupBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	group, _, err := o.client.GetGroup(ctx, resource.Id.Resource)
	if err != nil {
		return nil, "", nil, err
	}

	var rv []*v2.Grant
	for _, user := range group.Members {
		u, _, err := o.client.GetUser(ctx, user)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to get user")
		}

		userResource, err := userResource(u)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create user resource")
		}

		rv = append(rv, grant.NewGrant(resource, memberEntitlement, userResource))
	}

	return rv, "", nil, nil
}

func newGroupBuilder(client *panorama.Client) *groupBuilder {
	return &groupBuilder{
		resourceType: groupResourceType,
		client:       client,
	}
}
