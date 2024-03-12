package connector

import (
	"context"

	"github.com/conductorone/baton-network-security/pkg/panorama"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
)

type groupBuilder struct {
	resourceType *v2.ResourceType
	client       *panorama.Client
}

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
		return nil, "", nil, err // TODO: wrap error
	}

	var resources []*v2.Resource
	for _, group := range groups {
		resource, err := groupResource(&group)
		if err != nil {
			return nil, "", nil, err // TODO: wrap error
		}

		resources = append(resources, resource)
	}

	return resources, "", nil, nil
}

func (o *groupBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (o *groupBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func newGroupBuilder(client *panorama.Client) *groupBuilder {
	return &groupBuilder{
		resourceType: groupResourceType,
		client:       client,
	}
}
