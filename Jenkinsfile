pipeline {
    agent any
    
    triggers {
        pollSCM '*/1 * * * *'
    }

    stages {
        stage('Build') {
            steps {
                echo "Building.."
                sh 'export PATH=$PATH:/usr/local/go/bin'
                // export PATH=$PATH:/usr/local/go/bin
            }
        }
        stage('Test') {
            steps {
                echo "Testing.."
                // sh 'bash -c "go run main.go"'
                sh 'go version'
                echo "doing test stuff.."
                
            }
        }
        stage('Deliver') {
            steps {
                echo 'Deliver....'
                // go run main.go
                sh 'go run main.go'
              
                echo "doing delivery stuff.."
                
            }
        }
    }
}