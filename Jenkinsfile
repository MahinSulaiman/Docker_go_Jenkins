pipeline {
    agent any
    
    triggers {
        pollSCM '*/1 * * * *'
    }

    stages {
        stage('Build') {
            steps {
                echo "Building.."
                // sh 'bash -c "export PATH=$PATH:/usr/local/go/bin"'
                export PATH=$PATH:/usr/local/go/bin
            }
        }
        stage('Test') {
            steps {
                echo "Testing.."
                // sh 'bash -c "go run main.go"'
                sh 'bash -c "go version"'
                echo "doing test stuff.."
                
            }
        }
        stage('Deliver') {
            steps {
                echo 'Deliver....'
                go run main.go
                // sh 'bash -c "go run main.go"'
                sh '''
                echo "doing delivery stuff.."
                '''
            }
        }
    }
}