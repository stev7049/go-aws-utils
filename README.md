# go-aws-utils
CLI utilities written in go using aws go sdk

Makes use of a go routine and WaitGroup in order to get faster execution time against all 10 aws regions.

## instancelist

Currently outputs the following (print skips empty regions):

```
Oregon                  us-west-2 : 2   
Instance ID      State        OS         Type               VPC          Subnet        Public IP        Private IP       Name
----------------------------------------------------------------------------------------------------------------------------------------
i-xxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.xx      172.31.20.191                                  
i-xxxxxxxx      running      Linux     t2.micro        vpc-xxxxxxxx  subnet-xxxxxxxx   xx.xx.xx.x       172.31.20.224    test      
```

This should eventually go into awsresources as a swith, but good coding is hard and I needed the original version of this much faster.

## awsresources

Currently outputs the following:

* Snet = Subnets
* CD = CodeDeploy
* CF = Number of Cloudformation stacks

Rest should e fairly obvious

```
                           Region :  EC2  RDS  EBS  ELB  ASG VPC SNET   SG   CF Bean   CD
-----------------------------------------------------------------------------------------------------------
   N. California        us-west-1 :    3    0    4    2    2   4   11   15    4    2    0
          Oregon        us-west-2 :    2    0   10    0    0   2    3   12    1    2    0
     N. Virginia        us-east-1 :   12    3   46    8    5   6   19   82    8    1    6
         Ireland        eu-west-1 :    0    0    0    0    0   1    3    3    0    0    0
       Frankfurt     eu-central-1 :    0    0    0    0    0   1    2    1    0    0    0
       São Paulo        sa-east-1 :    0    0    0    0    0   1    3    1    0    0    0
           Tokyo   ap-northeast-1 :    0    0    0    0    0   1    2    1    0    0    0
           Seoul   ap-northeast-2 :    0    0    0    0    0   1    2    1    0    1    0
       Singapore   ap-southeast-1 :    0    0    0    0    0   1    2    2    0    1    0
          Sydney   ap-southeast-2 :    0    0    0    0    0   1    3    1    0    1    0
          Mumbai       ap-south-1 :    0    0    0    0    0   1    2    1    0    0    0
```
