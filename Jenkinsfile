pipeline {
    agent any

    triggers {
        pollSCM '*/1 * * * *'
    }

    stages {
        stage('Build') {
            steps {
                echo "Building..."
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
                sh '''
                echo "Running tests..."
                '''
            }
        }
        stage('Deliver') {
            steps {
                echo "Delivering..."
                sh '''
                export PATH=$PATH:/usr/local/go/bin
                echo "Running main.go..."
                go run main.go
                '''
            }
        }
    }
}
