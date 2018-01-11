# go-aws-utils
CLI utilities written in go using aws go sdk

Makes use of a go routine and WaitGroup in order to get faster execution time against all 15 normal public aws regions.

## instancelist

Currently outputs the following (print skips empty regions):

```
Oregon                  us-west-2 : 3   
Instance ID               State        OS        Type            VPC            Subnet           Public IP         Private IP        Backup  Name
-------------------------------------------------------------------------------------------------------------------------------------------------
i-xxxxxxxxxxxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.xx      172.31.20.191
i-xxxxxxxxxxxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.xx      172.31.20.224               test
i-xxxxxxxxxxxxxxxxx      stopped      Windows   t2.small        vpc-xxxxxxxx  subnet-xxxxxxxx                    172.18.33.19        False   test

```

This should eventually go into awsresources as a switch, but good coding is hard and I needed the original version of this much faster.

## awsresources

Currently outputs the following:

* Snet = Subnets
* CD = CodeDeploy
* CF = Number of Cloudformation stacks
* ECS = Number of ECS clusters
* L = Number of Lambda functions
* SS = Number of Snapshots
* DDB = Number of DynamoDB tables

Rest should be fairly obvious.

Additonal NOTES:

* Snapshots has a hidden dependency. If the **AWS_ACCOUNT** environment variable is not set it will return -1.

```
                           Region :  EC2  ECS  RDS  EBS   SS  ELB  ASG VPC SNET   SG   CF   EB   CD  DDB  EFS    L
-------------------------------------------------------------------------------------------------------------------
     N. Virginia        us-east-1 :   15    1    0   26    7    0    2   7   21   59   14    0    1    3    1   11
            Ohio        us-east-2 :    2    1    0    4    8    0    1   3   11   10    8    1    2    0    0    1
   N. California        us-west-1 :    0    0    0    0    0    0    0   1    2    1    1    1    0    0    0    0
Canada (Central)     ca-central-1 :    1    1    0    3    2    1    1   3   10    8    4    0    0    0    0    0
          Oregon        us-west-2 :    1    5    1    3    2    1    0   2    7    5    2    1    0    0    1    2
          London        eu-west-2 :    0    0    0    0    0    0    0   2    6    3    2    0    0    0    0    0
         Ireland        eu-west-1 :    0    1    0    0    0    0    0   2    7    2    3    0    1    0    0    0
           Paris        eu-west-3 :    0    0    0    0    0    0    0   1    3    1    0    0    0    0    0    0
       São Paulo        sa-east-1 :    0    0    0    0    0    0    0   1    3    1    1    0    0    0    0    0
       Frankfurt     eu-central-1 :    0    1    0    0    0    0    0   1    3    1    1    0    0    0    0    0
           Tokyo   ap-northeast-1 :    0    0    0    0    0    0    0   1    2    1    1    0    0    0    0    0
          Sydney   ap-southeast-2 :    0    0    0    0    0    0    0   1    3    1    1    1    2    0    0    0
          Mumbai       ap-south-1 :    0    0    0    0    0    0    0   1    2    1    1    0    0    0    0    0
       Singapore   ap-southeast-1 :    0    0    0    0    0    0    0   1    2    1    1    1    0    0    0    0
           Seoul   ap-northeast-2 :    0    0    0    0    0    0    0   1    2    1    1    1    0    0    0    0
```
