pipeline {
    agent any

    triggers {
        pollSCM '*/1 * * * *'
    }

    stages {
        stage('Build') {
            steps {
                echo "Building..."
                // Combine setting PATH and running commands in a single bash session
                sh '''
                export PATH=$PATH:/usr/local/go/bin
                echo "Go version:"
                go version
                '''
            }
        }
        stage('Test') {
            steps {
                echo "Testing..."
                // Combine setting PATH and running commands in a single bash session
                sh '''
                echo "Running tests..."
                '''
            }
        }
        stage('Deliver') {
            steps {
                echo "Delivering..."
                // Combine setting PATH and running commands in a single bash session
                sh '''
                echo "Running main.go..."
                go run main.go
                '''
            }
        }
    }
}
