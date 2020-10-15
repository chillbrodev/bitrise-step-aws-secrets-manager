// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the account identifier for the specified access key ID. Access keys
// consist of two parts: an access key ID (for example, AKIAIOSFODNN7EXAMPLE) and a
// secret access key (for example, wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY). For
// more information about access keys, see Managing Access Keys for IAM Users
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html)
// in the IAM User Guide. When you pass an access key ID to this operation, it
// returns the ID of the AWS account to which the keys belong. Access key IDs
// beginning with AKIA are long-term credentials for an IAM user or the AWS account
// root user. Access key IDs beginning with ASIA are temporary credentials that are
// created using STS operations. If the account in the response belongs to you, you
// can sign in as the root user and review your root user access keys. Then, you
// can pull a credentials report
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html)
// to learn which IAM user owns the keys. To learn who requested the temporary
// credentials for an ASIA access key, view the STS events in your CloudTrail logs
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html)
// in the IAM User Guide. This operation does not indicate the state of the access
// key. The key might be active, inactive, or deleted. Active keys might not have
// permissions to perform an operation. Providing a deleted access key might return
// an error that the key doesn't exist.
func (c *Client) GetAccessKeyInfo(ctx context.Context, params *GetAccessKeyInfoInput, optFns ...func(*Options)) (*GetAccessKeyInfoOutput, error) {
	stack := middleware.NewStack("GetAccessKeyInfo", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetAccessKeyInfoMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetAccessKeyInfoValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessKeyInfo(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetAccessKeyInfo",
			Err:           err,
		}
	}
	out := result.(*GetAccessKeyInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessKeyInfoInput struct {

	// The identifier of an access key. This parameter allows (through its regex
	// pattern) a string of characters that can consist of any upper- or lowercase
	// letter or digit.
	//
	// This member is required.
	AccessKeyId *string
}

type GetAccessKeyInfoOutput struct {

	// The number used to identify the AWS account.
	Account *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetAccessKeyInfoMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetAccessKeyInfo{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetAccessKeyInfo{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessKeyInfo(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sts",
		OperationName: "GetAccessKeyInfo",
	}
}