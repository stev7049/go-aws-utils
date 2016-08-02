package main

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/rds"
)

func getResourceCounts(region string, humanregion string, goGroup *sync.WaitGroup) {
	defer goGroup.Done()

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	var totalInstances int
	for _, res := range resp.Reservations {
		totalInstances += len(res.Instances)
	}

	// Get VPC count
	respVpc, err := svc.DescribeVpcs(nil)
	if err != nil {
		panic(err)
	}

	totalVPCs := len(respVpc.Vpcs)

	// Get Subnets count
	respSubnets, err := svc.DescribeSubnets(nil)
	if err != nil {
		panic(err)
	}

	totalSubnets := len(respSubnets.Subnets)

	// Get SecurityGroups count
	respSecurityGroups, err := svc.DescribeSecurityGroups(nil)
	if err != nil {
		panic(err)
	}

	totalSecurityGroups := len(respSecurityGroups.SecurityGroups)

	// Get Volumes count
	respEBS, err := svc.DescribeVolumes(nil)
	if err != nil {
		panic(err)
	}

	totalEBS := len(respEBS.Volumes)

	// Get AutoScalingGroup count``
	asg := autoscaling.New(session.New(), &aws.Config{Region: aws.String(region)})

	respAsg, err := asg.DescribeAutoScalingGroups(nil)
	if err != nil {
		panic(err)
	}

	totalASG := (len(respAsg.AutoScalingGroups))

	// Get CloudFormation Stack Counts
	cf := cloudformation.New(session.New(), &aws.Config{Region: aws.String(region)})

	respCf, err := cf.DescribeStacks(nil)
	if err != nil {
		panic(err)
	}

	totalCloudFormationStacks := len(respCf.Stacks)

	// Get ELB Counts
	awselb := elb.New(session.New(), &aws.Config{Region: aws.String(region)})

	respElb, err := awselb.DescribeLoadBalancers(nil)
	if err != nil {
		panic(err)
	}

	totalElb := len(respElb.LoadBalancerDescriptions)

	// Get RDS Counts
	awsrds := rds.New(session.New(), &aws.Config{Region: aws.String(region)})

	respRDS, err := awsrds.DescribeDBInstances(nil)
	if err != nil {
		panic(err)
	}

	totalRDS := len(respRDS.DBInstances)

	// Get Elastic BeanStalk Counts
	bean := elasticbeanstalk.New(session.New(), &aws.Config{Region: aws.String(region)})

	respBean, err := bean.DescribeApplications(nil)
	if err != nil {
		panic(err)
	}

	totalBean := len(respBean.Applications)

	// Get CodeDeploy Counts
	cd := codedeploy.New(session.New(), &aws.Config{Region: aws.String(region)})

	respCD, err := cd.ListApplications(nil)
	if err != nil {
		panic(err)
	}

	totalCD := len(respCD.Applications)

	//Print stuff
	fmt.Printf("%+16s %+16s : %4d %4d %4d %4d %4d", humanregion, region, totalInstances, totalRDS, totalEBS, totalElb, totalASG)
	fmt.Printf("%4d %4d %4d %4d %4d %4d\n", totalVPCs, totalSubnets, totalSecurityGroups, totalCloudFormationStacks, totalBean, totalCD)
}

func main() {

	awsregions := map[int][]string{
		0:  {"us-east-1", "N. Virginia"},
		1:  {"us-west-1", "N. California"},
		2:  {"us-west-2", "Oregon"},
		3:  {"eu-west-1", "Ireland"},
		4:  {"eu-central-1", "Frankfurt"},
		5:  {"ap-southeast-1", "Singapore"},
		6:  {"ap-southeast-2", "Sydney"},
		7:  {"ap-northeast-1", "Tokyo"},
		8:  {"ap-northeast-2", "Seoul"},
		9:  {"sa-east-1", "São Paulo"},
		10: {"ap-south-1", "Mumbai"},
	}
	goGroup := new(sync.WaitGroup)
	defer goGroup.Wait()

	fmt.Printf("%+33s : %4s %4s %4s %4s %4s", "Region", "EC2", "RDS", "EBS", "ELB", "ASG")
	fmt.Printf("%4s %4s %4s %4s %4s %4s\n", "VPC", "SNET", "SG", "CF", "Bean", "CD")
	fmt.Println("-----------------------------------------------------------------------------------------------------------")
	for i := 0; i < len(awsregions); i++ {
		goGroup.Add(1)
		go getResourceCounts(awsregions[i][0], awsregions[i][1], goGroup)
	}
}
