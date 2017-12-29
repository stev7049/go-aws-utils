package main

import (
	"fmt"
	"sync"

	"github.com/mdfilio/go-aws-utils/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func regionError(region string, err error) error {
	return fmt.Errorf("Error occured in region %s: %v\n", region, err)
}

func getInstances(region string, humanregion string, goGroup *sync.WaitGroup, errChan chan error) {
	defer goGroup.Done()

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		errChan <- regionError(region, err)
	}

	var totalInstances int
	for _, res := range resp.Reservations {
		totalInstances += len(res.Instances)
	}

	if totalInstances > 0 {
		fmt.Printf("%-25s %+16s : %-4d\n", humanregion, region, totalInstances)
		fmt.Println("Instance ID               State        OS        Type            VPC            Subnet           Public IP         Private IP        Backup  Name")
		fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	}

	instanceCounter := 0
	for idx, res := range resp.Reservations {
		if len(res.Instances) > 0 {
			for _, inst := range resp.Reservations[idx].Instances {
				instanceCounter++
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

				thisBackup := ""
				thisName := ""
				for tag := range inst.Tags {
					if *inst.Tags[tag].Key == "Name" {
						thisName = *inst.Tags[tag].Value
					}
					if *inst.Tags[tag].Key == "Backup" || *inst.Tags[tag].Key == "backup" {
						thisBackup = *inst.Tags[tag].Value
					}
				}

				fmt.Printf("%-25s %-12s %-9s %-15s %-14s %-16s %-17s %-17s %-7s %-30s\n", thisInstanceID, thisState, thisPlatform, thisInstanceType, thisVpcID, thisSubnetID, thisPublicIpAddress, thisPrivateIpAddress, thisBackup, thisName)
				if instanceCounter == totalInstances {
					fmt.Printf("\n")
				}
			}
		}
	}
}

func main() {

	goGroup := new(sync.WaitGroup)

	errChan := make(chan error)
	defer func() {
		close(errChan)
	}()
	defer goGroup.Wait()

	for region, pName := range common.RegionMap {
		goGroup.Add(1)
		go getInstances(region, pName, goGroup, errChan)
	}

	go func() {
		for err := range errChan {
			print(err.Error())
		}
	}()

}
