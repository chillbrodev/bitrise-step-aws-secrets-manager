// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	cryptorand "crypto/rand"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	smithyrand "github.com/awslabs/smithy-go/rand"
	"net/http"
	"time"
)

const ServiceID = "Secrets Manager"
const ServiceAPIVersion = "2017-10-17"

// AWS Secrets Manager API Reference AWS Secrets Manager provides a service to
// enable you to store, manage, and retrieve, secrets.  <p>This guide provides
// descriptions of the Secrets Manager API. For more information about using this
// service, see the <a
// href="https://docs.aws.amazon.com/secretsmanager/latest/userguide/introduction.html">AWS
// Secrets Manager User Guide</a>.</p> <p> <b>API Version</b> </p> <p>This version
// of the Secrets Manager API Reference documents the Secrets Manager API version
// 2017-10-17.</p> <note> <p>As an alternative to using the API, you can use one of
// the AWS SDKs, which consist of libraries and sample code for various programming
// languages and platforms such as Java, Ruby, .NET, iOS, and Android. The SDKs
// provide a convenient way to create programmatic access to AWS Secrets Manager.
// For example, the SDKs provide cryptographically signing requests, managing
// errors, and retrying requests automatically. For more information about the AWS
// SDKs, including downloading and installing them, see <a
// href="http://aws.amazon.com/tools/">Tools for Amazon Web Services</a>.</p>
// </note> <p>We recommend you use the AWS SDKs to make programmatic API calls to
// Secrets Manager. However, you also can use the Secrets Manager HTTP Query API to
// make direct calls to the Secrets Manager web service. To learn more about the
// Secrets Manager HTTP Query API, see <a
// href="https://docs.aws.amazon.com/secretsmanager/latest/userguide/query-requests.html">Making
// Query Requests</a> in the <i>AWS Secrets Manager User Guide</i>. </p> <p>Secrets
// Manager API supports GET and POST requests for all actions, and doesn't require
// you to use GET for some actions and POST for others. However, GET requests are
// subject to the limitation size of a URL. Therefore, for operations that require
// larger sizes, use a POST request.</p> <p> <b>Support and Feedback for AWS
// Secrets Manager</b> </p> <p>We welcome your feedback. Send your comments to <a
// href="mailto:awssecretsmanager-feedback@amazon.com">awssecretsmanager-feedback@amazon.com</a>,
// or post your feedback and questions in the <a
// href="http://forums.aws.amazon.com/forum.jspa?forumID=296">AWS Secrets Manager
// Discussion Forum</a>. For more information about the AWS Discussion Forums, see
// <a href="http://forums.aws.amazon.com/help.jspa">Forums Help</a>.</p> <p> <b>How
// examples are presented</b> </p> <p>The JSON that AWS Secrets Manager expects as
// your request parameters and the service returns as a response to HTTP query
// requests contain single, long strings without line breaks or white space
// formatting. The JSON shown in the examples displays the code formatted with both
// line breaks and white space to improve readability. When example input
// parameters can also cause long strings extending beyond the screen, you can
// insert line breaks to enhance readability. You should always submit the input as
// a single JSON text string.</p> <p> <b>Logging API Requests</b> </p> <p>AWS
// Secrets Manager supports AWS CloudTrail, a service that records AWS API calls
// for your AWS account and delivers log files to an Amazon S3 bucket. By using
// information that's collected by AWS CloudTrail, you can determine the requests
// successfully made to Secrets Manager, who made the request, when it was made,
// and so on. For more about AWS Secrets Manager and support for AWS CloudTrail,
// see <a
// href="http://docs.aws.amazon.com/secretsmanager/latest/userguide/monitoring.html#monitoring_cloudtrail">Logging
// AWS Secrets Manager Events with AWS CloudTrail</a> in the <i>AWS Secrets Manager
// User Guide</i>. To learn more about CloudTrail, including enabling it and find
// your log files, see the <a
// href="https://docs.aws.amazon.com/awscloudtrail/latest/userguide/what_is_cloud_trail_top_level.html">AWS
// CloudTrail User Guide</a>.</p>
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	resolveIdempotencyTokenProvider(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// Provides idempotency tokens values that will be automatically populated into
	// idempotent API operations.
	IdempotencyTokenProvider IdempotencyTokenProvider

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetHTTPSignerV4() HTTPSignerV4 {
	return o.HTTPSignerV4
}

func (o Options) GetIdempotencyTokenProvider() IdempotencyTokenProvider {
	return o.IdempotencyTokenProvider
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
		APIOptions:  cfg.APIOptions,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("secretsmanager")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}

func resolveIdempotencyTokenProvider(o *Options) {
	if o.IdempotencyTokenProvider != nil {
		return
	}
	o.IdempotencyTokenProvider = smithyrand.NewUUIDIdempotencyToken(cryptorand.Reader)
}

// IdempotencyTokenProvider interface for providing idempotency token
type IdempotencyTokenProvider interface {
	GetIdempotencyToken() (string, error)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) {
	awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) {
	awshttp.AddResponseErrorMiddleware(stack)
}