pipeline {
    agent any

    environment {
        IMAGE_NAME = 'metrics-app'
        DOCKER_REGISTRY = 'edy2010' // e.g. 'myregistry.com' if you use one
        IMAGE_TAG = "${env.BUILD_NUMBER}"
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
                    def fullImageName = DOCKER_REGISTRY ? "${DOCKER_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}" : "${IMAGE_NAME}:${IMAGE_TAG}"
                    docker.build(fullImageName)
                }
            }
        }

        stage('Push Docker Image') {
            when {
                expression { return env.DOCKER_REGISTRY?.trim() }
            }
            steps {
                script {
                    def fullImageName = "${DOCKER_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
                    docker.withRegistry("https://${DOCKER_REGISTRY}", 'docker-credentials-id') {
                        docker.image(fullImageName).push()
                    }
                }
            }
        }
    }

    post {
        success {
            echo "Docker image built successfully: ${IMAGE_NAME}:${IMAGE_TAG}"
        }
        failure {
            echo "Build failed!"
        }
    }
}

