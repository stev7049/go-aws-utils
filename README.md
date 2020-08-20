# go-aws-utils
Forked from https://github.com/mdfilio/go-aws-utils
#
https://docs.aws.amazon.com/sdk-for-go/api/
#
CLI utilities written in go using aws go sdk

Makes use of a Go Language to pull data against all 22 aws regions (as of 2020).
#
Data Univailable reported for Disabled or Regions in Error State.


## instancelist

Currently outputs the following (print reports "Data Univailable" for Disabled or Error State Regions):

```
Data Univailable for: ca-central-1
Oregon                  us-west-2 : 3   
Instance ID               State        OS        Type            VPC            Subnet           Public IP         Private IP        Backup  Name
-------------------------------------------------------------------------------------------------------------------------------------------------
i-xxxxxxxxxxxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.xx      172.31.20.191
i-xxxxxxxxxxxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.xx      172.31.20.224               test
i-xxxxxxxxxxxxxxxxx      stopped      Windows   t2.small        vpc-xxxxxxxx  subnet-xxxxxxxx                    172.18.33.19        False   test
Data Univailable for: ap-southeast-2
Data Univailable for: me-south-1
```


## awsresources

Currently outputs the following:

* EC2 = EC2 Instances
* ECS = Elastic Container Clusters
* RDS = RDS Databases
* EBS = Elastic Block Storage
* ELB = Load Balancer
* ASG = AutoScale Groups
* VPC = VPCs 
* SNET = Subnets
* SG = Security Groups
* CF = Number of Cloudformation stacks
* EB = Elastic BeanStalk
* CD = CodeDeploy
* DDB = Number of DynamoDB tables
* EFS = EFS Filesystem Count
* L = Number of Lambda functions


Additonal NOTES:

```
                           Region :  EC2  ECS  RDS  EBS    ELB  ASG VPC SNET   SG   CF   EB   CD  DDB  EFS   
---------------------------------------------------------------------------------------------------------------
Data Univailable for: me-south-1
Data Univailable for: af-south-1
     N. Virginia        us-east-1 :   15    1    0   26    0    2   7   21   59   14    0    1    3    1   11
            Ohio        us-east-2 :    2    1    0    4    0    1   3   11   10    8    1    2    0    0    1
   N. California        us-west-1 :    0    0    0    0    0    0   1    2    1    1    1    0    0    0    0
Canada (Central)     ca-central-1 :    1    1    0    3    1    1   3   10    8    4    0    0    0    0    0
          Oregon        us-west-2 :    1    5    1    3    1    0   2    7    5    2    1    0    0    1    2
          London        eu-west-2 :    0    0    0    0    0    0   2    6    3    2    0    0    0    0    0
         Ireland        eu-west-1 :    0    1    0    0    0    0   2    7    2    3    0    1    0    0    0
           Paris        eu-west-3 :    0    0    0    0    0    0   1    3    1    0    0    0    0    0    0
       São Paulo        sa-east-1 :    0    0    0    0    0    0   1    3    1    1    0    0    0    0    0
       Frankfurt     eu-central-1 :    0    1    0    0    0    0   1    3    1    1    0    0    0    0    0
           Tokyo   ap-northeast-1 :    0    0    0    0    0    0   1    2    1    1    0    0    0    0    0
          Sydney   ap-southeast-2 :    0    0    0    0    0    0   1    3    1    1    1    2    0    0    0
          Mumbai       ap-south-1 :    0    0    0    0    0    0   1    2    1    1    0    0    0    0    0
       Singapore   ap-southeast-1 :    0    0    0    0    0    0   1    2    1    1    1    0    0    0    0
           Seoul   ap-northeast-2 :    0    0    0    0    0    0   1    2    1    1    1    0    0    0    0
```
