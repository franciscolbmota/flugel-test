package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

	// look up the tags for the give instance ID
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)

	// check if the EC2 instance with a given tag and name is set.
	testingTag, containsTestingTag := instanceTags["Flugel"]
	assert.True(t, containsTestingTag)
	assert.Equal(t, "testing-tag-value", Flugel)
}
