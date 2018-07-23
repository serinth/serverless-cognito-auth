# Serverless Auth Test Repo
This is a spike on testing Cognito for authentication and AWS Lambda functions in Go.
Provider: AWS

This solution does not use the identity pool for giving the client access straight to the AWS resources.

It uses a custom authorizer function which looks at the JWT. The reason for this is that there was too much vendor lock in using Cognito and Identity Pools or took too much effort to get fine grained permissions.

Instead, this will allow us to swap out authenticators e.g. Auth0 by using standard JWT verification.

The User Pool needs to be deployed first and then the JWKS URL needs to be updated in the auth function.
# AWS Cognito

**User Pool** manages Two factor auth either via Phone number or Email address. It also manages confirmation code to verify those fields.

# Deployment

With MFA enabled on the AWS account, we need to first grab a temporary session token and use it. 

1 - Get an AWS Temporary Session Token:
```bash
pip install awsmfa
awsmfa -i <Profile>
```

2 - Optional Log Forwarder

If using the log forwarder, you can first implement and deploy using this template:
[Serverless Log Forwarder Template](https://github.com/serinth/serverless-log-forwarder)

3 - Run Serverless Deployment

```bash
    serverless deploy -v --aws-profile TEMPSESSION
```

# Testing The Authenticated Endpoint

Sign in using AWS Amplify to grab the token. The token can be seen from:

```javascript
let jwtToken = Auth.currentSession().then(function(result) {
    console.log(result);
});;
```

Then curl:

```bash
curl -v -H "Authorization: Bearer <TOKEN>" https://<redacted>.execute-api.ap-southeast-2.amazonaws.com/dev/authenticated
```


# Clean Up

```bash
    serverless remove -v --aws-profile TEMPSESSION
```

# References
[Serverless Framework](https://serverless.com/) - the main framework used for writing and deploying lambda functions

[AWS Amplify](https://github.com/aws/aws-amplify) - for the UI SDK to sign up and sign in users.

[Splunk Log Forwarder Blueprint](https://ap-southeast-2.console.aws.amazon.com/lambda/home?region=ap-southeast-2#/create/new?bp=splunk-logging) - log aggregation