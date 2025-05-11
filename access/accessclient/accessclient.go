package accessclient

import (
	"context"

	"github.com/tempest-org/tempest/access/access"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// AccessClient is a client for the access service
type AccessClient struct {
	client access.AccessClient
}

// NewAccessClient creates a new access client
func NewAccessClient(target string, opts ...grpc.DialOption) *AccessClient {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: target,
	})
	
	return &AccessClient{
		client: access.NewAccessClient(client.Conn()),
	}
}

// Check checks if a subject has permission to perform an action on an object in an organization
func (c *AccessClient) Check(ctx context.Context, organizationID, subjectID, object, action string) (bool, error) {
	resp, err := c.client.Check(ctx, &access.CheckAccessRequest{
		OrganizationId: organizationID,
		SubjectId:      subjectID,
		Object:         object,
		Action:         action,
	})
	if err != nil {
		return false, err
	}
	
	return resp.Allowed, nil
}

// Grant grants permission to a subject to perform an action on an object in an organization
func (c *AccessClient) Grant(ctx context.Context, organizationID, subjectID, object, action string) error {
	_, err := c.client.Grant(ctx, &access.GrantAccessRequest{
		OrganizationId: organizationID,
		SubjectId:      subjectID,
		Object:         object,
		Action:         action,
	})
	
	return err
}

// Revoke revokes permission from a subject to perform an action on an object in an organization
func (c *AccessClient) Revoke(ctx context.Context, organizationID, subjectID, object, action string) error {
	_, err := c.client.Revoke(ctx, &access.RevokeAccessRequest{
		OrganizationId: organizationID,
		SubjectId:      subjectID,
		Object:         object,
		Action:         action,
	})
	
	return err
}

// GetSubjectAccess returns all access permissions for a subject in an organization
func (c *AccessClient) GetSubjectAccess(ctx context.Context, organizationID, subjectID string) ([]*access.AccessPermission, error) {
	resp, err := c.client.GetSubjectAccess(ctx, &access.GetSubjectAccessRequest{
		OrganizationId: organizationID,
		SubjectId:      subjectID,
	})
	if err != nil {
		return nil, err
	}
	
	return resp.Permissions, nil
}

// GetObjectSubjects returns all subjects that have access to perform an action on an object in an organization
func (c *AccessClient) GetObjectSubjects(ctx context.Context, organizationID, object, action string) ([]string, error) {
	resp, err := c.client.GetObjectSubjects(ctx, &access.GetObjectSubjectsRequest{
		OrganizationId: organizationID,
		Object:         object,
		Action:         action,
	})
	if err != nil {
		return nil, err
	}
	
	return resp.SubjectIds, nil
}

// HealthCheck checks the health of the access service
func (c *AccessClient) HealthCheck(ctx context.Context) (bool, error) {
	resp, err := c.client.HealthCheck(ctx, &access.HealthCheckRequest{
		Service: "access",
	})
	if err != nil {
		return false, err
	}
	
	return resp.Status == access.HealthCheckResponse_SERVING, nil
}