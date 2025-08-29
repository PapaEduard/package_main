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
                    docker build -t "${REPO}:${PRJ_NAME}-${BUILD_NUMBER}" .
                    """
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    sh"""
                    docker login -u Edy2010 -p ${env.DOCKER_TOKEN}
                    docker push "${env.REPO}:${env.PRJ_NAME}-${env.BUILD_NUMBER}"
                    """
                    }
                }
            }
        }
    }
}

