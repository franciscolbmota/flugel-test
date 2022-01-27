package test

import (
	"testing"
)

func GetTagsForEc2Instance(t testing.TestingT, region string, instanceID string) map[string]string {
	tags, err := GetTagsForEc2InstanceE(t, region, instanceID)
	require.NoError(t, err)
	return tags
}
