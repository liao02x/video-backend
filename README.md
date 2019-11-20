# Video-backend

My test project for golang backend.

What we have now:

- Dockerized
- Github Action
    - Build the docker image
    - Push to AWS ECR
    - Trigger the Jenkins webhook
- Jenkins
    - Triggered by the webhook
    - pull image from AWS ECR
    - run the image in AWS EC2