pipeline {
    agent { label 'docker-agent' }  // connected jenkins agent

    environment {
        DOCKERHUB_CREDENTIALS = credentials('docker-hub')  // Docker Hub credentials
        DOCKER_IMAGE = "awet/go-api:${env.GIT_COMMIT}"  // image name based on commit hash
    }

    stages {
        stage('Clone Repository') {
            steps {
                // Clone the repository and checkout the latest commit from the 'main' branch
                git branch: 'main', url: 'https://github.com/avvvet/go-api-kubernetes.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Build the Docker image using the Dockerfile in the repository
                    sh "docker build -t $DOCKER_IMAGE ."
                }
            }
        }

        stage('Push Image to Docker Hub') {
            steps {
                script {
                    // Login to Docker Hub using Jenkins credentials and push the image
                    sh "echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
                    sh "docker push $DOCKER_IMAGE"
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                script {
                    // Deploy or upgrade the application in Kubernetes using Helm with the new image tag
                    sh "helm upgrade --install go-api ./helm --set image.repository=awet/go-api --set image.tag=${env.GIT_COMMIT}"
                }
            }
        }
    }
}
