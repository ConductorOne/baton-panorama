package connector

import (
	"context"

	"github.com/conductorone/baton-network-security/pkg/panorama"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
)

type userBuilder struct {
	resourceType *v2.ResourceType
	client       *panorama.Client
}

func (o *userBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return userResourceType
}

func getUserStatus(user *panorama.User) v2.UserTrait_Status_Status {
	if user.Disabled {
		return v2.UserTrait_Status_STATUS_ENABLED
	}

	return v2.UserTrait_Status_STATUS_DISABLED
}

func userResource(user *panorama.User) (*v2.Resource, error) {
	profile := map[string]interface{}{
		"name": user.Name,
	}

	userTraits := []rs.UserTraitOption{
		rs.WithUserProfile(profile),
		rs.WithUserLogin(user.Name),
		rs.WithStatus(getUserStatus(user)),
	}

	resource, err := rs.NewUserResource(user.Name, userResourceType, user.Name, userTraits)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

// List returns all the users from the database as resource objects.
// Users include a UserTrait because they are the 'shape' of a standard user.
func (o *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	users, _, err := o.client.ListUsers(ctx)
	if err != nil {
		return nil, "", nil, err // TODO: wrap error
	}

	var resources []*v2.Resource
	for _, user := range users {
		resource, err := userResource(&user)
		if err != nil {
			return nil, "", nil, err // TODO: wrap error
		}

		resources = append(resources, resource)
	}

	return resources, "", nil, nil
}

// Entitlements always returns an empty slice for users.
func (o *userBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (o *userBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func newUserBuilder(client *panorama.Client) *userBuilder {
	return &userBuilder{
		resourceType: userResourceType,
		client:       client,
	}
}
