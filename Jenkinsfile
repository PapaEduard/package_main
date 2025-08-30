pipeline {
    agent any

    environment {
        REPO = 'edy2010/main-package'
        DOCKER_TOKEN = credentials('dok-tok')
        PRJ_NAME = 'metric'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh"""
                    docker build -t "${REPO}:${PRJ_NAME}-${env.BUILD_NUMBER}" .
                    """
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    sh"""
                    docker login -u edy2010 -p ${env.DOCKER_TOKEN}
                    docker push "${REPO}:${PRJ_NAME}-${env.BUILD_NUMBER}"
                    """
                }
            }
        }
        stage('Deploy'){
            steps{
                script{
                    sh"""
                    docker pull "${REPO}:${PRJ_NAME}-${env.BUILD_NUMBER}"
                    docker stop "${PRJ_NAME}" || true
                    docker rm "${PJR_NAME}" || true
                    docker run -d -it --name "${PRJ_NAME}" 
                    """
                }
            }
        }
    }
}

