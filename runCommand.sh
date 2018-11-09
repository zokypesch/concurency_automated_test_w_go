#!/bin/sh
# file : /scripts/automation-test
#running automation test

echo "-------------------------- Run Automate Test --------------------------"
source ~/.zshrc
cd /Users/maulana/python/robot/KLBers/KalbeStore-Automation/case
workon env1
yourfilenames=`ls ./*.robot`
for eachfile in $yourfilenames
do
   PYTHONPATH=$PYTHON_LOCATION:. pybot --variable BROWSER:chrome --variable VERSION:latest --variable PLATFORM:MAC --outputdir ./report $eachfile
done
cp -r ./report/. ~/docker/phalcon/public/report_test
# curl --header "Content-Type: application/json" \
#   --request POST \
#   --data '{"username":"xyz","password":"xyz"}' \
#   http://localhost:3000/api/login

