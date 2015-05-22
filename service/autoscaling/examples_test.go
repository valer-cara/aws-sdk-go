package autoscaling_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/autoscaling"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleAutoScaling_AttachInstances() {
	svc := autoscaling.New(nil)

	params := &autoscaling.AttachInstancesInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		InstanceIDs: []*string{
			aws.String("XmlStringMaxLen16"), // Required
			// More values...
		},
	}
	resp, err := svc.AttachInstances(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_CompleteLifecycleAction() {
	svc := autoscaling.New(nil)

	params := &autoscaling.CompleteLifecycleActionInput{
		AutoScalingGroupName:  aws.String("ResourceName"),          // Required
		LifecycleActionResult: aws.String("LifecycleActionResult"), // Required
		LifecycleActionToken:  aws.String("LifecycleActionToken"),  // Required
		LifecycleHookName:     aws.String("AsciiStringMaxLen255"),  // Required
	}
	resp, err := svc.CompleteLifecycleAction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_CreateAutoScalingGroup() {
	svc := autoscaling.New(nil)

	params := &autoscaling.CreateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("XmlStringMaxLen255"), // Required
		MaxSize:              aws.Long(1),                      // Required
		MinSize:              aws.Long(1),                      // Required
		AvailabilityZones: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		DefaultCooldown:         aws.Long(1),
		DesiredCapacity:         aws.Long(1),
		HealthCheckGracePeriod:  aws.Long(1),
		HealthCheckType:         aws.String("XmlStringMaxLen32"),
		InstanceID:              aws.String("XmlStringMaxLen16"),
		LaunchConfigurationName: aws.String("ResourceName"),
		LoadBalancerNames: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		PlacementGroup: aws.String("XmlStringMaxLen255"),
		Tags: []*autoscaling.Tag{
			&autoscaling.Tag{ // Required
				Key:               aws.String("TagKey"), // Required
				PropagateAtLaunch: aws.Boolean(true),
				ResourceID:        aws.String("XmlString"),
				ResourceType:      aws.String("XmlString"),
				Value:             aws.String("TagValue"),
			},
			// More values...
		},
		TerminationPolicies: []*string{
			aws.String("XmlStringMaxLen1600"), // Required
			// More values...
		},
		VPCZoneIdentifier: aws.String("XmlStringMaxLen255"),
	}
	resp, err := svc.CreateAutoScalingGroup(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_CreateLaunchConfiguration() {
	svc := autoscaling.New(nil)

	params := &autoscaling.CreateLaunchConfigurationInput{
		LaunchConfigurationName:  aws.String("XmlStringMaxLen255"), // Required
		AssociatePublicIPAddress: aws.Boolean(true),
		BlockDeviceMappings: []*autoscaling.BlockDeviceMapping{
			&autoscaling.BlockDeviceMapping{ // Required
				DeviceName: aws.String("XmlStringMaxLen255"), // Required
				EBS: &autoscaling.EBS{
					DeleteOnTermination: aws.Boolean(true),
					IOPS:                aws.Long(1),
					SnapshotID:          aws.String("XmlStringMaxLen255"),
					VolumeSize:          aws.Long(1),
					VolumeType:          aws.String("BlockDeviceEbsVolumeType"),
				},
				NoDevice:    aws.Boolean(true),
				VirtualName: aws.String("XmlStringMaxLen255"),
			},
			// More values...
		},
		ClassicLinkVPCID: aws.String("XmlStringMaxLen255"),
		ClassicLinkVPCSecurityGroups: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		EBSOptimized:       aws.Boolean(true),
		IAMInstanceProfile: aws.String("XmlStringMaxLen1600"),
		ImageID:            aws.String("XmlStringMaxLen255"),
		InstanceID:         aws.String("XmlStringMaxLen16"),
		InstanceMonitoring: &autoscaling.InstanceMonitoring{
			Enabled: aws.Boolean(true),
		},
		InstanceType:     aws.String("XmlStringMaxLen255"),
		KernelID:         aws.String("XmlStringMaxLen255"),
		KeyName:          aws.String("XmlStringMaxLen255"),
		PlacementTenancy: aws.String("XmlStringMaxLen64"),
		RAMDiskID:        aws.String("XmlStringMaxLen255"),
		SecurityGroups: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		SpotPrice: aws.String("SpotPrice"),
		UserData:  aws.String("XmlStringUserData"),
	}
	resp, err := svc.CreateLaunchConfiguration(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_CreateOrUpdateTags() {
	svc := autoscaling.New(nil)

	params := &autoscaling.CreateOrUpdateTagsInput{
		Tags: []*autoscaling.Tag{ // Required
			&autoscaling.Tag{ // Required
				Key:               aws.String("TagKey"), // Required
				PropagateAtLaunch: aws.Boolean(true),
				ResourceID:        aws.String("XmlString"),
				ResourceType:      aws.String("XmlString"),
				Value:             aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateOrUpdateTags(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DeleteAutoScalingGroup() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DeleteAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		ForceDelete:          aws.Boolean(true),
	}
	resp, err := svc.DeleteAutoScalingGroup(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DeleteLaunchConfiguration() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DeleteLaunchConfigurationInput{
		LaunchConfigurationName: aws.String("ResourceName"), // Required
	}
	resp, err := svc.DeleteLaunchConfiguration(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DeleteLifecycleHook() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DeleteLifecycleHookInput{
		AutoScalingGroupName: aws.String("ResourceName"),         // Required
		LifecycleHookName:    aws.String("AsciiStringMaxLen255"), // Required
	}
	resp, err := svc.DeleteLifecycleHook(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DeleteNotificationConfiguration() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DeleteNotificationConfigurationInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		TopicARN:             aws.String("ResourceName"), // Required
	}
	resp, err := svc.DeleteNotificationConfiguration(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DeletePolicy() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DeletePolicyInput{
		PolicyName:           aws.String("ResourceName"), // Required
		AutoScalingGroupName: aws.String("ResourceName"),
	}
	resp, err := svc.DeletePolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DeleteScheduledAction() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DeleteScheduledActionInput{
		ScheduledActionName:  aws.String("ResourceName"), // Required
		AutoScalingGroupName: aws.String("ResourceName"),
	}
	resp, err := svc.DeleteScheduledAction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DeleteTags() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DeleteTagsInput{
		Tags: []*autoscaling.Tag{ // Required
			&autoscaling.Tag{ // Required
				Key:               aws.String("TagKey"), // Required
				PropagateAtLaunch: aws.Boolean(true),
				ResourceID:        aws.String("XmlString"),
				ResourceType:      aws.String("XmlString"),
				Value:             aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.DeleteTags(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeAccountLimits() {
	svc := autoscaling.New(nil)

	var params *autoscaling.DescribeAccountLimitsInput
	resp, err := svc.DescribeAccountLimits(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeAdjustmentTypes() {
	svc := autoscaling.New(nil)

	var params *autoscaling.DescribeAdjustmentTypesInput
	resp, err := svc.DescribeAdjustmentTypes(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeAutoScalingGroups() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		MaxRecords: aws.Long(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeAutoScalingGroups(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeAutoScalingInstances() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeAutoScalingInstancesInput{
		InstanceIDs: []*string{
			aws.String("XmlStringMaxLen16"), // Required
			// More values...
		},
		MaxRecords: aws.Long(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeAutoScalingInstances(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeAutoScalingNotificationTypes() {
	svc := autoscaling.New(nil)

	var params *autoscaling.DescribeAutoScalingNotificationTypesInput
	resp, err := svc.DescribeAutoScalingNotificationTypes(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeLaunchConfigurations() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeLaunchConfigurationsInput{
		LaunchConfigurationNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		MaxRecords: aws.Long(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeLaunchConfigurations(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeLifecycleHookTypes() {
	svc := autoscaling.New(nil)

	var params *autoscaling.DescribeLifecycleHookTypesInput
	resp, err := svc.DescribeLifecycleHookTypes(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeLifecycleHooks() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeLifecycleHooksInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		LifecycleHookNames: []*string{
			aws.String("AsciiStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeLifecycleHooks(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeMetricCollectionTypes() {
	svc := autoscaling.New(nil)

	var params *autoscaling.DescribeMetricCollectionTypesInput
	resp, err := svc.DescribeMetricCollectionTypes(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeNotificationConfigurations() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeNotificationConfigurationsInput{
		AutoScalingGroupNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		MaxRecords: aws.Long(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeNotificationConfigurations(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribePolicies() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribePoliciesInput{
		AutoScalingGroupName: aws.String("ResourceName"),
		MaxRecords:           aws.Long(1),
		NextToken:            aws.String("XmlString"),
		PolicyNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribePolicies(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeScalingActivities() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeScalingActivitiesInput{
		ActivityIDs: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		AutoScalingGroupName: aws.String("ResourceName"),
		MaxRecords:           aws.Long(1),
		NextToken:            aws.String("XmlString"),
	}
	resp, err := svc.DescribeScalingActivities(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeScalingProcessTypes() {
	svc := autoscaling.New(nil)

	var params *autoscaling.DescribeScalingProcessTypesInput
	resp, err := svc.DescribeScalingProcessTypes(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeScheduledActions() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeScheduledActionsInput{
		AutoScalingGroupName: aws.String("ResourceName"),
		EndTime:              aws.Time(time.Now()),
		MaxRecords:           aws.Long(1),
		NextToken:            aws.String("XmlString"),
		ScheduledActionNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		StartTime: aws.Time(time.Now()),
	}
	resp, err := svc.DescribeScheduledActions(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeTags() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DescribeTagsInput{
		Filters: []*autoscaling.Filter{
			&autoscaling.Filter{ // Required
				Name: aws.String("XmlString"),
				Values: []*string{
					aws.String("XmlString"), // Required
					// More values...
				},
			},
			// More values...
		},
		MaxRecords: aws.Long(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DescribeTerminationPolicyTypes() {
	svc := autoscaling.New(nil)

	var params *autoscaling.DescribeTerminationPolicyTypesInput
	resp, err := svc.DescribeTerminationPolicyTypes(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DetachInstances() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DetachInstancesInput{
		AutoScalingGroupName:           aws.String("ResourceName"), // Required
		ShouldDecrementDesiredCapacity: aws.Boolean(true),          // Required
		InstanceIDs: []*string{
			aws.String("XmlStringMaxLen16"), // Required
			// More values...
		},
	}
	resp, err := svc.DetachInstances(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_DisableMetricsCollection() {
	svc := autoscaling.New(nil)

	params := &autoscaling.DisableMetricsCollectionInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		Metrics: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.DisableMetricsCollection(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_EnableMetricsCollection() {
	svc := autoscaling.New(nil)

	params := &autoscaling.EnableMetricsCollectionInput{
		AutoScalingGroupName: aws.String("ResourceName"),       // Required
		Granularity:          aws.String("XmlStringMaxLen255"), // Required
		Metrics: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.EnableMetricsCollection(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_EnterStandby() {
	svc := autoscaling.New(nil)

	params := &autoscaling.EnterStandbyInput{
		AutoScalingGroupName:           aws.String("ResourceName"), // Required
		ShouldDecrementDesiredCapacity: aws.Boolean(true),          // Required
		InstanceIDs: []*string{
			aws.String("XmlStringMaxLen16"), // Required
			// More values...
		},
	}
	resp, err := svc.EnterStandby(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_ExecutePolicy() {
	svc := autoscaling.New(nil)

	params := &autoscaling.ExecutePolicyInput{
		PolicyName:           aws.String("ResourceName"), // Required
		AutoScalingGroupName: aws.String("ResourceName"),
		HonorCooldown:        aws.Boolean(true),
	}
	resp, err := svc.ExecutePolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_ExitStandby() {
	svc := autoscaling.New(nil)

	params := &autoscaling.ExitStandbyInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		InstanceIDs: []*string{
			aws.String("XmlStringMaxLen16"), // Required
			// More values...
		},
	}
	resp, err := svc.ExitStandby(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_PutLifecycleHook() {
	svc := autoscaling.New(nil)

	params := &autoscaling.PutLifecycleHookInput{
		AutoScalingGroupName:  aws.String("ResourceName"),         // Required
		LifecycleHookName:     aws.String("AsciiStringMaxLen255"), // Required
		DefaultResult:         aws.String("LifecycleActionResult"),
		HeartbeatTimeout:      aws.Long(1),
		LifecycleTransition:   aws.String("LifecycleTransition"),
		NotificationMetadata:  aws.String("XmlStringMaxLen1023"),
		NotificationTargetARN: aws.String("ResourceName"),
		RoleARN:               aws.String("ResourceName"),
	}
	resp, err := svc.PutLifecycleHook(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_PutNotificationConfiguration() {
	svc := autoscaling.New(nil)

	params := &autoscaling.PutNotificationConfigurationInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		NotificationTypes: []*string{ // Required
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		TopicARN: aws.String("ResourceName"), // Required
	}
	resp, err := svc.PutNotificationConfiguration(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_PutScalingPolicy() {
	svc := autoscaling.New(nil)

	params := &autoscaling.PutScalingPolicyInput{
		AdjustmentType:       aws.String("XmlStringMaxLen255"), // Required
		AutoScalingGroupName: aws.String("ResourceName"),       // Required
		PolicyName:           aws.String("XmlStringMaxLen255"), // Required
		ScalingAdjustment:    aws.Long(1),                      // Required
		Cooldown:             aws.Long(1),
		MinAdjustmentStep:    aws.Long(1),
	}
	resp, err := svc.PutScalingPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_PutScheduledUpdateGroupAction() {
	svc := autoscaling.New(nil)

	params := &autoscaling.PutScheduledUpdateGroupActionInput{
		AutoScalingGroupName: aws.String("ResourceName"),       // Required
		ScheduledActionName:  aws.String("XmlStringMaxLen255"), // Required
		DesiredCapacity:      aws.Long(1),
		EndTime:              aws.Time(time.Now()),
		MaxSize:              aws.Long(1),
		MinSize:              aws.Long(1),
		Recurrence:           aws.String("XmlStringMaxLen255"),
		StartTime:            aws.Time(time.Now()),
		Time:                 aws.Time(time.Now()),
	}
	resp, err := svc.PutScheduledUpdateGroupAction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_RecordLifecycleActionHeartbeat() {
	svc := autoscaling.New(nil)

	params := &autoscaling.RecordLifecycleActionHeartbeatInput{
		AutoScalingGroupName: aws.String("ResourceName"),         // Required
		LifecycleActionToken: aws.String("LifecycleActionToken"), // Required
		LifecycleHookName:    aws.String("AsciiStringMaxLen255"), // Required
	}
	resp, err := svc.RecordLifecycleActionHeartbeat(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_ResumeProcesses() {
	svc := autoscaling.New(nil)

	params := &autoscaling.ScalingProcessQuery{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		ScalingProcesses: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.ResumeProcesses(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_SetDesiredCapacity() {
	svc := autoscaling.New(nil)

	params := &autoscaling.SetDesiredCapacityInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		DesiredCapacity:      aws.Long(1),                // Required
		HonorCooldown:        aws.Boolean(true),
	}
	resp, err := svc.SetDesiredCapacity(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_SetInstanceHealth() {
	svc := autoscaling.New(nil)

	params := &autoscaling.SetInstanceHealthInput{
		HealthStatus:             aws.String("XmlStringMaxLen32"), // Required
		InstanceID:               aws.String("XmlStringMaxLen16"), // Required
		ShouldRespectGracePeriod: aws.Boolean(true),
	}
	resp, err := svc.SetInstanceHealth(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_SuspendProcesses() {
	svc := autoscaling.New(nil)

	params := &autoscaling.ScalingProcessQuery{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		ScalingProcesses: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.SuspendProcesses(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_TerminateInstanceInAutoScalingGroup() {
	svc := autoscaling.New(nil)

	params := &autoscaling.TerminateInstanceInAutoScalingGroupInput{
		InstanceID:                     aws.String("XmlStringMaxLen16"), // Required
		ShouldDecrementDesiredCapacity: aws.Boolean(true),               // Required
	}
	resp, err := svc.TerminateInstanceInAutoScalingGroup(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleAutoScaling_UpdateAutoScalingGroup() {
	svc := autoscaling.New(nil)

	params := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		AvailabilityZones: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		DefaultCooldown:         aws.Long(1),
		DesiredCapacity:         aws.Long(1),
		HealthCheckGracePeriod:  aws.Long(1),
		HealthCheckType:         aws.String("XmlStringMaxLen32"),
		LaunchConfigurationName: aws.String("ResourceName"),
		MaxSize:                 aws.Long(1),
		MinSize:                 aws.Long(1),
		PlacementGroup:          aws.String("XmlStringMaxLen255"),
		TerminationPolicies: []*string{
			aws.String("XmlStringMaxLen1600"), // Required
			// More values...
		},
		VPCZoneIdentifier: aws.String("XmlStringMaxLen255"),
	}
	resp, err := svc.UpdateAutoScalingGroup(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}