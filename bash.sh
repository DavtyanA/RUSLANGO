for i in $( aws s3 ls s3://ruslanbot | awk '{print $4}' ); do aws s3 mv s3://ruslanbot/$i s3://ruslanbot/`echo $i | ToLower.cmd`; done