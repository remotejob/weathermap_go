kubectl create -f  deployment.yaml

kubectl replace -f  deployment.yaml

kubectl expose deployment weathermap-go --type="LoadBalancer"
