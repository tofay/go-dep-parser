package bundler

import "github.com/aquasecurity/go-dep-parser/pkg/types"

var (
	// docker run --name bundler --rm -it ruby:2.6 bash
	// bundle init
	// bundle add dotenv json faker rubocop pry
	// bundler show | grep "*" | grep -v bundler | awk '{if(match($0, /\((.*)\)/)) printf("{\""$2"\", \""substr($0, RSTART+1, RLENGTH-2)"\", \"\"},\n");}'
	BundlerNormal = []types.Library{
		{Name: "ast", Version: "2.4.0"},
		{Name: "coderay", Version: "1.1.2"},
		{Name: "concurrent-ruby", Version: "1.1.5"},
		{Name: "dotenv", Version: "2.7.2"},
		{Name: "faker", Version: "1.9.3"},
		{Name: "i18n", Version: "1.6.0"},
		{Name: "jaro_winkler", Version: "1.5.2"},
		{Name: "json", Version: "2.2.0"},
		{Name: "method_source", Version: "0.9.2"},
		{Name: "parallel", Version: "1.17.0"},
		{Name: "parser", Version: "2.6.3.0"},
		{Name: "pry", Version: "0.12.2"},
		{Name: "psych", Version: "3.1.0"},
		{Name: "rainbow", Version: "3.0.0"},
		{Name: "rubocop", Version: "0.67.2"},
		{Name: "ruby-progressbar", Version: "1.10.0"},
		{Name: "unicode-display_width", Version: "1.5.0"},
	}

	// docker run --name bundler --rm -it ruby:2.6 bash
	// bundle init
	// bundle add dotenv json faker rubocop pry
	// bundle add rails
	// bundler show | grep "*" | grep -v bundler | awk '{if(match($0, /\((.*)\)/)) printf("{\""$2"\", \""substr($0, RSTART+1, RLENGTH-2)"\", \"\"},\n");}'
	BundlerRails = []types.Library{
		{Name: "actioncable", Version: "5.2.3"},
		{Name: "actionmailer", Version: "5.2.3"},
		{Name: "actionpack", Version: "5.2.3"},
		{Name: "actionview", Version: "5.2.3"},
		{Name: "activejob", Version: "5.2.3"},
		{Name: "activemodel", Version: "5.2.3"},
		{Name: "activerecord", Version: "5.2.3"},
		{Name: "activestorage", Version: "5.2.3"},
		{Name: "activesupport", Version: "5.2.3"},
		{Name: "arel", Version: "9.0.0"},
		{Name: "ast", Version: "2.4.0"},
		{Name: "builder", Version: "3.2.3"},
		{Name: "coderay", Version: "1.1.2"},
		{Name: "concurrent-ruby", Version: "1.1.5"},
		{Name: "crass", Version: "1.0.4"},
		{Name: "dotenv", Version: "2.7.2"},
		{Name: "erubi", Version: "1.8.0"},
		{Name: "faker", Version: "1.9.3"},
		{Name: "globalid", Version: "0.4.2"},
		{Name: "i18n", Version: "1.6.0"},
		{Name: "jaro_winkler", Version: "1.5.2"},
		{Name: "json", Version: "2.2.0"},
		{Name: "loofah", Version: "2.2.3"},
		{Name: "mail", Version: "2.7.1"},
		{Name: "marcel", Version: "0.3.3"},
		{Name: "method_source", Version: "0.9.2"},
		{Name: "mimemagic", Version: "0.3.3"},
		{Name: "mini_mime", Version: "1.0.1"},
		{Name: "mini_portile2", Version: "2.4.0"},
		{Name: "minitest", Version: "5.11.3"},
		{Name: "nio4r", Version: "2.3.1"},
		{Name: "nokogiri", Version: "1.10.3"},
		{Name: "parallel", Version: "1.17.0"},
		{Name: "parser", Version: "2.6.3.0"},
		{Name: "pry", Version: "0.12.2"},
		{Name: "psych", Version: "3.1.0"},
		{Name: "rack", Version: "2.0.7"},
		{Name: "rack-test", Version: "1.1.0"},
		{Name: "rails", Version: "5.2.3"},
		{Name: "rails-dom-testing", Version: "2.0.3"},
		{Name: "rails-html-sanitizer", Version: "1.0.4"},
		{Name: "railties", Version: "5.2.3"},
		{Name: "rainbow", Version: "3.0.0"},
		{Name: "rake", Version: "12.3.2"},
		{Name: "rubocop", Version: "0.67.2"},
		{Name: "ruby-progressbar", Version: "1.10.0"},
		{Name: "sprockets", Version: "3.7.2"},
		{Name: "sprockets-rails", Version: "3.2.1"},
		{Name: "thor", Version: "0.20.3"},
		{Name: "thread_safe", Version: "0.3.6"},
		{Name: "tzinfo", Version: "1.2.5"},
		{Name: "unicode-display_width", Version: "1.5.0"},
		{Name: "websocket-driver", Version: "0.7.0"},
		{Name: "websocket-extensions", Version: "0.1.3"},
	}
	// docker run --name bundler --rm -it ruby:2.6 bash
	// bundle init
	// bundle add dotenv json faker rubocop pry
	// bundle add rails
	// bundle add sinatra multi-json thor sass aws-sdk faraday
	// bundler show | grep "*" | grep -v bundler | awk '{if(match($0, /\((.*)\)/)) printf("{\""$2"\", \""substr($0, RSTART+1, RLENGTH-2)"\"}, \"\"},\n");}'
	BundlerMany = []types.Library{
		{Name: "actioncable", Version: "5.2.3"},
		{Name: "actionmailer", Version: "5.2.3"},
		{Name: "actionpack", Version: "5.2.3"},
		{Name: "actionview", Version: "5.2.3"},
		{Name: "activejob", Version: "5.2.3"},
		{Name: "activemodel", Version: "5.2.3"},
		{Name: "activerecord", Version: "5.2.3"},
		{Name: "activestorage", Version: "5.2.3"},
		{Name: "activesupport", Version: "5.2.3"},
		{Name: "arel", Version: "9.0.0"},
		{Name: "ast", Version: "2.4.0"},
		{Name: "aws-eventstream", Version: "1.0.3"},
		{Name: "aws-partitions", Version: "1.154.0"},
		{Name: "aws-sdk", Version: "3.0.1"},
		{Name: "aws-sdk-acm", Version: "1.19.0"},
		{Name: "aws-sdk-acmpca", Version: "1.13.0"},
		{Name: "aws-sdk-alexaforbusiness", Version: "1.20.0"},
		{Name: "aws-sdk-amplify", Version: "1.3.0"},
		{Name: "aws-sdk-apigateway", Version: "1.26.0"},
		{Name: "aws-sdk-apigatewaymanagementapi", Version: "1.3.0"},
		{Name: "aws-sdk-apigatewayv2", Version: "1.4.0"},
		{Name: "aws-sdk-applicationautoscaling", Version: "1.22.0"},
		{Name: "aws-sdk-applicationdiscoveryservice", Version: "1.15.0"},
		{Name: "aws-sdk-appmesh", Version: "1.6.0"},
		{Name: "aws-sdk-appstream", Version: "1.25.0"},
		{Name: "aws-sdk-appsync", Version: "1.12.0"},
		{Name: "aws-sdk-athena", Version: "1.12.0"},
		{Name: "aws-sdk-autoscaling", Version: "1.20.0"},
		{Name: "aws-sdk-autoscalingplans", Version: "1.12.0"},
		{Name: "aws-sdk-backup", Version: "1.3.0"},
		{Name: "aws-sdk-batch", Version: "1.17.0"},
		{Name: "aws-sdk-budgets", Version: "1.18.0"},
		{Name: "aws-sdk-chime", Version: "1.6.0"},
		{Name: "aws-sdk-cloud9", Version: "1.11.0"},
		{Name: "aws-sdk-clouddirectory", Version: "1.14.0"},
		{Name: "aws-sdk-cloudformation", Version: "1.19.0"},
		{Name: "aws-sdk-cloudfront", Version: "1.15.0"},
		{Name: "aws-sdk-cloudhsm", Version: "1.12.0"},
		{Name: "aws-sdk-cloudhsmv2", Version: "1.12.0"},
		{Name: "aws-sdk-cloudsearch", Version: "1.9.0"},
		{Name: "aws-sdk-cloudsearchdomain", Version: "1.9.0"},
		{Name: "aws-sdk-cloudtrail", Version: "1.11.0"},
		{Name: "aws-sdk-cloudwatch", Version: "1.20.0"},
		{Name: "aws-sdk-cloudwatchevents", Version: "1.17.0"},
		{Name: "aws-sdk-cloudwatchlogs", Version: "1.17.0"},
		{Name: "aws-sdk-codebuild", Version: "1.32.0"},
		{Name: "aws-sdk-codecommit", Version: "1.17.0"},
		{Name: "aws-sdk-codedeploy", Version: "1.18.0"},
		{Name: "aws-sdk-codepipeline", Version: "1.15.0"},
		{Name: "aws-sdk-codestar", Version: "1.11.0"},
		{Name: "aws-sdk-cognitoidentity", Version: "1.10.0"},
		{Name: "aws-sdk-cognitoidentityprovider", Version: "1.18.0"},
		{Name: "aws-sdk-cognitosync", Version: "1.9.0"},
		{Name: "aws-sdk-comprehend", Version: "1.18.0"},
		{Name: "aws-sdk-comprehendmedical", Version: "1.3.0"},
		{Name: "aws-sdk-configservice", Version: "1.26.0"},
		{Name: "aws-sdk-connect", Version: "1.13.0"},
		{Name: "aws-sdk-core", Version: "3.48.6"},
		{Name: "aws-sdk-costandusagereportservice", Version: "1.10.0"},
		{Name: "aws-sdk-costexplorer", Version: "1.21.0"},
		{Name: "aws-sdk-databasemigrationservice", Version: "1.20.0"},
		{Name: "aws-sdk-datapipeline", Version: "1.9.0"},
		{Name: "aws-sdk-datasync", Version: "1.3.0"},
		{Name: "aws-sdk-dax", Version: "1.11.0"},
		{Name: "aws-sdk-devicefarm", Version: "1.19.0"},
		{Name: "aws-sdk-directconnect", Version: "1.16.0"},
		{Name: "aws-sdk-directoryservice", Version: "1.15.0"},
		{Name: "aws-sdk-dlm", Version: "1.11.0"},
		{Name: "aws-sdk-docdb", Version: "1.4.0"},
		{Name: "aws-sdk-dynamodb", Version: "1.26.0"},
		{Name: "aws-sdk-dynamodbstreams", Version: "1.9.0"},
		{Name: "aws-sdk-ec2", Version: "1.80.0"},
		{Name: "aws-sdk-ecr", Version: "1.14.0"},
		{Name: "aws-sdk-ecs", Version: "1.36.0"},
		{Name: "aws-sdk-efs", Version: "1.13.0"},
		{Name: "aws-sdk-eks", Version: "1.15.0"},
		{Name: "aws-sdk-elasticache", Version: "1.14.0"},
		{Name: "aws-sdk-elasticbeanstalk", Version: "1.19.0"},
		{Name: "aws-sdk-elasticloadbalancing", Version: "1.12.0"},
		{Name: "aws-sdk-elasticloadbalancingv2", Version: "1.26.0"},
		{Name: "aws-sdk-elasticsearchservice", Version: "1.19.0"},
		{Name: "aws-sdk-elastictranscoder", Version: "1.11.0"},
		{Name: "aws-sdk-emr", Version: "1.14.0"},
		{Name: "aws-sdk-firehose", Version: "1.14.0"},
		{Name: "aws-sdk-fms", Version: "1.12.0"},
		{Name: "aws-sdk-fsx", Version: "1.4.0"},
		{Name: "aws-sdk-gamelift", Version: "1.16.0"},
		{Name: "aws-sdk-glacier", Version: "1.18.0"},
		{Name: "aws-sdk-globalaccelerator", Version: "1.4.0"},
		{Name: "aws-sdk-glue", Version: "1.30.0"},
		{Name: "aws-sdk-greengrass", Version: "1.17.0"},
		{Name: "aws-sdk-guardduty", Version: "1.14.0"},
		{Name: "aws-sdk-health", Version: "1.12.0"},
		{Name: "aws-sdk-iam", Version: "1.19.0"},
		{Name: "aws-sdk-importexport", Version: "1.9.0"},
		{Name: "aws-sdk-inspector", Version: "1.16.0"},
		{Name: "aws-sdk-iot", Version: "1.29.0"},
		{Name: "aws-sdk-iot1clickdevicesservice", Version: "1.11.0"},
		{Name: "aws-sdk-iot1clickprojects", Version: "1.10.0"},
		{Name: "aws-sdk-iotanalytics", Version: "1.16.0"},
		{Name: "aws-sdk-iotdataplane", Version: "1.9.0"},
		{Name: "aws-sdk-iotjobsdataplane", Version: "1.10.0"},
		{Name: "aws-sdk-kafka", Version: "1.4.0"},
		{Name: "aws-sdk-kinesis", Version: "1.13.1"},
		{Name: "aws-sdk-kinesisanalytics", Version: "1.12.0"},
		{Name: "aws-sdk-kinesisanalyticsv2", Version: "1.3.0"},
		{Name: "aws-sdk-kinesisvideo", Version: "1.12.0"},
		{Name: "aws-sdk-kinesisvideoarchivedmedia", Version: "1.11.0"},
		{Name: "aws-sdk-kinesisvideomedia", Version: "1.10.0"},
		{Name: "aws-sdk-kms", Version: "1.17.0"},
		{Name: "aws-sdk-lambda", Version: "1.22.0"},
		{Name: "aws-sdk-lambdapreview", Version: "1.9.0"},
		{Name: "aws-sdk-lex", Version: "1.12.0"},
		{Name: "aws-sdk-lexmodelbuildingservice", Version: "1.15.0"},
		{Name: "aws-sdk-licensemanager", Version: "1.3.0"},
		{Name: "aws-sdk-lightsail", Version: "1.18.0"},
		{Name: "aws-sdk-machinelearning", Version: "1.10.0"},
		{Name: "aws-sdk-macie", Version: "1.9.0"},
		{Name: "aws-sdk-marketplacecommerceanalytics", Version: "1.9.0"},
		{Name: "aws-sdk-marketplaceentitlementservice", Version: "1.9.0"},
		{Name: "aws-sdk-marketplacemetering", Version: "1.11.0"},
		{Name: "aws-sdk-mediaconnect", Version: "1.5.0"},
		{Name: "aws-sdk-mediaconvert", Version: "1.25.0"},
		{Name: "aws-sdk-medialive", Version: "1.28.0"},
		{Name: "aws-sdk-mediapackage", Version: "1.15.0"},
		{Name: "aws-sdk-mediastore", Version: "1.12.0"},
		{Name: "aws-sdk-mediastoredata", Version: "1.11.0"},
		{Name: "aws-sdk-mediatailor", Version: "1.14.0"},
		{Name: "aws-sdk-migrationhub", Version: "1.11.0"},
		{Name: "aws-sdk-mobile", Version: "1.9.0"},
		{Name: "aws-sdk-mq", Version: "1.13.0"},
		{Name: "aws-sdk-mturk", Version: "1.12.0"},
		{Name: "aws-sdk-neptune", Version: "1.11.0"},
		{Name: "aws-sdk-opsworks", Version: "1.13.0"},
		{Name: "aws-sdk-opsworkscm", Version: "1.16.0"},
		{Name: "aws-sdk-organizations", Version: "1.24.0"},
		{Name: "aws-sdk-pi", Version: "1.9.0"},
		{Name: "aws-sdk-pinpoint", Version: "1.19.0"},
		{Name: "aws-sdk-pinpointemail", Version: "1.6.0"},
		{Name: "aws-sdk-pinpointsmsvoice", Version: "1.6.0"},
		{Name: "aws-sdk-polly", Version: "1.19.0"},
		{Name: "aws-sdk-pricing", Version: "1.9.0"},
		{Name: "aws-sdk-quicksight", Version: "1.5.0"},
		{Name: "aws-sdk-ram", Version: "1.4.0"},
		{Name: "aws-sdk-rds", Version: "1.50.0"},
		{Name: "aws-sdk-rdsdataservice", Version: "1.4.0"},
		{Name: "aws-sdk-redshift", Version: "1.23.0"},
		{Name: "aws-sdk-rekognition", Version: "1.22.0"},
		{Name: "aws-sdk-resourcegroups", Version: "1.14.0"},
		{Name: "aws-sdk-resourcegroupstaggingapi", Version: "1.9.0"},
		{Name: "aws-sdk-resources", Version: "3.41.0"},
		{Name: "aws-sdk-robomaker", Version: "1.5.0"},
		{Name: "aws-sdk-route53", Version: "1.22.0"},
		{Name: "aws-sdk-route53domains", Version: "1.11.0"},
		{Name: "aws-sdk-route53resolver", Version: "1.4.0"},
		{Name: "aws-sdk-s3", Version: "1.36.1"},
		{Name: "aws-sdk-s3control", Version: "1.4.0"},
		{Name: "aws-sdk-sagemaker", Version: "1.33.0"},
		{Name: "aws-sdk-sagemakerruntime", Version: "1.10.0"},
		{Name: "aws-sdk-secretsmanager", Version: "1.24.0"},
		{Name: "aws-sdk-securityhub", Version: "1.4.0"},
		{Name: "aws-sdk-serverlessapplicationrepository", Version: "1.15.0"},
		{Name: "aws-sdk-servicecatalog", Version: "1.20.0"},
		{Name: "aws-sdk-servicediscovery", Version: "1.12.0"},
		{Name: "aws-sdk-ses", Version: "1.18.0"},
		{Name: "aws-sdk-shield", Version: "1.13.0"},
		{Name: "aws-sdk-signer", Version: "1.9.0"},
		{Name: "aws-sdk-simpledb", Version: "1.9.0"},
		{Name: "aws-sdk-sms", Version: "1.10.0"},
		{Name: "aws-sdk-snowball", Version: "1.14.0"},
		{Name: "aws-sdk-sns", Version: "1.13.0"},
		{Name: "aws-sdk-sqs", Version: "1.13.0"},
		{Name: "aws-sdk-ssm", Version: "1.43.0"},
		{Name: "aws-sdk-states", Version: "1.14.0"},
		{Name: "aws-sdk-storagegateway", Version: "1.21.0"},
		{Name: "aws-sdk-support", Version: "1.9.0"},
		{Name: "aws-sdk-swf", Version: "1.9.0"},
		{Name: "aws-sdk-textract", Version: "1.4.0"},
		{Name: "aws-sdk-transcribeservice", Version: "1.19.0"},
		{Name: "aws-sdk-transcribestreamingservice", Version: "1.2.0"},
		{Name: "aws-sdk-transfer", Version: "1.5.0"},
		{Name: "aws-sdk-translate", Version: "1.11.0"},
		{Name: "aws-sdk-waf", Version: "1.16.0"},
		{Name: "aws-sdk-wafregional", Version: "1.17.0"},
		{Name: "aws-sdk-workdocs", Version: "1.12.0"},
		{Name: "aws-sdk-worklink", Version: "1.4.0"},
		{Name: "aws-sdk-workmail", Version: "1.11.0"},
		{Name: "aws-sdk-workspaces", Version: "1.19.0"},
		{Name: "aws-sdk-xray", Version: "1.13.0"},
		{Name: "aws-sigv2", Version: "1.0.1"},
		{Name: "aws-sigv4", Version: "1.1.0"},
		{Name: "builder", Version: "3.2.3"},
		{Name: "coderay", Version: "1.1.2"},
		{Name: "concurrent-ruby", Version: "1.1.5"},
		{Name: "crass", Version: "1.0.4"},
		{Name: "dotenv", Version: "2.7.2"},
		{Name: "erubi", Version: "1.8.0"},
		{Name: "faker", Version: "1.9.3"},
		{Name: "faraday", Version: "0.15.4"},
		{Name: "ffi", Version: "1.10.0"},
		{Name: "globalid", Version: "0.4.2"},
		{Name: "i18n", Version: "1.6.0"},
		{Name: "jaro_winkler", Version: "1.5.2"},
		{Name: "jmespath", Version: "1.4.0"},
		{Name: "json", Version: "2.2.0"},
		{Name: "loofah", Version: "2.2.3"},
		{Name: "mail", Version: "2.7.1"},
		{Name: "marcel", Version: "0.3.3"},
		{Name: "method_source", Version: "0.9.2"},
		{Name: "mimemagic", Version: "0.3.3"},
		{Name: "mini_mime", Version: "1.0.1"},
		{Name: "mini_portile2", Version: "2.4.0"},
		{Name: "minitest", Version: "5.11.3"},
		{Name: "multi_json", Version: "1.13.1"},
		{Name: "multipart-post", Version: "2.0.0"},
		{Name: "mustermann", Version: "1.0.3"},
		{Name: "nio4r", Version: "2.3.1"},
		{Name: "nokogiri", Version: "1.10.3"},
		{Name: "parallel", Version: "1.17.0"},
		{Name: "parser", Version: "2.6.3.0"},
		{Name: "pry", Version: "0.12.2"},
		{Name: "psych", Version: "3.1.0"},
		{Name: "rack", Version: "2.0.7"},
		{Name: "rack-protection", Version: "2.0.5"},
		{Name: "rack-test", Version: "1.1.0"},
		{Name: "rails", Version: "5.2.3"},
		{Name: "rails-dom-testing", Version: "2.0.3"},
		{Name: "rails-html-sanitizer", Version: "1.0.4"},
		{Name: "railties", Version: "5.2.3"},
		{Name: "rainbow", Version: "3.0.0"},
		{Name: "rake", Version: "12.3.2"},
		{Name: "rb-fsevent", Version: "0.10.3"},
		{Name: "rb-inotify", Version: "0.10.0"},
		{Name: "rubocop", Version: "0.67.2"},
		{Name: "ruby-progressbar", Version: "1.10.0"},
		{Name: "sass", Version: "3.7.4"},
		{Name: "sass-listen", Version: "4.0.0"},
		{Name: "sinatra", Version: "2.0.5"},
		{Name: "sprockets", Version: "3.7.2"},
		{Name: "sprockets-rails", Version: "3.2.1"},
		{Name: "thor", Version: "0.20.3"},
		{Name: "thread_safe", Version: "0.3.6"},
		{Name: "tilt", Version: "2.0.9"},
		{Name: "tzinfo", Version: "1.2.5"},
		{Name: "unicode-display_width", Version: "1.5.0"},
		{Name: "websocket-driver", Version: "0.7.0"},
		{Name: "websocket-extensions", Version: "0.1.3"},
	}

	// docker run --name bundler --rm -it ruby:3 bash
	// bundle init
	// bundle add dotenv json faker rubocop pry
	// bundle add rails
	// bundler show | grep "*" | grep -v bundler | awk '{if(match($0, /\((.*)\)/)) printf("{Name: \""$2"\", Version: \""substr($0, RSTART+1, RLENGTH-2)"\"},\n");}'
	BundlerV2RailsV7 = []types.Library{
		{Name: "actioncable", Version: "7.0.3"},
		{Name: "actionmailbox", Version: "7.0.3"},
		{Name: "actionmailer", Version: "7.0.3"},
		{Name: "actionpack", Version: "7.0.3"},
		{Name: "actiontext", Version: "7.0.3"},
		{Name: "actionview", Version: "7.0.3"},
		{Name: "activejob", Version: "7.0.3"},
		{Name: "activemodel", Version: "7.0.3"},
		{Name: "activerecord", Version: "7.0.3"},
		{Name: "activestorage", Version: "7.0.3"},
		{Name: "activesupport", Version: "7.0.3"},
		{Name: "ast", Version: "2.4.2"},
		{Name: "builder", Version: "3.2.4"},
		{Name: "coderay", Version: "1.1.3"},
		{Name: "concurrent-ruby", Version: "1.1.10"},
		{Name: "crass", Version: "1.0.6"},
		{Name: "digest", Version: "3.1.0"},
		{Name: "dotenv", Version: "2.7.6"},
		{Name: "erubi", Version: "1.10.0"},
		{Name: "faker", Version: "2.21.0"},
		{Name: "globalid", Version: "1.0.0"},
		{Name: "i18n", Version: "1.10.0"},
		{Name: "json", Version: "2.6.2"},
		{Name: "loofah", Version: "2.18.0"},
		{Name: "mail", Version: "2.7.1"},
		{Name: "marcel", Version: "1.0.2"},
		{Name: "method_source", Version: "1.0.0"},
		{Name: "mini_mime", Version: "1.1.2"},
		{Name: "minitest", Version: "5.16.0"},
		{Name: "net-imap", Version: "0.2.3"},
		{Name: "net-pop", Version: "0.1.1"},
		{Name: "net-protocol", Version: "0.1.3"},
		{Name: "net-smtp", Version: "0.3.1"},
		{Name: "nio4r", Version: "2.5.8"},
		{Name: "nokogiri", Version: "1.13.6"},
		{Name: "parallel", Version: "1.22.1"},
		{Name: "parser", Version: "3.1.2.0"},
		{Name: "pry", Version: "0.14.1"},
		{Name: "racc", Version: "1.6.0"},
		{Name: "rack", Version: "2.2.3.1"},
		{Name: "rack-test", Version: "1.1.0"},
		{Name: "rails", Version: "7.0.3"},
		{Name: "rails-dom-testing", Version: "2.0.3"},
		{Name: "rails-html-sanitizer", Version: "1.4.3"},
		{Name: "railties", Version: "7.0.3"},
		{Name: "rainbow", Version: "3.1.1"},
		{Name: "rake", Version: "13.0.6"},
		{Name: "regexp_parser", Version: "2.5.0"},
		{Name: "rexml", Version: "3.2.5"},
		{Name: "rubocop", Version: "1.30.1"},
		{Name: "rubocop-ast", Version: "1.18.0"},
		{Name: "ruby-progressbar", Version: "1.11.0"},
		{Name: "strscan", Version: "3.0.3"},
		{Name: "thor", Version: "1.2.1"},
		{Name: "timeout", Version: "0.3.0"},
		{Name: "tzinfo", Version: "2.0.4"},
		{Name: "unicode-display_width", Version: "2.1.0"},
		{Name: "websocket-driver", Version: "0.7.5"},
		{Name: "websocket-extensions", Version: "0.1.5"},
		{Name: "zeitwerk", Version: "2.6.0"},
	}
)
