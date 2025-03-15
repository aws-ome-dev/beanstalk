# beanstalk

setting up a Web (Load-Balanced) Environment

1. Install eb cli `pip install awsebcli` \
    binary will be installed in `/Users/<USER>/Library/Python/3.9/bin/eb`
2. initialize EB application `eb init` \
    ```sh
    Select a default region
    9) ap-northeast-1 : Asia Pacific (Tokyo)
    
    Select an application to use
    1) beanstalk
    (default is 1): 1
    
    Select a platform.
    4) Go
    (make a selection): 4
    
    Select a platform branch.
    1) Go 1 running on 64bit Amazon Linux 2023
    (default is 1): 1
    
    Do you wish to continue with CodeCommit? (Y/n): n
    Do you want to set up SSH for your instances?
    (Y/n): Y
    
    Type a keypair name.
    (Default is aws-eb): 
    Generating public/private ed25519 key pair.
    Your identification has been saved in ...
    Your public key has been saved in ...
    The key fingerprint is:
    SHA256:XXX+XXX aws-eb
    The key's randomart image is:
    +--[ED25519 256]--+
    ...
    ```

3. create EB environment `eb create`
    ```
   ➜  beanstalk git:(main) ✗ /Users/<USER>/Library/Python/3.9/bin/eb create
    Enter Environment Name
    (default is beanstalk-dev):
    Enter DNS CNAME prefix
    (default is beanstalk-dev):
    
    Select a load balancer type
    1) classic
       2) application
       3) network
          (default is 2): 2
    
    
    Would you like to enable Spot Fleet requests for this environment? (y/N): y
    Enter a list of one or more valid EC2 instance types separated by commas (at least two instance types are recommended).
    (Defaults provided on Enter): t2.nano
    
    
    2.0+ Platforms require a service role. We will attempt to create one for you. You can specify your own role using the --service-role option.
    Type "view" to see the policy, or just press ENTER to continue: view
    {
    "Version": "2012-10-17",
    "Statement": [
    {
    "Effect": "Allow",
    "Action": [
    "elasticloadbalancing:DescribeInstanceHealth",
    "elasticloadbalancing:DescribeLoadBalancers",
    "elasticloadbalancing:DescribeTargetHealth",
    "ec2:DescribeInstances",
    "ec2:DescribeInstanceStatus",
    "ec2:GetConsoleOutput",
    "ec2:AssociateAddress",
    "ec2:DescribeAddresses",
    "ec2:DescribeSecurityGroups",
    "sqs:GetQueueAttributes",
    "sqs:GetQueueUrl",
    "autoscaling:DescribeAutoScalingGroups",
    "autoscaling:DescribeAutoScalingInstances",
    "autoscaling:DescribeScalingActivities",
    "autoscaling:DescribeNotificationConfigurations",
    "sns:Publish"
    ],
    "Resource": [
    "*"
    ]
    },
    {
    "Effect": "Allow",
    "Action": [
    "logs:DescribeLogStreams",
    "logs:CreateLogStream",
    "logs:PutLogEvents"
    ],
    "Resource": "arn:aws:logs:*:*:log-group:/aws/elasticbeanstalk/*:log-stream:*"
    }
    ]
    }
    {
    "Version": "2012-10-17",
    "Statement": [
    {
    "Sid": "ElasticBeanstalkPermissions",
    "Effect": "Allow",
    "Action": [
    "elasticbeanstalk:*"
    ],
    "Resource": "*"
    },
    {
    "Sid": "AllowPassRoleToElasticBeanstalkAndDownstreamServices",
    "Effect": "Allow",
    "Action": "iam:PassRole",
    "Resource": "arn:aws:iam::*:role/*",
    "Condition": {
    "StringEquals": {
    "iam:PassedToService": [
    "elasticbeanstalk.amazonaws.com",
    "ec2.amazonaws.com",
    "ec2.amazonaws.com.cn",
    "autoscaling.amazonaws.com",
    "elasticloadbalancing.amazonaws.com",
    "ecs.amazonaws.com",
    "cloudformation.amazonaws.com"
    ]
    }
    }
    },
    {
    "Sid": "ReadOnlyPermissions",
    "Effect": "Allow",
    "Action": [
    "autoscaling:DescribeAccountLimits",
    "autoscaling:DescribeAutoScalingGroups",
    "autoscaling:DescribeAutoScalingInstances",
    "autoscaling:DescribeLaunchConfigurations",
    "autoscaling:DescribeLoadBalancers",
    "autoscaling:DescribeNotificationConfigurations",
    "autoscaling:DescribeScalingActivities",
    "autoscaling:DescribeScheduledActions",
    "ec2:DescribeAccountAttributes",
    "ec2:DescribeAddresses",
    "ec2:DescribeAvailabilityZones",
    "ec2:DescribeImages",
    "ec2:DescribeInstanceAttribute",
    "ec2:DescribeInstances",
    "ec2:DescribeKeyPairs",
    "ec2:DescribeLaunchTemplates",
    "ec2:DescribeLaunchTemplateVersions",
    "ec2:DescribeSecurityGroups",
    "ec2:DescribeSnapshots",
    "ec2:DescribeSpotInstanceRequests",
    "ec2:DescribeSubnets",
    "ec2:DescribeVpcClassicLink",
    "ec2:DescribeVpcs",
    "elasticloadbalancing:DescribeInstanceHealth",
    "elasticloadbalancing:DescribeLoadBalancers",
    "elasticloadbalancing:DescribeTargetGroups",
    "elasticloadbalancing:DescribeTargetHealth",
    "logs:DescribeLogGroups",
    "rds:DescribeDBEngineVersions",
    "rds:DescribeDBInstances",
    "rds:DescribeOrderableDBInstanceOptions",
    "sns:ListSubscriptionsByTopic"
    ],
    "Resource": [
    "*"
    ]
    },
    {
    "Sid": "EC2BroadOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "ec2:AllocateAddress",
    "ec2:AssociateAddress",
    "ec2:AuthorizeSecurityGroupEgress",
    "ec2:AuthorizeSecurityGroupIngress",
    "ec2:CreateLaunchTemplate",
    "ec2:CreateLaunchTemplateVersion",
    "ec2:CreateSecurityGroup",
    "ec2:DeleteLaunchTemplate",
    "ec2:DeleteLaunchTemplateVersions",
    "ec2:DeleteSecurityGroup",
    "ec2:DisassociateAddress",
    "ec2:ReleaseAddress",
    "ec2:RevokeSecurityGroupEgress",
    "ec2:RevokeSecurityGroupIngress"
    ],
    "Resource": "*"
    },
    {
    "Sid": "EC2RunInstancesOperationPermissions",
    "Effect": "Allow",
    "Action": "ec2:RunInstances",
    "Resource": "*",
    "Condition": {
    "ArnLike": {
    "ec2:LaunchTemplate": "arn:aws:ec2:*:*:launch-template/*"
    }
    }
    },
    {
    "Sid": "EC2TerminateInstancesOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "ec2:TerminateInstances"
    ],
    "Resource": "arn:aws:ec2:*:*:instance/*",
    "Condition": {
    "StringLike": {
    "ec2:ResourceTag/aws:cloudformation:stack-id": [
    "arn:aws:cloudformation:*:*:stack/awseb-e-*",
    "arn:aws:cloudformation:*:*:stack/eb-*"
    ]
    }
    }
    },
    {
    "Sid": "ECSBroadOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "ecs:CreateCluster",
    "ecs:DescribeClusters",
    "ecs:RegisterTaskDefinition"
    ],
    "Resource": "*"
    },
    {
    "Sid": "ECSDeleteClusterOperationPermissions",
    "Effect": "Allow",
    "Action": "ecs:DeleteCluster",
    "Resource": "arn:aws:ecs:*:*:cluster/awseb-*"
    },
    {
    "Sid": "ASGOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "autoscaling:AttachInstances",
    "autoscaling:CreateAutoScalingGroup",
    "autoscaling:CreateLaunchConfiguration",
    "autoscaling:CreateOrUpdateTags",
    "autoscaling:DeleteLaunchConfiguration",
    "autoscaling:DeleteAutoScalingGroup",
    "autoscaling:DeleteScheduledAction",
    "autoscaling:DetachInstances",
    "autoscaling:DeletePolicy",
    "autoscaling:PutScalingPolicy",
    "autoscaling:PutScheduledUpdateGroupAction",
    "autoscaling:PutNotificationConfiguration",
    "autoscaling:ResumeProcesses",
    "autoscaling:SetDesiredCapacity",
    "autoscaling:SuspendProcesses",
    "autoscaling:TerminateInstanceInAutoScalingGroup",
    "autoscaling:UpdateAutoScalingGroup"
    ],
    "Resource": [
    "arn:aws:autoscaling:*:*:launchConfiguration:*:launchConfigurationName/awseb-e-*",
    "arn:aws:autoscaling:*:*:launchConfiguration:*:launchConfigurationName/eb-*",
    "arn:aws:autoscaling:*:*:autoScalingGroup:*:autoScalingGroupName/awseb-e-*",
    "arn:aws:autoscaling:*:*:autoScalingGroup:*:autoScalingGroupName/eb-*"
    ]
    },
    {
    "Sid": "CFNOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "cloudformation:*"
    ],
    "Resource": [
    "arn:aws:cloudformation:*:*:stack/awseb-*",
    "arn:aws:cloudformation:*:*:stack/eb-*"
    ]
    },
    {
    "Sid": "ELBOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "elasticloadbalancing:AddTags",
    "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
    "elasticloadbalancing:ConfigureHealthCheck",
    "elasticloadbalancing:CreateLoadBalancer",
    "elasticloadbalancing:DeleteLoadBalancer",
    "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
    "elasticloadbalancing:DeregisterTargets",
    "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
    "elasticloadbalancing:RegisterTargets"
    ],
    "Resource": [
    "arn:aws:elasticloadbalancing:*:*:targetgroup/awseb-*",
    "arn:aws:elasticloadbalancing:*:*:targetgroup/eb-*",
    "arn:aws:elasticloadbalancing:*:*:loadbalancer/awseb-*",
    "arn:aws:elasticloadbalancing:*:*:loadbalancer/eb-*",
    "arn:aws:elasticloadbalancing:*:*:loadbalancer/*/awseb-*/*",
    "arn:aws:elasticloadbalancing:*:*:loadbalancer/*/eb-*/*"
    ]
    },
    {
    "Sid": "CWLogsOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "logs:CreateLogGroup",
    "logs:DeleteLogGroup",
    "logs:PutRetentionPolicy"
    ],
    "Resource": "arn:aws:logs:*:*:log-group:/aws/elasticbeanstalk/*"
    },
    {
    "Sid": "S3ObjectOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "s3:DeleteObject",
    "s3:GetObject",
    "s3:GetObjectAcl",
    "s3:GetObjectVersion",
    "s3:GetObjectVersionAcl",
    "s3:PutObject",
    "s3:PutObjectAcl",
    "s3:PutObjectVersionAcl"
    ],
    "Resource": "arn:aws:s3:::elasticbeanstalk-*/*"
    },
    {
    "Sid": "S3BucketOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "s3:GetBucketLocation",
    "s3:GetBucketPolicy",
    "s3:ListBucket",
    "s3:PutBucketPolicy"
    ],
    "Resource": "arn:aws:s3:::elasticbeanstalk-*"
    },
    {
    "Sid": "SNSOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "sns:CreateTopic",
    "sns:GetTopicAttributes",
    "sns:SetTopicAttributes",
    "sns:Subscribe"
    ],
    "Resource": "arn:aws:sns:*:*:ElasticBeanstalkNotifications-*"
    },
    {
    "Sid": "SQSOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "sqs:GetQueueAttributes",
    "sqs:GetQueueUrl"
    ],
    "Resource": [
    "arn:aws:sqs:*:*:awseb-e-*",
    "arn:aws:sqs:*:*:eb-*"
    ]
    },
    {
    "Sid": "CWPutMetricAlarmOperationPermissions",
    "Effect": "Allow",
    "Action": [
    "cloudwatch:PutMetricAlarm"
    ],
    "Resource": [
    "arn:aws:cloudwatch:*:*:alarm:awseb-*",
    "arn:aws:cloudwatch:*:*:alarm:eb-*"
    ]
    },
    {
    "Sid": "AllowECSTagResource",
    "Effect": "Allow",
    "Action": [
    "ecs:TagResource"
    ],
    "Resource": "*",
    "Condition": {
    "StringEquals": {
    "ecs:CreateAction": [
    "CreateCluster",
    "RegisterTaskDefinition"
    ]
    }
    }
    },
    {
    "Sid": "LaunchTemplateTagPropagationPermissions",
    "Effect": "Allow",
    "Action": "ec2:createTags",
    "Resource": "*",
    "Condition": {
    "StringEquals": {
    "ec2:CreateAction": [
    "CreateLaunchTemplate",
    "RunInstances"
    ]
    }
    }
    }
    ]
    }
    Press enter to continue:
    Creating application version archive "app-469b-250315_202230347740".
    Uploading beanstalk/app-469b-250315_202230347740.zip to S3. This may take a while.
    Upload Complete.
    Environment details for: beanstalk-dev
    Application name: beanstalk
    Region: ap-northeast-1
    Deployed Version: app-469b-250315_202230347740
    Environment ID: e-ebhm2svymf
    Platform: arn:aws:elasticbeanstalk:ap-northeast-1::platform/Go 1 running on 64bit Amazon Linux 2023/4.2.4
    Tier: WebServer-Standard-1.0
    CNAME: beanstalk-dev.ap-northeast-1.elasticbeanstalk.com
    Updated: 2025-03-15 12:22:35.763000+00:00
    Printing Status:
    2025-03-15 12:22:33    INFO    createEnvironment is starting.
    2025-03-15 12:22:35    INFO    Using elasticbeanstalk-ap-northeast-1-69
    ```
   
4. After `eb create`, health should be grey (suspended):
    ```
    ➜  beanstalk git:(main) ✗ /Users/<USER>/Library/Python/3.9/bin/eb status
    Environment details for: beanstalk-dev
      Application name: beanstalk
      Region: ap-northeast-1
      Deployed Version: app-469b-250315_202230347740
      Environment ID: e-ebhm2svymf
      Platform: arn:aws:elasticbeanstalk:ap-northeast-1::platform/Go 1 running on 64bit Amazon Linux 2023/4.2.4
      Tier: WebServer-Standard-1.0
      CNAME: beanstalk-dev.ap-northeast-1.elasticbeanstalk.com
      Updated: 2025-03-15 12:27:07.024000+00:00
      Status: Ready
      Health: Grey
    ```
   
    Execute `eb deploy`