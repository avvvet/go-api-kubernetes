pipeline {
    agent { label 'docker-agent' }  // connected jenkins agent

    environment {
        DOCKERHUB_CREDENTIALS = credentials('docker-hub')  // Docker Hub credentials
        DOCKER_IMAGE = "awet/go-api:latest"  //  image name
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/avvvet/go-api-kubernetes.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t $DOCKER_IMAGE ."
                }
            }
        }

        stage('Push Image to Docker Hub') {
            steps {
                script {
                    sh "echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
                    sh "docker push $DOCKER_IMAGE"
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                script {
                    sh "helm upgrade --install go-api ./helm/go-api --set image.repository=awet/go-api --set image.tag=latest"
                }
            }
        }
    }
}
