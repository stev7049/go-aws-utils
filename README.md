# go-aws-utils
CLI utilities written in go using aws go sdk

Makes use of a go routine and WaitGroup in order to get faster execution time against all 12 normal public aws regions.

## instancelist

Currently outputs the following (print skips empty regions):

```
Oregon                  us-west-2 : 2   
Instance ID      State        OS         Type               VPC          Subnet        Public IP        Private IP       Name
----------------------------------------------------------------------------------------------------------------------------------------
i-xxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.xx      172.31.20.191                                  
i-xxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.x       172.31.20.224    test      
```

This should eventually go into awsresources as a switch, but good coding is hard and I needed the original version of this much faster.

## awsresources

Currently outputs the following:

* Snet = Subnets
* CD = CodeDeploy
* CF = Number of Cloudformation stacks

Rest should e fairly obvious

```
                           Region :  EC2  RDS  EBS  ELB  ASG VPC SNET   SG   CF Bean   CD
-----------------------------------------------------------------------------------------------------------
            Ohio        us-east-2 :    0    0    0    0    0   1    3    1    0    0    0
   N. California        us-west-1 :    6    0    6    0    0   5   15   13    3    1    0
          Oregon        us-west-2 :    3    0   11    2    1   4   11   23    5    2    0
     N. Virginia        us-east-1 :   20    2   56    7    5   7   22  104   10    3    5
         Ireland        eu-west-1 :    1    1    1    0    0   2    7   13    2    0    1
       Frankfurt     eu-central-1 :    0    0    0    0    0   1    2    1    0    0    0
           Tokyo   ap-northeast-1 :    0    0    0    0    0   1    2    1    0    0    0
       São Paulo        sa-east-1 :    0    0    0    0    0   1    3    1    0    0    0
          Sydney   ap-southeast-2 :    0    0    0    0    0   2    5    4    0    1    2
       Singapore   ap-southeast-1 :    0    0    0    0    0   1    2    2    0    1    0
          Mumbai       ap-south-1 :    0    0    0    0    0   1    2    1    0    0    0
           Seoul   ap-northeast-2 :    0    0    0    0    0   1    2    1    0    1    0
```
