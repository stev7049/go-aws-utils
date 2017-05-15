package main

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	//    "github.com/mgutz/ansi"
)

func getInstances(region string, humanregion string, goGroup *sync.WaitGroup) {
	defer goGroup.Done()
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	var totalInstances int = 0
	for _, res := range resp.Reservations {
		totalInstances = totalInstances + len(res.Instances)
	}

	if totalInstances > 0 {
		fmt.Printf("%-25s %+16s : %-4d\n", humanregion, region, totalInstances)
		fmt.Println("Instance ID               State        OS        Type            VPC            Subnet           Public IP         Private IP        Name")
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------")
	}

	instanceCounter := 0
	for idx, res := range resp.Reservations {
		if len(res.Instances) > 0 {
			for _, inst := range resp.Reservations[idx].Instances {
				instanceCounter += 1
				thisInstanceID := ""
				thisPrivateIpAddress := ""
				thisPublicIpAddress := ""
				thisState := ""
				thisPlatform := ""
				thisSubnetID := ""
				thisVpcID := ""
				thisInstanceType := ""

				if inst.InstanceId != nil {
					thisInstanceID = *inst.InstanceId
				}

				if inst.PublicIpAddress != nil {
					thisPublicIpAddress = *inst.PublicIpAddress
				}
				if inst.PrivateIpAddress != nil {
					thisPrivateIpAddress = *inst.PrivateIpAddress
				}
				if inst.State.Name != nil {
					thisState = *inst.State.Name
				}
				if inst.Platform != nil {
					thisPlatform = "Windows"
				} else {
					thisPlatform = "Linux"
				}
				if inst.VpcId != nil {
					thisVpcID = *inst.VpcId
				}
				if inst.SubnetId != nil {
					thisSubnetID = *inst.SubnetId
				}
				if inst.InstanceType != nil {
					thisInstanceType = *inst.InstanceType
				}

				thisName := ""
				for tag := range inst.Tags {
					if *inst.Tags[tag].Key == "Name" {
						thisName = *inst.Tags[tag].Value
					}
				}

				//fmt.Printf("\033[34m%-15s\033[0m \033[34m%-12s\033[0m \033[34m%-9s\033[0m \033[34m%-15s\033[0m \033[34m%-9s\033[0m  \033[34m%-9s\033[0m   \033[34m%-15s\033[0m  \033[34m%-15s\033[0m  \033[34m%-30s\033[0m\n", thisInstanceID, thisState, thisPlatform, thisInstanceType, thisVpcID, thisSubnetID, thisPublicIpAddress, thisPrivateIpAddress, thisName)
				fmt.Printf("%-25s %-12s %-9s %-15s %-14s %-16s %-17s %-17s %-30s\n", thisInstanceID, thisState, thisPlatform, thisInstanceType, thisVpcID, thisSubnetID, thisPublicIpAddress, thisPrivateIpAddress, thisName)
				if instanceCounter == totalInstances {
					fmt.Printf("\n")
				}
			}
		}
	}
}

func main() {

	awsregions := map[int][]string{
		0:  {"us-east-1", "N. Virginia"},
		1:  {"us-east-2", "Ohio"},
		2:  {"us-west-1", "N. California"},
		3:  {"us-west-2", "Oregon"},
		4:  {"ca-central-1", "Canada (Central)"},
		5:  {"eu-west-1", "Ireland"},
		6:  {"eu-central-1", "Frankfurt"},
		7:  {"eu-west-2", "London"},
		8:  {"ap-southeast-1", "Singapore"},
		9:  {"ap-southeast-2", "Sydney"},
		10: {"ap-northeast-1", "Tokyo"},
		11: {"ap-northeast-2", "Seoul"},
		12: {"sa-east-1", "São Paulo"},
		13: {"ap-south-1", "Mumbai"},
	}
	goGroup := new(sync.WaitGroup)
	defer goGroup.Wait()

	for i := 0; i < len(awsregions); i++ {
		goGroup.Add(1)
		go getInstances(awsregions[i][0], awsregions[i][1], goGroup)
	}

}
