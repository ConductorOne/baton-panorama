package connector

import (
	"context"

	"github.com/conductorone/baton-panorama/pkg/panorama"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
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
		return v2.UserTrait_Status_STATUS_DISABLED
	}

	return v2.UserTrait_Status_STATUS_ENABLED
}

func userResource(user *panorama.User) (*v2.Resource, error) {
	profile := map[string]interface{}{
		"name": user.Name,
	}

	userTraits := []rs.UserTraitOption{
		rs.WithUserLogin(user.Name),
	}

	r, err := rs.NewUserResource(user.Name, userResourceType, user.Name, userTraits,
		rs.WithResourceProfile(profile),
		rs.WithResourceStatus(v2.Status_ResourceStatus(getUserStatus(user)), ""))
	if err != nil {
		return nil, err
	}

	return r, nil
}

// List returns all the users from the database as resource objects.
// Users include a UserTrait because they are the 'shape' of a standard user.
func (o *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, opts rs.SyncOpAttrs) ([]*v2.Resource, *rs.SyncOpResults, error) {
	users, _, err := o.client.ListUsers(ctx)
	if err != nil {
		return nil, nil, wrapError(err, "failed to list users")
	}

	var resources []*v2.Resource
	for _, user := range users {
		r, err := userResource(&user) // #nosec G601
		if err != nil {
			return nil, nil, wrapError(err, "failed to create user resource")
		}

		resources = append(resources, r)
	}

	return resources, &rs.SyncOpResults{}, nil
}

// Entitlements always returns an empty slice for users.
func (o *userBuilder) Entitlements(_ context.Context, r *v2.Resource, _ rs.SyncOpAttrs) ([]*v2.Entitlement, *rs.SyncOpResults, error) {
	return nil, &rs.SyncOpResults{}, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (o *userBuilder) Grants(ctx context.Context, r *v2.Resource, opts rs.SyncOpAttrs) ([]*v2.Grant, *rs.SyncOpResults, error) {
	return nil, &rs.SyncOpResults{}, nil
}

func newUserBuilder(client *panorama.Client) *userBuilder {
	return &userBuilder{
		resourceType: userResourceType,
		client:       client,
	}
}
