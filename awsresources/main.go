package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/mdfilio/go-aws-utils/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/rds"
)

func regionError(service string, region string, err error) error {
	return fmt.Errorf("Error occured for service %s in region %s: %v\n", service, region, err)
}

func getResourceCounts(region string, humanregion string, goGroup *sync.WaitGroup, errChan chan error) {
	defer goGroup.Done()

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		errChan <- regionError("ec2 (instance count)", region, err)
		return
	}

	var totalInstances int
	for _, res := range resp.Reservations {
		totalInstances += len(res.Instances)
	}

	// Get VPC count
	respVpc, err := svc.DescribeVpcs(nil)
	if err != nil {
		errChan <- regionError("ec2 (vpc count)", region, err)
		return
	}

	totalVPCs := len(respVpc.Vpcs)

	// Get Subnets count
	respSubnets, err := svc.DescribeSubnets(nil)
	if err != nil {
		errChan <- regionError("ec2 (subnet count)", region, err)
		return
	}

	totalSubnets := len(respSubnets.Subnets)

	// Get SecurityGroups count
	respSecurityGroups, err := svc.DescribeSecurityGroups(nil)
	if err != nil {
		errChan <- regionError("ec2 (security group count)", region, err)
		return
	}

	totalSecurityGroups := len(respSecurityGroups.SecurityGroups)

	// Get Volumes count
	respEBS, err := svc.DescribeVolumes(nil)
	if err != nil {
		errChan <- regionError("ec2 (ebs volume count)", region, err)
		return
	}

	totalEBS := len(respEBS.Volumes)

	/* Get Snapshots count
	   Relies on env variable being set to filter snapshots to
	   account, rather than all public snapshots for AMIs
	*/
	var totalSS = -1
	var awsaccount string
	awsaccount = os.Getenv("AWS_ACCOUNT")
	if awsaccount != "" {
		ssinput := &ec2.DescribeSnapshotsInput{
			OwnerIds: []*string{
				aws.String(awsaccount),
			},
		}

		respSS, err := svc.DescribeSnapshots(ssinput)
		if err != nil {
			errChan <- regionError("ec2 (snapshot count)", region, err)
			return
		}
		totalSS = len(respSS.Snapshots)
	}

	// Get EFS filesystems  count
	efssvc := efs.New(session.New(), &aws.Config{Region: aws.String(region)})
	respEFS, err := efssvc.DescribeFileSystems(nil)
	/*if err != nil {
		errChan <- regionError(region, err)
		return
	}*/

	var totalFileSystems int
	totalFileSystems = len(respEFS.FileSystems)

	// DynamoDB tables
	ddb := dynamodb.New(session.New(&aws.Config{Region: aws.String(region)}))

	respddb, err := ddb.ListTables(nil)
	if err != nil {
		errChan <- regionError("dynamodb", region, err)
		return
	}

	var totalTables int
	totalTables = len(respddb.TableNames)

	// Get AutoScalingGroup count
	asg := autoscaling.New(session.New(), &aws.Config{Region: aws.String(region)})

	respAsg, err := asg.DescribeAutoScalingGroups(nil)
	if err != nil {
		errChan <- regionError("asg", region, err)
		return
	}

	totalASG := (len(respAsg.AutoScalingGroups))

	// Get CloudFormation Stack Counts
	cf := cloudformation.New(session.New(), &aws.Config{Region: aws.String(region)})

	respCf, err := cf.DescribeStacks(nil)
	if err != nil {
		errChan <- regionError("CloudFormation", region, err)
		return
	}

	totalCloudFormationStacks := len(respCf.Stacks)

	// Get ELB Counts
	awselb := elb.New(session.New(), &aws.Config{Region: aws.String(region)})

	respElb, err := awselb.DescribeLoadBalancers(nil)
	if err != nil {
		errChan <- regionError("elb", region, err)
		return
	}

	totalElb := len(respElb.LoadBalancerDescriptions)

	// Get ELBv2 Counts. Adds to exiting ELB counter.
	awselbv2 := elbv2.New(session.New(), &aws.Config{Region: aws.String(region)})

	respElbv2, err := awselbv2.DescribeLoadBalancers(nil)
	if err != nil {
		errChan <- regionError("elbv2", region, err)
		return
	}

	totalElb += len(respElbv2.LoadBalancers)


	// Get RDS Counts
	awsrds := rds.New(session.New(), &aws.Config{Region: aws.String(region)})

	respRDS, err := awsrds.DescribeDBInstances(nil)
	if err != nil {
		errChan <- regionError("rds", region, err)
		return
	}

	totalRDS := len(respRDS.DBInstances)

	// Get Elastic BeanStalk Counts
	eb := elasticbeanstalk.New(session.New(), &aws.Config{Region: aws.String(region)})

	respeb, err := eb.DescribeApplications(nil)
	if err != nil {
		errChan <- regionError("eb", region, err)
		return
	}

	totaleb := len(respeb.Applications)

	//Get ECS Cluster Counts
	contService := ecs.New(session.New(), &aws.Config{Region: aws.String(region)})

	ecsRes, err := contService.ListClusters(nil)
	if err != nil {
		errChan <- regionError("ecs", region, err)
		return
	}
	totalECS := len(ecsRes.ClusterArns)

	// Get CodeDeploy Counts
	cd := codedeploy.New(session.New(), &aws.Config{Region: aws.String(region)})

	respCD, err := cd.ListApplications(nil)
	if err != nil {
		errChan <- regionError("CodeDeploy", region, err)
		return
	}

	totalCD := len(respCD.Applications)

	// Get Lambda Functions
	l := lambda.New(session.New(), &aws.Config{Region: aws.String(region)})

	respl, err := l.ListFunctions(nil)
	if err != nil {
		errChan <- regionError("Lambda", region, err)
		return
	}

	totalFunctions := len(respl.Functions)

	//Print stuff
	fmt.Printf("%+16s %+16s : %4d %4d %4d %4d %4d %4d %4d", humanregion, region, totalInstances, totalECS, totalRDS, totalEBS, totalSS, totalElb, totalASG)
	fmt.Printf("%4d %4d %4d %4d %4d %4d %4d %4d %4d\n", totalVPCs, totalSubnets, totalSecurityGroups, totalCloudFormationStacks, totaleb, totalCD, totalTables, totalFileSystems, totalFunctions)

	return
}

func main() {
	goGroup := new(sync.WaitGroup)

	errChan := make(chan error)
	defer func() {
		close(errChan)
	}()
	defer goGroup.Wait()

	fmt.Printf("%+33s : %4s %4s %4s %4s %4s %4s %4s", "Region", "EC2", "ECS", "RDS", "EBS", "SS", "ELB", "ASG")
	fmt.Printf("%4s %4s %4s %4s %4s %4s %4s %4s %4s\n", "VPC", "SNET", "SG", "CF", "EB", "CD", "DDB", "EFS", "L")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------")
	for region, pName := range common.RegionMap {
		goGroup.Add(1)
		go getResourceCounts(region, pName, goGroup, errChan)
	}

	go func() {
		for err := range errChan {
			print(err.Error())
		}
	}()

}
